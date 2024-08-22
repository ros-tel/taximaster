package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	UpdateGlobalAttributeRequest struct {
		// ИД изменяемого атрибута
		ID int `json:"id" validate:"required"`

		// Значение, если тип атрибута "Логический"
		BoolValue bool `json:"bool_value,omitempty" validate:"omitempty"`
		//  Значение, если тип атрибута:
		//  - "Число" (непосредственное значение)
		//  - "Число (выбор из списка)" (непосредственное значение выбранного элемента списка)
		//  - "Перечисляемый" (значение выбранного элемента перечисления)
		//  - "Дата" и "Дата/время" (Unix-время, всегда целое число)
		NumValue float64 `json:"num_value,omitempty" validate:"omitempty"`
		// 	Значение, если тип атрибута "Строка"
		StrValue string `json:"str_value,omitempty" validate:"omitempty"`
	}
)

// Изменение глобального атрибута
func (cl *Client) UpdateGlobalAttribute(req UpdateGlobalAttributeRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100	Атрибут не найден
		101	Атрибут не глобальный
	*/
	e := errorMap{
		100: ErrAttributeNotFound,
		101: ErrAttributeIsNotGlobal,
	}

	err = cl.PostJson("update_global_attribute", e, req, &response)

	return
}
