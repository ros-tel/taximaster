package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	UpdateClientInfoRequest struct {
		// ИД клиента
		ClientID int `validate:"required"`

		// ФИО
		Name string `validate:"omitempty,max=60"`
		// Логин
		Login string `validate:"omitempty,max=60"`
		// Пароль
		Password string `validate:"omitempty,max=60"`
		// Номера телефонов (через запятую)
		Phones string `validate:"omitempty"`
		// ИД группы клиента
		ClientGroup int `validate:"omitempty"`
		// ИД клиента-родителя
		ParentID int `validate:"omitempty"`
		// Домашний адрес
		Address string `validate:"omitempty"`
		// Дата рождения
		Birthday string `validate:"omitempty,datetime=20060102150405"`
		// Пол. Может принимать значения:
		// - male - мужской
		// - female - женский
		Gender string `validate:"omitempty,eq=male|eq=female"`
		// E-mail
		Email string `validate:"omitempty,email"`
		// Использовать E-mail для отправки уведомлений по заказу
		UseEmailInforming *bool `validate:"omitempty"`
		// Комментарий
		Comment *string `validate:"omitempty"`
		// Использовать собственный счет для оплаты заказов
		UseOwnAccount *bool `validate:"omitempty"`
	}
)

// Изменение информации по клиенту
func (cl *Client) UpdateClientInfo(req UpdateClientInfoRequest) (EmptyResponse, error) {
	var response = EmptyResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("client_id", strconv.Itoa(req.ClientID))

	if req.Name != "" {
		v.Add("name", req.Name)
	}
	if req.Login != "" {
		v.Add("login", req.Login)
	}
	if req.Password != "" {
		v.Add("password", req.Password)
	}
	if req.Phones != "" {
		v.Add("phones", req.Phones)
	}
	if req.ClientGroup > 0 {
		v.Add("client_group", strconv.Itoa(req.ClientGroup))
	}
	if req.ParentID > 0 {
		v.Add("parent_id", strconv.Itoa(req.ParentID))
	}
	if req.Address != "" {
		v.Add("address", req.Address)
	}
	if req.Birthday != "" {
		v.Add("birthday", req.Birthday)
	}
	if req.Gender != "" {
		v.Add("gender", req.Gender)
	}
	if req.Email != "" {
		v.Add("email", req.Email)
	}
	if req.UseEmailInforming != nil {
		if *req.UseEmailInforming {
			v.Add("use_email_informing", "true")
		} else {
			v.Add("use_email_informing", "false")
		}
	}
	if req.Comment != nil {
		v.Add("comment", *req.Comment)
	}
	if req.UseOwnAccount != nil {
		if *req.UseOwnAccount {
			v.Add("use_own_account", "true")
		} else {
			v.Add("use_own_account", "false")
		}
	}

	err = cl.Post("update_client_info", v, &response)

	return response, err
}
