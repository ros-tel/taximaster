package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	DeleteFixedDriverShiftRequest struct {
		// ИД фиксированной смены
		FixedDriverShiftID int `validate:"required"`
	}
)

// Удаление фиксированной смены водителя
func (cl *Client) DeleteFixedDriverShift(req DeleteFixedDriverShiftRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("fixed_driver_shift_id", strconv.Itoa(req.FixedDriverShiftID))

	/*
		100 Фиксированная смена водителя не найдена
	*/
	e := errorMap{
		100: ErrFixedDriverShiftNotFound,
	}

	err = cl.Post("delete_fixed_driver_shift", e, v, &response)

	return
}
