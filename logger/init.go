package logger

import (
	"fmt"
)

var log Logger

func Get() Logger {
	var err error
	if log == nil {
		log, err = NewZaplogger()
		if err != nil {
			fmt.Println("failed to initialize logger, got error ", err)
		}
	}
	return log
}
