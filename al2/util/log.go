package util

import (
	"fmt"
	"io"
	"log"
)

func CreateLogger(out io.Writer, caller interface{}) *log.Logger {
	return log.New(
		out,
		fmt.Sprintf("al2\t%T\t", caller),
		log.LstdFlags|log.LUTC)
}
