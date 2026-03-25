package args

import "github.com/alimtvnetwork/core/internal/convertinternal"

func toString(i any) string {
	return convertinternal.AnyTo.SmartString(i)
}
