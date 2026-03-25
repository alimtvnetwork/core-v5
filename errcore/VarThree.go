package errcore

import "fmt"

func VarThree(
	isIncludeType bool,
	firstName string,
	firstValue any,
	secondName string,
	secondValue any,
	thirdName string,
	thirdValue any,
) string {
	if isIncludeType {
		return fmt.Sprintf(
			var3WithTypeFormat,
			firstName, firstValue,
			secondName, secondValue,
			thirdName, thirdValue,
			firstValue, secondValue,
			thirdValue)
	}

	return fmt.Sprintf(
		var3Format,
		firstName, secondName, thirdName,
		firstValue, secondValue, thirdValue)
}
