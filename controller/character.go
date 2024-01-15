package controller

import (
	"github.com/Mau005/TibiaDepotTools/db"
	"github.com/Mau005/TibiaDepotTools/model"
	"gorm.io/gorm"
)

type CharacterController struct{}

func (cc *CharacterController) GetAllCharacter() (characters []model.Character, err error) {
	if err = db.DB.Find(&characters).Error; err != nil {
		return
	}
	return
}

func (cc *CharacterController) GetCharacter(idCharacter uint) (character model.Character, err error) {
	if err = db.DB.Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Order("updated_at asc")
	}).Where("id = ?", idCharacter).First(&character).Error; err != nil {
		return
	}
	return
}

func (cc *CharacterController) SaveCharacter(character model.Character) (model.Character, error) {
	if err := db.DB.Save(&character).Error; err != nil {
		return character, err
	}
	return character, nil
}

func (cc *CharacterController) CreateCharacter(character model.Character) (model.Character, error) {
	if err := db.DB.Create(&character).Error; err != nil {
		return character, err
	}
	return character, nil
}

func (cc *CharacterController) DelCharacter(character model.Character) (model.Character, error) {
	if err := db.DB.Delete(&character).Error; err != nil {
		return character, err
	}
	return character, nil
}
