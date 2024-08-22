package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	CreateWayBillCheckRequest struct {
		// ИД путевого листа (должен быть задан либо ИД либо номер)
		WayBillID int `validate:"omitempty"`
		// Номер путевого листа (должен быть задан либо ИД либо номер)
		WayBillNumber string `validate:"omitempty"`
		// Тип осмотра ("med/tech")
		Kind string `validate:"required,eq=med|eq=tech"`
		// Имя пользователя
		UserName string `validate:"required"`
		// Результат осмотра
		Success bool `validate:"required"`

		// Номер осмотра
		Number string `validate:"omitempty"`
		// Комментарий
		Comment string `validate:"omitempty"`
	}
)

// Создание осмотра по путевому листу
func (cl *Client) CreateWayBillCheck(req CreateWayBillCheckRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	if req.WayBillID != 0 {
		v.Add("way_bill_id", strconv.Itoa(req.WayBillID))
	}
	if req.WayBillNumber != "" {
		v.Add("way_bill_number", req.WayBillNumber)
	}
	v.Add("kind", req.Kind)
	v.Add("user_name", req.UserName)
	v.Add("success", strconv.FormatBool(req.Success))
	if req.Number != "" {
		v.Add("number", req.Number)
	}
	if req.Comment != "" {
		v.Add("comment", req.Comment)
	}

	/*
		100 Нет лицензии на использование путевых листов
		101 Не найден путевой лист
	*/
	e := errorMap{
		100: ErrNoLicenseToUseWayBill,
		101: ErrWayBillNotFound,
	}

	err = cl.Post("create_way_bill_check", e, v, &response)

	return
}
