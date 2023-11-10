package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

func ParenthesisWrapIf(
	isSquareWrap bool,
	source interface{},
) string {
	if !isSquareWrap {
		return toString(source)
	}
	
	return fmt.Sprintf(
		constants.ParenthesisWrapFormat,
		toString(source))
}
