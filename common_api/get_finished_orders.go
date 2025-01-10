package common_api

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetFinishedOrdersRequest struct {
		// Начало периода
		StartTime string `validate:"required,datetime=20060102150405"`
		// Конец периода
		FinishTime string `validate:"required,datetime=20060102150405"`

		// ИД клиента
		ClientID int `validate:"omitempty"`
		// ИД сотрудника (только если указан ИД клиента)
		ClientEmployeeID int `validate:"omitempty"`
		// Телефон клиента
		Phone string `validate:"omitempty,max=30"`
		// ИД экипажа
		CrewID int `validate:"omitempty"`
		// ИД водителя
		DriverID int `validate:"omitempty"`
		// Тип состояния заказа
		// Может принимать значения:
		// - "all" - все
		// - "finished" - выполненные
		// - "aborted" - прекращенные
		StateType string `validate:"omitempty,eq=all|eq=finished|eq=aborted"`
		// Список ИД состояний заказа, пример: []int{1, 2, 3}
		StateIDs []int `validate:"omitempty"`
		// Список возвращаемых полей через запятую
		Fields string `validate:"omitempty"`
	}

	GetFinishedOrdersResponse struct {
		Orders []GetFinishedOrdersArray `json:"orders"`
	}

	GetFinishedOrdersArray struct {
		// ИД заказа
		ID int `json:"id"`
		// ИД состояния заказа
		StateID int `json:"state_id"`
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
		// Массив адресов остановок
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
		// Гос.номер автомобиля
		CarNumber string `json:"car_number"`
		// Массив параметров заказа экипажа. Устарело. Рекомендуется использовать параметр attribute_values. (Возвращается только если в списке фильтра полей fields запросили поле order_params)
		OrderParams []int `json:"order_params"`
		// Способ создания заказа. Может принимать значения:
		// - "operator" — заказ создан оператором
		// - "sms" — заказ создан через смс
		// - "market" — заказ из биржи
		// - "common_api" — заказ создан через api
		// - "t_api" — заказ создан через api
		// - "taxophone" — заказ создан через таксофон с телефона
		// - "driver" — заказ создан водителем
		// - "daily_order" — ежедневный заказ
		// - "taxophone_web" — заказ создан через таксофон с сайта
		// - "unknown" — неизвестный
		CreationWay string `json:"creation_way"`
		// Предварительный заказ
		IsPrior bool `json:"is_prior"`
		// Email для уведомлений
		Email string `json:"email"`
		// Номер рейса
		FlightNumber string `json:"flight_number"`
		// Признак того, что заказ составной
		IsCombined bool `json:"is_combined"`
		// Массив ИД заказов-частей, передается только для составного заказа
		CombinedOrderPartsIDs []int `json:"combined_order_parts_ids"`
		// Признак того, что заказ является частью составного
		IsPartOfCombined bool `json:"is_part_of_combined"`
		// ИД составного заказа, передается, только если заказ является частью составного
		CombinedOrderID int `json:"combined_order_id"`
		// Стоимость заказа без учета скидок (наценок)
		Sum float64 `json:"sum"`
		// Итоговая стоимость заказа
		TotalSum float64 `json:"total_sum"`
		// Сумма наличными
		CashSum float64 `json:"cash_sum"`
		// Сумма безналичными
		CashlessSum float64 `json:"cashless_sum"`
		// Сумма бонусами
		BonusSum float64 `json:"bonus_sum"`
		// Сумма банковской картой
		BankCardSum float64 `json:"bank_card_sum"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values"`
		// Чек TMDriver. Данный узел выводится только, если по заказу есть чек и если в списке фильтра полей fields запросили поле bill
		Bill []Bill `json:"bill"`
		// Признак заказа-аукциона
		IsAuction bool `json:"is_auction"`
		// Тип платежной системы ("card", "gpay", "apple_pay", "qr", "sber_pay", либо пусто, если не используется)
		PaymentPaySystem string `json:"payment_pay_system"`
		// Комментарий
		Comment string `json:"comment"`
	}
)

// Запрос выполненных заказов
func (cl *Client) GetFinishedOrders(req GetFinishedOrdersRequest) (response GetFinishedOrdersResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}

	v.Add("start_time", req.StartTime)
	v.Add("finish_time", req.FinishTime)
	if req.ClientID != 0 {
		v.Add("client_id", strconv.Itoa(req.ClientID))
	}
	if req.ClientEmployeeID != 0 {
		v.Add("client_employee_id", strconv.Itoa(req.ClientEmployeeID))
	}
	if req.Phone != "" {
		v.Add("phone", req.Phone)
	}
	if req.CrewID != 0 {
		v.Add("crew_id", strconv.Itoa(req.CrewID))
	}
	if req.DriverID != 0 {
		v.Add("driver_id", strconv.Itoa(req.DriverID))
	}
	if req.StateType != "" {
		v.Add("state_type", req.StateType)
	}
	if len(req.StateIDs) != 0 {
		stringSlice := make([]string, len(req.StateIDs))

		for i, num := range req.StateIDs {
			stringSlice[i] = strconv.Itoa(num)
		}

		v.Add("state_ids", strings.Join(stringSlice, ";"))
	}
	if req.Fields != "" {
		v.Add("fields", req.Fields)
	}

	err = cl.Get("get_finished_orders", nil, v, &response)

	return
}
