package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetCurrentOrdersRequest struct {
		// ИД клиента
		ClientID int `validate:"omitempty"`
		// ИД сотрудника (только если указан ИД клиента)
		ClientEmployeeID int `validate:"omitempty"`
		// Телефон клиента
		Phone string `validate:"omitempty,max=30"`
		// ИД экипажа
		CrewID int `validate:"omitempty"`
		// ИД водителя
		DriverID int `validate:"omitempty"`
	}

	GetCurrentOrdersResponse struct {
		// Массив с информацией по заказам
		Orders []GetOrderStateResponse `json:"orders"`
	}
)

// Запрос текущих заказов
func (cl *Client) GetCurrentOrders(req GetCurrentOrdersRequest) (GetCurrentOrdersResponse, error) {
	var response = GetCurrentOrdersResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}

	if req.ClientID > 0 {
		v.Add("client_id", strconv.Itoa(req.ClientID))
	}
	if req.ClientEmployeeID > 0 {
		v.Add("client_employee_id", strconv.Itoa(req.ClientEmployeeID))
	}
	if req.Phone != "" {
		v.Add("phone", req.Phone)
	}
	if req.CrewID > 0 {
		v.Add("crew_id", strconv.Itoa(req.CrewID))
	}
	if req.DriverID > 0 {
		v.Add("driver_id", strconv.Itoa(req.DriverID))
	}

	err = cl.Get("get_finished_orders", v, &response)

	return response, err
}
