package controller

import (
	"strconv"

	"github.com/Fi44er/ton-backend/dto"
	"github.com/Fi44er/ton-backend/service"
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	dto := new(dto.Req)
	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Не удается прочитать JSON"})
	}
	return service.Create(ctx, dto)
}

func GetByWallet(ctx *fiber.Ctx) error {
	wallet := ctx.Params("wallet")
	return service.GetByWallet(ctx, wallet)
}

func GetAll(ctx *fiber.Ctx) error {
	return service.GetAll(ctx)
}

func Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	return service.Delete(ctx, id)
}
