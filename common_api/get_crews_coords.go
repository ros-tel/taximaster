package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetCrewsCoordsRequest struct {
		// ИД экипажа, по которому нужно вернуть координаты. Если не задано, то будут возвращены координаты всех экипажей на линии
		CrewID int `validate:"omitempty"`
	}

	GetCrewsCoordsResponse struct {
		// Список координат экипажей
		CrewsCoords []struct {
			// ИД экипажа
			CrewID int `json:"crew_id"`
			// Позывной экипажа
			CrewCode string `json:"crew_code"`
			// Время получения координат
			CoordsTime string `json:"coords_time"`
			// Долгота
			Lat float64 `json:"lat"`
			// Широта
			Lon float64 `json:"lon"`
			// Тип состояния экипажа. Может принимать значения:
			// - "not_available" - экипаж не на линии
			// - "waiting" - экипаж свободен, ожидает заказы
			// - "on_order" - экипаж на заказе
			// - "on_break" - экипаж на перерыве
			StateKind string `json:"state_kind"`
		} `json:"crews_coords"`
	}
)

// Запрос координат экипажей
func (cl *Client) GetCrewsCoords(req GetCrewsCoordsRequest) (GetCrewsCoordsResponse, error) {
	var response = GetCrewsCoordsResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	if req.CrewID > 0 {
		v.Add("crew_id", strconv.Itoa(req.CrewID))
	}

	/*
		100 Координаты не найдены
	*/
	e := errorMap{
		100: ErrCoordsNotFound,
	}

	err = cl.Get("get_crews_coords", e, v, &response)

	return response, err
}
