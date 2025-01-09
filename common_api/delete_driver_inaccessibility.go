package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	DeleteDriverInaccessibilityRequest struct {
		// ИД недоступности
		DriverInaccessibilityID int `validate:"required"`
	}
)

// Регистрация клиента
func (cl *Client) DeleteDriverInaccessibility(req DeleteDriverInaccessibilityRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("driver_inaccessibility_id", strconv.Itoa(req.DriverInaccessibilityID))

	/*
		100 Недоступность водителя не найдена
	*/
	e := errorMap{
		100: ErrCarInaccessibilityNotFound,
	}

	err = cl.Post("delete_driver_inaccessibility", e, v, &response)

	return
}
