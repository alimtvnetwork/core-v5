package errcore

import "fmt"

func MessageVarTwo(
	message string,
	firstName string,
	firstValue any,
	secondName string,
	secondValue any,
) string {
	return fmt.Sprintf(
		messageVar2Format,
		message,
		firstName,
		secondName,
		firstValue,
		secondValue)
}
