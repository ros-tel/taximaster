package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	SaveClientFeedBackRequest struct {
		// Телефон
		Phone string `json:"phone" validate:"required"`
		// Рейтинг (от 1 до 5)
		Rating int `json:"rating" validate:"required"`
		// Текстовый отзыв
		Text string `json:"text" validate:"required"`

		// ИД заказа
		OrderID int `json:"order_id,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
	}
)

// Сохранение отзыва клиента
func (cl *Client) SaveClientFeedBack(req SaveClientFeedBackRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	err = cl.PostJson("save_client_feed_back", nil, req, &response)

	return
}
