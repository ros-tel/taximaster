package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	UpdateClientEmployeeInfoRequest struct {
		// ИД редактируемого сотрудника клиента
		ClientEmployeeID int `json:"client_employee_id" validate:"required"`

		// ФИО сотрудника
		Name string `json:"name,omitempty" validate:"omitempty"`
		// Признак удаленного сотрудника
		IsDeleted bool `json:"is_deleted,omitempty" validate:"omitempty"`
		// E-mail
		Email string `json:"email,omitempty" validate:"omitempty,email"`
		// Использовать E-mail для отправки уведомлений по заказу
		UseEmailInforming bool `json:"use_email_informing,omitempty" validate:"omitempty"`
		// Массив телефонов сотрудника
		Phones []Phone `json:"phones,omitempty" validate:"omitempty"`
	}
)

// Обновление информации о сотруднике клиента
func (cl *Client) UpdateClientEmployeeInfo(req UpdateClientEmployeeInfoRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100	Сотрудник с ИД=CLIENT_EMPLOYEE_ID не найден
		101	Клиент с ИД=ID имеет такой же номер телефона=PHONE
		102	Запрещено редактирование телефонов удаленного сотрудника
	*/
	e := errorMap{
		100: ErrCustomerClientNotFound,
		101: ErrClientConflictByPhone,
		102: ErrEditingRemoteEmployeePhoneNumbersIsProhibited,
	}

	err = cl.PostJson("update_client_employee_info", e, req, &response)

	return
}
