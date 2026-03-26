package corestr

import (
	"testing"
)

// Tests using unexported field 'items' — must remain in source package.

func TestHashmap_ParseInjectUsingJson_Unexported_C13(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	jr := hm.Json()
	hm2 := &Hashmap{items: map[string]string{}}
	result, err := hm2.ParseInjectUsingJson(&jr)
	if err != nil || result.Length() != 1 { t.Fatal("unexpected") }
}

func TestHashmap_ParseInjectUsingJsonMust_Unexported_C13(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	jr := hm.Json()
	hm2 := &Hashmap{items: map[string]string{}}
	result := hm2.ParseInjectUsingJsonMust(&jr)
	if result.Length() != 1 { t.Fatal("unexpected") }
}

func TestHashmap_JsonParseSelfInject_Unexported_C13(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	jr := hm.Json()
	hm2 := &Hashmap{items: map[string]string{}}
	err := hm2.JsonParseSelfInject(&jr)
	if err != nil { t.Fatal("unexpected") }
}

func TestHashmap_Dispose_Unexported_C13(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm.Dispose()
	if hm.items != nil { t.Fatal("expected nil") }
}
