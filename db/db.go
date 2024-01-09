package db

import (
	"github.com/Mau005/TibiaDepotTools/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Migration() {
	DB.AutoMigrate(&model.Character{}, &model.Items{})
}

func ConectionSqlite() error {
	var err error

	DB, err = gorm.Open(sqlite.Open("TibiaDepotTools.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	Migration()
	return err
}
