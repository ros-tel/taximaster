package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	SetCrewsCoordsRequest struct {
		// Массив координат экипажей
		CrewCoords []struct {
			// ИД экипажа
			CrewID int `json:"crew_id"`
			// GPS идентификатор (если не задан ИД экипажа). Данный параметр устарел - не применяется, начиная с версии 3.9
			GpsID int `json:"gps_id"`
			// Широта
			Lat float64 `json:"lat"`
			// Долгота
			Lon float64 `json:"lon"`
		} `json:"crew_coords" validate:"required"`

		// Скорость
		Speed float64 `json:"speed,omitempty" validate:"omitempty"`
		// Направление движения (0-Север, 90-Восток, 180-Юг, 270-Запад, -1-не задано)
		Direction int `json:"direction,omitempty" validate:"omitempty"`
	}
)

// Задание координат экипажей
func (cl *Client) SetCrewsCoords(req SetCrewsCoordsRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	err = cl.PostJson("set_crews_coords", nil, req, &response)

	return
}
