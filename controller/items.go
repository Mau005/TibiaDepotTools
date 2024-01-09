package controller

import (
	"github.com/Mau005/TibiaDepotTools/db"
	"github.com/Mau005/TibiaDepotTools/model"
)

type ItemsController struct{}

func (ic *ItemsController) GetAllItems() (Items []model.Items, err error) {
	if err = db.DB.Find(&Items).Error; err != nil {
		return
	}
	return
}

func (ic *ItemsController) GetItems(idItems uint) (Items model.Items, err error) {
	if err = db.DB.Where("id = ?", idItems).First(&Items).Error; err != nil {
		return
	}
	return
}

func (ic *ItemsController) SaveItems(Items model.Items) (model.Items, error) {
	if err := db.DB.Save(&Items).Error; err != nil {
		return Items, err
	}
	return Items, nil
}

func (ic *ItemsController) CreateItems(Items model.Items) (model.Items, error) {
	if err := db.DB.Create(&Items).Error; err != nil {
		return Items, err
	}
	return Items, nil
}

func (ic *ItemsController) DelItems(Items model.Items) (model.Items, error) {
	if err := db.DB.Delete(&Items).Error; err != nil {
		return Items, err
	}
	return Items, nil
}
