package service

import (
	"fmt"
	"github.com/Xhofe/QRPay/models"
)

type CreateReq struct {
	Link      string `json:"link" validate:"required"`
	AliPay    string `json:"ali_pay" validate:"required"`
	QQPay     string `json:"qq_pay" validate:"required"`
	WechatPay string `json:"wechat_pay" validate:"required"`
	Template  string `json:"template" validate:"required"`
	Title     string `json:"title" validate:"required"`
	QQ        string `json:"qq" validate:"required"`
}

func (req *CreateReq) Create() error {
	exist, _ := isLinkExist(req.Link)
	if exist {
		return fmt.Errorf("链接'%s'已存在，请更换。", req.Link)
	}
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
