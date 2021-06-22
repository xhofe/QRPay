package service

import "github.com/Xhofe/QRPay/models"

func isLinkExist(link string) (bool, error) {
	q, err := models.Get(link)
	return q != nil && q.Link != "", err
}
