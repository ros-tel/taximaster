package tm_tapi

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	SetRequestStateRequest struct {
		// Cостояние заказа до отзвона
		StateID int `validate:"required"`
		// Признак, что звонили клиенту (если 0 - водителю)
		PhoneType int `validate:"required"`
		// ИД заказа
		OrderID int `validate:"required"`
		// Состояние отзвона:
		// 0 - начальное
		// 1 - в процессе
		// 2 - успешно
		// 3 - занято
		// 4 - нет ответа
		// 5 - ошибка
		State int `validate:"required"`
	}
)

// Смена состояния заказа по результату автодозвона
func (cl *Client) SetRequestState(req SetRequestStateRequest) (EmptyResponse, error) {
	var response = EmptyResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("STATE_ID", strconv.Itoa(req.StateID))
	v.Add("PHONE_TYPE", strconv.Itoa(req.PhoneType))
	v.Add("ORDER_ID", strconv.Itoa(req.OrderID))
	v.Add("STATE", strconv.Itoa(req.State))

	err = cl.Post("set_request_state", v, &response)

	return response, err
}
