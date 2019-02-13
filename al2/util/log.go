package util

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CreateLogger(out io.Writer, caller interface{}) *log.Logger {
	if out == nil {
		out = os.Stderr
	}
	return log.New(
		out,
		fmt.Sprintf("al2\t%T\t", caller),
		log.LstdFlags|log.LUTC)
}
