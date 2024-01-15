package model

import "gorm.io/gorm"

type Items struct {
	gorm.Model
	Name         string
	CostItem     uint
	Balance      uint
	CharacterID  uint
	HistoryItems []HistoryItems `gorm:"foreignKey:ItemsID"`
}

type HistoryItems struct {
	gorm.Model
	BalanceOld uint
	BalanceNew uint
	ItemsID    uint
}
