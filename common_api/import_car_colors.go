package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	ImportCarColorsRequest struct {
		// Импортируемые цвета автомобилей
		Colors []string `json:"colors" validate:"required"`
	}
)

// Импорт цветов автомобилей в БД
func (cl *Client) ImportCarColors(req ImportCarColorsRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	err = cl.PostJson("import_car_colors", nil, req, &response)

	return
}
