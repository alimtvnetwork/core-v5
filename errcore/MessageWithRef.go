package errcore

import "fmt"

func MessageWithRef(msg string, reference any) string {
	return fmt.Sprintf(
		messageMapFormat,
		msg,
		reference)
}
