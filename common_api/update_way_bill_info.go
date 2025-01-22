package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	UpdateWayBillInfoRequest struct {
		// ИД путевого листа
		WayBillID int `validate:"required"`

		// Время начала
		StartTime string `validate:"omitempty,datetime=20060102150405"`
		// Время завершения
		FinishTime string `validate:"omitempty,datetime=20060102150405"`
		// ИД водителя
		DriverID int `validate:"omitempty"`
		// ИД автомобиля
		CarID int `validate:"omitempty"`
		// Номер путевого листа
		Number *string `validate:"omitempty"`
		// Комментарий
		Comment *string `validate:"omitempty"`
	}
)

// Изменение путевого листа
func (cl *Client) UpdateWayBillInfo(req UpdateWayBillInfoRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("way_bill_id", strconv.Itoa(req.WayBillID))

	if req.StartTime != "" {
		v.Add("start_time", req.StartTime)
	}
	if req.FinishTime != "" {
		v.Add("finish_time", req.FinishTime)
	}
	if req.DriverID > 0 {
		v.Add("driver_id", strconv.Itoa(req.DriverID))
	}
	if req.CarID > 0 {
		v.Add("car_id", strconv.Itoa(req.CarID))
	}
	if req.Number != nil {
		v.Add("number", *req.Number)
	}
	if req.Comment != nil {
		v.Add("comment", *req.Comment)
	}

	/*
		100	Нет лицензии на использование путевых листов
		101	Не найден водитель
		102	Не найден автомобиль
		103	Не найден путевой лист
	*/
	e := errorMap{
		100: ErrNoLicenseToUseWayBill,
		101: ErrDriverNotFound,
		102: ErrCarNotFound,
		103: ErrWayBillNotFound,
	}

	err = cl.Post("update_way_bill_info", e, v, &response)

	return
}
