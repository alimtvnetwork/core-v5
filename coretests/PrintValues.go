package coretests

import (
	"fmt"
)

func PrintValue(header string, any interface{}) {
	toString := ToStringValues(any)

	fmt.Printf(printValuesFormat, header, any, toString)
}
