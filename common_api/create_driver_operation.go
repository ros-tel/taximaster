package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateDriverOperationRequest struct {
		// ИД водителя
		DriverID int `json:"driver_id" validate:"required"`
		// Сумма
		OperSum float64 `json:"oper_sum" validate:"required"`
		// Тип операции:
		// - receipt - приход
		// - expense - расход
		OperType string `json:"oper_type" validate:"required,eq=receipt|eq=expense"`

		// Наименование операции
		Name string `json:"name,omitempty" validate:"omitempty"`
		// Время создания операции (если не задано, текущее)  !! Не используется с ТМ 3.7
		OperTime string `json:"oper_time,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Комментарий
		Comment string `json:"comment,omitempty" validate:"omitempty"`
		// ИД типа счета (0 - основной счет), по умолчанию 0
		AccountKind int `json:"account_kind,omitempty" validate:"omitempty,min=0"`
	}

	CreateDriverOperationResponse struct {
		OperID int `json:"oper_id"`
	}
)

// Проведение операции по водителю
func (cl *Client) CreateDriverOperation(req CreateDriverOperationRequest) (response CreateDriverOperationResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100 Водитель не найден
		101 Не найден тип счета ИД=ACCOUNT_KIND
	*/
	e := errorMap{
		100: ErrDriverNotFound,
		101: ErrAccountTypeNotFound,
	}

	err = cl.PostJson("create_driver_operation", e, req, &response)

	return
}
