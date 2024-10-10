package service

import (
	"github.com/Fi44er/ton-backend/database"
	"github.com/Fi44er/ton-backend/dto"
	"github.com/Fi44er/ton-backend/model"
	"github.com/gofiber/fiber/v2"
)

func Update(ctx *fiber.Ctx, req *dto.Record) error {
	var records []model.Record

	if err := database.DB.Db.Find(&records).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	for i := range records {
		updateData := req.Data[i]
		if err := database.DB.Db.Model(&records[i]).Updates(&updateData).Error; err != nil {
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
	return ctx.Status(200).JSON(records)
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
