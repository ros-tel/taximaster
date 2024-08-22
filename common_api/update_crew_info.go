package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	UpdateCrewInfoRequest struct {
		// ИД экипажа
		CrewID int `json:"crew_id" validate:"required"`

		// ИД автомобиля
		CarID int `json:"car_id,omitempty" validate:"omitempty"`
		// ИД водителя
		DriverID int `json:"driver_id,omitempty" validate:"omitempty"`
		// ИД группы экипажа
		CrewGroupID int `json:"crew_group_id,omitempty" validate:"omitempty"`
		// Позывной экипажа
		Code string `json:"code,omitempty" validate:"omitempty"`
		// Сумма, списываемая за смену
		WorkShiftSum float64 `json:"work_shift_sum,omitempty" validate:"omitempty"`
		// Минимальный баланс, при котором можно выйти на смену
		MinBalance int `json:"min_balance,omitempty" validate:"omitempty"`
		// Время работы, формат: “6.00-10.30, 23:00-00:48”
		WorkTime string `json:"work_time,omitempty" validate:"omitempty"`
		// Шашка
		HasLightHouse bool `json:"has_light_house,omitempty" validate:"omitempty"`
		// Наклейка
		HasLabel bool `json:"has_label,omitempty" validate:"omitempty"`
		// GPS идентификатор экипажа
		CrewGpsID int `json:"crew_gps_id,omitempty" validate:"omitempty"`
		// Запрет работы вне запланированных смен
		UsePlanShifts bool `json:"use_plan_shifts,omitempty" validate:"omitempty"`
		// Массив параметров экипажа
		OrderParams []int `json:"order_params,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
	}
)

// Обновление информации об экипаже
func (cl *Client) UpdateCrewInfo(req UpdateCrewInfoRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100 Автомобиль с ИД=ID не найден
		101 Водитель с ИД=ID не найден
		102 Группа экипажа с ИД=ID не найдена
		103 Параметр с ИД=ID не найден или не может быть привязан к экипажу
		104 Экипаж с таким водителем и автомобилем уже существует
		105 Служба ЕДС автомобиля и водителя не совпадает
		106 Экипаж с ИД=ID не найден
		107 Экипаж на линии, запрещено редактирование полей: водитель, автомобиль, позывной, группа экипажа, сумма за смену, минимальный баланс, запрет выхода вне запланированной смены.
	*/
	e := errorMap{
		100: ErrCarNotFound,
		101: ErrDriverNotFound,
		102: ErrCrewNotFound,
		103: ErrParameterNotFoundOrCannotBeBoundCrew,
		104: ErrCrewConflictByDriverAndCar,
		105: ErrUdsCarAndDriverDoesNotMatch,
		106: ErrCrewNotFound,
		107: ErrForbiddenEditCrewOnLine,
	}

	err = cl.PostJson("update_crew_info", e, req, &response)

	return
}
