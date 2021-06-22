package models

type QRCode struct {
	Link      string `gorm:"primaryKey"`
	AliPay    string
	QQPay     string
	WechatPay string
	Template  string
	Title     string
	QQ        string
}

func (q *QRCode) Create() error {
	return DB.Create(q).Error
}

func Get(link string) (*QRCode, error) {
	var q QRCode
	err := DB.Where(&QRCode{Link: link}).Limit(1).Find(&q).Error
	//err := DB.Where(&QRCode{Link: link}).First(&q).Error
	return &q, err
}
