package coretests

import "fmt"

func PrintNameValue(header string, any interface{}) {
	toString := ToStringNameValues(any)

	fmt.Printf(printValuesFormat, header, any, toString)
}
