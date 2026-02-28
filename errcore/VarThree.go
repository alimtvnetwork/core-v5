package errcore

import "fmt"

func VarThree(
	isIncludeType bool,
	var1 string,
	val1 any,
	var2 string,
	val2 any,
	var3 string,
	val3 any,
) string {
	if isIncludeType {
		return fmt.Sprintf(
			var3WithTypeFormat,
			var1, val1,
			var2, val2,
			var3, val3,
			val1, val2,
			val3)
	}

	return fmt.Sprintf(
		var3Format,
		var1, var2, var3,
		val1, val2, val3)
}
