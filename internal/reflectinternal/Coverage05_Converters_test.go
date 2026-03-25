package reflectinternal

import (
	"reflect"
	"testing"
)

// ======= reflectConverter =======

func TestConverter_ArgsToReflectValues(t *testing.T) {
	rv := Converter.ArgsToReflectValues([]any{"a", 1})
	if len(rv) != 2 {
		t.Fatal("expected 2")
	}
	rv2 := Converter.ArgsToReflectValues([]any{})
	if len(rv2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestConverter_ReflectValuesToInterfaces(t *testing.T) {
	rvs := []reflect.Value{reflect.ValueOf("a"), reflect.ValueOf(1)}
	ifs := Converter.ReflectValuesToInterfaces(rvs)
	if len(ifs) != 2 {
		t.Fatal("expected 2")
	}
	ifs2 := Converter.ReflectValuesToInterfaces([]reflect.Value{})
	if len(ifs2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestConverter_ReflectValueToAnyValue(t *testing.T) {
	v := Converter.ReflectValueToAnyValue(reflect.ValueOf("hello"))
	if v != "hello" {
		t.Fatal("unexpected")
	}
	v2 := Converter.ReflectValueToAnyValue(reflect.ValueOf(42))
	if v2 != int64(42) {
		t.Fatal("unexpected")
	}
	s := "hello"
	v3 := Converter.ReflectValueToAnyValue(reflect.ValueOf(&s))
	if v3 != "hello" {
		t.Fatal("unexpected")
	}
}

func TestConverter_InterfacesToTypes(t *testing.T) {
	types := Converter.InterfacesToTypes([]any{"a", 1})
	if len(types) != 2 {
		t.Fatal("expected 2")
	}
	types2 := Converter.InterfacesToTypes([]any{})
	if len(types2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestConverter_InterfacesToTypesNames(t *testing.T) {
	names := Converter.InterfacesToTypesNames([]any{"a", 1})
	if len(names) != 2 {
		t.Fatal("expected 2")
	}
	names2 := Converter.InterfacesToTypesNames([]any{})
	if len(names2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestConverter_InterfacesToTypesNamesWithValues(t *testing.T) {
	names := Converter.InterfacesToTypesNamesWithValues([]any{"a", 1})
	if len(names) != 2 {
		t.Fatal("expected 2")
	}
	names2 := Converter.InterfacesToTypesNamesWithValues([]any{})
	if len(names2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestConverter_ReflectValueToPointerReflectValue(t *testing.T) {
	rv := reflect.ValueOf(42)
	prv := Converter.ReflectValueToPointerReflectValue(rv)
	if prv.Kind() != reflect.Ptr {
		t.Fatal("expected ptr")
	}
}

func TestConverter_ToPtrRvIfNotAlready(t *testing.T) {
	rv := reflect.ValueOf(42)
	prv := Converter.ToPtrRvIfNotAlready(rv)
	if prv.Kind() != reflect.Ptr {
		t.Fatal("expected ptr")
	}

	s := "hello"
	ptrRv := reflect.ValueOf(&s)
	prv2 := Converter.ToPtrRvIfNotAlready(ptrRv)
	if prv2.Kind() != reflect.Ptr {
		t.Fatal("expected ptr")
	}
}

func TestConverter_ReducePointer(t *testing.T) {
	s := "hello"
	r := Converter.ReducePointer(&s, 3)
	if r == nil || !r.IsValid {
		t.Fatal("expected valid")
	}
}

func TestConverter_ReducePointerRv(t *testing.T) {
	s := "hello"
	r := Converter.ReducePointerRv(reflect.ValueOf(&s), 3)
	if r == nil || !r.IsValid {
		t.Fatal("expected valid")
	}
}

func TestConverter_ReducePointerDefault(t *testing.T) {
	r := Converter.ReducePointerDefault("hello")
	if r == nil || !r.IsValid {
		t.Fatal("expected valid")
	}
}

func TestConverter_ReducePointerRvDefault(t *testing.T) {
	r := Converter.ReducePointerRvDefault(reflect.ValueOf("hello"))
	if r == nil || !r.IsValid {
		t.Fatal("expected valid")
	}
}

func TestConverter_ReducePointerDefaultToType(t *testing.T) {
	rt := Converter.ReducePointerDefaultToType("hello")
	if rt == nil {
		t.Fatal("expected non-nil")
	}
}

func TestConverter_ReducePointerRvDefaultToType(t *testing.T) {
	rt := Converter.ReducePointerRvDefaultToType(reflect.ValueOf("hello"))
	if rt == nil {
		t.Fatal("expected non-nil")
	}
}

func TestConverter_ReflectValToInterfaces(t *testing.T) {
	slice := []string{"a", "b"}
	ifs := Converter.ReflectValToInterfaces(false, reflect.ValueOf(slice))
	if len(ifs) != 2 {
		t.Fatal("expected 2")
	}

	// ptr
	ifs2 := Converter.ReflectValToInterfaces(false, reflect.ValueOf(&slice))
	if len(ifs2) != 2 {
		t.Fatal("expected 2")
	}

	// non-slice
	ifs3 := Converter.ReflectValToInterfaces(false, reflect.ValueOf(42))
	if len(ifs3) != 0 {
		t.Fatal("expected 0")
	}

	// empty
	ifs4 := Converter.ReflectValToInterfaces(false, reflect.ValueOf([]string{}))
	if len(ifs4) != 0 {
		t.Fatal("expected 0")
	}
}

func TestConverter_ReflectValToInterfacesAsync(t *testing.T) {
	slice := []string{"a", "b"}
	ifs := Converter.ReflectValToInterfacesAsync(reflect.ValueOf(slice))
	if len(ifs) != 2 {
		t.Fatal("expected 2")
	}
}

func TestConverter_ReflectInterfaceVal(t *testing.T) {
	v := Converter.ReflectInterfaceVal("hello")
	if v != "hello" {
		t.Fatal("unexpected")
	}
	s := "hello"
	v2 := Converter.ReflectInterfaceVal(&s)
	if v2 != "hello" {
		t.Fatal("unexpected")
	}
}

func TestConverter_ToPointerRv(t *testing.T) {
	rv := Converter.ToPointerRv("hello")
	if rv == nil {
		t.Fatal("expected non-nil")
	}
	rv2 := Converter.ToPointerRv(nil)
	if rv2 != nil {
		t.Fatal("expected nil")
	}
}

func TestConverter_ToPointer(t *testing.T) {
	p := Converter.ToPointer("hello")
	if p == nil {
		t.Fatal("expected non-nil")
	}
	p2 := Converter.ToPointer(nil)
	if p2 != nil {
		t.Fatal("expected nil")
	}
}

// ======= SliceConverter =======

func TestSliceConverter_Length(t *testing.T) {
	if SliceConverter.Length([]int{1, 2}) != 2 {
		t.Fatal("expected 2")
	}
	if SliceConverter.Length(nil) != 0 {
		t.Fatal("expected 0")
	}
	if SliceConverter.Length(42) != 0 {
		t.Fatal("expected 0 for non-slice")
	}
	if SliceConverter.Length(map[string]int{"a": 1}) != 1 {
		t.Fatal("expected 1 for map")
	}
}

func TestSliceConverter_ToStrings(t *testing.T) {
	s, err := SliceConverter.ToStrings([]int{1, 2})
	if err != nil || len(s) != 2 {
		t.Fatal("unexpected")
	}
}

func TestSliceConverter_ToStringsMust(t *testing.T) {
	s := SliceConverter.ToStringsMust([]int{1, 2})
	if len(s) != 2 {
		t.Fatal("expected 2")
	}
}

func TestSliceConverter_ToStringsRv(t *testing.T) {
	s, err := SliceConverter.ToStringsRv(reflect.ValueOf([]int{1}))
	if err != nil || len(s) != 1 {
		t.Fatal("unexpected")
	}

	// not slice
	_, err2 := SliceConverter.ToStringsRv(reflect.ValueOf(42))
	if err2 == nil {
		t.Fatal("expected error")
	}

	// empty
	s2, _ := SliceConverter.ToStringsRv(reflect.ValueOf([]int{}))
	if len(s2) != 0 {
		t.Fatal("expected 0")
	}

	// ptr
	items := []int{1}
	s3, _ := SliceConverter.ToStringsRv(reflect.ValueOf(&items))
	if len(s3) != 1 {
		t.Fatal("expected 1")
	}
}

func TestSliceConverter_ToAnyItemsAsync(t *testing.T) {
	items := SliceConverter.ToAnyItemsAsync([]int{1, 2})
	if len(items) != 2 {
		t.Fatal("expected 2")
	}
	items2 := SliceConverter.ToAnyItemsAsync(nil)
	if len(items2) != 0 {
		t.Fatal("expected 0")
	}
}

// ======= MapConverter =======

func TestMapConverter_Length(t *testing.T) {
	if MapConverter.Length(map[string]int{"a": 1}) != 1 {
		t.Fatal("expected 1")
	}
}

func TestMapConverter_ToStrings(t *testing.T) {
	keys, err := MapConverter.ToStrings(map[string]int{"a": 1})
	if err != nil || len(keys) != 1 {
		t.Fatal("unexpected")
	}
	keys2, _ := MapConverter.ToStrings(nil)
	if len(keys2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestMapConverter_ToKeysStrings(t *testing.T) {
	keys, err := MapConverter.ToKeysStrings(map[string]int{"a": 1})
	if err != nil || len(keys) != 1 {
		t.Fatal("unexpected")
	}
}

func TestMapConverter_ToSortedStrings(t *testing.T) {
	keys, err := MapConverter.ToSortedStrings(map[string]int{"b": 1, "a": 2})
	if err != nil || keys[0] != "a" {
		t.Fatal("unexpected")
	}
	keys2, _ := MapConverter.ToSortedStrings(nil)
	if len(keys2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestMapConverter_ToStringsMust(t *testing.T) {
	keys := MapConverter.ToStringsMust(map[string]int{"a": 1})
	if len(keys) != 1 {
		t.Fatal("expected 1")
	}
}

func TestMapConverter_ToSortedStringsMust(t *testing.T) {
	keys := MapConverter.ToSortedStringsMust(map[string]int{"b": 1, "a": 2})
	if keys[0] != "a" {
		t.Fatal("expected sorted")
	}
	keys2 := MapConverter.ToSortedStringsMust(nil)
	if len(keys2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestMapConverter_ToValuesAny(t *testing.T) {
	vals, err := MapConverter.ToValuesAny(map[string]int{"a": 1})
	if err != nil || len(vals) != 1 {
		t.Fatal("unexpected")
	}
	vals2, _ := MapConverter.ToValuesAny(nil)
	if len(vals2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestMapConverter_ToKeysAny(t *testing.T) {
	keys, err := MapConverter.ToKeysAny(map[string]int{"a": 1})
	if err != nil || len(keys) != 1 {
		t.Fatal("unexpected")
	}
	keys2, _ := MapConverter.ToKeysAny(nil)
	if len(keys2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestMapConverter_ToKeysValuesAny(t *testing.T) {
	keys, vals, err := MapConverter.ToKeysValuesAny(map[string]int{"a": 1})
	if err != nil || len(keys) != 1 || len(vals) != 1 {
		t.Fatal("unexpected")
	}
	k2, v2, _ := MapConverter.ToKeysValuesAny(nil)
	if len(k2) != 0 || len(v2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestMapConverter_ToMapStringAny(t *testing.T) {
	m, err := MapConverter.ToMapStringAny(map[string]int{"a": 1})
	if err != nil || len(m) != 1 {
		t.Fatal("unexpected")
	}
	m2, _ := MapConverter.ToMapStringAny(nil)
	if len(m2) != 0 {
		t.Fatal("expected 0")
	}

	// non-string key
	m3, err3 := MapConverter.ToMapStringAny(map[int]string{1: "a"})
	if err3 != nil || len(m3) != 1 {
		t.Fatal("unexpected")
	}
}

func TestMapConverter_ToMapStringAnyRv(t *testing.T) {
	// not a map
	_, err := MapConverter.ToMapStringAnyRv(reflect.ValueOf(42))
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestMapConverter_ToStringsRv_NotMap(t *testing.T) {
	_, err := MapConverter.ToStringsRv(reflect.ValueOf(42))
	if err == nil {
		t.Fatal("expected error")
	}
}
