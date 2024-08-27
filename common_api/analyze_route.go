package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	AnalyzeRouteRequest struct {
		// Адрес подачи Параметр source не является обязательным к заполнению, если заданы source_lon и source_lat.
		Source string `validate:"omitempty"`
		// Адрес назначения Параметр dest не является обязательным к заполнению, если заданы dest_lon и dest_lat.
		Dest string `validate:"omitempty"`

		// Долгота адреса подачи
		SourceLon *float64 `validate:"omitempty"`
		// Широта адреса подачи
		SourceLat *float64 `validate:"omitempty"`
		// Долгота адреса назначения
		DestLon *float64 `validate:"omitempty"`
		// Широта адреса назначения
		DestLat *float64 `validate:"omitempty"`
	}

	AnalyzeRouteResponse struct {
		// Широта адреса подачи
		SourceLat float64 `json:"source_lat"`
		// Долгота адреса подачи
		SourceLon float64 `json:"source_lon"`
		// ИД района подачи
		SourceZoneID int `json:"source_zone_id"`
		// Широта адреса назначения
		DestLat float64 `json:"dest_lat"`
		// Долгота адреса назначения
		DestLon float64 `json:"dest_lon"`
		// ИД района назначения
		DestZoneID int `json:"dest_zone_id"`
		// Километраж по городу
		CityDist float64 `json:"city_dist"`
		// Километраж за городом
		CountryDist float64 `json:"country_dist"`
		// Километраж до адреса подачи, если адрес подачи за городом
		SourceCountryDist float64 `json:"source_country_dist"`
	}
)

// Анализ маршрута
func (cl *Client) AnalyzeRoute(req AnalyzeRouteRequest) (response AnalyzeRouteResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	if req.Source != "" {
		v.Add("source", req.Source)
	}
	if req.Dest != "" {
		v.Add("dest", req.Dest)
	}
	if req.SourceLon != nil {
		v.Add("source_lon", strconv.FormatFloat(*req.SourceLon, 'g', -1, 64))
	}
	if req.SourceLat != nil {
		v.Add("source_lat", strconv.FormatFloat(*req.SourceLat, 'g', -1, 64))
	}
	if req.DestLon != nil {
		v.Add("dest_lon", strconv.FormatFloat(*req.DestLon, 'g', -1, 64))
	}
	if req.DestLat != nil {
		v.Add("dest_lat", strconv.FormatFloat(*req.DestLat, 'g', -1, 64))
	}

	/*
		100	Адрес подачи не распознан
		101	Адрес назначения не распознан
		102	Маршрут не распознан
	*/
	e := errorMap{
		100: ErrSourceNotFound,
		101: ErrDestNotFound,
		102: ErrRouteNotRecognized,
	}

	err = cl.Get("analyze_route", e, v, &response)

	return
}
