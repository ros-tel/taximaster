package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetClientsInfoRequest struct {
		// Текст для поиска по названию или по номеру договора клиента
		Text string `validate:"omitempty"`
		// Максимальное количество клиентов, которое надо вернуть. Если не указано, то 10.
		MaxClientsCount int `validate:"omitempty"`
		// Фильтр по группе клиентов
		ClientGroupID int `validate:"omitempty"`
		// Фильтр по вышестоящему подразделению, возвращаются все подчиненные отделы и сотрудники на всю глубину иерархии
		ParentID int `validate:"omitempty"`
		// Список возвращаемых полей через запятую. По умолчанию возвращаются поля "name" и "number".
		// Поле "client_id" возвращается всегда
		Fields string `validate:"omitempty"`
	}

	GetClientsInfoResponse struct {
		ClientsInfo []GetClientInfoResponse `json:"clients_info"`
	}
)

// Запрос информации по клиентам
func (cl *Client) GetClientsInfo(req GetClientsInfoRequest) (response GetClientsInfoResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	if req.Text != "" {
		v.Add("text", req.Text)
	}
	if req.MaxClientsCount != 0 {
		v.Add("max_clients_count", strconv.Itoa(req.MaxClientsCount))
	}
	if req.ClientGroupID != 0 {
		v.Add("client_group_id", strconv.Itoa(req.ClientGroupID))
	}
	if req.ParentID != 0 {
		v.Add("parent_id", strconv.Itoa(req.ParentID))
	}
	if req.Fields != "" {
		v.Add("fields", req.Fields)
	}

	err = cl.Get("get_clients_info", nil, v, &response)

	return
}
