package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	err := sendRequest(&http.Client{Timeout: time.Second * 15})
	if err != nil {
		fmt.Print(err.Error())
	}
}
