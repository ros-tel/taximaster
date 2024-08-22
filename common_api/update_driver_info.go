package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	UpdateDriverInfoRequest struct {
		// ИД редактируемого водителя
		DriverID int `json:"driver_id" validate:"required"`

		// ФИО водителя
		Name string `json:"name,omitempty" validate:"omitempty"`
		// ИД основного автомобиля
		CarID int `json:"car_id,omitempty" validate:"omitempty"`
		// Пароль
		Password string `json:"password,omitempty" validate:"omitempty"`
		// ИД службы ЕДС (обязательное поле, если используется ЕДС, иначе можно не указывать)
		UdsID int `json:"uds_id,omitempty" validate:"omitempty"`
		// Неосновной телефон водителя (устаревший параметр)
		HomePhone string `json:"home_phone,omitempty" validate:"omitempty"`
		// Основной телефон водителя (устаревший параметр)
		MobilePhone string `json:"mobile_phone,omitempty" validate:"omitempty"`
		// Паспортные данные
		Passport string `json:"passport,omitempty" validate:"omitempty"`
		// Водительское удостоверение
		DriverLicense string `json:"driver_license,omitempty" validate:"omitempty"`
		// Разрешение на перевозку
		License string `json:"license,omitempty" validate:"omitempty"`
		// Тип работника (0 - работник компании, 1 - частник)
		EmployeeType int `json:"employee_type" validate:"omitempty,eq=0|eq=1"`
		// День рождения
		Birthday string `json:"birthday,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Табельный номер
		Number string `json:"number,omitempty" validate:"omitempty"`
		// Дата приема на работу
		StartDate string `json:"start_date,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Дата окончания договора
		LicDate string `json:"lic_date,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Терминальный аккаунт (если не указан, будет сгенерирован автоматически), должен состоять из 5 цифр
		TermAccount string `json:"term_account,omitempty" validate:"omitempty"`
		// Описание
		Comment string `json:"comment,omitempty" validate:"omitempty"`
		// Временная блокировка до
		TimeBlock string `json:"time_block,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Заблокирован
		IsLocked bool `json:"is_locked,omitempty" validate:"omitempty"`
		// Причина блокировки
		LockDescription string `json:"lock_description,omitempty" validate:"omitempty"`
		// Уволен
		IsDismissed bool `json:"is_dismissed" validate:"omitempty"`
		// Причина увольнения
		DismissDescription string `json:"dismiss_description,omitempty" validate:"omitempty"`
		// Водитель самозанятый
		SelfEmployed bool `json:"self_employed,omitempty" validate:"omitempty"`
		// ИНН водителя
		Inn string `json:"inn,omitempty" validate:"omitempty"`
		// СНИЛС водителя
		InsuranceNumber string `json:"insurance_number,omitempty" validate:"omitempty"`
		// Массив параметров водителя. Устарело. Рекомендуется использовать параметр attribute_values
		OrderParams []int `json:"order_params,omitempty" validate:"omitempty"`
		// Фотография водителя
		DriverPhoto string `json:"driver_photo,omitempty" validate:"omitempty,base64"`
		// Массив телефонов водителя
		Phones []Phone `json:"phones,omitempty" validate:"omitempty"`
		// Имя для TaxoPhone
		NameForTaxophone string `json:"name_for_taxophone,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
	}
)

// Обновление информации о водителе
func (cl *Client) UpdateDriverInfo(req UpdateDriverInfoRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100 Автомобиль с ИД=ID не найден
		101 Служба ЕДС с ИД=ID не найдена
		102 Параметр с ИД=ID не найден или не может быть привязан к водителю
		103 Терминальный аккаунт не уникален
		104 Некорректный терминальный аккаунт
		105 Водитель с ИД=ID не найден
		106 Экипаж на линии, запрещено редактирование полей: основной автомобиль, тип работника, служба ЕДС.
		107 Основной телефон может быть только один
		108 Водитель должен иметь основной телефон
		109	Пароль водителя не соответствует политике паролей
	*/
	e := errorMap{
		100: ErrCarNotFound,
		101: ErrUdsNotFound,
		102: ErrParameterNotFoundOrCannotBeBoundDriver,
		103: ErrDriverConflictByTerminalAccount,
		104: ErrTerminalAccountIncorrect,
		105: ErrDriverNotFound,
		106: ErrForbiddenEditCrewOnLine,
		107: ErrConflictByPrimaryPhone,
		108: ErrDriverRequiredPrimaryPhone,
		109: ErrPasswordDoesNotComplyWithPasswordPolicy,
	}

	err = cl.PostJson("update_driver_info", e, req, &response)

	return
}
