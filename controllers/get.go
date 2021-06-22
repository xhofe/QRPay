package controllers

import (
	"github.com/Xhofe/QRPay/service"
	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	link := c.Params("link")
	if link == "" {
		return c.JSON(Resp{
			Code: 400,
			Msg:  "link is required",
		})
	}
	q, err := service.GetQR(link)
	if err != nil {
		return c.JSON(Resp{
			Code: 500,
			Msg: err.Error(),
		})
	}
	return c.JSON(Resp{
		Code: 200,
		Msg:  "success",
		Data: q,
	})
}