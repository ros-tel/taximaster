package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateCrewRequest struct {
		// ИД автомобиля
		CarID int `json:"car_id" validate:"required"`
		// ИД водителя
		DriverID int `json:"driver_id" validate:"required"`
		// ИД группы экипажа
		CrewGroupID int `json:"crew_group_id" validate:"required"`

		// Позывной экипажа
		Code string `json:"code,omitempty" validate:"omitempty"`
		// Сумма, списываемая за смену
		WorkShiftSum float64 `json:"work_shift_sum,omitempty" validate:"omitempty"`
		// Минимальный баланс, при котором можно выйти на смену
		MinBalance float64 `json:"min_balance,omitempty" validate:"omitempty"`
		// Время работы, формат: “6.00-10.30, 23:00-00:48”
		WorkTime string `json:"work_time,omitempty" validate:"omitempty"`
		// Шашка
		HasLightHouse bool `json:"has_light_house,omitempty" validate:"omitempty"`
		// Наклейка
		HasLabel bool `json:"has_label,omitempty" validate:"omitempty"`
		// Запрет работы вне запланированных смен
		UsePlanShifts bool `json:"use_plan_shifts,omitempty" validate:"omitempty"`
		// Массив параметров экипажа. Устарело. Рекомендуется использовать параметр attribute_values
		OrderParams []int `json:"order_params,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
	}

	CreateCrewResponse struct {
		// ИД созданного экипажа
		CrewID int `json:"crew_id"`
	}
)

// Создание экипажа
func (cl *Client) CreateCrew(req CreateCrewRequest) (response CreateCrewResponse, err error) {
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
	*/
	e := errorMap{
		100: ErrCarNotFound,
		101: ErrDriverNotFound,
		102: ErrCrewNotFound,
		103: ErrParameterNotFoundOrCannotBeBoundCrew,
		104: ErrCrewConflictByDriverAndCar,
		105: ErrUdsCarAndDriverDoesNotMatch,
	}

	err = cl.PostJson("create_crew", e, req, &response)

	return
}
