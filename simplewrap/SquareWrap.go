package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

func SquareWrap(
	source any,
) string {
	return fmt.Sprintf(
		constants.SquareWrapFormat,
		toString(source))
}
