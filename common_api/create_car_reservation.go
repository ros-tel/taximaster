package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateCarReservationRequest struct {
		// ИД автомобиля
		CarID int `json:"car_id" validate:"required"`
		// ИД водителя
		DriverID int `json:"driver_id" validate:"required"`
		// ИД типа резервирования
		CarReservationTypeID int `json:"car_reservation_type_id" validate:"required"`
		// Дата или дата/время начала.
		// Если указывается дата без времени, то планируемое время начала резервирования устанавливается исходя из времени начала,
		// указанного в соответствующем типе резервирования
		StartTime string `json:"start_time" validate:"required"`

		// Время окончания резервирования. Если не указано, рассчитывается на основании времени планируемого начала резервирования
		// и продолжительности резервирования, указанного в соответствующем типе резервирования
		FinishTime string `json:"finish_time,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Комментарий
		Comment string `json:"comment,omitempty" validate:"omitempty"`
	}

	CreateCarReservationResponse struct {
		// ИД резервирования
		CarReservationID int `json:"car_reservation_id"`
	}
)

// Создание резервирования автомобиля
func (cl *Client) CreateCarReservation(req CreateCarReservationRequest) (response CreateCarReservationResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100	Не найден автомобиль
		101	Не найден водитель
		102	Не найден тип резервирования
		103	Время начала резервирования должно быть меньше времени завершения
		104	Автомобиль уже зарезервирован в указанный период времени
		105	Водитель уже имеет зарезервированный автомобиль в указанный период времени
	*/
	e := errorMap{
		100: ErrCarNotFound,
		101: ErrDriverNotFound,
		102: ErrReservationTypeNotFound,
		103: ErrTimeRange,
		104: ErrCarAlreadyReservedInThisTime,
		105: ErrDriverAlreadyHaveCarInThisTime,
	}

	err = cl.PostJson("create_car_reservation", e, req, &response)

	return
}
