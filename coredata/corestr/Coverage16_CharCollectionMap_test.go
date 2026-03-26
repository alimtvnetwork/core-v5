package corestr

import (
	"testing"
)

// Tests using unexported symbols — must remain in source package.

func Test_CharCollectionMap_GetChar_EmptyChar_C16(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	if cm.GetChar("") != emptyChar {
		t.Fatal("expected emptyChar for empty string")
	}
}

func Test_CharCollectionMap_Dispose_Unexported_C16(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	cm.Dispose()
	if cm.items != nil {
		t.Fatal("expected nil items after dispose")
	}
}
