package utils

import (
	"fmt"
)

func SprintError(code, message, extra string, origErr error) string {
	msg := fmt.Sprintf("%s: %s", code, message)
	if extra != "" {
		msg = fmt.Sprintf("%s\n\t%s", msg, extra)
	}
	if origErr != nil {
		msg = fmt.Sprintf("%s: %s\ncaused by: %s", code, msg, origErr.Error())
	}
	return msg
}