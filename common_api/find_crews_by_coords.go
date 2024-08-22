package common_api

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/ros-tel/taximaster/validator"
)

type (
	FindCrewsByCoordsRequest struct {
		// 	Широта точки центра поиска
		Lat float64 `validate:"required"`
		// Долгота точки центра поиска
		Lon float64 `validate:"required"`
		// Радиус поиска, км или мили
		Radius float64 `validate:"required"`

		// Искать экипажи без реальных координат по координатам стоянок
		CrewsWithoutCoords bool `validate:"omitempty"`
		// Допустимое время до освобождения в зоне поиска, если надо возвращать занятые экипажи (если 0, то искать только свободные экипажи), мин
		CrewsReleaseIn int `validate:"omitempty"`
		// ИД группы экипажей, заказы из которой должны видеть подходящие экипажи
		CrewGroupID int `validate:"omitempty"`
		// ИД службы ЕДС, к которой должны принадлежать подходящие экипажи
		UdsID int `validate:"omitempty"`
		// Список ИД атрибутов заказа, которыми должны обладать подходящие экипажи (пример: []int{1,2,3}).
		// Устарело. Рекомендуется использовать параметр attribute_values.
		Attributes []int `validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `validate:"omitempty"`
	}

	FindCrewsByCoordsResponse struct {
		// Количество свободных экипажей в зоне поиска
		WaitingCount int `json:"waiting_count"`
		// Количество занятых экипажей, освобождающихся в зоне поиска
		OnOrderCount int `json:"on_order_count"`
	}
)

// Поиск экипажей по координатам
func (cl *Client) FindCrewsByCoords(req FindCrewsByCoordsRequest) (response FindCrewsByCoordsResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("lat", strconv.FormatFloat(req.Lat, 'g', -1, 64))
	v.Add("lon", strconv.FormatFloat(req.Lon, 'g', -1, 64))
	v.Add("radius", strconv.FormatFloat(req.Radius, 'g', -1, 64))
	if req.CrewsWithoutCoords {
		v.Add("crews_without_coords", "true")
	}
	if req.CrewsReleaseIn != 0 {
		v.Add("crews_release_in", strconv.Itoa(req.CrewsReleaseIn))
	}
	if req.CrewGroupID != 0 {
		v.Add("crew_group_id", strconv.Itoa(req.CrewGroupID))
	}
	if req.UdsID != 0 {
		v.Add("uds_id", strconv.Itoa(req.UdsID))
	}
	if len(req.Attributes) != 0 {
		stringSlice := make([]string, len(req.Attributes))

		for i, num := range req.Attributes {
			stringSlice[i] = strconv.Itoa(num)
		}

		v.Add("attributes", strings.Join(stringSlice, ";"))
	}
	if len(req.AttributeValues) != 0 {
		for i := 0; i < len(req.AttributeValues); i++ {
			attrValue := req.AttributeValues[i]
			if attrValue.ID != 0 {
				v.Add(fmt.Sprintf("attribute_values[%d][id]", i), strconv.Itoa(attrValue.ID))
			}
			if attrValue.BoolValue != nil {
				v.Add(fmt.Sprintf("attribute_values[%d][bool_value]", i), strconv.FormatBool(*attrValue.BoolValue))
			}
			if attrValue.NumValue != nil {
				v.Add(fmt.Sprintf("attribute_values[%d][num_value]", i), strconv.FormatFloat(*attrValue.NumValue, 'g', -1, 64))
			}
			if attrValue.StrValue != nil {
				v.Add(fmt.Sprintf("attribute_values[%d][str_value]", i), *attrValue.StrValue)
			}
		}
	}

	err = cl.Get("find_crews_by_coords", nil, v, &response)

	return
}
