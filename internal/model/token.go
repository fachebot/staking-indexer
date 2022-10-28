package model

import (
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

// Token ERC20代币信息
type Token struct {
	gorm.Model
	Contract common.Address `gorm:"column:contract;uniqueIndex;not null"`
	Name     string         `gorm:"column:symbol;size:32;not null"`
	Symbol   string         `gorm:"column:symbol;size:32;not null"`
	Decimals uint8          `gorm:"column:decimals;not null"`
}

func (Token) TableName() string {
	return "tokens"
}
