package common_api

type (
	GetUdsListResponse struct {
		// Список служб ЕДС
		Uds []struct {
			// ИД службы ЕДС
			ID int `json:"id"`
			// Название службы ЕДС
			Name string `json:"name"`
		} `json:"uds"`
	}
)

// Запрос списка служб ЕДС
func (cl *Client) GetUdsList() (response GetUdsListResponse, err error) {
	err = cl.Get("get_uds_list", nil, nil, &response)
	return
}
