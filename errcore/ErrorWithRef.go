package errcore

import (
	"fmt"
)

func ErrorWithRef(err error, reference any) string {
	if err == nil {
		return ""
	}

	if reference == nil || reference == "" {
		return err.Error()
	}

	return fmt.Sprintf(
		messageMapFormat,
		err.Error(),
		reference)
}
