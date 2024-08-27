package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetDriverOperationsRequest struct {
		// ИД водителя
		DriverID int `validate:"required"`
		// Начало периода
		StartTime string `validate:"required,datetime=20060102150405"`
		//Конец периода
		FinishTime string `validate:"required,datetime=20060102150405"`

		// ИД типа счета (0 - основной счет), по умолчанию 0
		AccountKind int `validate:"omitempty"`
	}

	GetDriverOperationsResponse struct {
		// ИД операции
		OperID int `json:"oper_id"`
		// Время создания операции
		OperTime string `json:"oper_time"`
		// Сумма
		Sum float64 `json:"sum"`
		// Заказ, связанный с операцией
		OrderID int `json:"order_id"`
		// Тип операции:
		// - "receipt" - приход
		// - "expense" - расход
		OperType int `json:"oper_type"`
		// Наименование
		Name string `json:"name"`
		// Комментарий
		Comment string `json:"comment"`
		// ИД типа счета
		AccountKind int `json:"account_kind"`
		// ИД операции отмены (для отмененной операции)
		CancelledByOperID int `json:"cancelled_by_oper_id"`
		// ИД отмененной операции (для операции отмены)
		CancelledOperID int `json:"cancelled_oper_id"`
	}
)

// Запрос операций по водителю
func (cl *Client) GetDriverOperations(req GetDriverOperationsRequest) (response GetDriverOperationsResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("driver_id", strconv.Itoa(req.DriverID))
	v.Add("start_time", req.StartTime)
	v.Add("finish_time", req.FinishTime)
	if req.AccountKind != 0 {
		v.Add("account_kind", strconv.Itoa(req.AccountKind))
	}

	/*
		100	Не найден водитель ИД=DRIVER_ID
		101	Не найден тип счета ИД=ACCOUNT_KIND
	*/
	e := errorMap{
		100: ErrDriverNotFound,
		101: ErrAccountTypeNotFound,
	}

	err = cl.Get("get_driver_operations", e, v, &response)

	return
}
