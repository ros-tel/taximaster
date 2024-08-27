package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetAddressesLike2Request struct {
		// Искать улицы
		GetStreets bool `validate:"required"`
		// Искать пункты
		GetPoints bool `validate:"required"`
		// Искать дома
		GetHouses bool `validate:"required"`
		// Строка для поиска адреса
		Address string `validate:"required"`

		// Города, разделенные запятой, в которых искать адреса
		City string `validate:"omitempty"`
		// Максимальное количество адресов в ответе
		MaxAddressesCount int `validate:"omitempty"`
		// Искать адреса в ТМ (по умолчанию = true)
		SearchInTm *bool `validate:"omitempty"`
		// Искать адреса в Яндекс (по умолчанию = false)
		SearchInYandex *bool `validate:"omitempty"`
		// Искать адреса в Google (по умолчанию = false)
		SearchInGoogle *bool `validate:"omitempty"`
		// Искать адреса в 2GIS (по умолчанию = false)
		SearchIn2Gis *bool `validate:"omitempty"`
		// Искать адреса в TMGeoService (по умолчанию = false)
		SearchInTmGeoService *bool `validate:"omitempty"`
		// Искать адреса в Map.md (по умолчанию = false)
		SearchInMapMd *bool `validate:"omitempty"`
	}

	GetAddressesLike2Response struct {
		// Список подходящих адресов
		Addresses []struct {
			// Источник адреса. Может принимать значения:
			// - "tm" - Такси Мастер (адрес из базы данных или из карты)
			// - "yandex" - Яндекс
			// - "google" - Google
			// - "2gis" - 2GIS
			// - "tmgeoservice" - TMGeoService
			AddressSource string `json:"address_source"`
			// Название города
			City string `json:"city"`
			// Название пункта
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
		} `json:"addresses"`
	}
)

// Запрос адресов, содержащих нужную строку 2
func (cl *Client) GetAddressesLike2(req GetAddressesLike2Request) (response GetAddressesLike2Response, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("get_streets", strconv.FormatBool(req.GetStreets))
	v.Add("get_points", strconv.FormatBool(req.GetPoints))
	v.Add("get_houses", strconv.FormatBool(req.GetHouses))
	v.Add("address", req.Address)
	if req.City != "" {
		v.Add("city", req.City)
	}
	if req.MaxAddressesCount > 0 {
		v.Add("max_addresses_count", strconv.Itoa(req.MaxAddressesCount))
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
	if req.SearchIn2Gis != nil {
		v.Add("search_in_2gis", strconv.FormatBool(*req.SearchIn2Gis))
	}
	if req.SearchInTmGeoService != nil {
		v.Add("search_in_tmgeoservice", strconv.FormatBool(*req.SearchInTmGeoService))
	}
	if req.SearchInMapMd != nil {
		v.Add("search_in_mapmd", strconv.FormatBool(*req.SearchInMapMd))
	}

	/*
		100 Подходящие адреса не найдены
		101 Не указано место для поиска адресов
	*/
	e := errorMap{
		100: ErrNoMatchingAddressesFound,
		101: ErrSearchLocationNotSpecified,
	}

	err = cl.Get("get_addresses_like2", e, v, &response)

	return
}
