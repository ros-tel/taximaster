package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateOrder2Request struct {
		// Номер телефонам (необязателен, если client_id присутствует)
		Phone string `json:"phone,omitempty" validate:"omitempty,max=30"`
		// ИД клиента (необязателен, если phone присутствует)
		ClientID int `json:"client_id,omitempty" validate:"omitempty"`
		// Массив адресов. Первый элемент — адрес подачи(обязательно), последний — адрес назначения, между ними — остановки
		Addresses []Address `json:"addresses" validate:"required"`
		// Смещения относительно серверного времени
		ServerTimeOffset int `json:"server_time_offset" validate:"omitempty"`
		// Время подачи
		SourceTime string `json:"source_time" validate:"required,datetime=20060102150405"`

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
		Email string `json:"email,omitempty" validate:"omitempty"`
		// Время перехода из предварительного в текущие заказы, мин
		PriorToCurrentBeforeMinutes int `json:"prior_to_current_before_minutes,omitempty" validate:"omitempty"`
		// Номер рейса
		FlightNumber string `json:"flight_number,omitempty" validate:"omitempty"`
		// Использовать специальную проверку перед созданием заказа
		NeedCustomValidate bool `json:"need_custom_validate,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues *[]AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
	}

	CreateOrder2Response struct {
		// ИД созданного заказа
		OrderID int `json:"order_id"`
	}
)

// Создание нового заказа 2
func (cl *Client) CreateOrder2(req CreateOrder2Request) (CreateOrder2Response, error) {
	var response = CreateOrder2Response{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	err = cl.PostJson("create_order2", req, &response)

	return response, err
}
