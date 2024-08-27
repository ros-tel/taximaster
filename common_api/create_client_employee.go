package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateClientEmployeeRequest struct {
		// ИД клиента
		ClientID int `json:"client_id" validate:"required"`
		// ФИО сотрудника
		Name string `json:"name" validate:"required"`

		// E-mail
		Email string `json:"email,omitempty" validate:"omitempty,email"`
		// Использовать E-mail для отправки уведомлений по заказу
		UseEmailInforming bool `json:"use_email_informing,omitempty" validate:"omitempty"`
		// Массив телефонов сотрудника
		Phones []Phone `json:"phones,omitempty" validate:"omitempty"`
	}
)

// Создание сотрудника клиента
func (cl *Client) CreateClientEmployee(req CreateClientEmployeeRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100	Клиент с ИД=CLIENT_ID не найден
		101	Клиент с ИД=ID имеет такой же номер телефона=PHONE
	*/
	e := errorMap{
		100: ErrClientNotFound,
		101: ErrClientConflictByPhone,
	}

	err = cl.PostJson("create_client_employee", e, req, &response)

	return
}
