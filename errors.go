package xk6_youzan_nsq

import (
	"fmt"
)

func ReportError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
	}
}
