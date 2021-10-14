package errcore

import "fmt"

func ExpectingSimpleNoTypeError(title, wasExpecting, actual interface{}) error {
	return fmt.Errorf(
		expectingSimpleNoTypeMessageFormat,
		title,
		wasExpecting,
		actual)
}
