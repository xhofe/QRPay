package controllers

import (
	"github.com/Xhofe/QRPay/service"
	"github.com/Xhofe/QRPay/utils"
	"github.com/gofiber/fiber/v2"
)

func View(c *fiber.Ctx) error {
	link := c.Params("link")
	q, err := service.GetQR(link)
	if err != nil {
		return c.Render("error",Resp{Code: 400,Msg: err.Error(),Data: q})
	}
	if q == nil || q.Link == "" {
		q.Title = "Not Exist"
		return c.Render("error",Resp{Code: 400,Msg: "该链接不存在",Data: q})
	}
	ua := c.Get("User-Agent")
	// Render index template
	switch utils.JudgeUA(ua) {
	case utils.Ali:
		if q.AliPay == "" {
			return c.Render("error",Resp{Code: 400,Msg: "该用户未开通支付宝收款"})
		}
		return c.Redirect(q.AliPay)
	case utils.Wechat:
		return c.Render("wechat",Resp{Code: 200,Msg: "success",Data: q})
	case utils.QQ:
		return c.Render("qq",Resp{Code: 200,Msg: "success",Data: q})
	default:
		return c.Render("themes/"+q.Template,Resp{Code: 200,Msg: "success",Data: q})
	}
}