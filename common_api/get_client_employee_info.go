package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetClientEmployeeInfoRequest struct {
		// ИД сотрудника клиента
		ClientEmployeeID int `validate:"required"`
	}

	GetClientEmployeeInfoResponse struct {
		// ИД сотрудника
		ClientEmployeeID int `json:"client_employee_id"`
		// ИД клиента
		ClientID int `json:"client_id"`
		// ФИО сотрудника
		Name string `json:"name"`
		// Признак удаленного сотрудника
		IsDeleted bool `json:"is_deleted"`
		// E-mail
		Email string `json:"email"`
		// Использовать E-mail для отправки уведомлений по заказу
		UseEmailInforming bool `json:"use_email_informing"`
		// Группа экипажей по умолчанию
		DefaultCrewGroup int `json:"default_crew_group"`
		// Массив телефонов сотрудника клиента
		Phones []Phone `json:"phones"`
	}
)

// Запрос информации по сотруднику клиента
func (cl *Client) GetClientEmployeeInfo(req GetClientEmployeeInfoRequest) (response GetClientEmployeeInfoResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("client_employee_id", strconv.Itoa(req.ClientEmployeeID))

	/*
		100	Не найден сотрудник клиента ИД=CLIENT_EMPLOYEE_ID
	*/
	e := errorMap{
		100: ErrCustomerClientNotFound,
	}

	err = cl.Get("get_client_employee_info", e, v, &response)

	return
}
