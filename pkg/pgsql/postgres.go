package pgsql

import (
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgConfig struct {
	Host               string
	Port               int
	Database           *string
	User               string
	Password           string
	Passfile           *string `json:",optional"`
	ConnectTimeout     *int64  `json:",optional"`
	SslMode            *bool   `json:",optional"`
	SslKey             *string `json:",optional"`
	SslCert            *string `json:",optional"`
	SslRootCert        *string `json:",optional"`
	TargetSessionAttrs *string `json:",optional"`
	Service            *string `json:",optional"`
	Servicefile        *string `json:",optional"`
	TimeZone           *string `json:",optional"`
}

func (p PgConfig) ToDSN() string {
	var elems []string
	elems = append(elems, fmt.Sprintf(`host=%s`, p.Host))
	elems = append(elems, fmt.Sprintf(`port=%d`, p.Port))
	elems = append(elems, fmt.Sprintf(`user=%s`, p.User))
	elems = append(elems, fmt.Sprintf(`password=%s`, p.Password))

	if p.Database != nil {
		elems = append(elems, fmt.Sprintf(`database=%s`, *p.Database))
	}
	if p.Passfile != nil {
		elems = append(elems, fmt.Sprintf(`passfile="%s"`, *p.Passfile))
	}
	if p.ConnectTimeout != nil {
		elems = append(elems, fmt.Sprintf(`connect_timeout=%d`, *p.ConnectTimeout))
	}
	if p.SslMode != nil {
		elems = append(elems, fmt.Sprintf(`sslmode=%v`, *p.SslMode))
	}
	if p.SslKey != nil {
		elems = append(elems, fmt.Sprintf(`sslkey="%s"`, *p.SslKey))
	}
	if p.SslCert != nil {
		elems = append(elems, fmt.Sprintf(`sslcert="%s"`, *p.SslCert))
	}
	if p.SslRootCert != nil {
		elems = append(elems, fmt.Sprintf(`sslrootcert="%s"`, *p.SslRootCert))
	}
	if p.SslRootCert != nil {
		elems = append(elems, fmt.Sprintf(`target_session_attrs="%s"`, *p.TargetSessionAttrs))
	}
	if p.SslRootCert != nil {
		elems = append(elems, fmt.Sprintf(`service="%s"`, *p.Service))
	}
	if p.SslRootCert != nil {
		elems = append(elems, fmt.Sprintf(`servicefile="%s"`, *p.Servicefile))
	}
	if p.SslRootCert != nil {
		elems = append(elems, fmt.Sprintf(`TimeZone="%s"`, *p.TimeZone))
	}

	return strings.Join(elems, " ")
}

func Open(conf PgConfig, opts ...gorm.Option) (*gorm.DB, error) {
	d := conf.Database
	conf.Database = nil

	db, err := gorm.Open(postgres.Open(conf.ToDSN()), opts...)
	if err != nil {
		return nil, err
	}

	defer func() {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}()

	err = createDatabaseIfNotExists(db, *d)
	if err != nil {
		return nil, err
	}

	conf.Database = d
	return gorm.Open(postgres.Open(conf.ToDSN()), opts...)
}

func createDatabaseIfNotExists(db *gorm.DB, database string) error {
	var result struct {
		Count int `json:"count"`
	}
	sql := "SELECT COUNT(*) FROM pg_catalog.pg_database WHERE datname='%s' LIMIT 1;"
	err := db.Raw(fmt.Sprintf(sql, database)).Scan(&result).Error
	if err != nil {
		return err
	}

	if result.Count > 0 {
		return nil
	}

	sql = fmt.Sprintf("CREATE DATABASE %s ENCODING = 'UTF8';", database)
	return db.Exec(sql).Error
}
