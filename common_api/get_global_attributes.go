package common_api

type (
	GetGlobalAttributesResponse struct {
		// Массив значений глобальных атрибутов
		GlobalAttributes []AttributeValue `json:"global_attributes"`
	}
)

// Запрос глобальных атрибутов
func (cl *Client) GetGlobalAttributes() (response GetGlobalAttributesResponse, err error) {
	err = cl.Get("get_global_attributes", nil, nil, &response)
	return
}
