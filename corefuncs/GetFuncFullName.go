package corefuncs

import "gitlab.com/auk-go/core/internal/reflectinternal"

// GetFuncFullName
//
// Get the function name, passing non function may result panic
func GetFuncFullName(i any) string {
	return reflectinternal.GetFuncFullName(i)
}
