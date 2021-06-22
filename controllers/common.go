package controllers

import "github.com/go-playground/validator"

type Resp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

var validate = validator.New()