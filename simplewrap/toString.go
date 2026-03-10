package simplewrap

import "github.com/alimtvnetwork/core/internal/convertinternal"

func toString(
	source any,
) string {
	return convertinternal.AnyTo.SmartString(source)
}
