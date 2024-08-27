package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	RegisterClientRequest struct {
		// ФИО
		Name string `validate:"required,max=60"`
		// Логин
		Login string `validate:"required,max=60"`
		// Пароль
		Password string `validate:"required,max=60"`
		// Номера телефонов (через запятую)
		Phones string `validate:"required"`

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
		Comment string `validate:"omitempty"`
		// Использовать собственный счет для оплаты заказов
		UseOwnAccount *bool `validate:"omitempty"`
	}

	RegisterClientResponse struct {
		ClientID int `json:"client_id"`
	}
)

// Регистрация клиента
func (cl *Client) RegisterClient(req RegisterClientRequest) (response RegisterClientResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("name", req.Name)
	v.Add("login", req.Login)
	v.Add("password", req.Password)
	v.Add("phones", req.Phones)
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
		v.Add("use_email_informing", strconv.FormatBool(*req.UseEmailInforming))
	}
	if req.Comment != "" {
		v.Add("comment", req.Comment)
	}
	if req.UseOwnAccount != nil {
		v.Add("use_own_account", strconv.FormatBool(*req.UseOwnAccount))
	}

	/*
		100 Клиент с ИД=ID имеет такой же номер телефона=PHONE
		101 Клиент с логином=LOGIN уже существует
		102 Группа клиента с ИД=CLIENT_GROUP не найдена
		103 Клиент указанный в качестве родителя с ИД=PARENT_ID не найден
		109	Пароль клиента не соответствует политике паролей
	*/
	e := errorMap{
		100: ErrClientConflictByPhone,
		101: ErrClientExistsWithLogin,
		102: ErrClientGroupNotFound,
		103: ErrParentClientNotFound,
		109: ErrPasswordDoesNotComplyWithPasswordPolicy,
	}

	err = cl.Post("register_client", e, v, &response)

	return
}
