package reflectinternal

import "reflect"

// TypeNames
//
// isFullName:
//   - Generates : reflect.TypeOf(item).String()
//   - for array pointer it will still output []Type, *typeName
//
// Or else,
//   - reflect.TypeOf(item).Name()
//   - for array, pointer it will become empty string.
func TypeNames(
	isFullName bool,
	anyItems ...any,
) []string {
	slice := make([]string, len(anyItems))

	if isFullName {
		for i, item := range anyItems {
			slice[i] = reflect.TypeOf(item).String()
		}

		return slice
	}

	for i, item := range anyItems {
		slice[i] = reflect.TypeOf(item).Name()
	}

	return slice
}
