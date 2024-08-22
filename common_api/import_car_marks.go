package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	ImportCarMarksRequest struct {
		// Импортируемые марки автомобилей
		Marks []string `json:"marks" validate:"required"`
	}
)

// Импорт марок автомобилей в БД
func (cl *Client) ImportCarMarks(req ImportCarMarksRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	err = cl.PostJson("import_car_marks", nil, req, &response)

	return
}
