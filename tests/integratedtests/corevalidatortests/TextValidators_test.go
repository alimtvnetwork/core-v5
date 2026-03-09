package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

// ==========================================
// TextValidators — collection operations
// ==========================================

func Test_TextValidators_NewEmpty(t *testing.T) {
	v := corevalidator.NewTextValidators(5)
	if !v.IsEmpty() {
		t.Error("new should be empty")
	}
	if v.Length() != 0 {
		t.Errorf("length should be 0, got %d", v.Length())
	}
}

func Test_TextValidators_Add(t *testing.T) {
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	v.Add(corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal})
	if v.Length() != 2 {
		t.Errorf("expected 2, got %d", v.Length())
	}
}

func Test_TextValidators_Adds(t *testing.T) {
	v := corevalidator.NewTextValidators(2)
	v.Adds(
		corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
		corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal},
	)
	if v.Length() != 2 {
		t.Errorf("expected 2, got %d", v.Length())
	}
}

func Test_TextValidators_Adds_Empty(t *testing.T) {
	v := corevalidator.NewTextValidators(2)
	v.Adds()
	if v.Length() != 0 {
		t.Error("adds with nothing should stay empty")
	}
}

func Test_TextValidators_AddSimple(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.AddSimple("test", stringcompareas.Contains)
	if v.Length() != 1 {
		t.Error("AddSimple should add one")
	}
}

func Test_TextValidators_HasIndex(t *testing.T) {
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	if !v.HasIndex(0) {
		t.Error("should have index 0")
	}
	if v.HasIndex(1) {
		t.Error("should not have index 1")
	}
}

func Test_TextValidators_LastIndex(t *testing.T) {
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	v.Add(corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal})
	if v.LastIndex() != 1 {
		t.Errorf("expected 1, got %d", v.LastIndex())
	}
}

// ==========================================
// TextValidators.IsMatch
// ==========================================

func Test_TextValidators_IsMatch_EmptyValidators(t *testing.T) {
	v := corevalidator.NewTextValidators(0)
	if !v.IsMatch("anything", true) {
		t.Error("empty validators should match anything")
	}
}

func Test_TextValidators_IsMatch_AllPass(t *testing.T) {
	v := corevalidator.NewTextValidators(2)
	v.AddSimple("hello", stringcompareas.Contains)
	v.AddSimple("world", stringcompareas.Contains)
	if !v.IsMatch("hello world", true) {
		t.Error("content containing both should match")
	}
}

func Test_TextValidators_IsMatch_OneFails(t *testing.T) {
	v := corevalidator.NewTextValidators(2)
	v.AddSimple("hello", stringcompareas.Contains)
	v.AddSimple("xyz", stringcompareas.Contains)
	if v.IsMatch("hello world", true) {
		t.Error("missing substring should fail")
	}
}

// ==========================================
// TextValidators.IsMatchMany
// ==========================================

func Test_TextValidators_IsMatchMany_EmptyValidators(t *testing.T) {
	v := corevalidator.NewTextValidators(0)
	if !v.IsMatchMany(false, true, "a", "b") {
		t.Error("empty validators should match many")
	}
}

// ==========================================
// TextValidators.VerifyFirstError
// ==========================================

func Test_TextValidators_VerifyFirstError_AllPass(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	err := v.VerifyFirstError(0, "hello", true)
	if err != nil {
		t.Errorf("should pass: %v", err)
	}
}

func Test_TextValidators_VerifyFirstError_Fails(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	err := v.VerifyFirstError(0, "world", true)
	if err == nil {
		t.Error("mismatch should return error")
	}
}

func Test_TextValidators_VerifyFirstError_Empty(t *testing.T) {
	v := corevalidator.NewTextValidators(0)
	err := v.VerifyFirstError(0, "anything", true)
	if err != nil {
		t.Error("empty validators should return nil")
	}
}

// ==========================================
// TextValidators.AllVerifyError
// ==========================================

func Test_TextValidators_AllVerifyError_AllPass(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	err := v.AllVerifyError(0, "hello", true)
	if err != nil {
		t.Errorf("should pass: %v", err)
	}
}

func Test_TextValidators_AllVerifyError_MultipleFail(t *testing.T) {
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{
		Search:    "x",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	v.Add(corevalidator.TextValidator{
		Search:    "y",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	err := v.AllVerifyError(0, "z", true)
	if err == nil {
		t.Error("both mismatches should return error")
	}
}

// ==========================================
// TextValidators.Dispose
// ==========================================

func Test_TextValidators_Dispose(t *testing.T) {
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	v.Dispose()
	if v.Items != nil {
		t.Error("Dispose should nil out Items")
	}
}

// (nil receiver tests migrated to TextValidators_NilReceiver_testcases.go)
