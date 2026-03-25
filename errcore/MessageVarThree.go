package errcore

import "fmt"

func MessageVarThree(
	message string,
	firstName string,
	firstValue any,
	secondName string,
	secondValue any,
	thirdName string,
	thirdValue any,
) string {
	return fmt.Sprintf(
		messageVar3Format,
		message,
		firstName,
		secondName,
		firstValue,
		secondValue,
		thirdName,
		thirdValue,
	)
}
