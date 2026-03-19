package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/issetter"
)

func Test_Cov_IsSetter_NamesNotContaining(t *testing.T) {
	// This exercises the toHashset path via issetter's public API
	s := issetter.Names("a", "b", "c")
	if !s.Has("a") {
		t.Error("expected has a")
	}
	if s.Has("z") {
		t.Error("expected not has z")
	}
}
