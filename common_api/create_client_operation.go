package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	CreateClientOperationRequest struct {
		// ИД клиента
		ClientID int `validate:"required"`
		// Сумма
		Sum float64 `validate:"required"`
		// Тип операции:
		// - receipt - приход
		// - expense - расход
		OperType string `validate:"required,eq=receipt|eq=expense"`

		// Время создания операции (если не указано, текущее)
		OperTime string `validate:"omitempty,datetime=20060102150405"`
		// Комментарий
		Comment string `validate:"omitempty"`
		// Тип оплаты:
		// - cash - наличный
		// - nocash - безналичный
		PayType string `validate:"omitempty,eq=cash|eq=nocash"`
		// Операция по бонусному счёту. Данный параметр устарел - рекомендуется использовать "account_kind"
		BonusOper *bool `validate:"omitempty"`
		// Тип счета:
		// - 0 - Основной счет
		// - 1 - Бонусный счет
		// - Остальные - нестандартные счета
		AccountKind *int `validate:"omitempty,min=0"`
	}

	CreateClientOperationResponse struct {
		OperID int `json:"oper_id"`
	}
)

// Проведение операции по клиенту
func (cl *Client) CreateClientOperation(req CreateClientOperationRequest) (response CreateClientOperationResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("client_id", strconv.Itoa(req.ClientID))
	v.Add("sum", strconv.FormatFloat(req.Sum, 'g', -1, 64))
	v.Add("oper_type", req.OperType)
	if req.OperTime != "" {
		v.Add("oper_time", req.OperTime)
	}
	if req.Comment != "" {
		v.Add("comment", req.Comment)
	}
	if req.PayType != "" {
		v.Add("pay_type", req.PayType)
	}
	if req.BonusOper != nil {
		v.Add("bonus_oper", strconv.FormatBool(*req.BonusOper))
	}
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

	err = cl.Post("create_client_operation", e, v, &response)

	return
}
