package pgsql

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	conn, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}))
	return db, mock, err
}
