package model

import "gorm.io/gorm"

type Items struct {
	gorm.Model
	Name        string
	CostItem    uint
	Balance     uint
	CharacterID uint
}
