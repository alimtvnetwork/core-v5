package corestr

import (
	"testing"
)

// Tests using unexported field 'items' — must remain in source package.

func TestCollection_Dispose_ItemsNil_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c.Dispose()
	if c.items != nil { t.Fatal("expected nil") }
}
