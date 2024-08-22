package common_api

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	ShowTmMessageRequest struct {
		// Текст сообщения
		Text string `validate:"required"`

		// Тип сообщения ("warning", "error", "information", "confirmation"), по умолчанию "information"
		Type string `validate:"omitempty,eq=warning|eq=error|eq=information|eq=confirmation"`
		// Заголовок сообщения
		Header string `validate:"omitempty"`
		// Скрывать сообщение через, сек. (0 — не скрывать)
		Timeout int `validate:"omitempty"`
		// Массив пользователей (если не указаны — отправлять всем)
		Users []int `validate:"omitempty"`
		// Цвет уведомления в формате RGB: #FFFFFF
		Color string `validate:"omitempty,iscolor=hexcolor"`
		// ИД заказа для кнопки открытия карточки в уведомлении
		OrderID int `validate:"omitempty"`
		// ИД автомобиля для кнопки открытия карточки в уведомлении
		CarID int `validate:"omitempty"`
		// ИД водителя для кнопки открытия карточки в уведомлении
		DriverID int `validate:"omitempty"`
		// ИД экипажа для кнопки открытия карточки в уведомлении
		CrewID int `validate:"omitempty"`
		// ИД клиента для кнопки открытия карточки в уведомлении
		ClientID int `validate:"omitempty"`
	}
)

// Показать сообщение в ТМ
func (cl *Client) ShowTmMessage(req ShowTmMessageRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("text", req.Text)
	if req.Type != "" {
		v.Add("type", req.Type)
	}
	if req.Header != "" {
		v.Add("header", req.Header)
	}
	if req.Timeout != 0 {
		v.Add("timeout", strconv.Itoa(req.Timeout))
	}
	if len(req.Users) != 0 {
		for _, user := range req.Users {
			v.Add("users", fmt.Sprint(user))
		}
	}
	if req.Color != "" {
		v.Add("color", req.Color)
	}
	if req.OrderID != 0 {
		v.Add("order_id", strconv.Itoa(req.OrderID))
	}
	if req.CarID != 0 {
		v.Add("car_id", strconv.Itoa(req.CarID))
	}
	if req.DriverID != 0 {
		v.Add("driver_id", strconv.Itoa(req.DriverID))
	}
	if req.CrewID != 0 {
		v.Add("crew_id", strconv.Itoa(req.CrewID))
	}
	if req.ClientID != 0 {
		v.Add("client_id", strconv.Itoa(req.ClientID))
	}

	/*
		100 Пользователи для отправки сообщения не найдены
	*/
	e := errorMap{
		100: ErrUsersNotFound,
	}

	err = cl.Post("show_tm_message", e, v, &response)

	return
}
