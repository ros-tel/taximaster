package main

import (
	"log"

	tm "github.com/ros-tel/taximaster/common_api"
	pay "github.com/ros-tel/taximaster/pay_term_api"
)

func main() {
	tmcl := tm.NewClient("192.168.0.33:8089", "very-secret-api-key", nil)
	paycl := pay.NewClient("192.168.0.33:8089", "very-secret-pay-api-key")

	// Получение терминального аккаунта водителя
	res1, err := tmcl.GetDriverInfo(
		tm.GetDriverInfoRequest{
			DriverID: 2258,
		},
	)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("%#v", res1)

	// Проведение зачисления на счет
	res2, err := paycl.AddTermOperation(
		pay.AddTermOperationRequest{
			PaySystemType: 8,
			CityID:        "02422",
			TermAccount:   res1.TermAccount,
			Sum:           0.1,
			OperID:        "203152597448",
			OperTime:      "20211023010101",
		},
	)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("%#v", res2)
}
