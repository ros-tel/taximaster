package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateCarRequest struct {
		// Позывной
		Code string `json:"code" validate:"required"`
		// Марка
		Mark string `json:"mark" validate:"required"`
		// Цвет
		Color string `json:"color" validate:"required"`
		// Государственный номер
		GosNumber string `json:"gos_number" validate:"required"`
		// ИД службы ЕДС (обязательное поле, если используется ЕДС, иначе можно не указывать)
		UdsID int `json:"uds_id,omitempty" validate:"omitempty"`

		// Модель
		Model string `json:"model,omitempty" validate:"omitempty"`
		// Краткое название
		ShortName string `json:"short_name,omitempty" validate:"omitempty"`
		// Год выпуска
		ProductionYear int `json:"production_year,omitempty" validate:"omitempty"`
		// Класс автомобиля (A, B, C, ...)
		CarClass string `json:"car_class,omitempty" validate:"omitempty"`
		// VIN
		Vin string `json:"vin,omitempty" validate:"omitempty"`
		// Номер кузова
		BodyNumber string `json:"body_number,omitempty" validate:"omitempty"`
		// Номер двигателя
		EngineNumber string `json:"engine_number,omitempty" validate:"omitempty"`
		// Разрешение на перевозку
		Permit string `json:"permit,omitempty" validate:"omitempty"`
		// Описание
		Comment string `json:"comment,omitempty" validate:"omitempty"`
		// Массив параметров автомобиля. Устарело. Рекомендуется использовать параметр attribute_values
		OrderParams []int `json:"order_params,omitempty" validate:"omitempty"`
		// Фотография автомобиля
		CarPhoto string `json:"car_photo,omitempty" validate:"omitempty,base64"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
	}

	CreateCarResponse struct {
		CarID int `json:"car_id"`
	}
)

// Создание автомобиля
func (cl *Client) CreateCar(req CreateCarRequest) (response CreateCarResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100 Автомобиль с ИД=ID имеет такой же позывной=CODE
		101 Служба ЕДС не найдена
		0	Параметр с ИД=ID не найден или не может быть привязан к автомобилю
	*/
	e := errorMap{
		100: ErrCarConflictByCode,
		101: ErrUdsNotFound,
		0:   ErrParameterNotFoundOrCannotBeBoundCar,
	}

	err = cl.PostJson("create_car", e, req, &response)

	return
}
