package errcore

import (
	"fmt"
)

func VarTwo(
	isIncludeType bool,
	var1 string,
	val1 any,
	var2 string,
	val2 any,
) string {
	if isIncludeType {
		return fmt.Sprintf(
			var2WithTypeFormat,
			var1,
			val1,
			var2,
			val2,
			val1,
			val2)
	}

	return fmt.Sprintf(
		var2Format,
		var1,
		var2,
		val1,
		val2)
}
