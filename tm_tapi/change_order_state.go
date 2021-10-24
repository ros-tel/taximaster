package tm_tapi

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
		NeedState int `validate:"required"`
	}

	ChangeOrderStateResponse struct {
		// ИД заказа
		OrderID int `xml:"ORDER_ID"`
		// Новое состояние заказа
		NewState int `xml:"NEW_STATE"`
	}
)

// Смена состояния заказа
func (cl *Client) ChangeOrderState(req ChangeOrderStateRequest) (ChangeOrderStateResponse, error) {
	var response = ChangeOrderStateResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("ORDER_ID", strconv.Itoa(req.OrderID))
	v.Add("NEED_STATE", strconv.Itoa(req.NeedState))

	err = cl.Post("change_order_state", v, &response)

	return response, err
}
