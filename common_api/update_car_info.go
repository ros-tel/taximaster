package common_api

type (
	UpdateCarInfoRequest struct {
		// ИД автомобиля
		CarID int `json:"car_id" validate:"required"`

		// Позывной
		Code *string `json:"code,omitempty" validate:"omitempty"`
		// Марка
		Mark *string `json:"mark,omitempty" validate:"omitempty"`
		// Цвет
		Color *string `json:"color,omitempty" validate:"omitempty"`
		// Модель
		Model *string `json:"model,omitempty" validate:"omitempty"`
		// Краткое название
		ShortName *string `json:"short_name,omitempty" validate:"omitempty"`
		// Государственный номер
		GosNumber *string `json:"gos_number,omitempty" validate:"omitempty"`
		// ИД службы ЕДС
		UdsID *int `json:"uds_id,omitempty" validate:"omitempty" validate:"omitempty"`
		// Фотография автомобиля
		CarPhoto *string `json:"car_photo,omitempty" validate:"omitempty,base64"`
		// Год выпуска
		ProductionYear *int `json:"production_year,omitempty" validate:"omitempty"`
		// Класс автомобиля (A, B, C, ...)
		CarClass *int `json:"car_class,omitempty" validate:"omitempty"`
		// VIN
		Vin *int `json:"vin,omitempty" validate:"omitempty"`
		// Номер кузова
		BodyNumber *int `json:"body_number,omitempty" validate:"omitempty"`
		// Номер двигателя
		EngineNumber *int `json:"engine_number,omitempty" validate:"omitempty"`
		// Разрешение на перевозку
		Permit *int `json:"permit,omitempty" validate:"omitempty"`
		// Описание
		Comment *string `json:"comment,omitempty" validate:"omitempty"`
		// Автомобиль заблокирован
		IsLocked *bool `json:"is_locked,omitempty" validate:"omitempty"`
		// Причина блокировки
		LockDescription *string `json:"lock_description,omitempty" validate:"omitempty"`
		// Массив параметров автомобиля
		OrderParams *[]int `json:"order_params,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues *[]AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
	}
)

// Обновить информацию об автомобиле
func (cl *Client) UpdateCarInfo(req UpdateCarInfoRequest) (EmptyResponse, error) {
	var response = EmptyResponse{}

	err := validate(req)
	if err != nil {
		return response, err
	}

	err = cl.PostJson("update_car_info", req, &response)

	return response, err
}
