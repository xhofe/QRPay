package controllers

import (
	"github.com/Xhofe/QRPay/service"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	var req service.CreateReq
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(Resp{
			Code: 400,
			Msg: err.Error(),
		})
	}
	if err := validate.Struct(req);err != nil {
		return c.JSON(Resp{
			Code: 400,
			Msg: err.Error(),
		})
	}
	if err := req.Create(); err != nil {
		return c.JSON(Resp{
			Code: 500,
			Msg: err.Error(),
		})
	}
	return c.JSON(Resp{Code: 200, Msg: "success"})
}