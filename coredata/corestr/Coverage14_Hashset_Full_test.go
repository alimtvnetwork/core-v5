package corestr

import (
	"testing"
)

// Tests using unexported fields 'items', 'hasMapUpdated' — must remain in source package.

func TestHashset_Dispose_Unexported_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.Dispose()
	if hs.items != nil { t.Fatal("expected nil") }
}

func TestHashset_Join_Unexported_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.Join(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_NonEmptyJoins_Unexported_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.NonEmptyJoins(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_NonWhitespaceJoins_Unexported_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.NonWhitespaceJoins(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_JsonModel_Unexported_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	m := hs.JsonModel()
	if len(m) != 1 { t.Fatal("expected 1") }
}

func TestHashset_JsonModelAny_Unexported_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	_ = hs.JsonModelAny()
}

func TestHashset_MarshalJSON_Unexported_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	b, err := hs.MarshalJSON()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestHashset_Json_Unexported_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	r := hs.Json()
	if r.HasError() { t.Fatal("unexpected") }
}

func TestHashset_ParseInjectUsingJson_Unexported_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	jr := hs.Json()
	hs2 := New.Hashset.Cap(5)
	_, err := hs2.ParseInjectUsingJson(&jr)
	if err != nil { t.Fatal("unexpected") }
}

func TestHashset_JoinLine_Unexported_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.JoinLine()
	if s == "" { t.Fatal("expected non-empty") }
}
