package common_api

import (
	"net/url"

	"github.com/ros-tel/taximaster/validator"
)

type (
	SendSmsRequest struct {
		// Номер телефона
		Phone string `validate:"required,max=30"`
		// Текст СМС
		Message string `validate:"required"`
	}
)

// Создание задачи СМС серверу
func (cl *Client) SendSms(req SendSmsRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("phone", req.Phone)
	v.Add("message", req.Message)

	err = cl.Post("send_sms", nil, v, &response)

	return
}
