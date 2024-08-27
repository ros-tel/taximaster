package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CalcOrderCost2Request struct {
		// ИД заказа
		OrderID int `json:"order_id,omitempty" validate:"omitempty"`
		// ИД тарифа
		TariffID int `json:"tariff_id,omitempty" validate:"omitempty"`
		// Время подачи
		SourceTime string `json:"source_time,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Предварительный заказ
		IsPrior bool `json:"is_prior,omitempty" validate:"omitempty"`
		// ИД клиента
		ClientID int `json:"client_id,omitempty" validate:"omitempty"`
		// Телефон клиента
		Phone string `json:"phone,omitempty" validate:"omitempty,max=30"`
		// ИД скидки
		DiscountID int `json:"discount_id,omitempty" validate:"omitempty"`
		// ИД дисконтной карты
		DiscCardID int `json:"disc_card_id,omitempty" validate:"omitempty"`
		// ИД района подачи
		SourceZoneID int `json:"source_zone_id,omitempty" validate:"omitempty"`
		// Долгота адреса подачи
		SourceLon float64 `json:"source_lon,omitempty" validate:"omitempty"`
		// Широта адреса подачи
		SourceLat float64 `json:"source_lat,omitempty" validate:"omitempty"`
		// ИД района назначения
		DestZoneID int `json:"dest_zone_id,omitempty" validate:"omitempty"`
		// Долгота адреса назначения
		DestLon float64 `json:"dest_lon,omitempty" validate:"omitempty"`
		// Широта адреса назначения
		DestLat float64 `json:"dest_lat,omitempty" validate:"omitempty"`
		// Километраж по городу
		DistanceCity float64 `json:"distance_city,omitempty" validate:"omitempty"`
		// Километраж за городом
		DistanceCountry float64 `json:"distance_country,omitempty" validate:"omitempty"`
		// Километраж до подачи за городом
		SourceDistanceCountry float64 `json:"source_distance_country,omitempty" validate:"omitempty"`
		// Загородный заказ
		IsCountry bool `json:"is_country,omitempty" validate:"omitempty"`
		// Время ожидания посадки клиента в минутах
		WaitingMinutes int `json:"waiting_minutes,omitempty" validate:"omitempty"`
		// Почасовой заказ
		IsHourly bool `json:"is_hourly,omitempty" validate:"omitempty"`
		// Длительность почасового заказа в минутах
		HourlyMinutes int `json:"hourly_minutes,omitempty" validate:"omitempty"`
		// Призовой заказ
		IsPrize bool `json:"is_prize,omitempty" validate:"omitempty"`
		// Обратный путь за городом
		BackWay bool `json:"back_way,omitempty" validate:"omitempty"`
		// Массив параметров заказа. Устарело. Рекомендуется использовать параметр attribute_values
		OrderParams []int `json:"order_params,omitempty" validate:"omitempty"`
		// Признак безналичного заказа
		Cashless bool `json:"cashless,omitempty" validate:"omitempty"`
		// Список остановок
		Stops []struct {
			// ИД района остановки
			ZoneID int `json:"zone_id"`
			// Широта адреса остановки
			Lat float64 `json:"lat"`
			// Долгота адреса остановки
			Lon float64 `json:"lon"`
		} `json:"stops,omitempty" validate:"omitempty"`
		// ИД группы экипажа
		CrewGroupID int `json:"crew_group_id,omitempty" validate:"omitempty"`
		// ИД службы ЕДС
		UdsID int `json:"uds_id,omitempty" validate:"omitempty"`
		// Нужно ли выполнять анализ адресов и маршрута.
		// Если данный флаг установлен (analyze_route=true), то значения параметров: distance_city, distance_country, source_distance_country, переданные в данном запросе будут игнорироваться.
		// Они автоматически будут рассчитаны в ходе выполнения запроса в результате анализа адресов и маршрута.
		// Также перед анализом адресов будут автоматически найдены районы (по справочнику "Районы") для тех адресов, у которых район не указан явно (zone_id=0).
		// Также по результатам анализа адресов автоматически будут определены флаги "Загородный заказ" (is_country) и "Межгород".
		AnalyzeRoute bool `json:"analyze_route,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
	}

	CalcOrderCost2Response struct {
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

// Расчет суммы заказа 2
func (cl *Client) CalcOrderCost2(req CalcOrderCost2Request) (response CalcOrderCost2Response, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
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
		111	Сотрудник клиента не найден
		112	Атрибут не найден
		113	Атрибут не может быть привязан к заказу
		114	Заказ не найден
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
		112: ErrAttributeNotFound,
		113: ErrAttributeCannotBeBoundOrder,
		114: ErrOrderNotFound,
	}

	err = cl.PostJson("calc_order_cost2", e, req, &response)

	return
}
