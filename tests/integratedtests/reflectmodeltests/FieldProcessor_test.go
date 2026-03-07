package reflectmodeltests

import (
	"reflect"
	"testing"
)

// ===== FieldProcessor Tests =====

func Test_FieldProcessor_IsFieldType_Match(t *testing.T) {
	fp := newFieldProcessor("Name", 0)
	if fp == nil {
		t.Fatal("failed to create FieldProcessor for Name")
	}

	if !fp.IsFieldType(reflect.TypeOf("")) {
		t.Error("expected IsFieldType(string) = true for Name field")
	}
}

func Test_FieldProcessor_IsFieldType_NoMatch(t *testing.T) {
	fp := newFieldProcessor("Name", 0)
	if fp == nil {
		t.Fatal("failed to create FieldProcessor for Name")
	}

	if fp.IsFieldType(reflect.TypeOf(0)) {
		t.Error("expected IsFieldType(int) = false for Name (string) field")
	}
}

func Test_FieldProcessor_IsFieldType_NilReceiver(t *testing.T) {
	var fp *FieldProcessorAlias

	if fp.IsFieldType(reflect.TypeOf("")) {
		t.Error("expected IsFieldType = false on nil receiver")
	}
}

func Test_FieldProcessor_IsFieldKind_Match(t *testing.T) {
	fp := newFieldProcessor("Age", 1)
	if fp == nil {
		t.Fatal("failed to create FieldProcessor for Age")
	}

	if !fp.IsFieldKind(reflect.Int) {
		t.Error("expected IsFieldKind(Int) = true for Age field")
	}
}

func Test_FieldProcessor_IsFieldKind_NoMatch(t *testing.T) {
	fp := newFieldProcessor("Age", 1)
	if fp == nil {
		t.Fatal("failed to create FieldProcessor for Age")
	}

	if fp.IsFieldKind(reflect.String) {
		t.Error("expected IsFieldKind(String) = false for Age (int) field")
	}
}

func Test_FieldProcessor_IsFieldKind_NilReceiver(t *testing.T) {
	var fp *FieldProcessorAlias

	if fp.IsFieldKind(reflect.String) {
		t.Error("expected IsFieldKind = false on nil receiver")
	}
}

func Test_FieldProcessor_BoolField(t *testing.T) {
	fp := newFieldProcessor("Active", 2)
	if fp == nil {
		t.Fatal("failed to create FieldProcessor for Active")
	}

	if !fp.IsFieldKind(reflect.Bool) {
		t.Error("expected Active field to be Bool kind")
	}

	if !fp.IsFieldType(reflect.TypeOf(true)) {
		t.Error("expected Active field to match bool type")
	}
}

func Test_FieldProcessor_FieldData(t *testing.T) {
	fp := newFieldProcessor("Name", 0)
	if fp == nil {
		t.Fatal("failed to create FieldProcessor for Name")
	}

	if fp.Name != "Name" {
		t.Errorf("Name = %q, want %q", fp.Name, "Name")
	}

	if fp.Index != 0 {
		t.Errorf("Index = %d, want 0", fp.Index)
	}

	if fp.Field.Name != "Name" {
		t.Errorf("Field.Name = %q, want %q", fp.Field.Name, "Name")
	}
}
