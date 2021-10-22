package common_api

import (
	"net/url"
	"strconv"
)

type (
	GetCarInfoRequest struct {
		// ИД автомобиля
		CarID int `validate:"required"`

		// Нужна ли фотография автомобиля
		NeedPhoto bool `validate:"omitempty"`
	}

	GetCarInfoResponse struct {
		// ИД автомобиля
		CarID int `json:"car_id"`
		// Позывной автомобиля
		Code string `json:"code"`
		// Наименование автомобиля
		Name string `json:"name"`
		// Государственный номер автомобиля
		GosNumber string `json:"gos_number"`
		// Цвет автомобиля
		Color string `json:"color"`
		// Марка автомобиля
		Mark string `json:"mark"`
		// Модель автомобиля
		Model string `json:"model"`
		// Краткое название автомобиля
		ShortName string `json:"short_name"`
		// Год выпуска автомобиля
		ProductionYear int `json:"production_year"`
		// Автомобиль заблокирован
		IsLocked bool `json:"is_locked"`
		// Массив параметров автомобиля. Устарело. Рекомендуется использовать параметр attribute_values
		OrderParams []int `json:"order_params"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values"`
	}
)

// Запрос информации об автомобиле
func (cl *Client) GetCarInfo(req GetCarInfoRequest) (GetCarInfoResponse, error) {
	var response = GetCarInfoResponse{}

	err := validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("car_id", strconv.Itoa(req.CarID))
	if req.NeedPhoto {
		v.Add("need_photo", "true")
	}

	err = cl.Get("get_car_info", v, &response)

	return response, err
}
