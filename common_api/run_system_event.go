package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	RunSystemEventRequest struct {
		// ИД системного события с типом "По запросу CommonAPI"
		SystemEventID int `json:"system_event_id" validate:"required"`
	}
)

func (cl *Client) RunSystemEvent(req interface{}) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100 У системного события тип не "По запросу CommonAPI"
		101 Системное событие не найдено
		102 Системное событие не активно
	*/
	e := errorMap{
		100: ErrSystemEventBadType,
		101: ErrSystemEventNotFound,
		102: ErrSystemEventNotActive,
	}

	err = cl.PostJson("run_system_event", e, req, &response)

	return
}
