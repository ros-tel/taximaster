package common_api

type (
	GetCrewStatesListResponse struct {
		// Список параметров заказа
		OrderParams []struct {
			// ИД параметра
			ID int `json:"id"`
			// Название параметра
			Name string `json:"name"`
			// Тип состояния. Может принимать значения:
			// - "waiting" — экипаж свободен
			// - "not_available" — экипаж не на линии
			// - "on_order" — экипаж на заказе
			// - "on_break" — экипаж на перерыве
			StateType string `json:"state_type"`
		} `json:"order_params"`
	}
)

// Запрос списка состояний экипажа
func (cl *Client) GetCrewStatesList() (response GetCrewStatesListResponse, err error) {
	err = cl.Get("get_crew_states_list", nil, nil, &response)
	return
}
