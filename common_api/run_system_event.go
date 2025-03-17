package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	RunSystemEventRequest struct {
		// ИД системного события с типом "По запросу CommonAPI".
		// Если не передан, то system_event_id будет определен исходя из переданного кода системного события.
		// Обязательно должен передаваться system_event_id или system_event_code
		SystemEventID int `json:"system_event_id,omitempty" validate:"omitempty"`
		// Код системного события с типом "По запросу CommonAPI". Обязательно должен передаваться system_event_id или system_event_code
		SystemEventCode string `json:"system_event_code,omitempty" validate:"omitempty"`

		// В запросе можно передавать дополнительные произвольные параметры с любыми названиями и значениями.
		// Эти параметры могут использоваться в системном событии.
		// Зарезервированный параметр "json_data": может быть любого типа, в т.ч. объект или массив,
		// внутри которого могут быть в т.ч. вложенные объекты и массивы.
		JsonData any `json:"json_data,omitempty" validate:"omitempty"`
		// Признак необходимости ожидать завершения действий системного события.
		WaitForCompletion *bool `json:"wait_for_completion,omitempty" validate:"omitempty"`
	}
)

// Вызвать системное событие
func (cl *Client) RunSystemEvent(req any) (response EmptyResponse, err error) {
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
