package common_api

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetFinishedOrdersRequest struct {
		// Начало периода
		StartTime string `validate:"required,datetime=20060102150405"`
		// Конец периода
		FinishTime string `validate:"required,datetime=20060102150405"`

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
		// Тип состояния заказа
		// Может принимать значения:
		// - "all" - все
		// - "finished" - выполненные
		// - "aborted" - прекращенные
		StateType string `validate:"omitempty,eq=all|eq=finished|eq=aborted"`
		// Список ИД состояний заказа, пример: []int{1, 2, 3}
		StateIDs []int `validate:"omitempty"`
		// Список возвращаемых полей через запятую
		Fields string `validate:"omitempty"`
	}

	GetFinishedOrdersResponse struct {
		// Массив с информацией по заказам
		Orders []GetCurrentOrdersArray `json:"orders"`
	}
)

// Запрос выполненных заказов
func (cl *Client) GetFinishedOrders(req GetFinishedOrdersRequest) (response GetFinishedOrdersResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}

	v.Add("start_time", req.StartTime)
	v.Add("finish_time", req.FinishTime)
	if req.ClientID != 0 {
		v.Add("client_id", strconv.Itoa(req.ClientID))
	}
	if req.ClientEmployeeID != 0 {
		v.Add("client_employee_id", strconv.Itoa(req.ClientEmployeeID))
	}
	if req.Phone != "" {
		v.Add("phone", req.Phone)
	}
	if req.CrewID != 0 {
		v.Add("crew_id", strconv.Itoa(req.CrewID))
	}
	if req.DriverID != 0 {
		v.Add("driver_id", strconv.Itoa(req.DriverID))
	}
	if req.StateType != "" {
		v.Add("state_type", req.StateType)
	}
	if len(req.StateIDs) != 0 {
		stringSlice := make([]string, len(req.StateIDs))

		for i, num := range req.StateIDs {
			stringSlice[i] = strconv.Itoa(num)
		}

		v.Add("state_ids", strings.Join(stringSlice, ";"))
	}
	if req.Fields != "" {
		v.Add("fields", req.Fields)
	}

	err = cl.Get("get_finished_orders", nil, v, &response)

	return
}
