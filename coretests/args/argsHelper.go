package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// getByIndex safely retrieves an item from a slice by its index.
// Returns nil if the index is out of bounds.
func getByIndex(slice []any, index int) any {
	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

// buildToString builds a formatted string representation for an arg type,
// using the given type name and cached SimpleStringOnce for memoization.
//
// Example output: "Three { first-val, second-val, third-val }"
func buildToString(
	typeName string,
	slice []any,
	cache *corestr.SimpleStringOnce,
) string {
	if cache.IsInitialized() {
		return cache.String()
	}

	var items []string

	for _, item := range slice {
		items = append(items, toString(item))
	}

	result := fmt.Sprintf(
		selfToStringFmt,
		typeName,
		strings.Join(items, constants.CommaSpace),
	)

	return cache.GetSetOnce(result)
}

// appendIfDefined appends the given value to the slice only if it is
// non-nil and non-zero as determined by reflectinternal.Is.Defined.
func appendIfDefined(args []any, value any) []any {
	if reflectinternal.Is.Defined(value) {
		return append(args, value)
	}

	return args
}

// invokeMustHelper invokes the given FuncWrapAny with args, panicking on error.
// This eliminates duplicate InvokeMust patterns across all Func arg types.
func invokeMustHelper(fw *FuncWrapAny, args ...any) []any {
	results, err := fw.Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}
