package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	AnalyzeRoute2Request struct {
		// Массив адресов. Первый элемент — адрес подачи, последний — адрес назначения, между ними — остановки
		Addresses []Address `json:"addresses" validate:"required"`

		// Возвращать координаты точек полного маршрута (по умолчанию false)
		GetFullRouteCoords bool `json:"get_full_route_coords,omitempty" validate:"omitempty"`
		// ИД группы экипажей
		CrewGroupID int `json:"crew_group_id,omitempty" validate:"omitempty"`
	}

	AnalyzeRoute2Response struct {
		// Массив адресов. Первый элемент — адрес подачи, последний — адрес назначения, между ними — остановки
		Addresses []Address `json:"addresses"`
		// Километраж по городу
		CityDist float64 `json:"city_dist"`
		// Километраж за городом
		CountryDist float64 `json:"country_dist"`
		// Километраж до адреса подачи, если адрес подачи за городом
		SourceCountryDist float64 `json:"source_country_dist"`
		// Массив координат точек маршрута
		FullRouteCoords []Point `json:"full_route_coords"`
	}
)

// Анализ маршрута 2
func (cl *Client) AnalyzeRoute2(req AnalyzeRoute2Request) (AnalyzeRoute2Response, error) {
	var response = AnalyzeRoute2Response{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	err = cl.PostJson("analyze_route2", req, &response)

	return response, err
}
