package service

import "github.com/Xhofe/QRPay/models"

type CreateReq struct {
	Link      string `json:"link" validate:"required"`
	AliPay    string `json:"ali_pay"`
	QQPay     string `json:"qq_pay"`
	WechatPay string `json:"wechat_pay"`
	Template  string `json:"template" validate:"required"`
	Title     string `json:"title"`
	QQ        string `json:"qq" validate:"required"`
}

func (req *CreateReq) Create() error {
	q := models.QRCode{
		Link:      req.Link,
		AliPay:    req.AliPay,
		QQPay:     req.QQPay,
		WechatPay: req.WechatPay,
		Template:  req.Template,
		Title:     req.Title,
		QQ:        req.QQ,
	}
	return q.Create()
}
