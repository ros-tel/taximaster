package common_api

type (
	UpdateCrewInfoRequest struct {
		// ИД экипажа
		CrewID int `json:"crew_id" validate:"required"`

		// ИД автомобиля
		CarID *int `json:"car_id,omitempty" validate:"omitempty"`
		// ИД водителя
		DriverID *int `json:"driver_id,omitempty" validate:"omitempty"`
		// ИД группы экипажа
		CrewGroupID *int `json:"crew_group_id,omitempty" validate:"omitempty"`
		// Позывной экипажа
		Code *string `json:"code,omitempty" validate:"omitempty"`
		Name *string `json:"name,omitempty" validate:"omitempty"`
		// Сумма, списываемая за смену
		WorkShiftSum *float64 `json:"work_shift_sum,omitempty" validate:"omitempty"`
		// Минимальный баланс, при котором можно выйти на смену
		MinBalance *int `json:"min_balance,omitempty" validate:"omitempty"`
		// Время работы, формат: “6.00-10.30, 23:00-00:48”
		WorkTime *string `json:"work_time,omitempty" validate:"omitempty"`
		// Шашка
		HasLightHouse *bool `json:"has_light_house,omitempty" validate:"omitempty"`
		// Наклейка
		HasLabel *bool `json:"has_label,omitempty" validate:"omitempty"`
		// Запрет работы вне запланированных смен
		UsePlanShifts *bool `json:"use_plan_shifts,omitempty" validate:"omitempty"`
		// Массив параметров экипажа
		OrderParams *[]int `json:"order_params,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues *[]AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
	}
)

// Обновление информации об экипаже
func (cl *Client) UpdateCrewInfo(req UpdateCrewInfoRequest) (EmptyResponse, error) {
	var response = EmptyResponse{}

	err := validate(req)
	if err != nil {
		return response, err
	}

	err = cl.PostJson("update_crew_info", req, &response)

	return response, err
}
