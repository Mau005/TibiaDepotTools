package controller

import (
	"github.com/Mau005/TibiaDepotTools/db"
	"github.com/Mau005/TibiaDepotTools/model"
	"gorm.io/gorm"
)

type ItemsController struct{}

func (ic *ItemsController) GetAllItems() (Items []model.Items, err error) {
	if err = db.DB.Find(&Items).Error; err != nil {
		return
	}
	return
}

func (ic *ItemsController) GetItems(idItems uint) (Items model.Items, err error) {
	if err = db.DB.Preload("HistoryItems", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at desc")
	}).Where("id = ?", idItems).First(&Items).Error; err != nil {
		return
	}
	return
}

func (ic *ItemsController) CreateHistoryItems(history model.HistoryItems) error {
	if err := db.DB.Create(&history).Error; err != nil {
		return err
	}
	return nil
}

func (ic *ItemsController) SaveItems(items model.Items) (model.Items, error) {
	itemOld, err := ic.GetItems(items.ID)
	if err != nil {
		return items, err
	}
	var itemsHistory model.HistoryItems
	itemsHistory.BalanceOld = itemOld.Balance
	itemsHistory.BalanceNew = items.Balance
	itemsHistory.ItemsID = items.ID

	if err := db.DB.Save(&items).Error; err != nil {
		return items, err
	}

	ic.CreateHistoryItems(itemsHistory)
	return items, nil
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

func (ic *ItemsController) GetTotalItemsCharacter(characterID uint) (count int64) {
	if err := db.DB.Where("character_id = ?", characterID).Model(&model.Items{}).Count(&count).Error; err != nil {
		return 0
	}
	return count
}

func (ic *ItemsController) OffsetItems(characterid uint, pag int) (model.Items, error) {

	var items model.Items
	index := (pag - 1) * 1
	if err := db.DB.Where("character_id = ?", characterid).Offset(index).Limit(1).First(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}
