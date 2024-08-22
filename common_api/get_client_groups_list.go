package common_api

type (
	GetClientGroupsListResponse struct {
		// Список групп клиентов
		ClientGroups []struct {
			// ИД группы клиентов
			ID int `json:"id"`
			// Название группы клиентов
			Name string `json:"name"`
		} `json:"client_groups"`
	}
)

// Запрос списка групп клиентов
func (cl *Client) GetClientGroupsList() (response GetClientGroupsListResponse, err error) {
	err = cl.Get("get_client_groups_list", nil, nil, &response)
	return
}
