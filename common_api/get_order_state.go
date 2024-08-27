package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetOrderStateRequest struct {
		// ИД заказа
		OrderID int `validate:"required"`

		// Список возвращаемых полей через запятую
		Fields string `validate:"omitempty"`
	}

	GetOrderStateResponse struct {
		// ИД заказа
		OrderID int `json:"order_id"`
		// ИД состояния заказа
		StateID int `json:"state_id"`
		// Тип состояния заказа
		// Может принимать значения:
		// - "new_order" - новый заказ
		// - "driver_assigned" - водитель назначен
		// - "car_at_place" - машина подъехала на место
		// - "client_inside" - клиент в машине
		// - "finished" - заказ успешно завершен
		// - "aborted" - заказ прекращен
		StateKind string `json:"state_kind"`
		// ИД экипажа
		CrewID int `json:"crew_id"`
		// ИД предварительного экипажа
		PriorCrewID int `json:"prior_crew_id"`
		// ИД водителя
		DriverID int `json:"driver_id"`
		// ИД автомобиля
		CarID int `json:"car_id"`
		// Смещение относительно серверного времени
		ServerTimeOffset int `json:"server_time_offset"`
		// Время создания заказа
		StartTime string `json:"start_time"`
		// Время подачи
		SourceTime string `json:"source_time"`
		// Время завершения заказа
		FinishTime string `json:"finish_time"`
		// Адрес подачи
		Source string `json:"source"`
		// Широта адреса подачи
		SourceLat float64 `json:"source_lat"`
		// Долгота адреса подачи
		SourceLon float64 `json:"source_lon"`
		// Адрес назначения
		Destination string `json:"destination"`
		// Широта адреса назначения
		DestinationLat float64 `json:"destination_lat"`
		// Долгота адреса назначения
		DestinationLon float64 `json:"destination_lon"`
		// Информация по остановкам заказа
		Stops []Stop `json:"stops"`
		// Фактический километраж
		TripDistance float64 `json:"trip_distance"`
		// Фактическое время в пути
		TripTime int `json:"trip_time"`
		// Заказчик
		Customer string `json:"customer"`
		// Пассажир
		Passenger string `json:"passenger"`
		// Номер телефона
		Phone string `json:"phone"`
		// Номер телефона для отзвона
		PhoneToDial string `json:"phone_to_dial"`
		// ИД клиента
		ClientID int `json:"client_id"`
		// Имя клиента
		ClientName string `json:"client_name"`
		// ИД группы клиента
		ClientGroupID int `json:"client_group_id"`
		// Название группы клиента
		ClientGroupName string `json:"client_group_name"`
		// ИД сотрудника клиента
		ClientEmployeeID int `json:"client_employee_id"`
		// ИД группы экипажей, которая указана в заказе
		OrderCrewGroupID int `json:"order_crew_group_id"`
		// ИД тарифа
		TariffID int `json:"tariff_id"`
		// Марка автомобиля
		CarMark string `json:"car_mark"`
		// Модель автомобиля
		CarModel string `json:"car_model"`
		// Цвет автомобиля
		CarColor string `json:"car_color"`
		// Государственный номер автомобиля
		CarNumber string `json:"car_number"`
		// Состояние подтвержденности заказа водителем или оператором.
		// Может принимать значения:
		// - "not_confirmed" - не подтверждено
		// - "confirmed_by_driver" - заказ принят водителем
		// - "confirmed_by_oper" - заказ подтвержден оператором
		Confirmed string `json:"confirmed"`
		// Координаты экипажа
		CrewCoords Point `json:"crew_coords"`
		// Массив параметров заказа экипажа
		OrderParams []int `json:"order_params"`
		// Способ создания заказа. Может принимать значения:
		// - "operator" - заказ создан оператором
		// - "sms" - заказ создан через смс
		// - "market" - заказ из биржи
		// - "common_api" - заказ создан через api
		// - "t_api" - заказ создан через api
		// - "taxophone" - заказ создан через таксофон с телефона
		// - "driver" - заказ создан водителем
		// - "daily_order" - ежедневный заказ
		// - "taxophone_web" - заказ создан через таксофон с сайта
		// - "unknown" - неизвестный
		CreationWay string `json:"creation_way"`
		// Предварительный заказ
		IsPrior bool `json:"is_prior"`
		// Предварительный заказ на вкладке "Предварительные"
		IsReallyPrior bool `json:"is_really_prior"`
		// Email для уведомлений
		Email string `json:"email"`
		// Время перехода из предварительного в текущие заказы, мин
		PriorToCurrentBeforeMinutes int `json:"prior_to_current_before_minutes"`
		// Номер рейса
		FlightNumber string `json:"flight_number"`
		// Сумма без скидки
		Sum float64 `json:"sum"`
		// Итоговая сумма заказа
		TotalSum float64 `json:"total_sum"`
		// Сумма наличными
		CashSum float64 `json:"cash_sum"`
		// 	Сумма безналичными
		CashlessSum float64 `json:"cashless_sum"`
		// Сумма бонусами
		BonusSum float64 `json:"bonus_sum"`
		// Сумма банковской картой
		BankCardSum float64 `json:"bank_card_sum"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values"`
		// Чек TMDriver. Данный узел выводится только, если по заказу есть чек
		Bill []struct {
			// Код элемента расчета
			Code string `json:"code"`
			// Наименование элемента расчета
			Text string `json:"text"`
			// Значение элемента расчета (количество)
			Value string `json:"value"`
			// Стоимость элемента расчета
			Sum string `json:"sum"`
		} `json:"bill"`
		// Фактический маршрут по заказу, выводится только если поле fact_route передали в списке фильтра полей fields
		FactRoute []struct {
			// Широта точки маршрута
			Lat float64 `json:"lat"`
			// Долгота точки маршрута
			Lon float64 `json:"lon"`
			// Время данной точки
			Time string `json:"time"`
			// Скорость в данный момент, км/ч
			Speed int `json:"speed"`
			// Направление движения, градусы (0 - север, 90 - восток, 180 - юг, 270 - запад)
			Direction int `json:"direction"`
		} `json:"fact_route"`
		// Признак заказа-аукциона
		IsAuction bool `json:"is_auction"`
		// Тип платежной системы ("card", "gpay", "apple_pay", "qr", "sber_pay", либо пусто, если не используется).
		PaymentPaySystem string `json:"payment_pay_system"`
	}
)

// Запрос информации о состоянии заказа
func (cl *Client) GetOrderState(req GetOrderStateRequest) (response GetOrderStateResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("order_id", strconv.Itoa(req.OrderID))
	if req.Fields != "" {
		v.Add("fields", req.Fields)
	}

	/*
		100 Заказ не найден
	*/
	e := errorMap{
		100: ErrOrderNotFound,
	}

	err = cl.Get("get_order_state", e, v, &response)

	return
}
