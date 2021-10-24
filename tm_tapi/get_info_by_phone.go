package tm_tapi

import (
	"net/url"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetInfoByPhoneRequest struct {
		// Номер телефона
		Phone string `validate:"required,max=16"`
		// Список полей, которые необходимо вернуть. Поля перечисляются через символ "-"
		Fields string `validate:"required"`
	}

	GetInfoByPhoneResponse struct {
		// Тип телефона звонящего:
		// 1 - если звонит водитель с основного телефона
		// 2 - если звонит физическое лицо
		// 3 - если звонит юридическое лицо
		// 4 - если звонит номер из справочника Телефоны
		// 5 - если звонит водитель с неосновного телефона
		// 6 - если звонит ЦОЗ водитель
		// 0 - неизвестный номер
		PhoneType int `xml:"PHONE_TYPE"`
		// Номер телефона для отзвона по заказу
		PhoneToDial string `xml:"PHONE_TO_DIAL"`
		// ИД экипажа
		CrewID int `xml:"CREW_ID"`
		// Признак предварительного заказа 0 или 1 (false или true) (БИТ Мастер, чтоб вам всю жизнь с таким API работать)
		IsPrior string `xml:"IS_PRIOR"`
		// Признак призового заказа 0 или 1 (false или true)
		IsPrize string `xml:"IS_PRIZE"`
		// ИД клиента из заказа
		OrderClientID int `xml:"ORDER_CLIENT_ID"`
		// Номер телефона водителя
		DriverPhone string `xml:"DRIVER_PHONE"`
		// Состояние экипажа
		CrewSystemState int `xml:"CREW_SYSTEMSTATE"`
		// ИД клиента
		ClientID int `xml:"CLIENT_ID"`
		// Тип клиента
		ClientType int `xml:"CLIENT_TYPE"`
		// ИД категории телефона
		CategoryID int `xml:"CATEGORYID"`
		// Системное значение категории телефона:
		// 0 - обычный
		// 1 - черный
		// 2 - белый
		// 3 - серый
		PhoneSystemCategory int `xml:"PHONE_SYSTEM_CATEGORY"`
		// ИД заказа
		OrderID int `xml:"ORDER_ID"`
		// ИД водителя
		DriverID int `xml:"DRIVER_ID"`
		// Состояние заказа
		OrderState string `xml:"ORDER_STATE"`
		// Баланс счета водителя
		DriverRemainder float64 `xml:"DRIVER_REMAINDER"`
		// Сумма заказа с учетом всех скидок
		DiscountedSumm float64 `xml:"DISCOUNTEDSUMM"`
		// ИД группы экипажа (если в заказе указан экипаж, то берется ИД группы экипажа, если нет — из карты заказа)
		CrewGroupID int `xml:"CREW_GROUP_ID"`
		// ИД первой группы экипажа
		FirstCrewGroupID int `xml:"FIRST_CREW_GROUP_ID"`
		// Баланс клиента
		ClientBalance float64 `xml:"CLIENT_BALANCE"`
		// Время пути водителя до адреса подачи в минутах
		DriverTimeCount int `xml:"DRIVER_TIMECOUNT"`
		// Время оставшееся до подачи в минутах
		SourceTimecount int `xml:"SOURCE_TIMECOUNT"`
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
		// Координаты места подачи. Порядок координат: долгота адреса, широта адреса.
		OrderCoords string `xml:"ORDER_COORDS"`
		// Координаты назначенного экипажа. Порядок координат: долгота адреса, широта адреса.
		CrewCoords string `xml:"CREW_COORDS"`
		// Тип биржи заказов:
		// 2 — Яндекс
		// 3 — ЦОЗ
		MarketType int `xml:"MARKET_TYPE"`
		// Для телефона клиента — количество текущих и предварительных заказов по телефону
		// Для телефона водителя — количество текущих заказов, на которые назначен данный водитель
		OrdersCount int `xml:"ORDERS_COUNT"`
		// ИД группы клиента
		ClientGroupID int `xml:"CLIENT_GROUP_ID"`
		// Бонусный баланс клиента
		ClientBonusBalance float64 `xml:"CLIENT_BONUS_BALANCE"`
		// Признак наличия шашек. Пустая строка – номер клиента или водителя без экипажа.
		// 1 – есть шашки
		// 0 – нет шашек.
		AdLightHouse string `xml:"AD_LIGHTHOUSE"`
		// Состояние экипажа
		Crew_State int `xml:"CREW_STATE"`
		// Дата и время подачи
		SourceTime string `xml:"SOURCE_TIME"`
		// Начало фактической смены водителя
		DrvShiftStartTime string `xml:"DRV_SHIFT_START_TIME"`
		// Список логинов SIP–аккаунтов пользователя, принявшего заказ
		StartUserSipAccounts string `xml:"START_USER_SIP_ACCOUNTS"`
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
		// Признак блокировки клиента
		ClientIsLocked string `xml:"CLIENT_IS_LOCKED"`
	}
)

// Запрос информации по номеру телефона
func (cl *Client) GetInfoByPhone(req GetInfoByPhoneRequest) (GetInfoByPhoneResponse, error) {
	var response = GetInfoByPhoneResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("PHONE", req.Phone)
	v.Add("FIELDS", req.Fields)

	err = cl.Get("get_info_by_phone", v, &response)

	return response, err
}
