package coretests

import (
	"fmt"

	"gitlab.com/evatix-go/core/internal/msgformats"
)

func PrintValue(header string, any interface{}) {
	toString := ToStringValues(any)

	fmt.Printf(msgformats.PrintValuesFormat, header, any, toString)
}
