package corepropertytests

import (
	"testing"

	"github.com/alimtvnetwork/core/codegen/coreproperty"
)

func Test_QW_Write_Map(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	result := coreproperty.Writer.Write(m)
	if result == "" {
		t.Fatal("expected non-empty for map")
	}
}

func Test_QW_Write_NilPointer(t *testing.T) {
	var p *string
	result := coreproperty.Writer.Write(p)
	if result != "nil" {
		t.Fatalf("expected 'nil', got '%s'", result)
	}
}

func Test_QW_Write_NonNilPointer(t *testing.T) {
	s := "hello"
	result := coreproperty.Writer.Write(&s)
	if result == "" {
		t.Fatal("expected non-empty")
	}
}
