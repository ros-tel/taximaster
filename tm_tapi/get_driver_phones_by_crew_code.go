package tm_tapi

import (
	"net/url"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetDriverPhonesByCrewCodeRequest struct {
		// Позывной экипажа
		CrewCode string `validate:"required"`
	}

	GetDriverPhonesByCrewCodeResponse struct {
		// Основной телефон водителя
		MobilePhone string `xml:"mobile_phone"`
		// Неосновной телефон водителя
		HomePhone string `xml:"home_phone"`
	}
)

// Запрос телефонов водителя по позывному экипажа
func (cl *Client) GetDriverPhonesByCrewCode(req GetDriverPhonesByCrewCodeRequest) (GetDriverPhonesByCrewCodeResponse, error) {
	var response = GetDriverPhonesByCrewCodeResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("crew_code", req.CrewCode)

	err = cl.Get("get_driver_phones_by_crew_code", v, &response)

	return response, err
}
