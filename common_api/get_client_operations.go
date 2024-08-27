package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetClientOperationsRequest struct {
		// ИД клиента
		ClientID int `validate:"required"`
		// Начало периода
		StartTime string `validate:"required,datetime=20060102150405"`
		// Конец периода
		FinishTime string `validate:"required,datetime=20060102150405"`

		// Тип счета:
		// - 0 - Основной счет
		// - 1 - Бонусный счет
		// - Остальные - нестандартные счета
		AccountKind *int `validate:"omitempty,min=0"`
	}

	GetClientOperationsResponse struct {
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
		// Тип оплаты:
		// - "cash" - наличный
		// - "nocash" - безналичный
		PayType int `json:"pay_type"`
		// Тип счета:
		// 0 - Основной счет
		// 1- Бонусный счет
		// Остальные - нестандартные счета
		AccountKind int `json:"account_kind"`
		// Наименование
		Name string `json:"name"`
		// Комментарий
		Comment string `json:"comment"`
		// ИД операции отмены (для отмененной операции)
		CancelledByOperID int `json:"cancelled_by_oper_id"`
		// ИД отмененной операции (для операции отмены)
		CancelledOperID int `json:"cancelled_oper_id"`
	}
)

// Запрос операций по клиенту
func (cl *Client) GetClientOperations(req GetClientOperationsRequest) (response GetClientOperationsResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("client_id", strconv.Itoa(req.ClientID))
	v.Add("start_time", req.StartTime)
	v.Add("finish_time", req.FinishTime)
	if req.AccountKind != nil {
		v.Add("account_kind", strconv.Itoa(*req.AccountKind))
	}

	/*
		100 Не найден клиент
		101 Не найден тип счета ИД=ACCOUNT_KIND
	*/
	e := errorMap{
		100: ErrClientNotFound,
		101: ErrAccountTypeNotFound,
	}

	err = cl.Get("get_client_operations", e, v, &response)

	return
}
