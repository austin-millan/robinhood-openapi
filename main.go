package main

import (
	"encoding/json"
	"fmt"

	robinhood "gitlab.com/trade-hawk/robinhood-openapi/openapi"
)

func main() {
	msg := robinhood.AccountInfo{
		AccountNumber: "12345",
	}
	res, _ := json.Marshal(msg)
	fmt.Println(string(res))
}
