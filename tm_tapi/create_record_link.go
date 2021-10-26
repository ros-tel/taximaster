package tm_tapi

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	CreateRecordLinkRequest struct {
		// Дата записи
		RecordDate string `validate:"required,datetime=20060102150405"`
		// Продолжительность записи (в секундах)
		RecordLength int `validate:"required"`
		// Путь к файлу записи
		FilePath string `validate:"required,max=255"`

		// 0 - Исходящий
		// 1 - Входящий
		CallType int `validate:"omitempty,eq=0|eq=1"`
		// Логин пользователя Такси-Мастер
		UserLogin string `validate:"omitempty"`
		// ИД звонка (необязателен, если указан PHONE)
		CallID string `validate:"omitempty,max=60"`
		// 	Номер телефона (необязателен, если указан CALL_ID)
		Phone string `validate:"omitempty,max=16"`
	}

	CreateRecordLinkResponse struct {
		// ИД созданной записи
		RecordID int `xml:"RECORD_ID"`
	}
)

// Запись пути к файлу разговора в базу данных
func (cl *Client) CreateRecordLink(req CreateRecordLinkRequest) (CreateRecordLinkResponse, error) {
	var response = CreateRecordLinkResponse{}

	err := validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	if req.UserLogin != "" {
		v.Add("USER_LOGIN", req.UserLogin)
	}
	v.Add("RECORD_DATE", req.RecordDate)
	v.Add("RECORD_LENGTH", strconv.Itoa(req.RecordLength))
	if req.CallID != "" {
		v.Add("CALL_ID", req.CallID)
	}
	if req.Phone != "" {
		v.Add("PHONE", req.Phone)
	}
	v.Add("FILE_PATH", req.FilePath)
	v.Add("CALL_TYPE", strconv.Itoa(req.CallType))

	err = cl.Post("create_record_link", v, &response)

	return response, err
}
