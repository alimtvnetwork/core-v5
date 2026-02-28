package simplewrap

import "gitlab.com/auk-go/core/internal/convertinternal"

func toString(
	source any,
) string {
	return convertinternal.AnyTo.SmartString(source)
}
