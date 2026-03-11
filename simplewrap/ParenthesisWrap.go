package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func ParenthesisWrap(
	source any,
) string {
	return fmt.Sprintf(
		constants.ParenthesisWrapFormat,
		toString(source))
}
