package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetCarInfoRequest struct {
		// ИД автомобиля
		CarID int `validate:"required"`

		// Нужна ли фотография автомобиля
		NeedPhoto *bool `validate:"omitempty"`
		// Список возвращаемых полей через запятую
		Fields string `validate:"omitempty"`
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
		// Уровень топлива в автомобиле
		FuelLevel float64 `json:"fuel_level"`
		// Массив параметров автомобиля. Устарело. Рекомендуется использовать параметр attribute_values
		OrderParams []int `json:"order_params"`
		// Фото автомобиля (только если need_photo = true или поле driver_photo указано в списке фильтра полей fields)
		CarPhoto string `json:"car_photo"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values"`
	}
)

// Запрос информации об автомобиле
func (cl *Client) GetCarInfo(req GetCarInfoRequest) (response GetCarInfoResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("car_id", strconv.Itoa(req.CarID))
	if req.NeedPhoto != nil {
		v.Add("need_photo", strconv.FormatBool(*req.NeedPhoto))
	}
	if req.Fields != "" {
		v.Add("fields", req.Fields)
	}

	/*
		100 Автомобиль не найден
	*/
	e := errorMap{
		100: ErrCarNotFound,
	}

	err = cl.Get("get_car_info", e, v, &response)

	return
}
