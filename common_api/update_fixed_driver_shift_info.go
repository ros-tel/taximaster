package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	UpdateFixedDriverShiftInfoRequest struct {
		// ИД фиксированной смены
		FixedDriverShiftID int `validate:"required"`

		// ИД водителя
		DriverID int `validate:"omitempty"`
		// Время начала
		StartTime string `validate:"omitempty,datetime=20060102150405"`
		// Время завершения
		FinishTime string `validate:"omitempty,datetime=20060102150405"`
	}
)

// Изменение фиксированной смены водителя
func (cl *Client) UpdateFixedDriverShiftInfo(req UpdateFixedDriverShiftInfoRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("fixed_driver_shift_id", strconv.Itoa(req.FixedDriverShiftID))

	if req.DriverID > 0 {
		v.Add("driver_id", strconv.Itoa(req.DriverID))
	}
	if req.StartTime != "" {
		v.Add("start_time", req.StartTime)
	}
	if req.FinishTime != "" {
		v.Add("finish_time", req.FinishTime)
	}

	/*
		100	Не найден водитель
		101	Время начала фиксированной смены должно быть меньше времени завершения
		102	Создаваемая смена пересекается по времени с уже существующей сменой данного водителя
		103	Не найдена фиксированная смена водителя
	*/
	e := errorMap{
		100: ErrDriverNotFound,
		101: ErrTimeRange,
		102: ErrIntersectionDriverShift,
		103: ErrFixedDriverShiftNotFound,
	}

	err = cl.Post("update_fixed_driver_shift_info", e, v, &response)

	return
}
