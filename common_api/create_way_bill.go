package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	CreateWayBillRequest struct {
		// Время начала
		StartTime string `validate:"required,datetime=20060102150405"`
		// Время завершения
		FinishTime string `validate:"required,datetime=20060102150405"`
		// ИД водителя
		DriverID int `validate:"required"`
		// ИД автомобиля
		CarID int `validate:"required"`

		// Номер путевого листа
		Number string `validate:"omitempty"`
		// Комментарий
		Comment string `validate:"omitempty"`
	}

	CreateWayBillResponse struct {
		// ИД созданного путевого листа
		WayBillID int `json:"way_bill_id"`
	}
)

// Создание путевого листа
func (cl *Client) CreateWayBill(req CreateWayBillRequest) (response CreateWayBillResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("start_time", req.StartTime)
	v.Add("finish_time", req.FinishTime)
	v.Add("driver_id", strconv.Itoa(req.DriverID))
	v.Add("car_id", strconv.Itoa(req.CarID))
	if req.Number != "" {
		v.Add("number", req.Number)
	}
	if req.Comment != "" {
		v.Add("comment", req.Comment)
	}

	/*
		100 Нет лицензии на использование путевых листов
		101 Не найден водитель
		102 Не найден автомобиль
	*/
	e := errorMap{
		100: ErrNoLicenseToUseWayBill,
		101: ErrDriverNotFound,
		102: ErrCarNotFound,
	}

	err = cl.Post("create_way_bill", e, v, &response)

	return
}
