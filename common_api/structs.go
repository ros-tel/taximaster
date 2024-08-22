package common_api

type (
	Phone struct {
		// Номер телефона
		Phone string `json:"phone"`
		// Признак основного телефона
		IsDefault bool `json:"is_default"`
		// Использовать для отзвона
		UseForCall bool `json:"use_for_call"`
	}

	Account struct {
		// Тип счета
		AccountKind int `json:"account_kind"`
		// Баланс счета
		Balance float64 `json:"balance"`
		// Баланс счета клиента с учетом вложенных клиентов
		BalanceWithChildren float64 `json:"balance_with_children"`
	}

	AttributeValue struct {
		// Идентификатор атрибута
		ID int `json:"id"`
		// Значение, если тип атрибута "Логический"
		BoolValue *bool `json:"bool_value,omitempty"`
		//  Значение, если тип атрибута:
		//  - "Число" (непосредственное значение)
		//  - "Число (выбор из списка)" (непосредственное значение выбранного элемента списка)
		//  - "Перечисляемый" (значение выбранного элемента перечисления)
		//  - "Дата" и "Дата/время" (Unix-время, всегда целое число)
		NumValue *float64 `json:"num_value,omitempty"`
		// 	Значение, если тип атрибута "Строка"
		StrValue *string `json:"str_value,omitempty"`
	}

	Point struct {
		// Широта
		Lat float64 `json:"lat"`
		// Долгота
		Lon float64 `json:"lon"`
	}

	Stop struct {
		// Адрес остановки
		Address string `json:"address"`
		// Широта адреса остановки
		Lat float64 `json:"lat"`
		// Долгота адреса остановки
		Lon float64 `json:"lon"`
	}

	Address struct {
		// Адрес
		Address string `json:"address,omitempty"`
		// Широта адреса
		Lat float64 `json:"lat"`
		// Долгота адреса
		Lon float64 `json:"lon"`
		// ИД района
		ZoneID int `json:"zone_id"`
		// ИД стоянки
		ParkingID int `json:"parking_id"`
	}
)
