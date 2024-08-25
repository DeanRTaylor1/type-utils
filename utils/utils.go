package utils

import (
	"fmt"
	"strings"
)

var DEBUG = false

func Debug(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format, args...)
	}
}

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
