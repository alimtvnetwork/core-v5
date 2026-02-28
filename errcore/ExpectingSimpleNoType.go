package errcore

import "fmt"

// ExpectingSimpleNoType
//
// returns
//
//	"%s - Expect [\"%v\"] != [\"%v\"] Left"
func ExpectingSimpleNoType(title, wasExpecting, actual any) string {
	return fmt.Sprintf(
		expectingSimpleNoTypeMessageFormat,
		title,
		wasExpecting,
		actual)
}
