package controllers

import (
	"github.com/Xhofe/QRPay/models"
	"github.com/gofiber/fiber/v2"
)

func Theme(c *fiber.Ctx) error {
	return c.Render("themes/"+c.Params("theme"),Resp{Code: 200,Msg: "success",Data: models.QRCode{
		Link:      "Xhofe",
		AliPay:    "https://qr.alipay.com/fkx06806bwayjre5fdrdd8e",
		QQPay:     "https://i.qianbao.qq.com/wallet/sqrcode.htm?m=tenpay&a=1&u=927625802&ac=CAEQyuSpugMYoPzT8gU%3D_xxx_sign&n=%E5%BE%AE%E5%87%89&f=wallet",
		WechatPay: "wxp://f2f0PXC9T62cioz1XPDS4PvUY6Ucy5pMrD94",
		Title:     "Xhofe",
		QQ:        "927625802",
	}})
}