package common_api

type (
	CreateDriverOperationRequest struct {
		// ИД водителя
		DriverID int `json:"driver_id" validate:"required"`
		// Сумма
		OperSum float64 `json:"oper_sum" validate:"required"`
		// Тип операции:
		// - receipt - приход
		// - expense - расход
		OperType string `json:"oper_type" validate:"required,eq=receipt|eq=expense"`

		// Наименование операции
		Name string `json:"name,omitempty" validate:"omitempty"`
		// Комментарий
		Comment string `json:"comment,omitempty" validate:"omitempty"`
		// ИД типа счета (0 - основной счет), по умолчанию 0
		AccountKind int `json:"account_kind,omitempty" validate:"omitempty,min=0"`
	}

	CreateDriverOperationResponse struct {
		OperID int `json:"oper_id"`
	}
)

// Проведение операции по водителю
func (cl *Client) CreateDriverOperation(req CreateDriverOperationRequest) (CreateDriverOperationResponse, error) {
	var response = CreateDriverOperationResponse{}

	err := validate(req)
	if err != nil {
		return response, err
	}

	err = cl.PostJson("create_driver_operation", req, &response)

	return response, err
}
