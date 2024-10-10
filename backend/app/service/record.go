package service

import (
	"github.com/Fi44er/ton-backend/database"
	"github.com/Fi44er/ton-backend/dto"
	"github.com/Fi44er/ton-backend/model"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func Update(ctx *fiber.Ctx, req *dto.Record) error {
	var records []model.Record

	if err := database.DB.Db.Find(&records).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	for i := range records {
		item := req.Data[i]
		updateRecord := model.Record{
			Total:   pq.Int64Array(item.Total),
			Percent: item.Percent,
		}
		if err := database.DB.Db.Model(&records[i]).Updates(&updateRecord).Error; err != nil {
			return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "OK"})
}

func GetRecords(ctx *fiber.Ctx) error {
	var records []model.Record

	if err := database.DB.Db.Find(&records).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if len(records) == 0 {
		return ctx.Status(404).JSON(fiber.Map{"error": "Ни одной записи не найдено"})
	}

	var transformedRecords []dto.ItemRes

	for _, record := range records {
		var total int64
		for _, val := range record.Total {
			total += val
		}

		var transformedRecord dto.ItemRes
		if len(record.Total) > 1 {
			transformedRecord = dto.ItemRes{
				Name:    record.Name,
				Total:   record.Total,
				Percent: record.Percent,
			}
		} else {
			transformedRecord = dto.ItemRes{
				Name:    record.Name,
				Total:   total,
				Percent: record.Percent,
			}
		}

		transformedRecords = append(transformedRecords, transformedRecord)
	}

	return ctx.Status(200).JSON(transformedRecords)
}

func UrlObj(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(Obj{
		Url:     "https://anytapton.ru/",
		Name:    "AnyTap",
		IconUrl: "https://imgur.com/a/hoURnBO",
	})
}

type Obj struct {
	Url     string `json:"url"`
	Name    string `json:"name"`
	IconUrl string `json:"iconUrl"`
}
