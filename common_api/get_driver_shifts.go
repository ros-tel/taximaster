package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetDriverShiftsRequest struct {
		// Начало периода
		StartTime string `validate:"required,datetime=20060102150405"`
		//Конец периода
		FinishTime string `validate:"required,datetime=20060102150405"`

		// ИД водителя
		DriverID int `validate:"omitempty"`
		// Включить в ответ новые смены (по умолчанию true)
		NewShifts *bool `validate:"omitempty"`
		// Включить в ответ смены в работе (по умолчанию true)
		InWorkShifts *bool `validate:"omitempty"`
		// Включить в ответ выполненные смены (по умолчанию true)
		FinishedShifts *bool `validate:"omitempty"`
		// Включить в ответ неуспешно завершенные смены (по умолчанию true)
		FailedShifts *bool `validate:"omitempty"`
		// Включить в ответ возвращенные смены (по умолчанию false)
		ReturnedShifts *bool `validate:"omitempty"`
	}

	GetDriverShiftsResponse struct {
		// Массив купленных смен
		Shifts []DriverShift `json:"shifts"`
	}

	DriverShift struct {
		// ИД купленной смены
		ShiftID int `json:"shift_id"`
		// ИД водителя
		DriverID int `json:"driver_id"`
		// ФИО водителя
		DriverName string `json:"driver_name"`
		// ИД запланированной смены
		PlanShiftID string `json:"plan_shift_id"`
		// Наименование запланированной смены
		PlanShiftName string `json:"plan_shift_name"`
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
		// Состояние смены («new» - новая, «in_work» - в работе, «finished» - завершена успешно, «failed» - завершена неуспешно)
		ShiftState string `json:"shift_state"`
		// Время продажи смены водителю
		BuyTime string `json:"buy_time"`
		// Признак возвращенной смены
		IsReturned bool `json:"is_returned"`
		// Время возврата смены
		ReturnTime string `json:"return_time"`
		// Количество заказов, выполненных водителем за смену
		OrdersCount int `json:"orders_count"`
		// Сумма заказов, выполненных водителем за смену
		OrdersSum float64 `json:"orders_sum"`
		// Фактическая продолжительность смены, ч.
		FactLength float64 `json:"fact_length"`
	}
)

// Запрос списка купленных смен водителей
func (cl *Client) GetDriverShifts(req GetDriverShiftsRequest) (response GetDriverShiftsResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("start_time", req.StartTime)
	v.Add("finish_time", req.FinishTime)
	if req.DriverID != 0 {
		v.Add("driver_id", strconv.Itoa(req.DriverID))
	}
	if req.NewShifts != nil {
		v.Add("new_shifts", strconv.FormatBool(*req.NewShifts))
	}
	if req.InWorkShifts != nil {
		v.Add("in_work_shifts", strconv.FormatBool(*req.InWorkShifts))
	}
	if req.FinishedShifts != nil {
		v.Add("finished_shifts", strconv.FormatBool(*req.FinishedShifts))
	}
	if req.FailedShifts != nil {
		v.Add("failed_shifts", strconv.FormatBool(*req.FailedShifts))
	}
	if req.ReturnedShifts != nil {
		v.Add("returned_shifts", strconv.FormatBool(*req.ReturnedShifts))
	}

	err = cl.Get("get_driver_shifts", nil, v, &response)

	return
}
