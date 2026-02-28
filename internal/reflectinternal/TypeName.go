package reflectinternal

import (
	"reflect"
)

// TypeName
//
// isFullName:
//   - for array pointer it will still output []Type, *typeName
func TypeName(anyItem any) string {
	rfType := reflect.TypeOf(anyItem)

	if rfType == nil {
		return ""
	}

	return rfType.String()
}
