package errcore

import (
	"testing"
)

// ── Unexported field access — must remain in source package ──

func Test_Cov10_ExpectationMessageDef_ExpectedSafeString_Cached(t *testing.T) {
	cached := "pre-cached"
	def := ExpectationMessageDef{
		Expected:       "test",
		expectedString: &cached,
	}
	result := def.ExpectedSafeString()
	if result != "pre-cached" {
		t.Errorf("got %q, want %q", result, "pre-cached")
	}
}
