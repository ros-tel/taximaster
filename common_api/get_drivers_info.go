package common_api

import (
	"net/url"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetDriversInfoRequest struct {
		// Включить в ответ запроса заблокированных водителей
		LockedDrivers bool `validate:"omitempty"`
		// Включить в ответ запроса уволенных водителей
		DismissedDrivers bool `validate:"omitempty"`
	}

	GetDriversInfoResponse struct {
		// Массив с информацией о водителях
		DriversInfo []GetDriverInfoResponse `json:"drivers_info"`
	}
)

// Запрос списка водителей
func (cl *Client) GetDriversInfo(req GetDriversInfoRequest) (GetDriversInfoResponse, error) {
	var response = GetDriversInfoResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	if req.LockedDrivers {
		v.Add("locked_drivers", "true")
	}
	if req.DismissedDrivers {
		v.Add("dismissed_drivers", "true")
	}

	err = cl.Get("get_drivers_info", errorMap{}, v, &response)

	return response, err
}
