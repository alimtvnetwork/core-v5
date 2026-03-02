package strutilinternal

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
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
