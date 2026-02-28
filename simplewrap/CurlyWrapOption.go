package simplewrap

import (
	"gitlab.com/auk-go/core/internal/convertinternal"
)

func CurlyWrapOption(
	isSkipIfExists bool,
	source any,
) string {
	toStr := convertinternal.
		AnyTo.
		SmartString(source)
	
	if isSkipIfExists {
		return ConditionalWrapWith(
			'{',
			toStr,
			'}')
	}
	
	return CurlyWrap(source)
}
