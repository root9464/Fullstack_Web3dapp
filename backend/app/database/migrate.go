package database

import (
	"github.com/Fi44er/ton-backend/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.Record{}, &model.Body{}, &model.Header{}); err != nil {
		return err
	}

	defaultRecords := []model.Record{
		{Name: "TotalSupply", Total: 100, Percent: 5},
		{Name: "TotalEarn", Total: 200, Percent: 10},
		{Name: "TonValue", Total: 300, Percent: 15},
	}

	for _, record := range defaultRecords {
		var count int64
		db.Model(&model.Record{}).Where("name = ?", record.Name).Count(&count)
		if count == 0 {
			db.Create(&record)
		}
	}

	return nil
}
