package simplewrap

import (
	"fmt"
	
	"github.com/alimtvnetwork/core/constants"
)

func SquareWrap(
	source any,
) string {
	return fmt.Sprintf(
		constants.SquareWrapFormat,
		toString(source))
}
