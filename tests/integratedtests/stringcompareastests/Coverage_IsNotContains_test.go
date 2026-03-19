package stringcompareastests

import (
	"testing"

	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

func Test_Cov_NotContains_CaseSensitive(t *testing.T) {
	nc := stringcompareas.NotContains
	result := nc.Compare("Hello World", "hello", true)
	if result {
		t.Error("case sensitive: should not contain lowercase")
	}
}

func Test_Cov_NotContains_CaseInsensitive(t *testing.T) {
	nc := stringcompareas.NotContains
	result := nc.Compare("Hello World", "hello", false)
	if !result {
		t.Error("case insensitive: should contain")
	}
}
