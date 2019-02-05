package util

import (
	"fmt"
	"log"
	"os"
)

func CreateLogger(caller interface{}) *log.Logger {
	return log.New(
		os.Stderr,
		fmt.Sprintf("admin-mate-cli\t%T\t", caller),
		log.LstdFlags|log.LUTC)
}
