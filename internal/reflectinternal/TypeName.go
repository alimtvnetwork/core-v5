package reflectinternal

import (
	"reflect"
)

// TypeName
//
// isFullName:
//   - for array pointer it will still output []Type, *typeName
func TypeName(any interface{}) string {
	rfType := reflect.TypeOf(any)

	if rfType == nil {
		return ""
	}

	return rfType.String()
}
