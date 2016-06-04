package log

import (
	. "log"
	"os"
)

var (
	format int = Ldate | Ltime | Lshortfile

	Info  *Logger = New(os.Stdout, "INFO: ", format)
	Error *Logger = New(os.Stderr, "ERROR: ", format)
)
