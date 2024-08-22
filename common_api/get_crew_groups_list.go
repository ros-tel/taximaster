package common_api

type (
	GetCrewGroupsListResponse struct {
		// Список групп экипажей
		CrewGroups []struct {
			// ИД группы экипажей
			ID int `json:"id"`
			// Название группы экипажей
			Name string `json:"name"`
		} `json:"crew_groups"`
	}
)

// Запрос списка групп экипажей
func (cl *Client) GetCrewGroupsList() (response GetCrewGroupsListResponse, err error) {
	err = cl.Get("get_crew_groups_list", nil, nil, &response)
	return
}
