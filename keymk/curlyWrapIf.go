package keymk

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func curlyWrapIf(
	isCurly bool,
	source any,
) string {
	if !isCurly {
		return fmt.Sprintf(
			constants.SprintValueFormat,
			source)
	}

	return fmt.Sprintf(
		constants.CurlyWrapFormat,
		source)
}
