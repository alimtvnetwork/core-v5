package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/corevalidator"
)

// ==========================================
// Parameter.IsIgnoreCase
// ==========================================

func Test_Parameter_IsIgnoreCase_WhenCaseSensitive(t *testing.T) {
	p := corevalidator.Parameter{IsCaseSensitive: true}
	if p.IsIgnoreCase() {
		t.Error("case-sensitive should not ignore case")
	}
}

func Test_Parameter_IsIgnoreCase_WhenNotCaseSensitive(t *testing.T) {
	p := corevalidator.Parameter{IsCaseSensitive: false}
	if !p.IsIgnoreCase() {
		t.Error("not case-sensitive should ignore case")
	}
}

func Test_Parameter_DefaultValues(t *testing.T) {
	p := corevalidator.Parameter{}
	if p.CaseIndex != 0 {
		t.Error("default CaseIndex should be 0")
	}
	if p.Header != "" {
		t.Error("default Header should be empty")
	}
	if p.IsSkipCompareOnActualEmpty {
		t.Error("default IsSkipCompareOnActualEmpty should be false")
	}
	if p.IsAttachUserInputs {
		t.Error("default IsAttachUserInputs should be false")
	}
	if p.IsCaseSensitive {
		t.Error("default IsCaseSensitive should be false")
	}
}
