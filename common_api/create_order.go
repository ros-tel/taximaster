package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	CreateOrderRequest struct {
		// Номер телефона
		Phone string `json:"phone" validate:"required,max=30"`
		// Адрес подачи
		Source string `json:"source" validate:"required"`
		// Время подачи
		SourceTime string `json:"source_time" validate:"required,datetime=20060102150405"`

		// Адрес назначения
		Dest string `json:"dest,omitempty" validate:"omitempty"`
		// Заказчик
		Customer string `json:"customer,omitempty" validate:"omitempty"`
		// Комментарий
		Comment string `json:"comment,omitempty" validate:"omitempty"`
		// ИД группы экипажей
		CrewGroupID int `json:"crew_group_id,omitempty" validate:"omitempty"`
		// ИД службы ЕДС
		UdsID int `json:"uds_id,omitempty" validate:"omitempty"`
		// ИД тарифа
		TariffID int `json:"tariff_id,omitempty" validate:"omitempty"`
		// Предварительный заказ
		IsPrior *bool `json:"is_prior,omitempty" validate:"omitempty"`
		// Долгота адреса подачи
		SourceLon *float64 `json:"source_lon,omitempty" validate:"omitempty"`
		// Широта адреса подачи
		SourceLat *float64 `json:"source_lat,omitempty" validate:"omitempty"`
		// Долгота адреса назначения
		DestLon *float64 `json:"dest_lon,omitempty" validate:"omitempty"`
		// Широта адреса назначения
		DestLat *float64 `json:"dest_lat,omitempty" validate:"omitempty"`
	}

	CreateOrderResponse struct {
		// ИД созданного заказа
		OrderID int `json:"order_id"`
	}
)

// Создание нового заказа
func (cl *Client) CreateOrder(req CreateOrderRequest) (response CreateOrderResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("phone", req.Phone)
	v.Add("source", req.Source)
	v.Add("source_time", req.SourceTime)
	if req.Dest != "" {
		v.Add("dest", req.Dest)
	}
	if req.Customer != "" {
		v.Add("customer", req.Customer)
	}
	if req.Comment != "" {
		v.Add("comment", req.Comment)
	}
	if req.CrewGroupID != 0 {
		v.Add("crew_group_id", strconv.Itoa(req.CrewGroupID))
	}
	if req.UdsID != 0 {
		v.Add("uds_id", strconv.Itoa(req.UdsID))
	}
	if req.TariffID != 0 {
		v.Add("tariff_id", strconv.Itoa(req.TariffID))
	}
	if req.IsPrior != nil {
		v.Add("uds_id", strconv.FormatBool(*req.IsPrior))
	}
	if req.SourceLon != nil {
		v.Add("source_lon", strconv.FormatFloat(*req.SourceLon, 'g', -1, 64))
	}
	if req.SourceLat != nil {
		v.Add("source_lat", strconv.FormatFloat(*req.SourceLat, 'g', -1, 64))
	}
	if req.DestLon != nil {
		v.Add("dest_lon", strconv.FormatFloat(*req.DestLon, 'g', -1, 64))
	}
	if req.DestLat != nil {
		v.Add("dest_lat", strconv.FormatFloat(*req.DestLat, 'g', -1, 64))
	}

	/*
		100	Заказ с такими параметрами уже создан
		101	Тариф не найден
		102	Группа экипажа не найдена
		103	Служба ЕДС не найдена
		110	Клиент заблокирован
		111	Не найден клиент, который может использовать собственный счет для оплаты заказов
		114	Недостаточно средств на безналичном счете клиента в ТМ
		115	Отрицательный баланс на безналичном счете клиента в ТМ
		116	Для клиента запрещена оплата заказа наличными. Клиент должен максимально использовать в заказе безналичную оплату (оплату с основного счета)
	*/

	e := errorMap{
		100: ErrOrderExistsWithParametrs,
		101: ErrTariffNotFound,
		102: ErrCrewGroupsNotFound,
		103: ErrUdsNotFound,
		110: ErrClientBlocked,
		111: ErrClientWhoCanUseTheirOwnNotFound,
		114: ErrInsufficientFundsCashless,
		115: ErrNegativeBalanceCashless,
		116: ErrCashPaymentNotAllowed,
	}

	err = cl.Post("create_order", e, v, &response)

	return
}
