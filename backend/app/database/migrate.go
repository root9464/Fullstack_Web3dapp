package database

import (
	"github.com/Fi44er/ton-backend/model"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.Record{}, &model.Body{}, &model.Header{}); err != nil {
		return err
	}

	defaultRecords := []model.Record{
		{Name: "TotalSupply", Total: []int64{100}, Percent: 5},
		{Name: "TotalEarn", Total: []int64{200}, Percent: 10},
		{Name: "TonValue", Total: []int64{300, 250, 200}, Percent: 15},
	}
	for _, record := range defaultRecords {
		var count int64
		db.Model(&model.Record{}).Where("name = ?", record.Name).Count(&count)
		if count == 0 {
			totalArray := pq.Array(record.Total)

			db.Exec(`INSERT INTO "records" ("name", "total", "percent") VALUES ($1, $2, $3)`, record.Name, totalArray, record.Percent)

			db.Create(&record)
		}
	}

	return nil
}
