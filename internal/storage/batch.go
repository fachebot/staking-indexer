package storage

import "gorm.io/gorm"

// Tx 事务操作
type Tx func(tx *gorm.DB) error

// BatchExec 批量执行
func BatchExec(db *gorm.DB, txs []Tx) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, fn := range txs {
			if err := fn(tx); err != nil {
				return err
			}
		}
		return nil
	})
}
