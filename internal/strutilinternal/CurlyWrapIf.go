package strutilinternal

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func CurlyWrapIf(
	isCurly bool,
	source any,
) string {
	isNoCurly := !isCurly

	if isNoCurly {
		return fmt.Sprintf(
			constants.SprintValueFormat,
			source)
	}

	return fmt.Sprintf(
		constants.CurlyWrapFormat,
		source)
}
