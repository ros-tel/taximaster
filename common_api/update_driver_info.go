package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	UpdateDriverInfoRequest struct {
		// ИД редактируемого водителя
		DriverID int `json:"driver_id" validate:"required"`

		// ФИО водителя
		Name *string `json:"name,omitempty" validate:"omitempty"`
		// ИД основного автомобиля
		CarID *int `json:"car_id,omitempty" validate:"omitempty"`
		// Пароль
		Password *string `json:"password,omitempty" validate:"omitempty"`
		// ИД службы ЕДС (обязательное поле, если используется ЕДС, иначе можно не указывать)
		UdsID *int `json:"uds_id,omitempty" validate:"omitempty"`
		// Паспортные данные
		Passport *string `json:"passport,omitempty" validate:"omitempty"`
		// Водительское удостоверение
		DriverLicense *string `json:"driver_license,omitempty" validate:"omitempty"`
		// Тип работника (0 - работник компании, 1 - частник)
		EmployeeType *int `json:"employee_type" validate:"omitempty,eq=0|eq=1"`
		// День рождения
		Birthday *string `json:"birthday,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Табельный номер
		Number *string `json:"number,omitempty" validate:"omitempty"`
		// Удостоверение
		License *string `json:"license,omitempty" validate:"omitempty"`
		// Дата приема на работу
		StartDate *string `json:"start_date,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Дата окончания договора
		LicDate *string `json:"lic_date,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Терминальный аккаунт (если не указан, будет сгенерирован автоматически), должен состоять из 5 цифр
		TermAccount *string `json:"term_account,omitempty" validate:"omitempty"`
		// Описание
		Comment *string `json:"comment,omitempty" validate:"omitempty"`
		// Временная блокировка до
		TimeBlock *string `json:"time_block,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Заблокирован
		IsLocked *bool `json:"is_locked,omitempty" validate:"omitempty"`
		// Причина блокировки
		LockDescription *string `json:"lock_description,omitempty" validate:"omitempty"`
		// Уволен
		IsDismissed *bool `json:"is_dismissed" validate:"omitempty"`
		// Причина увольнения
		DismissDescription *string `json:"dismiss_description,omitempty" validate:"omitempty"`
		// Имя для TaxoPhone
		NameForTaxophone *string `json:"name_for_taxophone,omitempty" validate:"omitempty"`
		// Фотография водителя
		DriverPhoto *string `json:"driver_photo,omitempty" validate:"omitempty,base64"`
		// Массив телефонов водителя
		Phones *[]Phone `json:"phones,omitempty" validate:"omitempty"`
		// Массив параметров водителя. Устарело. Рекомендуется использовать параметр attribute_values
		OrderParams *[]int `json:"order_params,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues *[]AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
	}
)

// Обновление информации о водителе
func (cl *Client) UpdateDriverInfo(req UpdateDriverInfoRequest) (EmptyResponse, error) {
	var response = EmptyResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	err = cl.PostJson("update_driver_info", req, &response)

	return response, err
}
