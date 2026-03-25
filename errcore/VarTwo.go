package errcore

import (
	"fmt"
)

func VarTwo(
	isIncludeType bool,
	firstName string,
	firstValue any,
	secondName string,
	secondValue any,
) string {
	if isIncludeType {
		return fmt.Sprintf(
			var2WithTypeFormat,
			firstName,
			firstValue,
			secondName,
			secondValue,
			firstValue,
			secondValue)
	}

	return fmt.Sprintf(
		var2Format,
		firstName,
		secondName,
		firstValue,
		secondValue)
}
