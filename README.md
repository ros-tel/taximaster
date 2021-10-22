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
	cl := tm.NewClient("192.168.0.33:8089", "very-secret-api-key")

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
