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
		// Домашний адрес
		Address string `validate:"omitempty"`
		// Дата рождения
		Birthday string `validate:"omitempty,datetime=20060102150405"`
		// Пол. Может принимать значения:
		// - male - мужской
		// - female - женский
		Gender string `validate:"omitempty,eq=male|eq=female"`
		// ИД клиента-родителя
		ParentID int `validate:"omitempty"`
		// ИД группы клиента
		ClientGroupID int `validate:"omitempty"`
		// E-mail
		Email string `validate:"omitempty,email"`
		// Использовать E-mail для отправки уведомлений по заказу
		UseEmailInforming *bool `validate:"omitempty"`
		// Комментарий
		Comment string `validate:"omitempty"`
		// Использовать собственный счет для оплаты заказов
		UseOwnAccount *bool `validate:"omitempty"`
	}
)

// Изменение информации по клиенту
func (cl *Client) UpdateClientInfo(req UpdateClientInfoRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
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
	if req.ClientGroupID > 0 {
		v.Add("client_group_id", strconv.Itoa(req.ClientGroupID))
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
		v.Add("use_email_informing", strconv.FormatBool(*req.UseEmailInforming))
	}
	if req.Comment != "" {
		v.Add("comment", req.Comment)
	}
	if req.UseOwnAccount != nil {
		v.Add("use_own_account", strconv.FormatBool(*req.UseOwnAccount))
	}

	/*
		100 Клиент с номером телефона=PHONE уже существует
		101 Клиент с ИД=ID имеет такой же номер телефона=PHONE
		102 Клиент с логином=LOGIN уже существует
		103 Группа клиента с ИД=CLIENT_GROUP_ID не найдена
		104 Клиент указанный в качестве родителя с ИД=PARENT_ID не найден
		109 Пароль клиента не соответствует политике паролей
	*/
	e := errorMap{
		100: ErrClientExistsWithPhone,
		101: ErrClientConflictByPhone,
		102: ErrClientExistsWithLogin,
		103: ErrClientGroupNotFound,
		104: ErrParentClientNotFound,
		109: ErrPasswordDoesNotComplyWithPasswordPolicy,
	}

	err = cl.Post("update_client_info", e, v, &response)

	return
}
