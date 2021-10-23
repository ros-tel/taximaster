package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	ShowTmMessageRequest struct {
		// Текст сообщения
		Text string `json:"text" validate:"required"`

		// Заголовок сообщения
		Header string `json:"header,omitempty" validate:"omitempty"`
		// Скрывать сообщение через, сек. (0 — не скрывать)
		Timeout int `json:"timeout" validate:"omitempty"`
		// Массив пользователей (если не указаны — отправлять всем)
		Users *[]int `json:"users,omitempty" validate:"omitempty"`
	}
)

// Показать сообщение в ТМ
func (cl *Client) ShowTmMessage(req ShowTmMessageRequest) (EmptyResponse, error) {
	var response = EmptyResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	err = cl.PostJson("show_tm_message", req, &response)

	return response, err
}
