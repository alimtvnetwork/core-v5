package corestr

import (
	"testing"
)

// ── reflectInterfaceVal (unexported — must remain in source package) ──

func TestReflectInterfaceVal_Nil_C01(t *testing.T) {
	if reflectInterfaceVal(nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestReflectInterfaceVal_Value_C01(t *testing.T) {
	v := reflectInterfaceVal(42)
	if v != 42 {
		t.Fatal("expected 42")
	}
}

func TestReflectInterfaceVal_Ptr_C01(t *testing.T) {
	val := "hello"
	v := reflectInterfaceVal(&val)
	if v != "hello" {
		t.Fatal("expected hello")
	}
}
