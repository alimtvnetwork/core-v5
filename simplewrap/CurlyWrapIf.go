package simplewrap

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
		return toString(source)
	}
	
	return fmt.Sprintf(
		constants.CurlyWrapFormat,
		toString(source))
}
