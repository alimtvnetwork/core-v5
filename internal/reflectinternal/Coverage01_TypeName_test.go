package reflectinternal

import (
	"testing"
)

func TestTypeName(t *testing.T) {
	s := TypeName("hello")
	if s != "string" {
		t.Fatal("expected string")
	}

	s2 := TypeName(42)
	if s2 != "int" {
		t.Fatal("expected int")
	}

	s3 := TypeName(nil)
	if s3 != "" {
		t.Fatal("expected empty for nil")
	}
}

func TestTypeNames(t *testing.T) {
	names := TypeNames(true, "a", 1)
	if len(names) != 2 {
		t.Fatal("expected 2")
	}
	if names[0] != "string" {
		t.Fatal("expected string")
	}

	names2 := TypeNames(false, "a", 1)
	if len(names2) != 2 {
		t.Fatal("expected 2")
	}
}

func TestTypeNamesString(t *testing.T) {
	s := TypeNamesString(true, "a", 1)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTypeNamesReferenceString(t *testing.T) {
	s := TypeNamesReferenceString(true, "a", 1)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTypeNameToValidVariableName(t *testing.T) {
	// simple
	s := TypeNameToValidVariableName("string")
	if s != "string" {
		t.Fatal("expected string, got:", s)
	}

	// empty
	s2 := TypeNameToValidVariableName("")
	if s2 != "" {
		t.Fatal("expected empty")
	}

	// with dot (complex)
	s3 := TypeNameToValidVariableName("*pkg.Type")
	_ = s3

	// slice type
	s4 := TypeNameToValidVariableName("[]string")
	_ = s4
}
