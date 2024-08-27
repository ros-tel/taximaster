package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetCrewTrackRequest struct {
		// ИД экипажа
		CrewID int `validate:"required"`
		//Начало периода
		StartTime string `validate:"required,datetime=20060102150405"`
		// Конец периода, должен отличаться о начала периода не более чем на 7 дней
		FinishTime string `validate:"required,datetime=20060102150405"`
	}

	GetCrewTrackResponse struct {
		// Массив, трек экипажа за период времени
		Track []struct {
			// Широта точки маршрута
			Lat float64 `json:"lat"`
			// Долгота точки маршрута
			Lon float64 `json:"lon"`
			// Время данной точки
			Time string `json:"time"`
			// Скорость в данный момент, км/ч
			Speed int `json:"speed"`
			// Направление движения, градусы (0 - север, 90 - восток, 180 - юг, 270 - запад)
			Direction int `json:"direction"`
		} `json:"track"`
	}
)

// Запрос трека экипажа
func (cl *Client) GetCrewTrack(req GetCrewTrackRequest) (response GetCrewTrackResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("crew_id", strconv.Itoa(req.CrewID))
	v.Add("start_time", req.StartTime)
	v.Add("finish_time", req.FinishTime)

	/*
		100	Не найден экипаж ИД=crew_id
		101	Задан период времени более 7 дней
	*/
	e := errorMap{
		100: ErrCrewNotFound,
		101: ErrTimePeriodMore7Days,
	}

	err = cl.Get("get_crew_track", e, v, &response)

	return
}
