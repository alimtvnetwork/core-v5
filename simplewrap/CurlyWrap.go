package simplewrap

import (
	"fmt"
	
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/internal/convertinternal"
)

func CurlyWrap(
	source any,
) string {
	toStr := convertinternal.
		AnyTo.
		SmartString(source)
	
	return fmt.Sprintf(
		constants.CurlyWrapFormat,
		toStr)
}
