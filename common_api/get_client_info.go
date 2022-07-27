package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetClientInfoRequest struct {
		// ИД клиента
		ClientID int `validate:"required"`
	}

	GetClientInfoResponse struct {
		// ИД клиента
		ClientID int `json:"client_id"`
		// ИД клиента-родителя
		ParentID int `json:"parent_id"`
		// ФИО
		Name string `json:"name"`
		// Номер договора
		Number string `json:"number"`
		// Домашний адрес
		Address string `json:"address"`
		// Пол. Может принимать значения:
		// - "" - не указан
		// - male - мужской
		// - female - женский
		Gender string `json:"gender"`
		// Дата рождения
		Birthday string `json:"birthday"`
		// Массив телефонов клиента
		Phones []string `json:"phones"`
		// Баланс
		Balance float64 `json:"balance"`
		// Бонусный баланс
		BonusBalance int `json:"bonus_balance"`
		// Логин
		Login string `json:"login"`
		// Пароль
		Password string `json:"password"`
		// ИД группы клиента
		ClientGroupID int `json:"client_group_id"`
		// ИД тарифа клиента или группы клиентов
		TariffID int `json:"tariff_id"`
		// ИД призового тарифа клиента или группы клиентов
		PrizeTariffID int `json:"prize_tariff_id"`
		// ИД смены тарифов клиента или группы клиентов
		TariffShiftID int `json:"tariff_shift_id"`
		// ИД скидки клиента-сотрудника или группы клиентов
		DiscountID int `json:"discount_id"`
		// ИД призовой скидки клиента-сотрудника или группы клиентов
		PrizeDiscountID int `json:"prize_discount_id"`
		// Порог, ниже которого не может опускаться баланс клиента-сотрудника
		MinBalance float64 `json:"min_balance"`
		// Минимальный баланс для использования безналичного счета клиента-сотрудника
		MinBalanceForUseCashless float64 `json:"min_balance_for_use_cashless"`
		// Минимальный баланс для использования бонусного счета клиента-сотрудника
		MinBonusBalanceForUseBonus float64 `json:"min_bonus_balance_for_use_bonus"`
		// Клиент-сотрудник заблокирован
		IsLocked bool `json:"is_locked"`
		// Причина блокировка клиента-сотрудника
		LockDescription string `json:"lock_description"`
		// Признак использования безналичного расчета по умолчанию для клиента-сотрудника
		UseCashless bool `json:"use_cashless"`
		// Сколько заказов осталось до призового
		RemainPrize int `json:"remain_prize"`
		// E-mail клиента-сотрудника
		Email string `json:"email"`
		// Использовать E-mail для отправки уведомлений по заказу
		UseEmailInforming bool `json:"use_email_informing"`
		// Группа экипажей по умолчанию
		DefaultCrewGroup int `json:"default_crew_group"`
		// Комментарий
		Comment   string `json:"comment"`
		Employees []struct {
			// ИД клиента-сотрудника
			ClientID int `json:"client_id"`
			// ИД родителя клиента-сотрудника
			ParentID int `json:"parent_id"`
			// ФИО клиента-сотрудника
			Name string `json:"name"`
			// Номер договора клиента-сотрудника
			Number string `json:"number"`
			// Домашний адрес клиента-сотрудника
			Address string `json:"address"`
			// Пол клиента-сотрудника. Может принимать значения:
			// - "" - не указан
			// - male - мужской
			// - female - женский
			Gender string `json:"gender"`
			// Дата рождения клиента-сотрудника
			Birthday string `json:"birthday"`
			// Массив телефонов клиента-сотрудника
			Phones []string `json:"phones"`
			// Баланс клиента-сотрудника
			Balance float64 `json:"balance"`
			// Бонусный баланс клиента-сотрудника
			BonusBalance int `json:"bonus_balance"`
			// Логин клиента-сотрудника
			Login string `json:"login"`
			// Пароль клиента-сотрудника
			Password string `json:"password"`
			// ИД группы клиента клиента-сотрудника
			ClientGroupID int `json:"client_group_id"`
			// ИД тарифа клиента или группы клиентов клиента-сотрудника
			TariffID int `json:"tariff_id"`
			// ИД призового тарифа клиента-сотрудника или группы клиентов
			PrizeTariffID int `json:"prize_tariff_id"`
			// ИД смены тарифов клиента-сотрудника или группы клиентов
			TariffShiftID int `json:"tariff_shift_id"`
			// ИД скидки клиента-сотрудника или группы клиентов
			DiscountID int `json:"discount_id"`
			// ИД призовой скидки клиента-сотрудника или группы клиентов
			PrizeDiscountID int `json:"prize_discount_id"`
			// Порог, ниже которого не может опускаться баланс клиента-сотрудника
			MinBalance float64 `json:"min_balance"`
			// Минимальный баланс для использования безналичного счета клиента-сотрудника
			MinBalanceForUseCashless float64 `json:"min_balance_for_use_cashless"`
			// Минимальный баланс для использования бонусного счета клиента-сотрудника
			MinBonusBalanceForUseBonus float64 `json:"min_bonus_balance_for_use_bonus"`
			// Клиент-сотрудник заблокирован
			IsLocked bool `json:"is_locked"`
			// Причина блокировка клиента-сотрудника
			LockDescription string `json:"lock_description"`
			// Признак использования безналичного расчета по умолчанию для клиента-сотрудника
			UseCashless bool `json:"use_cashless"`
			// Сколько заказов осталось до призового
			RemainPrize int `json:"remain_prize"`
			// E-mail клиента-сотрудника
			Email string `json:"email"`
			// Использовать E-mail для отправки уведомлений по заказу
			UseEmailInforming bool `json:"use_email_informing"`
			// Группа экипажей по умолчанию
			DefaultCrewGroup int `json:"default_crew_group"`
			// Комментарий
			Comment string `json:"comment"`
		} `json:"employees"`
		// Массив счетов клиента
		Accounts []Account `json:"accounts"`
	}
)

// Запрос информации по клиенту
func (cl *Client) GetClientInfo(req GetClientInfoRequest) (GetClientInfoResponse, error) {
	var response = GetClientInfoResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	v.Add("client_id", strconv.Itoa(req.ClientID))

	/*
		100 Не найден клиент
	*/
	e := errorMap{
		100: ErrClientNotFound,
	}

	err = cl.Get("get_client_info", e, v, &response)

	return response, err
}
