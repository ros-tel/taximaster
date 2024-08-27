package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	ChangeOrderStateRequest struct {
		// ИД заказа
		OrderID int `validate:"required"`
		// Новое состояние заказа
		NewState int `validate:"required"`

		// Сумма штрафа клиента за отмену заказа.
		// Если свойство new_state имеет тип "Прекращен" (выполняется отмена заказа) и при отмене заказа должен быть назначен штраф клиенту,
		// то если данное значение штрафа указано (даже если указано значение 0), то оно имеет приоритет.
		// Если данное значение не указано, то сумма штрафа будет определена автоматически по группе клиентов.
		CancelOrderPenaltySum *float64 `validate:"omitempty"`
	}

	ChangeOrderStateResponse struct {
		// ИД заказа
		OrderID int `json:"order_id"`
		// Новое состояние заказа
		NewState int `json:"new_state"`
	}
)

// Изменение состояния заказа
func (cl *Client) ChangeOrderState(req ChangeOrderStateRequest) (response ChangeOrderStateResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("order_id", strconv.Itoa(req.OrderID))
	v.Add("new_state", strconv.Itoa(req.NewState))
	if req.CancelOrderPenaltySum != nil {
		v.Add("cancel_order_penalty_sum", strconv.FormatFloat(*req.CancelOrderPenaltySum, 'g', -1, 64))
	}

	/*
		100	Не найден заказ ИД=ORDER_ID
		101	Не найдено состояние заказа ИД=NEW_STATE
		102	Изменение состояния не соответствует необходимым условиям.
	*/
	e := errorMap{
		100: ErrOrderNotFound,
		101: ErrOrderStateNotFound,
		102: ErrStateCannotBeChanged,
	}

	err = cl.Post("change_order_state", e, v, &response)

	return
}
