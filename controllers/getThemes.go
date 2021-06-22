package controllers

import (
	"github.com/Xhofe/QRPay/service"
	"github.com/gofiber/fiber/v2"
)

func GetThemes(c *fiber.Ctx) error {
	themes := service.GetThemes()
	return c.JSON(Resp{Code: 200,Msg: "success",Data: themes})
}