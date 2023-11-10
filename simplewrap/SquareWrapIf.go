package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

func SquareWrapIf(
	isSquareWrap bool,
	source interface{},
) string {
	if !isSquareWrap {
		return toString(source)
	}
	
	return fmt.Sprintf(
		constants.SquareWrapFormat,
		toString(source))
}
