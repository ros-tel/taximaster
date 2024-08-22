package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	FindNearestAddressRequest struct {
		// Широта
		Lat float64 `validate:"required"`
		// Долгота
		Lon float64 `validate:"required"`

		// Радиус в метрах
		Radius int `validate:"omitempty"`
		// Искать адреса в ТМ (по умолчанию = true)
		SearchInTm *bool `validate:"omitempty"`
		// Искать адреса в Яндекс (по умолчанию = false)
		SearchInYandex *bool `validate:"omitempty"`
		// Искать адреса в Google (по умолчанию = false)
		SearchInGoogle *bool `validate:"omitempty"`
		// Искать адреса в TMGeoService (по умолчанию = false)
		SearchInTmGeoService *bool `validate:"omitempty"`
		// Искать адреса в Map.md (по умолчанию = false)
		SearchInMapMd *bool `validate:"omitempty"`
		// Искать адреса в 2ГИС (по умолчанию = false)
		SearchIn2Gis *bool `validate:"omitempty"`
	}

	FindNearestAddressResponse struct {
		// Источник адреса. Может принимать значения:
		// - "tm" - Такси Мастер (адрес из базы данных или из карты)
		// - "google" - Google
		// - "tmgeoservice" - TMGeoService
		// - "2gis" - 2GIS
		// - "mapmd" — Map.md
		AddressSource string `json:"address_source"`
		// Название города
		City string `json:"city"`
		// 	Название пункта
		Point string `json:"point"`
		// Название улицы или пункта
		Street string `json:"street"`
		// Номер дома
		House string `json:"house"`
		// Тип адреса. Может принимать значения:
		// - "street" - улица
		// - "house" - дом
		// - "point" - пункт
		Kind string `json:"kind"`
		// Комментарий
		Comment string `json:"comment"`
		// Координаты дома или пункта
		Coords Point `json:"coords"`
	}
)

// Поиск ближайшего адреса
func (cl *Client) FindNearestAddress(req FindNearestAddressRequest) (response FindNearestAddressResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("lat", strconv.FormatFloat(req.Lat, 'g', -1, 64))
	v.Add("lon", strconv.FormatFloat(req.Lon, 'g', -1, 64))
	if req.Radius != 0 {
		v.Add("radius", strconv.Itoa(req.Radius))
	}
	if req.SearchInTm != nil {
		v.Add("search_in_tm", strconv.FormatBool(*req.SearchInTm))
	}
	if req.SearchInYandex != nil {
		v.Add("search_in_yandex", strconv.FormatBool(*req.SearchInYandex))
	}
	if req.SearchInGoogle != nil {
		v.Add("search_in_google", strconv.FormatBool(*req.SearchInGoogle))
	}
	if req.SearchInTmGeoService != nil {
		v.Add("search_in_tmgeoservice", strconv.FormatBool(*req.SearchInTmGeoService))
	}
	if req.SearchInMapMd != nil {
		v.Add("search_in_mapmd", strconv.FormatBool(*req.SearchInMapMd))
	}
	if req.SearchIn2Gis != nil {
		v.Add("search_in_2gis", strconv.FormatBool(*req.SearchIn2Gis))
	}

	/*
		100 Подходящие адреса не найдены
		101 Не указано место для поиска адресов
	*/
	e := errorMap{
		100: ErrNoMatchingAddressesFound,
		101: ErrSearchLocationNotSpecified,
	}

	err = cl.Get("find_nearest_address", e, v, &response)

	return
}
