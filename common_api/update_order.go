package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	UpdateOrderRequest struct {
		// ИД заказа
		OrderID int `json:"order_id" validate:"required"`

		// Номер телефона
		Phone *string `json:"phone,omitempty" validate:"omitempty,max=30"`
		// Время подачи
		SourceTime *string `json:"source_time" validate:"omitempty,datetime=20060102150405"`
		// Предварительный заказ
		IsPrior *bool `json:"is_prior,omitempty" validate:"omitempty"`
		// Заказчик
		Customer *string `json:"customer,omitempty" validate:"omitempty"`
		// Пассажир
		Passenger *string `json:"passenger,omitempty" validate:"omitempty"`
		// Комментарий
		Comment *string `json:"comment,omitempty" validate:"omitempty"`
		// ИД группы экипажей
		CrewGroupID *int `json:"crew_group_id,omitempty" validate:"omitempty"`
		// ИД клиента (необязателен, если phone присутствует)
		ClientID *int `json:"client_id,omitempty" validate:"omitempty"`
		// ИД службы ЕДС
		UdsID *int `json:"uds_id,omitempty" validate:"omitempty"`
		// ИД тарифа
		TariffID *int `json:"tariff_id,omitempty" validate:"omitempty"`
		// Массив адресов. Первый элемент — адрес подачи(обязательно), последний — адрес назначения, между ними — остановки
		Addresses *[]Address `json:"addresses,omitempty" validate:"omitempty"`
		// Массив параметров заказа. Устарело. Рекомендуется использовать параметр attribute_values
		OrderParams *[]int `json:"order_params,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues *[]AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
		// Сумма заказа
		CostOrder *float64 `json:"cost_order,omitempty" validate:"omitempty"`
		// ИД состояния заказа
		StateID *int `json:"state_id" validate:"omitempty"`
		// ИД скидки
		DiscountID *int `json:"discount_id,omitempty" validate:"omitempty"`
		// Автоматически подобрать скидку, если не указана явно
		AutoSelectDiscount *bool `json:"auto_select_discount,omitempty" validate:"omitempty"`
		// Автоматически подобрать тариф, если не указан явно
		AutoSelectTariff *bool `json:"auto_select_tariff,omitempty" validate:"omitempty"`
		// Автоматически пересчитать сумму заказа
		AutoRecalcCost *bool `json:"auto_recalc_cost,omitempty" validate:"omitempty"`
		// Автоматически обновить параметры заказа по клиенту и группе клиента
		AutoUpdateOrderParams *bool `json:"auto_update_order_params,omitempty" validate:"omitempty"`
		// Email для уведомлений
		Email *string `json:"email,omitempty" validate:"omitempty,email"`
		// Время перехода из предварительного в текущие заказы, мин
		PriorToCurrentBeforeMinutes *int `json:"prior_to_current_before_minutes,omitempty" validate:"omitempty"`
		// Номер рейса
		FlightNumber *string `json:"flight_number,omitempty" validate:"omitempty"`
		// Использовать специальную проверку перед изменением заказа
		NeedCustomValidate *bool `json:"need_custom_validate,omitempty" validate:"omitempty"`
	}
)

// Изменение информации по заказу
func (cl *Client) UpdateOrder(req UpdateOrderRequest) (EmptyResponse, error) {
	var response = EmptyResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	err = cl.PostJson("update_order", req, &response)

	return response, err
}
