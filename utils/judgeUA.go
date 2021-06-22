package utils

import "strings"

const (
	Default = iota
	Ali
	Wechat
	QQ
)

func JudgeUA(ua string) int {
	if strings.Contains(ua, "Alipay") {
		return Ali
	}
	if strings.Contains(ua, "MicroMessenger") {
		return Wechat
	}
	if strings.Contains(ua, "QQ") {
		return QQ
	}
	return Default
}