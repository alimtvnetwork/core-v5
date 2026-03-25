package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/issetter"
)

func Test_Cov_IsSetter_OnlySupportedErr_ExercisesToHashset(t *testing.T) {
	// OnlySupportedErr internally calls toHashset
	v := issetter.True
	err := v.OnlySupportedErr("True", "False")
	if err == nil {
		t.Error("expected error for unsupported names")
	}
}

func Test_Cov_IsSetter_OnlySupportedErr_AllSupported(t *testing.T) {
	v := issetter.True
	err := v.OnlySupportedErr("Uninitialized", "True", "False", "Unset", "Set", "Wildcard")
	if err != nil {
		t.Errorf("expected nil got %v", err)
	}
}

func Test_Cov_IsSetter_OnlySupportedErr_Empty(t *testing.T) {
	v := issetter.True
	err := v.OnlySupportedErr()
	if err != nil {
		t.Errorf("expected nil got %v", err)
	}
}
