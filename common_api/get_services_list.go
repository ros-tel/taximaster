package common_api

type (
	GetServicesListResponse struct {
		// Список услуг
		Services []struct {
			// ИД услуги
			ID int `json:"id"`
			// Название услуги
			Name string `json:"name"`
			// Абсолютная сумма услуги, руб
			Sum float64 `json:"sum"`
			// Процент услуги от стоимости заказа, %
			Percent float64 `json:"percent"`
		} `json:"services"`
	}
)

// Запрос списка услуг
func (cl *Client) GetServicesList() (response GetServicesListResponse, err error) {
	err = cl.Get("get_services_list", nil, nil, &response)
	return
}
