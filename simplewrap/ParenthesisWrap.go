package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

func ParenthesisWrap(
	source any,
) string {
	return fmt.Sprintf(
		constants.ParenthesisWrapFormat,
		toString(source))
}
