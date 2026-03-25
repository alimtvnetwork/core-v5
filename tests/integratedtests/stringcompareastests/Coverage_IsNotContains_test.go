package stringcompareastests

import (
	"testing"

	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

func Test_Cov_NotContains_CaseSensitive(t *testing.T) {
	nc := stringcompareas.NotContains
	// Case-sensitive: "Hello World" does NOT contain "hello" (different case)
	// so NotContains returns true (it is indeed not contained).
	result := nc.IsCompareSuccess(false, "Hello World", "hello")
	if !result {
		t.Error("case sensitive: 'hello' not in 'Hello World', NotContains should be true")
	}
}

func Test_Cov_NotContains_CaseInsensitive(t *testing.T) {
	nc := stringcompareas.NotContains
	result := nc.IsCompareSuccess(true, "Hello World", "hello")
	if result {
		t.Error("case insensitive: should contain")
	}
}
