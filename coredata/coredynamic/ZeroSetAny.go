package coredynamic

import (
	"reflect"

	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// ZeroSetAny
//
// Sets empty bytes to the struct or the value but don't make it nil.
//
// It only makes all fields to nil or zero values.
//
// Warning :
//   - Must be set as a pointer any.
func ZeroSetAny(anyItem any) {
	if reflectinternal.Is.Null(anyItem) {
		return
	}

	SafeZeroSet(reflect.ValueOf(anyItem))
}
