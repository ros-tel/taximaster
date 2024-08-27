package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateCarInaccessibilityRequest struct {
		// ИД автомобиля
		CarID int `json:"car_id" validate:"required"`
		// ИД типа резервирования
		CarInaccessibilityTypeID int `json:"car_inaccessibility_type_id" validate:"required"`
		// Время начала
		StartTime string `json:"start_time" validate:"required,datetime=20060102150405"`

		// Время завершения
		FinishTime string `json:"finish_time,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Комментарий
		Comment string `json:"comment,omitempty" validate:"omitempty"`
	}

	CreateCarInaccessibilityResponse struct {
		// ИД недоступности
		CarInaccessibilityID int `json:"car_inaccessibility_id"`
	}
)

// Создание недоступности автомобиля
func (cl *Client) CreateCarInaccessibility(req CreateCarInaccessibilityRequest) (response CreateCarInaccessibilityResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100	Не найден автомобиль
		101	Не найден тип недоступности
		102	Время начала недоступности должно быть меньше времени завершения
	*/
	e := errorMap{
		100: ErrCarNotFound,
		101: ErrInaccessibilityTypeNotFound,
		102: ErrTimeRange,
	}

	err = cl.PostJson("create_car_inaccessibility", e, req, &response)

	return
}
