package xk6_youzan_nsq

import (
	"errors"
	"fmt"
)

var (
	ErrorState = errors.New("invalid state")
)

func ReportError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
	}
}
