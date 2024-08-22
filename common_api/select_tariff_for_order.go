package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	SelectTariffForOrderRequest struct {
		// ИД клиента
		ClientID int `json:"client_id,omitempty" validate:"omitempty"`
		// ИД группы экипажей
		CrewGroupID int `json:"crew_group_id,omitempty" validate:"omitempty"`
		// ИД службы ЕДС
		UdsID int `json:"uds_id,omitempty" validate:"omitempty"`
		// Время подачи
		SourceTime string `json:"source_time,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Призовой заказ
		IsPrize bool `json:"is_prize,omitempty" validate:"omitempty"`
		// Массив координат адресов. Первый элемент — адрес подачи,
		// последний — адрес назначения, между ними — остановки. Заполняется если определены координаты всех адресов.
		Addresses Point `json:"addresses,omitempty" validate:"omitempty"`
	}

	SelectTariffForOrderResponse struct {
		// Тариф
		TariffID int `json:"tariff_id"`
	}
)

// Подбор тарифа для заказа
func (cl *Client) SelectTariffForOrder(req SelectTariffForOrderRequest) (response SelectTariffForOrderResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100	Клиент не найден
	*/
	e := errorMap{
		100: ErrClientNotFound,
	}

	err = cl.PostJson("select_tariff_for_order", e, req, &response)

	return
}
