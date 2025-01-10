package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateDriverInaccessibilityRequest struct {
		// ИД водителя
		DriverID int `json:"driver_id" validate:"required"`
		// ИД типа недоступности
		DriverInaccessibilityTypeID int `json:"driver_inaccessibility_type_id" validate:"required"`
		// Время начала
		StartTime string `json:"start_time" validate:"required,datetime=20060102150405"`

		// Время завершения
		FinishTime string `json:"finish_time,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Комментарий
		Comment string `json:"comment,omitempty" validate:"omitempty"`
	}

	CreateDriverInaccessibilityResponse struct {
		// ИД недоступности
		DriverInaccessibilityID int `json:"driver_inaccessibility_id"`
	}
)

// Создание недоступности водителя
func (cl *Client) CreateDriverInaccessibility(req CreateDriverInaccessibilityRequest) (response CreateDriverInaccessibilityResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100	Не найден водитель
		101	Не найден тип недоступности
		102	Время начала недоступности должно быть меньше времени завершения
	*/
	e := errorMap{
		100: ErrDriverNotFound,
		101: ErrInaccessibilityTypeNotFound,
		102: ErrTimeRange,
	}

	err = cl.PostJson("create_driver_inaccessibility", e, req, &response)

	return
}
