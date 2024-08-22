package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetDriverInfoRequest struct {
		// ИД водителя
		DriverID int `validate:"required"`

		// Нужна ли фотография водителя
		NeedPhoto *bool `validate:"omitempty"`
		// Список возвращаемых полей через запятую
		Fields string `validate:"omitempty"`
	}

	GetDriverInfoResponse struct {
		// ИД водителя
		DriverID int `json:"driver_id"`
		// ФИО водителя
		Name string `json:"name"`
		// Баланс основного счета водителя
		Balance float64 `json:"balance"`
		// День рождения водителя
		Birthday string `json:"birthday"`
		// ИД основного автомобиля водителя
		CarID int `json:"car_id"`
		// Водительское удостоверение
		DriverLicense string `json:"driver_license"`
		// Разрешение на перевозку
		License string `json:"license"`
		// Любой неосновной телефон водителя (устаревшее поле)
		HomePhone string `json:"home_phone"`
		// Основной телефон водителя (устаревшее поле)
		MobilePhone string `json:"mobile_phone"`
		// Водитель заблокирован
		IsLocked bool `json:"is_locked"`
		// Водитель уволен
		IsDismissed bool `json:"is_dismissed"`
		// Водитель самозанятый
		SelfEmployed bool `json:"self_employed"`
		// ИНН водителя
		Inn string `json:"inn"`
		// СНИЛС водителя
		InsuranceNumber string `json:"insurance_number"`
		// Фото водителя (только если need_photo = true или поле driver_photo указано в списке фильтра полей fields)
		DriverPhoto string `json:"driver_photo"`
		// Массив параметров водителя
		OrderParams []int `json:"order_params"`
		// Массив телефонов водителя
		Phones []Phone `json:"phones"`
		// Терминальный аккаунт
		TermAccount string `json:"term_account"`
		// Имя для TaxoPhone
		NameForTaxophone string `json:"name_for_taxophone"`
		// Массив балансов счетов
		Accounts []Account `json:"accounts"`
		// Массив значений атрибутов
		AttributeValues []AttributeValue `json:"attribute_values"`
		// Показатель эффективности (KPI) водителя
		Kpi float64 `json:"kpi"`
	}
)

// Запрос информации о водителе
func (cl *Client) GetDriverInfo(req GetDriverInfoRequest) (response GetDriverInfoResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("driver_id", strconv.Itoa(req.DriverID))
	if req.NeedPhoto != nil {
		v.Add("need_photo", strconv.FormatBool(*req.NeedPhoto))
	}
	if req.Fields != "" {
		v.Add("fields", req.Fields)
	}

	/*
		100 Водитель не найден
	*/
	e := errorMap{
		100: ErrDriverNotFound,
	}

	err = cl.Get("get_driver_info", e, v, &response)

	return
}
