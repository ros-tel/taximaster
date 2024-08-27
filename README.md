## Установка

```sh
$ go get -u github.com/ros-tel/taximaster
```

## Использование в вашем коде

1. TM API

```go
import "github.com/ros-tel/taximaster/common_api"
```

## Примеры

1. Анализ номера телефона

```go
package main

import (
	"log"

	tm "github.com/ros-tel/taximaster/common_api"
)

func main() {
	cl := tm.NewClient("192.168.0.33:8089", "very-secret-api-key", nil)

	res, err := cl.AnalyzePhone(
		tm.AnalyzePhoneRequest{
			Phone:                 "89876543210",
			SearchInDriversMobile: true,
		},
	)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("%#v", res)
}
```

2. Терминальная операция

```go
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
```

3. Запрос информации по номеру телефона

```go
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
```