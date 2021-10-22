package common_api

import (
	"net/url"
	"strconv"
)

type (
	GetCrewInfoRequest struct {
		// ИД экипажа
		CrewID int `validate:"required"`
	}

	GetCrewInfoResponse struct {
		// ИД экипажа
		CrewID int `json:"crew_id"`
		// 	Позывной экипажа
		Code string `json:"code"`
		// Наименование экипажа
		Name string `json:"name"`
		// ИД водителя
		DriverID int `json:"driver_id"`
		// ИД автомобиля
		CarID int `json:"car_id"`
		// ИД группы экипажа
		CrewGroupID int `json:"crew_group_id"`
		// ИД состояния экипажа
		CrewStateID int `json:"crew_state_id"`
		// Водитель подключен к серверу "Связи с водителями"
		Online bool `json:"online"`
		// Сумма, списываемая за смену
		WorkShiftSum float64 `json:"work_shift_sum"`
		// Минимальный баланс, при котором можно выйти на смену
		MinBalance int `json:"min_balance"`
		// Общий приоритет
		CommonPriority int `json:"common_priority"`
		// Статический приоритет
		StaticPriority int `json:"static_priority"`
		// Динамический приоритет
		DynamicPriority int `json:"dynamic_priority"`
		// Приоритет по рейтингу
		RatingPriority int `json:"rating_priority"`
		// Индивидуальная сдача с заказа
		OrderChangeID int `json:"order_change_id"`
		// Шашка
		HasLightHouse bool `json:"has_light_house"`
		// Наклейка
		HasLabel bool `json:"has_label"`
		// Запрет работы вне запланированных смен
		UsePlanShifts bool `json:"use_plan_shifts"`
		// Массив параметров экипажа
		OrderParams []int `json:"order_params"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values"`
	}
)

// Запрос информации об экипаже
func (cl *Client) GetCrewInfo(req GetCrewInfoRequest) (GetCrewInfoResponse, error) {
	var response = GetCrewInfoResponse{}

	err := validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("crew_id", strconv.Itoa(req.CrewID))

	err = cl.Get("get_crew_info", v, &response)

	return response, err
}
