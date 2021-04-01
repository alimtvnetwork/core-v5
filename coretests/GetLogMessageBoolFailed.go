package coretests

import (
	"fmt"
)

func LogOnFail(
	isPass bool,
	expected, actual interface{},
) {
	if isPass {
		return
	}

	logMessage := fmt.Sprintf(logFormat, expected, actual)
	fmt.Println(logMessage)
}
