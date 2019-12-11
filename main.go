package main

import (
	"fmt"
	// "github.com/sirupsen/logrus"
	"encoding/json"
	robinhood "gitlab.com/austin-millan/robinhood-openapi/openapi"
)

func main() {
	msg := robinhood.AccountInfo{
		AccountNumber: "12345",
	}
	res, _ := json.Marshal(msg)
	fmt.Println(string(res))
}
