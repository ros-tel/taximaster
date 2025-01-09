package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	DeleteCarInaccessibilityRequest struct {
		// ИД недоступности
		CarInaccessibilityID int `validate:"required"`
	}
)

// Регистрация клиента
func (cl *Client) DeleteCarInaccessibility(req DeleteCarInaccessibilityRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("car_inaccessibility_id", strconv.Itoa(req.CarInaccessibilityID))

	/*
		100 Недоступность автомобиля не найдена
	*/
	e := errorMap{
		100: ErrCarInaccessibilityNotFound,
	}

	err = cl.Post("delete_car_inaccessibility", e, v, &response)

	return
}
