package common_api

import (
	"net/url"

	"github.com/ros-tel/taximaster/validator"
)

type (
	AnalyzePhoneRequest struct {
		// Номер телефона
		Phone string `validate:"required"`

		// Искать среди телефонов водителей
		SearchInDriversMobile bool `validate:"omitempty"`
		// Искать среди телефонов клиентов
		SearchInClients bool `validate:"omitempty"`
		// Искать в справочнике телефонов
		SearchInPhones bool `validate:"omitempty"`
	}

	AnalyzePhoneResponse struct {
		// Может принимать значения: "driver_mobile", "client", "phone"
		PhoneType string `json:"phone_type"`
		// ИД водителя, клиента, телефона из справочника
		ID int `json:"id"`
		// ИД сотрудника клиента (если телефон найден среди телефонов клиента)
		ClientEmployeeID int `json:"client_employee_id"`
	}
)

// Анализ телефона
func (cl *Client) AnalyzePhone(req AnalyzePhoneRequest) (response AnalyzePhoneResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("phone", req.Phone)
	if req.SearchInDriversMobile {
		v.Add("search_in_drivers_mobile", "true")
	}
	if req.SearchInClients {
		v.Add("search_in_clients", "true")
	}
	if req.SearchInPhones {
		v.Add("search_in_phones", "true")
	}

	/*
		100 Телефон не найден
	*/
	e := errorMap{
		100: ErrPhoneNotFound,
	}

	err = cl.Get("analyze_phone", e, v, &response)

	return
}
