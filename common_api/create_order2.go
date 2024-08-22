package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateOrder2Request struct {
		// Номер телефона (необязателен, если client_id присутствует)
		Phone string `json:"phone,omitempty" validate:"omitempty,max=30"`
		// ИД клиента (необязателен, если phone присутствует)
		ClientID int `json:"client_id,omitempty" validate:"omitempty"`
		// Массив адресов. Первый элемент — адрес подачи(обязательно), последний — адрес назначения, между ними — остановки
		Addresses []Address `json:"addresses" validate:"required"`
		// Время подачи
		SourceTime string `json:"source_time" validate:"required,datetime=20060102150405"`

		// Смещения относительно серверного времени
		ServerTimeOffset int `json:"server_time_offset,omitempty" validate:"omitempty"`
		// Пассажир
		Passenger string `json:"passenger,omitempty" validate:"omitempty"`
		// Телефон для отзвона
		PhoneToDial string `json:"phone_to_dial,omitempty" validate:"omitempty,max=30"`
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
		IsPrior bool `json:"is_prior,omitempty" validate:"omitempty"`
		// Проверка на дубликат
		CheckDuplicate bool `json:"check_duplicate,omitempty" validate:"omitempty"`
		// Массив услуг. Устарело. Рекомендуется использовать параметр attribute_values
		Services *[]int `json:"services,omitempty" validate:"omitempty"`
		// Массив признаков экипажей. Устарело. Рекомендуется использовать параметр attribute_values
		CrewProps *[]int `json:"crew_props,omitempty" validate:"omitempty"`
		// Массив параметров заказа. Устарело. Рекомендуется использовать параметр attribute_values
		OrderParams *[]int `json:"order_params,omitempty" validate:"omitempty"`
		// Сумма заказа
		TotalCost float64 `json:"total_cost,omitempty" validate:"omitempty"`
		// Оплата по возможности всей суммы заказа с безналичного счета клиента (насколько хватает средств на счете)
		UseCashless bool `json:"use_cashless,omitempty" validate:"omitempty"`
		// Оплата по возможности всей суммы заказа с бонусного счета клиента (насколько хватает средств на бонусном счете)
		UseBonus bool `json:"use_bonus,omitempty" validate:"omitempty"`
		// Фиксированная сумма оплаты заказа с безналичного счета клиента (не используется, если use_cashless = true)
		CashlessSum int `json:"cashless_sum,omitempty" validate:"omitempty"`
		// Фиксированная сумма оплаты заказа с бонусного счета клиента (не используется, если use_bonus = true)
		BonusSum int `json:"bonus_sum,omitempty" validate:"omitempty"`
		// ИД сотрудника клиента (если задан client_id)
		ClientEmployeeID int `json:"client_employee_id,omitempty" validate:"omitempty"`
		// Email для отправки уведомлений
		Email string `json:"email,omitempty" validate:"omitempty,email"`
		// Время перехода из предварительного в текущие заказы, мин
		PriorToCurrentBeforeMinutes int `json:"prior_to_current_before_minutes,omitempty" validate:"omitempty"`
		// Номер рейса
		FlightNumber string `json:"flight_number,omitempty" validate:"omitempty"`
		// Использовать специальную проверку перед созданием заказа
		NeedCustomValidate bool `json:"need_custom_validate,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues *[]AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
		// Тип платежной системы ("qr", либо пусто, если не используется)
		PaymentPaySystem string `json:"payment_pay_system,omitempty" validate:"omitempty"`
	}

	CreateOrder2Response struct {
		// ИД созданного заказа
		OrderID int `json:"order_id"`
		// Текст ошибки для пользователя
		Message string `json:"message"`
	}
)

// Создание нового заказа 2
func (cl *Client) CreateOrder2(req CreateOrder2Request) (response CreateOrder2Response, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100 Заказ с такими параметрами уже создан
		101 Тариф не найден
		102 Группа экипажа не найдена
		103 Служба ЕДС не найдена
		104 Клиент не найден
		105 Район не найден
		106 Стоянка не найдена
		107 Сотрудник клиента не найден
		108 Параметр заказа не найден
		109 Атрибут не может быть привязан к заказу
		110 Клиент заблокирован
		111 Не найден клиент, который может использовать собственный счет для оплаты заказов
		112 Сотрудник клиента заблокирован
		113 Ошибка специальной проверки заказа перед созданием. В ответе будет возвращаться:
		 "data": {
		   "message":"Текст ошибки для пользователя."
		 }
		114 Недостаточно средств на безналичном счете клиента в ТМ
		115 Отрицательный баланс на безналичном счете клиента в ТМ
		116 Для клиента запрещена оплата заказа наличными. Клиент должен максимально использовать в заказе безналичную оплату (оплату с основного счета)
	*/
	e := errorMap{
		100: ErrOrderExistsWithParametrs,
		101: ErrTariffNotFound,
		102: ErrCrewNotFound,
		103: ErrUdsNotFound,
		104: ErrClientNotFound,
		105: ErrZoneNotFound,
		106: ErrStopNotFound,
		107: ErrCustomerClientNotFound,
		108: ErrOrderParameterNotFound,
		109: ErrAttributeCannotBeBoundOrder,
		110: ErrClientBlocked,
		111: ErrClientWhoCanUseTheirOwnNotFound,
		112: ErrCustomerClientBlocked,
		113: ErrSpecialOrderCheck,
		114: ErrInsufficientFundsCashless,
		115: ErrNegativeBalanceCashless,
		116: ErrCashPaymentNotAllowed,
	}

	err = cl.PostJson("create_order2", e, req, &response)

	return
}
