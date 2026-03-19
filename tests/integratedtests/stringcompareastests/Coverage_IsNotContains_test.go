package stringcompareastests

import (
	"testing"

	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

func Test_Cov_NotContains_CaseSensitive(t *testing.T) {
	nc := stringcompareas.NotContains
	result := nc.IsCompareSuccess(false, "Hello World", "hello")
	if result {
		t.Error("case sensitive: should contain lowercase match")
	}
}

func Test_Cov_NotContains_CaseInsensitive(t *testing.T) {
	nc := stringcompareas.NotContains
	result := nc.IsCompareSuccess(true, "Hello World", "hello")
	if result {
		t.Error("case insensitive: should contain")
	}
}
