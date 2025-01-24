package common_api

import (
	"github.com/ros-tel/taximaster/validator"
)

type (
	ShowTmMessageRequest struct {
		// Текст сообщения
		Text string `json:"text" validate:"required"`

		// Тип сообщения ("warning", "error", "information", "confirmation"), по умолчанию "information"
		Type string `json:"type,omitempty" validate:"omitempty,eq=warning|eq=error|eq=information|eq=confirmation"`
		// Заголовок сообщения
		Header string `json:"header,omitempty" validate:"omitempty"`
		// Скрывать сообщение через, сек. (0 — не скрывать)
		Timeout int `json:"timeout,omitempty" validate:"omitempty"`
		// Массив пользователей (если не указаны — отправлять всем)
		Users []int `json:"users,omitempty" validate:"omitempty"`
		// Цвет уведомления в формате RGB: #FFFFFF
		Color string `json:"color,omitempty" validate:"omitempty,hexcolor"`
		// ИД заказа для кнопки открытия карточки в уведомлении
		OrderID int `json:"order_id,omitempty" validate:"omitempty"`
		// ИД автомобиля для кнопки открытия карточки в уведомлении
		CarID int `json:"car_id,omitempty" validate:"omitempty"`
		// ИД водителя для кнопки открытия карточки в уведомлении
		DriverID int `json:"driver_id,omitempty" validate:"omitempty"`
		// ИД экипажа для кнопки открытия карточки в уведомлении
		CrewID int `json:"crew_id,omitempty" validate:"omitempty"`
		// ИД клиента для кнопки открытия карточки в уведомлении
		ClientID int `json:"client_id,omitempty" validate:"omitempty"`
	}
)

// Показать сообщение в ТМ
func (cl *Client) ShowTmMessage(req ShowTmMessageRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100 Пользователи для отправки сообщения не найдены
	*/
	e := errorMap{
		100: ErrUsersNotFound,
	}

	err = cl.PostJson("show_tm_message", e, req, &response)

	return
}
