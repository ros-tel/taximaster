package tm_tapi

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetInfoByOrderIDRequest struct {
		// ИД заказа
		OrderID int `validate:"required"`
		// Список полей, которые необходимо вернуть. Поля перечисляются через символ "-"
		Fields string `validate:"required"`
	}

	GetInfoByOrderIDResponse struct {
		// Время пути водителя до адреса подачи в минутах
		DriverTimeCount int `xml:"DRIVER_TIMECOUNT"`
		// Запись с информацией о цвете
		SoundColor string `xml:"SOUND_COLOR"`
		// Запись с информацией о марке автомобиля
		SoundMark string `xml:"SOUND_MARK"`
		// Государственный номер автомобиля
		GosNumber string `xml:"GOSNUMBER"`
		// Цвет автомобиля
		CarColor string `xml:"CAR_COLOR"`
		// Марка автомобиля
		CarMark string `xml:"CAR_MARK"`
		// ИД группы экипажа (если в заказе указан экипаж, то берется ИД группы экипажа, если нет — из карты заказа)
		CrewGroupID int `xml:"CREW_GROUP_ID"`
		// ИД первой группы экипажа
		FirstCrewGroupID int `xml:"FIRST_CREW_GROUP_ID"`
		// Признак предварительного заказа 0 или 1 (false или true) (БИТ Мастер, чтоб вам всю жизнь с таким API работать)
		IsPrior string `xml:"IS_PRIOR"`
		// Признак призового заказа 0 или 1 (false или true)
		IsPrize string `xml:"IS_PRIZE"`
		// Сумма заказа с учетом всех скидок
		DiscountedSumm float64 `xml:"DISCOUNTEDSUMM"`
		// Номер телефона водителя
		DriverPhone string `xml:"DRIVER_PHONE"`
		// ИД категории телефона
		CategoryID int `xml:"CATEGORYID"`
		// Время оставшееся до подачи в минутах
		SourceTimecount int `xml:"SOURCE_TIMECOUNT"`
		// Координаты места подачи. Порядок координат: долгота адреса, широта адреса.
		OrderCoords string `xml:"ORDER_COORDS"`
		// Координаты назначенного экипажа. Порядок координат: долгота адреса, широта адреса.
		CrewCoords string `xml:"CREW_COORDS"`
		// Состояние заказа
		OrderState string `xml:"ORDER_STATE"`
		// Тип биржи заказов:
		// 2 — Яндекс
		// 3 — ЦОЗ
		MarketType int `xml:"MARKET_TYPE"`
		// Признак наличия шашек. Пустая строка – номер клиента или водителя без экипажа.
		// 1 – есть шашки
		// 0 – нет шашек.
		AdLightHouse string `xml:"AD_LIGHTHOUSE"`
		// Состояние экипажа
		Crew_State int `xml:"CREW_STATE"`
		// Дата и время подачи
		SourceTime string `xml:"SOURCE_TIME"`
		// Список логинов SIP–аккаунтов пользователя, принявшего заказ
		StartUserSipAccounts string `xml:"START_USER_SIP_ACCOUNTS"`
		// Список ИД параметров заказа
		OrderParams string `xml:"ORDER_PARAMS"`
		// Способ создания заказа:
		// - operator
		// - sms
		// - market
		// - common_api
		// - t_api
		// - taxophone
		// - driver
		// - daily_order
		// - taxophone_web
		// - unknown
		CreationWay string `xml:"CREATION_WAY"`
		// Телефон службы-создателя для заказов, принятых из ЦОЗ
		CreatorTaxiPhone string `xml:"CREATOR_TAXI_PHONE"`
		// Телефон службы-исполнителя для заказов, отданных в ЦОЗ
		PerformerTaxiPhone string `xml:"PERFORMER_TAXI_PHONE"`
		// ИД водителя
		// DriverID int `xml:"DRIVER_ID"`
	}
)

// Запрос информации по ИД заказа
func (cl *Client) GetInfoByOrderID(req GetInfoByOrderIDRequest) (GetInfoByOrderIDResponse, error) {
	var response = GetInfoByOrderIDResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("ORDER_ID", strconv.Itoa(req.OrderID))
	v.Add("FIELDS", req.Fields)

	err = cl.Get("get_info_by_order_id", v, &response)

	return response, err
}
