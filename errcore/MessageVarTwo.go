package errcore

import "fmt"

func MessageVarTwo(
	message string,
	var1 string,
	val1 any,
	var2 string,
	val2 any,
) string {
	return fmt.Sprintf(
		messageVar2Format,
		message,
		var1,
		var2,
		val1,
		val2)
}
