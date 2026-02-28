package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/convertinternal"
)

// WithCurly
//
// {%v}
func WithCurly(
	source interface{},
) string {
	toStr := convertinternal.
		AnyTo.
		SmartString(source)
	
	return fmt.Sprintf(
		constants.CurlyWrapFormat,
		toStr)
}
