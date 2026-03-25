package isany

import "github.com/alimtvnetwork/core/internal/reflectinternal"

// Zero
//
//	returns true if the current value is null
//	or reflect value is zero
//
// Reference:
//   - Stackoverflow Example : https://stackoverflow.com/a/23555352
func Zero(anyItem any) bool {
	return reflectinternal.Is.Zero(anyItem)
}
