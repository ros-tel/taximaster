package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	DriverBuyShiftRequest struct {
		// ИД экипажа
		CrewID int `validate:"required"`
		// ИД запланированной смены
		PlanShiftID int `validate:"required"`
	}
)

// Продажа смены водителю
func (cl *Client) DriverBuyShift(req DriverBuyShiftRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("crew_id", strconv.Itoa(req.CrewID))
	v.Add("plan_shift_id", strconv.Itoa(req.PlanShiftID))

	/*
		100	Запланированная смена не найдена
		101	Экипаж не найден
		102	Водитель не найден
		103	Недостаточно денег на счете водителя
		104	Водитель уволен либо заблокирован
		105	Запланированная смена устарела
		106	Не подходит группа экипажа
		107	Превышено максимальное число покупок смены
		108	Повторная покупка смены
		109	Экипажу не назначен атрибут для доступа к смене
	*/
	e := errorMap{
		100: ErrPlanShiftNotFound,
		101: ErrCrewNotFound,
		102: ErrDriverNotFound,
		103: ErrInsufficientFundsDriver,
		104: ErrDriverFiredOrBlocked,
		105: ErrPlanShiftOutdated,
		106: ErrCrewGroupsNotSuitable,
		107: ErrExceededMaxPurchases,
		108: ErrDuplicatePurchases,
		109: ErrCrewNotAssignedAttributeForShiftAccess,
	}

	err = cl.Post("driver_buy_shift", e, v, &response)

	return
}
