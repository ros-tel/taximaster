package common_api

type (
	GetOrderStatesListResponse struct {
		// Список параметров заказа
		OrderParams []struct {
			// ИД параметра
			ID int `json:"id"`
			// Название параметра
			Name string `json:"name"`
			// Тип состояния. Может принимать значения:
			// - "accepted" — заказ принят
			// - "in_work" — заказ в работе
			// - "finished" — заказ выполнен
			// - "aborted" — заказ прекращен
			StateType string `json:"state_type"`
		} `json:"order_params"`
	}
)

// Запрос списка состояний заказа
func (cl *Client) GetOrderStatesList() (response GetOrderStatesListResponse, err error) {
	err = cl.Get("get_order_states_list", nil, nil, &response)
	return
}
