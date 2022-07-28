package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetAddressesLikeRequest struct {
		// Искать улицы
		GetStreets bool `validate:"omitempty"`
		// Искать пункты
		GetPoints bool `validate:"omitempty"`
		// Искать дома. Не может быть равно true, если get_streets = true или get_points = true.
		GetHouses bool `validate:"omitempty"`
		// Часть названия улицы или пункта, если идет поиск улиц или пунктов, или полное название улицы, если идет поиск домов
		Street string `validate:"required"`

		// Часть номера дома. Нужно только если get_houses = true
		House string `validate:"omitempty"`
		// Город, в котором искать адреса
		City string `validate:"omitempty"`
		// Максимальное количество адресов в ответе
		MaxAddressesCount int `validate:"omitempty"`
		// Искать адреса в ТМ (по умолчанию = true)
		SearchInTm *bool `validate:"omitempty"`
		// Искать адреса в Яндекс (по умолчанию = false)
		SearchInYandex bool `validate:"omitempty"`
		// Искать адреса в Google (по умолчанию = false)
		SearchInGoogle bool `validate:"omitempty"`
		// Искать адреса в 2GIS (по умолчанию = false)
		SearchIn2Gis bool `validate:"omitempty"`
		// Искать адреса в TMGeoService (по умолчанию = false)
		SearchInTmGeoService bool `validate:"omitempty"`
		// Искать адреса в Map.md (по умолчанию = false)
		SearchInMapMd bool `validate:"omitempty"`
	}

	GetAddressesLikeResponse struct {
		// Список подходящих адресов
		Addresses []struct {
			// Источник адреса. Может принимать значения:
			// - "tm" - Такси Мастер (адрес из базы данных или из карты)
			// - "yandex" - Яндекс
			// - "google" - Google
			// - "2gis" - 2GIS
			// - "tmgeoservice" - TMGeoService
			AddressSource string `json:"address_source"`
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

// Запрос адресов, содержащих нужную строку
func (cl *Client) GetAddressesLike(req GetAddressesLikeRequest) (GetAddressesLikeResponse, error) {
	var response = GetAddressesLikeResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	if req.GetStreets {
		v.Add("get_streets", "true")
	} else {
		v.Add("get_streets", "false")
	}
	if req.GetPoints {
		v.Add("get_points", "true")
	} else {
		v.Add("get_points", "false")
	}
	if req.GetHouses {
		v.Add("get_houses", "true")
	} else {
		v.Add("get_houses", "false")
	}
	v.Add("street", req.Street)
	if req.House != "" {
		v.Add("house", req.House)
	}
	if req.City != "" {
		v.Add("city", req.City)
	}
	if req.MaxAddressesCount > 0 {
		v.Add("max_addresses_count", strconv.Itoa(req.MaxAddressesCount))
	}
	if req.SearchInTm != nil {
		if *req.SearchInTm {
			v.Add("search_in_tm", "true")
		} else {
			v.Add("search_in_tm", "false")
		}
	}
	if req.SearchInYandex {
		v.Add("search_in_yandex", "true")
	}
	if req.SearchInGoogle {
		v.Add("search_in_google", "true")
	}
	if req.SearchIn2Gis {
		v.Add("search_in_2gis", "true")
	}
	if req.SearchInTmGeoService {
		v.Add("search_in_tmgeoservice", "true")
	}
	if req.SearchInMapMd {
		v.Add("search_in_mapmd", "true")
	}

	/*
		100 Подходящие адреса не найдены
		101 Не указано место для поиска адресов
	*/
	e := errorMap{
		100: ErrNoMatchingAddressesFound,
		101: ErrSearchLocationNotSpecified,
	}

	err = cl.Get("get_addresses_like", e, v, &response)

	return response, err
}
