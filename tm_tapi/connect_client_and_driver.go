package tm_tapi

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	ConnectClientAndDriverRequest struct {
		// ИД заказа
		OrderID int `validate:"required"`
	}
)

// Соединить клиента и водителя
func (cl *Client) ConnectClientAndDriver(req ConnectClientAndDriverRequest) (EmptyResponse, error) {
	var response = EmptyResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("order_id", strconv.Itoa(req.OrderID))

	err = cl.Post("connect_client_and_driver", v, &response)

	return response, err
}
