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
		NeedPhoto bool `validate:"omitempty"`
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
		// Удостоверение водителя
		License string `json:"license"`
		// Любой неосновной телефон водителя (устаревшее поле)
		HomePhone string `json:"home_phone"`
		// Основной телефон водителя (устаревшее поле)
		MobilePhone string `json:"mobile_phone"`
		// Водитель заблокирован
		IsLocked bool `json:"is_locked"`
		// Водитель уволен
		IsDismissed bool `json:"is_dismissed"`
		// Фото водителя (только если need_photo = true)
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
	}
)

// Запрос информации о водителе
func (cl *Client) GetDriverInfo(req GetDriverInfoRequest) (GetDriverInfoResponse, error) {
	var response = GetDriverInfoResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("driver_id", strconv.Itoa(req.DriverID))
	if req.NeedPhoto {
		v.Add("need_photo", "true")
	}

	err = cl.Get("get_driver_info", v, &response)

	return response, err
}
