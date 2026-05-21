package logger

import (
	"fmt"

	"github.com/SHAIK14/pulsemon/checker"
)

func Log(res checker.Result) {

	if !res.IsUp {
		fmt.Printf("[DOWN] %s  %s\n", res.Name, res.Error)
	} else {
		fmt.Printf("[UP] %s  %d  %s\n", res.Name, res.StatusCode, res.ResponseTime)
	}

}
