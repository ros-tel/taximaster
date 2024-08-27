package common_api

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/ros-tel/taximaster/validator"
)

type (
	CalcOrderCostRequest struct {
		// ИД тарифа
		TariffID int `validate:"required"`

		// Время подачи
		SourceTime string `validate:"omitempty,datetime=20060102150405"`
		// Предварительный заказ
		IsPrior *bool `validate:"omitempty"`
		// ИД клиента
		ClientID int `validate:"omitempty"`
		// ИД сотрудника клиента
		ClientEmployeeID int `validate:"omitempty"`
		// ИД скидки
		DiscountID int `validate:"omitempty"`
		// ИД дисконтной карты
		DiscCardID int `validate:"omitempty"`
		// ИД района подачи
		SourceZoneID int `validate:"omitempty"`
		// ИД района назначения
		DestZoneID int `validate:"omitempty"`
		// Километраж по городу
		DistanceCity float64 `validate:"omitempty"`
		// Километраж за городом
		DistanceCountry float64 `validate:"omitempty"`
		// Километраж до подачи за городом
		SourceDistanceCountry float64 `validate:"omitempty"`
		// Загородный заказ
		IsCountry *bool `validate:"omitempty"`
		// Время ожидания посадки клиента в минутах
		WaitingMinutes int `validate:"omitempty"`
		// Почасовой заказ
		IsHourly *bool `validate:"omitempty"`
		// Длительность почасового заказа в минутах
		HourlyMinutes int `validate:"omitempty"`
		// Призовой заказ
		IsPrize *bool `validate:"omitempty"`
		// Обратный путь за городом
		BackWay *bool `validate:"omitempty"`
		// Список ИД услуг, пример: []int{1, 2, 3} Устарело. Рекомендуется использовать параметр order_params.
		Services []int `validate:"omitempty"`
		// Список ИД параметров заказа, пример: []int{1, 2, 3}
		OrderParams []int `validate:"omitempty"`
		// Признак безналичного заказа
		Cashless *bool `validate:"omitempty"`
	}

	CalcOrderCostResponse struct {
		// Рассчитанная общая сумма заказа
		Sum float64 `json:"sum"`
		// Дополнительная информация по расчету суммы заказа
		Info []struct {
			// Описание позиции дополнительной информации по расчету суммы заказа
			Comment string `json:"comment"`
			// Сумма позиции дополнительной информации по расчету суммы заказа
			Sum string `json:"sum"`
		} `json:"info"`
	}
)

// Расчет суммы заказа
func (cl *Client) CalcOrderCost(req CalcOrderCostRequest) (response CalcOrderCostResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("tariff_id", strconv.Itoa(req.TariffID))
	if req.SourceTime != "" {
		v.Add("source_time", req.SourceTime)
	}
	if req.IsPrior != nil {
		v.Add("is_prior", strconv.FormatBool(*req.IsPrior))
	}
	if req.SourceTime != "" {
		v.Add("source_time", req.SourceTime)
	}
	if req.ClientID != 0 {
		v.Add("client_id", strconv.Itoa(req.ClientID))
	}
	if req.ClientEmployeeID != 0 {
		v.Add("client_employee_id", strconv.Itoa(req.ClientEmployeeID))
	}
	if req.DiscountID != 0 {
		v.Add("discount_id", strconv.Itoa(req.DiscountID))
	}
	if req.DiscCardID != 0 {
		v.Add("disc_card_id", strconv.Itoa(req.DiscCardID))
	}
	if req.SourceZoneID != 0 {
		v.Add("source_zone_id", strconv.Itoa(req.SourceZoneID))
	}
	if req.DestZoneID != 0 {
		v.Add("dest_zone_id", strconv.Itoa(req.DestZoneID))
	}
	if req.DistanceCity != 0 {
		v.Add("distance_city", strconv.FormatFloat(req.DistanceCity, 'g', -1, 64))
	}
	if req.DistanceCountry != 0 {
		v.Add("distance_country", strconv.FormatFloat(req.DistanceCountry, 'g', -1, 64))
	}
	if req.SourceDistanceCountry != 0 {
		v.Add("source_distance_country", strconv.FormatFloat(req.SourceDistanceCountry, 'g', -1, 64))
	}
	if req.IsCountry != nil {
		v.Add("is_country", strconv.FormatBool(*req.IsCountry))
	}
	if req.WaitingMinutes != 0 {
		v.Add("waiting_minutes", strconv.Itoa(req.WaitingMinutes))
	}
	if req.IsHourly != nil {
		v.Add("is_hourly", strconv.FormatBool(*req.IsHourly))
	}
	if req.HourlyMinutes != 0 {
		v.Add("hourly_minutes", strconv.Itoa(req.HourlyMinutes))
	}
	if req.IsPrize != nil {
		v.Add("is_prize", strconv.FormatBool(*req.IsPrize))
	}
	if req.BackWay != nil {
		v.Add("back_way", strconv.FormatBool(*req.BackWay))
	}
	if len(req.Services) != 0 {
		stringSlice := make([]string, len(req.Services))

		for i, num := range req.Services {
			stringSlice[i] = strconv.Itoa(num)
		}

		v.Add("services", strings.Join(stringSlice, ";"))
	}
	if len(req.OrderParams) != 0 {
		stringSlice := make([]string, len(req.OrderParams))

		for i, num := range req.Services {
			stringSlice[i] = strconv.Itoa(num)
		}

		v.Add("order_params", strings.Join(stringSlice, ";"))
	}
	if req.Cashless != nil {
		v.Add("cashless", strconv.FormatBool(*req.Cashless))
	}

	/*
		100	Тариф не найден
		101	Ошибка при расчете по тарифу
		102	Скидка не найдена
		103	Клиент не найден
		104	Район подачи не найден
		105	Район назначения не найден
		106	Дисконтная карта не найдена
		107	Район остановки не найден
		108	Группа экипажа не найдена
		109	Служба ЕДС не найдена
		110	Дисконтная карта не действительна
		111	Не найден сотрудник клиента
	*/
	e := errorMap{
		100: ErrTariffNotFound,
		101: ErrCalculationByTariff,
		102: ErrDiscountNotFound,
		103: ErrClientNotFound,
		104: ErrZoneSourceNotFound,
		105: ErrZoneDestinationNotFound,
		106: ErrDiscountCardNotFound,
		107: ErrZoneStopNotFound,
		108: ErrCrewNotFound,
		109: ErrUdsNotFound,
		110: ErrDiscountCardIsNotValid,
		111: ErrCustomerClientNotFound,
	}

	err = cl.Get("calc_order_cost", e, v, &response)

	return
}
