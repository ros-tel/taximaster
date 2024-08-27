package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	UpdateCarInfoRequest struct {
		// ИД автомобиля
		CarID int `json:"car_id" validate:"required"`

		// Позывной
		Code string `json:"code,omitempty" validate:"omitempty"`
		// Марка
		Mark string `json:"mark,omitempty" validate:"omitempty"`
		// Цвет
		Color string `json:"color,omitempty" validate:"omitempty"`
		// Государственный номер
		GosNumber string `json:"gos_number,omitempty" validate:"omitempty"`
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
		// Массив параметров автомобиля
		OrderParams []int `json:"order_params,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
		// Автомобиль заблокирован
		IsLocked bool `json:"is_locked,omitempty" validate:"omitempty"`
		// Причина блокировки
		LockDescription string `json:"lock_description,omitempty" validate:"omitempty"`
		// ИД службы ЕДС
		UdsID int `json:"uds_id,omitempty" validate:"omitempty"`
		// Фотография автомобиля
		CarPhoto string `json:"car_photo,omitempty" validate:"omitempty,base64"`
	}
)

// Обновить информацию об автомобиле
func (cl *Client) UpdateCarInfo(req UpdateCarInfoRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		101 Служба ЕДС не найдена
		102 Автомобиль с ИД=ID не найден
		103 Экипаж на линии, запрещено редактирование полей: марка, модель, краткое наименование, цвет, гос. номер, служба ЕДС.
		104 Параметр с ИД=ID не найден или не может быть привязан к автомобилю
	*/
	e := errorMap{
		101: ErrUdsNotFound,
		102: ErrCarNotFound,
		103: ErrForbiddenEditCrewOnLine,
		104: ErrParameterNotFoundOrCannotBeBoundCar,
	}

	err = cl.PostJson("update_car_info", e, req, &response)

	return
}
