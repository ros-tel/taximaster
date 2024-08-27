package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetDriverPlanShiftsRequest struct {
		// Начало периода
		StartTime string `validate:"required,datetime=20060102150405"`
		//Конец периода
		FinishTime string `validate:"required,datetime=20060102150405"`

		// Включить в ответ срочные смены (по умолчанию true)
		LimitedShifts *bool `validate:"omitempty"`
		// Включить в ответ бессрочные смены (по умолчанию true)
		UnlimitedShifts *bool `validate:"omitempty"`
	}

	GetDriverPlanShiftsResponse struct {
		// Массив купленных смен
		PlanShifts []DriverPlanShift `json:"plan_shifts"`
	}

	DriverPlanShift struct {
		// ИД запланированной смены
		PlanShiftID int `json:"plan_shift_id"`
		// Наименование запланированной смены
		PlanShiftName string `json:"plan_shift_name"`
		// Комментарий запланированной смены
		PlanShiftComment string `json:"plan_shift_comment"`
		// 	Цена смены
		PlanShiftCost float64 `json:"plan_shift_cost"`
		// Тип смены («limited» - срочная, «unlimited» - бессрочная)
		PlanShiftType string `json:"plan_shift_type"`
		// План-начало смены (для срочных)
		PlanShiftStartTime string `json:"plan_shift_start_time"`
		// План-конец смены (для срочных)
		PlanShiftFinishTime string `json:"plan_shift_finish_time"`
		// План продолжительность смены, ч. (для бессрочных)
		PlanShiftLength int `json:"plan_shift_length"`
		// ИД первой группы экипажей, которая может купить смену
		PlanShiftCrewGroupID int `json:"plan_shift_crew_group_id"`
		// ИД групп экипажей, которые могут купить смену
		PlanShiftCrewGroups []int `json:"plan_shift_crew_groups"`
		// Максимальное количество продаж смены
		MaxSellCount int `json:"max_sell_count"`
		// Количество продаж смены
		SoldCount int `json:"sold_count"`
	}
)

// Запрос списка запланированных смен водителей
func (cl *Client) GetDriverPlanShifts(req GetDriverPlanShiftsRequest) (response GetDriverPlanShiftsResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("start_time", req.StartTime)
	v.Add("finish_time", req.FinishTime)
	if req.LimitedShifts != nil {
		v.Add("limited_shifts", strconv.FormatBool(*req.LimitedShifts))
	}
	if req.UnlimitedShifts != nil {
		v.Add("unlimited_shifts", strconv.FormatBool(*req.UnlimitedShifts))
	}

	err = cl.Get("get_driver_plan_shifts", nil, v, &response)

	return
}
