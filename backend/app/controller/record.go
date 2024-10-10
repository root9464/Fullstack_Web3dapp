package controller

import (
	"github.com/Fi44er/ton-backend/dto"
	"github.com/Fi44er/ton-backend/service"
	"github.com/gofiber/fiber/v2"
)

func Update(ctx *fiber.Ctx) error {
	req := new(dto.Record)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Не удается прочитать JSON"})
	}
	return service.Update(ctx, req)
}

func GetAllRecords(ctx *fiber.Ctx) error {
	return service.GetRecords(ctx)
}

func UrlObj(ctx *fiber.Ctx) error {
	return service.UrlObj(ctx)
}
