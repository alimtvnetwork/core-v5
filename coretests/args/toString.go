package args

import "gitlab.com/auk-go/core/internal/convertinternal"

func toString(i any) string {
	return convertinternal.AnyTo.SmartString(i)
}
