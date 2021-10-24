package main

import (
	"log"

	tmt "github.com/ros-tel/taximaster/tm_tapi"
)

func main() {
	cl := tmt.NewClient("192.168.0.33:8089", "very-secret-tmapi-key")

	res, err := cl.GetInfoByPhone(
		tmt.GetInfoByPhoneRequest{
			Phone:  "89876543210",
			Fields: "PHONE_TYPE-PHONE_SYSTEM_CATEGORY-CATEGORYID-SOURCE_TIMECOUNT",
		},
	)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("%#v", res)
}
