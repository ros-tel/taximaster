package pay_term_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	AddTermOperationRequest struct {
		// Тип платежной системы
		PaySystemType int `validate:"required,min=0"`
		// ИД Города
		CityID string `validate:"required,len=5"`
		// Терминальный аккаунт
		TermAccount string `validate:"required,len=5"`
		// ИД операции
		OperID string `validate:"required"`
		// Сумма
		Sum float64 `validate:"required,min=0"`
		// Время операции
		OperTime string `validate:"required,datetime=20060102150405"`
	}
)

// Проведение терминальной операции
func (cl *Client) AddTermOperation(req AddTermOperationRequest) (EmptyResponse, error) {
	var response = EmptyResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("pay_system_type", strconv.Itoa(req.PaySystemType))
	v.Add("account", req.CityID+req.TermAccount)
	v.Add("oper_id", req.OperID)
	v.Add("sum", strconv.FormatFloat(req.Sum, 'g', -1, 64))
	v.Add("oper_time", req.OperTime)

	err = cl.Post("add_term_operation", v, &response)

	return response, err
}
