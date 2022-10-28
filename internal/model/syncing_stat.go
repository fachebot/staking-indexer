package model

import (
	"gorm.io/gorm"
)

// SyncingStat 区块同步状态
type SyncingStat struct {
	gorm.Model
	Number  int64 `gorm:"column:number;not null"`
	TxIndex int   `gorm:"column:tx_index;not null"`
}

func (SyncingStat) TableName() string {
	return "syncing_stat"
}
