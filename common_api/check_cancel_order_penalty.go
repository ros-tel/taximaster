package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	CheckCancelOrderPenaltyRequest struct {
		// ИД заказа
		OrderID int `validate:"required"`
		// ИД состояния заказа, в которое переходит заказ при отмене
		CancelOrderStateID int `validate:"required"`
	}

	CheckCancelOrderPenaltyResponse struct {
		// Величина штрафа, которая будет начислена клиенту в случае отмены заказа или 0, если штрафа нет
		CancelOrderPenaltySum float64 `json:"cancel_order_penalty_sum"`
	}
)

// Проверить штраф клиента за отмену заказа
func (cl *Client) CheckCancelOrderPenalty(req CheckCancelOrderPenaltyRequest) (response CheckCancelOrderPenaltyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("order_id", strconv.Itoa(req.OrderID))
	v.Add("cancel_order_state_id", strconv.Itoa(req.CancelOrderStateID))

	/*
		100	Не найден заказ ИД=order_id
		101	Не найдено состояние заказа ИД=cancel_order_state_id
		102	Состояние заказа не соответствует необходимым условиям
	*/
	e := errorMap{
		100: ErrCarNotFound,
		101: ErrOrderStateNotFound,
		102: ErrOrderStateNotMeetConditions,
	}

	err = cl.Get("check_cancel_order_penalty", e, v, &response)

	return
}
