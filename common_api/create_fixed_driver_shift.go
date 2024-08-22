package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	CreateFixedDriverShiftRequest struct {
		// ИД водителя
		DriverID int `json:"driver_id" validate:"required"`
		// Время начала
		StartTime string `json:"start_time" validate:"required,datetime=20060102150405"`
		// Время завершения
		FinishTime string `json:"finish_time" validate:"required,datetime=20060102150405"`
	}

	CreateFixedDriverShiftResponse struct {
		// ИД фиксированной смены
		FixedDriverShiftID int `json:"fixed_driver_shift_id"`
	}
)

// Создание фиксированной смены водителя
func (cl *Client) CreateFixedDriverShift(req CreateFixedDriverShiftRequest) (response CreateFixedDriverShiftResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("driver_id", strconv.Itoa(req.DriverID))
	v.Add("start_time", req.StartTime)
	v.Add("finish_time", req.FinishTime)

	/*
		100	Не найден водитель
		101	Время начала фиксированной смены должно быть меньше времени завершения
		102	Создаваемая смена пересекается по времени с уже существующей сменой данного водителя
	*/
	e := errorMap{
		100: ErrDriverNotFound,
		101: ErrTimeRange,
		103: ErrIntersectionDriverShift,
	}

	err = cl.Post("create_fixed_driver_shift", e, v, &response)

	return
}
