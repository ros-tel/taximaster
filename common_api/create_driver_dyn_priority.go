package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateDriverDynPriorityRequest struct {
		// ИД водителя (должно быть что-то одно: либо driver_id, либо crew_id)
		DriverID int `json:"driver_id,omitempty" validate:"omitempty"`
		// ИД экипажа (должно быть что-то одно: либо driver_id, либо crew_id)
		CrewID int `json:"crew_id,omitempty" validate:"omitempty"`
		// Приоритет
		Priority int `json:"priority" validate:"required"`
		// Время начала действия приоритета
		StartTime string `json:"start_time" validate:"required,datetime=20060102150405"`
		// Время окончания действия приоритета
		FinishTime string `json:"finish_time" validate:"required,datetime=20060102150405"`
		// Наименование приоритета
		Name string `json:"name" validate:"required"`
	}

	CreateDriverDynPriorityResponse struct {
		// ИД созданного динамического приоритета
		DynPriorityID int `json:"dyn_priority_id"`
	}
)

// Назначение динамического приоритета водителю
func (cl *Client) CreateDriverDynPriority(req CreateDriverDynPriorityRequest) (response CreateDriverDynPriorityResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100	Водитель не найден ИД=DRIVER_ID
		101	Экипаж не найден ИД=CREW_ID
		102	Время начала действия приоритета должно быть меньше времени окончания
		103	Время действия приоритета уже истекло
	*/
	e := errorMap{
		100: ErrDriverNotFound,
		101: ErrCrewNotFound,
		102: ErrTimeRange,
		103: ErrTimeExpired,
	}

	err = cl.PostJson("create_driver_dyn_priority", e, req, &response)

	return
}
