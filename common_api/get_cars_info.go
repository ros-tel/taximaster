package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetCarsInfoRequest struct {
		// Включить в ответ заблокированных автомобилей (по умолчанию false)
		LockedCars *bool `validate:"omitempty"`
		// Список возвращаемых полей через запятую
		Fields string `validate:"omitempty"`
	}

	GetCarsInfoResponse struct {
		// Массив автомобилей
		CrewsInfo []GetCarInfoRequest `json:"crews_info"`
	}
)

// Запрос списка автомобилей
func (cl *Client) GetCarsInfo(req GetCarsInfoRequest) (response GetCarsInfoResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	if req.LockedCars != nil {
		v.Add("locked_cars", strconv.FormatBool(*req.LockedCars))
	}
	if req.Fields != "" {
		v.Add("fields", req.Fields)
	}

	err = cl.Get("get_cars_info", nil, v, &response)

	return
}
