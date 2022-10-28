package storage

import (
	"gorm.io/gorm"
	"staking-indexer/internal/model"
)

var tables = []interface{}{
	new(model.SyncingStat),
	new(model.Token),
}

// InitModels 初始化数据库表
// 如果表不存在将自动创建表和索引
func InitModels(db *gorm.DB) (err error) {
	for _, table := range tables {
		if db.Migrator().HasTable(table) {
			err := db.Migrator().AutoMigrate(table)
			if err != nil {
				return err
			}
			continue
		}
		if err = db.Migrator().CreateTable(table); err != nil {
			return err
		}
	}
	return nil
}
