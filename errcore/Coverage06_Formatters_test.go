package errcore

import (
	"testing"
)

// ── Unexported functions — must remain in source package ──

func TestGetReferenceMessage_Nil(t *testing.T) {
	if getReferenceMessage(nil) != "" {
		t.Fatal("expected empty")
	}
}

func TestGetReferenceMessage_EmptyString(t *testing.T) {
	if getReferenceMessage("") != "" {
		t.Fatal("expected empty")
	}
}

func TestGetReferenceMessage_WithRef(t *testing.T) {
	s := getReferenceMessage("ref")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTypesNamesString(t *testing.T) {
	s := typesNamesString("a", 1)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}
