package common_api

import (
	"net/url"

	"github.com/ros-tel/taximaster/validator"
)

type (
	CheckAuthorizationRequest struct {
		// Логин
		Login string `validate:"required,max=60"`
		// Пароль
		Password string `validate:"required,max=60"`
	}

	CheckAuthorizationResponse struct {
		// ИД клиента
		ClientID int `json:"client_id"`
	}
)

// Проверка авторизации
func (cl *Client) CheckAuthorization(req CheckAuthorizationRequest) (response CheckAuthorizationResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("login", req.Login)
	v.Add("password", req.Password)

	/*
		100	Не найден клиент с логином LOGIN и/или неверный пароль
	*/
	e := errorMap{
		100: ErrClientNotFound,
	}

	err = cl.Get("check_authorization", e, v, &response)

	return
}
