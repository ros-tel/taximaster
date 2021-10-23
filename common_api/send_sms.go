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
func (cl *Client) SendSms(req SendSmsRequest) (EmptyResponse, error) {
	var response = EmptyResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("phone", req.Phone)
	v.Add("message", req.Message)

	err = cl.Post("send_sms", v, &response)

	return response, err
}
