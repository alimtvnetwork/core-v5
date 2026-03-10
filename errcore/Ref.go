package errcore

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func Ref(reference any) string {
	if reference == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.ReferenceWrapFormat,
		reference)

}
