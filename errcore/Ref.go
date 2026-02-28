package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func Ref(reference any) string {
	if reference == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.ReferenceWrapFormat,
		reference)

}
