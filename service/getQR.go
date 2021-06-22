package service

import "github.com/Xhofe/QRPay/models"

func GetQR(link string) (*models.QRCode,error) {
	return models.Get(link)
}