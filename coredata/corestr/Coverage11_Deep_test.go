package corestr

import (
	"testing"
)

// Tests using unexported symbols — must remain in source package.

func TestIsCollectionPrecheckEqual_BothNil_C11(t *testing.T) {
	var a, b *Collection
	result, handled := isCollectionPrecheckEqual(a, b)
	if !handled || !result { t.Fatal("expected true") }
}

func TestIsCollectionPrecheckEqual_OneNil_C11(t *testing.T) {
	a := New.Collection.Strings([]string{"a"})
	result, handled := isCollectionPrecheckEqual(a, nil)
	if !handled || result { t.Fatal("expected false") }
}
