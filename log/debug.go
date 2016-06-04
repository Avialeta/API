package log

import (
	. "log"
	"os"
)

var Debug *Logger = New(os.Stdout, "DEBUG: ", format)
