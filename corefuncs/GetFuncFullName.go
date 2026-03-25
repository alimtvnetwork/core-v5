package corefuncs

import "github.com/alimtvnetwork/core/internal/reflectinternal"

// GetFuncFullName
//
// Get the function name, passing non function may result panic
func GetFuncFullName(i any) string {
	return reflectinternal.GetFunc.FullName(i)
}
