package errcore

import "fmt"

func ExpectingSimpleNoTypeError(title, wasExpecting, actual any) error {
	return fmt.Errorf(
		expectingSimpleNoTypeMessageFormat,
		title,
		wasExpecting,
		actual)
}
