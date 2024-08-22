package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	UpdateClientInfo2Request struct {
		// ИД клиента
		ClientID int `json:"client_id" validate:"required"`

		// ФИО
		Name string `json:"name,omitempty" validate:"omitempty,max=60"`
		// Логин
		Login string `json:"login,omitempty" validate:"omitempty,max=60"`
		// Пароль
		Password string `json:"password,omitempty" validate:"omitempty,max=60"`
		// Массив телефонов клиента
		Phones []Phone `json:"phones,omitempty" validate:"omitempty"`
		// ИД группы клиента
		ClientGroupID int `json:"client_group_id,omitempty" validate:"omitempty"`
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
		AttributeValues []AttributeValue `json:"attribute_values,omitempty" validate:"omitempty"`
	}
)

// Изменение информации по клиенту 2
func (cl *Client) UpdateClientInfo2(req UpdateClientInfo2Request) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100 Клиент не найден
		101 Дублирование номера телефона=PHONE в списке
		102 Клиент с логином=LOGIN уже существует
		103 Группа клиента с ИД=CLIENT_GROUP_ID не найдена
		104 Клиент указанный в качестве родителя с ИД=PARENT_ID не найден
		105 Основной телефон может быть только один
		106 Клиент должен иметь основной телефон
		107 Атрибут с ИД=ID не найден
		108 Атрибут с ИД=ID не может быть привязан к клиенту
		109 Пароль клиента не соответствует политике паролей
	*/
	e := errorMap{
		100: ErrClientNotFound,
		101: ErrDuplicatePhoneNumberInTheList,
		102: ErrClientExistsWithLogin,
		103: ErrClientGroupNotFound,
		104: ErrParentClientNotFound,
		105: ErrConflictByPrimaryPhone,
		106: ErrClientRequiredPrimaryPhone,
		107: ErrAttributeNotFound,
		108: ErrAttributeCannotBeBoundClient,
		109: ErrPasswordDoesNotComplyWithPasswordPolicy,
	}

	err = cl.PostJson("update_client_info2", e, req, &response)

	return
}
