package errcore

import "fmt"

func ExpectingNotEqualSimpleNoType(
	title,
	wasExpecting,
	actual any,
) string {
	return fmt.Sprintf(
		expectingNotMatchingSimpleNoTypeMessageFormat,
		title,
		wasExpecting,
		actual)
}
