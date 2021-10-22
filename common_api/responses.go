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
		NumValue *float64 `json:"num_value,omitempty"`
		// 	Значение, если тип атрибута "Строка"
		StrValue *string `json:"str_value,omitempty"`
	}
)
