package common_api

import (
	"github.com/ros-tel/taximaster/validator"
)

type (
	RegisterClient2Request struct {
		// ФИО
		Name string `json:"name" validate:"required,max=60"`
		// Логин
		Login string `json:"login" validate:"required,max=60"`
		// Пароль
		Password string `json:"password" validate:"required,max=60"`
		// Номера телефонов (через запятую)
		Phones []Phone `json:"phones" validate:"required"`

		// ИД группы клиента
		ClientGroup int `json:"client_group_id,omitempty" validate:"omitempty"`
		// ИД клиента-родителя
		ParentID int `json:"parent_id,omitempty" validate:"omitempty"`
		// Домашний адрес
		Address string `json:"address,omitempty" validate:"omitempty"`
		// Дата рождения
		Birthday string `json:"birthday,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Пол. Может принимать значения:
		// - male - мужской
		// - female - женский
		Gender string `json:"gender,omitempty" validate:"omitempty,eq=male|eq=female"`
		// E-mail
		Email string `json:"email,omitempty" validate:"omitempty,email"`
		// Использовать E-mail для отправки уведомлений по заказу
		UseEmailInforming bool `json:"use_email_informing,omitempty" validate:"omitempty"`
		// Комментарий
		Comment string `json:"comment,omitempty" validate:"omitempty"`
		// Использовать собственный счет для оплаты заказов
		UseOwnAccount bool `json:"use_own_account,omitempty" validate:"omitempty"`
		// Массив значений атрибутов
		AttributeValue []AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
	}

	RegisterClient2Response struct {
		ClientID int `json:"client_id"`
	}
)

// Регистрация клиента 2
func (cl *Client) RegisterClient2(req RegisterClient2Request) (response RegisterClient2Response, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100	Дублирование номера телефона=PHONE в списке
		101	Клиент с логином=LOGIN уже существует
		102	Группа клиента с ИД=CLIENT_GROUP не найдена
		103	Клиент указанный в качестве родителя с ИД=PARENT_ID не найден
		105	Основной телефон может быть только один
		106	Клиент должен иметь основной телефон
		107	Атрибут с ИД=ID не найден
		108	Атрибут с ИД=ID не может быть привязан к клиенту
		109	Пароль клиента не соответствует политике паролей
	*/
	e := errorMap{
		100: ErrClientConflictByPhone,
		101: ErrClientExistsWithLogin,
		102: ErrClientGroupNotFound,
		103: ErrParentClientNotFound,
		105: ErrConflictByPrimaryPhone,
		106: ErrClientRequiredPrimaryPhone,
		107: ErrAttributeNotFound,
		108: ErrAttributeCannotBeBoundClient,
		109: ErrPasswordDoesNotComplyWithPasswordPolicy,
	}

	err = cl.PostJson("register_client2", e, req, &response)

	return
}
