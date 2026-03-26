package corestr

import (
	"testing"
)

// Tests using unexported symbols — must remain in source package.

func Test_CharHashsetMap_GetChar_EmptyChar_C17(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	if hsm.GetChar("") != emptyChar {
		t.Fatal("expected emptyChar")
	}
}

func Test_CharHashsetMap_GetCharOf_EmptyChar_C17(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	if hsm.GetCharOf("") != emptyChar {
		t.Fatal("expected emptyChar")
	}
}
