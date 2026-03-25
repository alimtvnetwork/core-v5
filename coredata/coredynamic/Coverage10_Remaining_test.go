package coredynamic

import (
	"reflect"
	"testing"
)

// === CollectionDistinct ===

func TestCollectionDistinct(t *testing.T) {
	c := NewCollection[int](5)
	c.AddMany(1, 2, 2, 3, 3)
	dc := c.Distinct(func(a, b int) bool { return a == b })
	if dc.Length() != 3 { t.Fatal("expected 3") }
	ec := EmptyCollection[int]()
	if ec.Distinct(func(a, b int) bool { return a == b }).Length() != 0 { t.Fatal("expected 0") }
}

// === CollectionGroupBy ===

func TestCollectionGroupBy(t *testing.T) {
	c := NewCollection[string](5)
	c.AddMany("ab", "ac", "bc", "bd", "ca")
	groups := c.GroupBy(func(s string) string { return string(s[0]) })
	if len(groups) != 3 { t.Fatal("expected 3 groups") }
}

func TestCollectionGroupBy_Empty(t *testing.T) {
	c := EmptyCollection[string]()
	groups := c.GroupBy(func(s string) string { return s })
	if len(groups) != 0 { t.Fatal("expected 0") }
}

// === CollectionSearch ===

func TestCollectionSearch_FindFirst(t *testing.T) {
	c := NewCollection[int](5)
	c.AddMany(1, 2, 3, 4, 5)
	item, found := c.FindFirst(func(i int) bool { return i == 3 })
	if !found || *item != 3 { t.Fatal("expected 3") }
	_, found2 := c.FindFirst(func(i int) bool { return i == 9 })
	if found2 { t.Fatal("expected not found") }
}

func TestCollectionSearch_FindLast(t *testing.T) {
	c := NewCollection[int](5)
	c.AddMany(1, 2, 3, 2, 1)
	item, found := c.FindLast(func(i int) bool { return i == 2 })
	if !found || *item != 2 { t.Fatal("expected 2") }
}

func TestCollectionSearch_FindAll(t *testing.T) {
	c := NewCollection[int](5)
	c.AddMany(1, 2, 3, 4, 5)
	found := c.FindAll(func(i int) bool { return i > 3 })
	if found.Length() != 2 { t.Fatal("expected 2") }
}

// === CollectionSort ===

func TestCollectionSort(t *testing.T) {
	c := NewCollection[int](5)
	c.AddMany(3, 1, 4, 1, 5)
	c.Sort(func(a, b int) bool { return a < b })
	if c.At(0) != 1 { t.Fatal("expected 1") }
}

func TestCollectionSort_Empty(t *testing.T) {
	c := EmptyCollection[int]()
	c.Sort(func(a, b int) bool { return a < b })
}

// === CollectionMap ===

func TestCollectionMap(t *testing.T) {
	c := NewCollection[int](3)
	c.AddMany(1, 2, 3)
	mapped := CollectionMap[int, string](c, func(i int) string {
		return "x"
	})
	if mapped.Length() != 3 { t.Fatal("expected 3") }
}

func TestCollectionFlatMap(t *testing.T) {
	c := NewCollection[int](2)
	c.AddMany(1, 2)
	flat := CollectionFlatMap[int, string](c, func(i int) []string {
		return []string{"a", "b"}
	})
	if flat.Length() != 4 { t.Fatal("expected 4") }
}

func TestCollectionReduce(t *testing.T) {
	c := NewCollection[int](3)
	c.AddMany(1, 2, 3)
	sum := CollectionReduce[int, int](c, 0, func(acc int, i int) int {
		return acc + i
	})
	if sum != 6 { t.Fatal("expected 6") }
}

// === CollectionLock ===

func TestCollectionLock(t *testing.T) {
	c := NewCollection[int](3)
	c.AddMany(1, 2, 3)
	c.LockAdd(4)
	if c.Length() != 4 { t.Fatal("expected 4") }
	c.LockAddMany(5, 6)
	if c.Length() != 6 { t.Fatal("expected 6") }
	c.LockRemoveAt(0)
	if c.Length() != 5 { t.Fatal("expected 5") }
}

// === Misc functions ===

func TestIsAnyTypesOf(t *testing.T) {
	st := reflect.TypeOf("")
	it := reflect.TypeOf(0)
	if !IsAnyTypesOf(st, st, it) { t.Fatal("expected true") }
	if IsAnyTypesOf(st, it) { t.Fatal("expected false") }
	if IsAnyTypesOf(st) { t.Fatal("expected false") }
}

func TestTypeMustBeSame(t *testing.T) {
	ts := TypeSameStatus("a", "b")
	ts.MustBeSame()
}

func TestTypeNotEqualErr(t *testing.T) {
	err := TypeNotEqualErr("a", "b")
	if err != nil { t.Fatal("expected nil for same types") }
	err2 := TypeNotEqualErr("a", 1)
	if err2 == nil { t.Fatal("expected error") }
}

func TestTypesIndexOf(t *testing.T) {
	types := []reflect.Type{reflect.TypeOf(""), reflect.TypeOf(0)}
	idx := TypesIndexOf(types, reflect.TypeOf(0))
	if idx != 1 { t.Fatal("expected 1") }
	idx2 := TypesIndexOf(types, reflect.TypeOf(true))
	if idx2 != -1 { t.Fatal("expected -1") }
	idx3 := TypesIndexOf(nil, reflect.TypeOf(""))
	if idx3 != -1 { t.Fatal("expected -1") }
}

func TestSafeTypeName(t *testing.T) {
	s := SafeTypeName("hello")
	if s == "" { t.Fatal("expected non-empty") }
	s2 := SafeTypeName(nil)
	if s2 == "" { t.Fatal("expected non-empty") }
}

func TestPointerOrNonPointer(t *testing.T) {
	val, ok := PointerOrNonPointer(false, "hello")
	if !ok { t.Fatal("expected ok") }
	_ = val
}

func TestPointerOrNonPointerUsingReflectValue(t *testing.T) {
	rv := reflect.ValueOf("hello")
	val, ok := PointerOrNonPointerUsingReflectValue(false, rv)
	if !ok { t.Fatal("expected ok") }
	_ = val
}

func TestReflectKindValidation(t *testing.T) {
	err := ReflectKindValidation(reflect.String, "hello")
	if err != nil { t.Fatal("unexpected") }
	err2 := ReflectKindValidation(reflect.Int, "hello")
	if err2 == nil { t.Fatal("expected error") }
}

func TestReflectTypeValidation(t *testing.T) {
	err := ReflectTypeValidation(false, reflect.TypeOf(""), "hello")
	if err != nil { t.Fatal("unexpected") }
	err2 := ReflectTypeValidation(true, reflect.TypeOf(""), nil)
	if err2 == nil { t.Fatal("expected error") }
}

func TestReflectInterfaceVal(t *testing.T) {
	rv := reflect.ValueOf("hello")
	val := ReflectInterfaceVal(rv)
	if val != "hello" { t.Fatal("expected hello") }
}

func TestMustBeAcceptedTypes(t *testing.T) {
	st := reflect.TypeOf("")
	MustBeAcceptedTypes(st, st) // should not panic
}

func TestNotAcceptedTypesErr(t *testing.T) {
	st := reflect.TypeOf("")
	it := reflect.TypeOf(0)
	err := NotAcceptedTypesErr(st, it)
	if err == nil { t.Fatal("expected error") }
	err2 := NotAcceptedTypesErr(st, st)
	if err2 != nil { t.Fatal("expected nil") }
}

func TestLengthOfReflect(t *testing.T) {
	rv := reflect.ValueOf([]int{1, 2, 3})
	if LengthOfReflect(rv) != 3 { t.Fatal("expected 3") }
	rv2 := reflect.ValueOf("hello")
	if LengthOfReflect(rv2) != 5 { t.Fatal("expected 5") }
	rv3 := reflect.ValueOf(42)
	if LengthOfReflect(rv3) != 0 { t.Fatal("expected 0") }
}

func TestSafeZeroSet(t *testing.T) {
	s := "hello"
	SafeZeroSet(&s)
	if s != "" { t.Fatal("expected empty") }
}

func TestZeroSet(t *testing.T) {
	s := "hello"
	ZeroSet(&s)
	if s != "" { t.Fatal("expected empty") }
}

func TestZeroSetAny(t *testing.T) {
	s := "hello"
	ZeroSetAny(&s)
}

func TestMapAsKeyValSlice(t *testing.T) {
	m := map[string]int{"a": 1}
	rv := reflect.ValueOf(m)
	kvc, err := MapAsKeyValSlice(rv)
	if err != nil || kvc.Length() != 1 { t.Fatal("unexpected") }
}

func TestAnyTypeMapToMapStringAny(t *testing.T) {
	m := map[string]int{"a": 1}
	result, err := AnyTypeMapToMapStringAny(m)
	if err != nil || len(result) != 1 { t.Fatal("unexpected") }
	_, err2 := AnyTypeMapToMapStringAny(nil)
	if err2 == nil { t.Fatal("expected error") }
}

// === newCreator constructors ===

func TestNewCreator(t *testing.T) {
	_ = New.AnyCollection(5)
	_ = New.EmptyAnyCollection()
	_ = New.DynamicCollection(5)
	_ = New.EmptyDynamicCollection()
	_ = New.KeyValCollection(5)
	_ = New.EmptyKeyValCollection()
	_ = New.MapAnyItems(5)
	_ = New.EmptyMapAnyItems()
	_ = New.BytesConverter([]byte(`"x"`))
	_ = New.Dynamic("x", true)
	_ = New.DynamicPtr("x", true)
	_ = New.DynamicValid("x")
	_ = New.InvalidDynamic()
	_ = New.InvalidDynamicPtr()
}

// === Typed collection creators ===

func TestNewStringCollectionCreator(t *testing.T) {
	_ = NewStringCollectionCreator.Default()
	_ = NewStringCollectionCreator.UsingCap(5)
	_ = NewStringCollectionCreator.UsingItems("a", "b")
}

func TestNewIntCollectionCreator(t *testing.T) {
	_ = NewIntCollectionCreator.Default()
	_ = NewIntCollectionCreator.UsingCap(5)
	_ = NewIntCollectionCreator.UsingItems(1, 2)
}

func TestNewInt64CollectionCreator(t *testing.T) {
	_ = NewInt64CollectionCreator.Default()
	_ = NewInt64CollectionCreator.UsingCap(5)
	_ = NewInt64CollectionCreator.UsingItems(int64(1))
}

func TestNewByteCollectionCreator(t *testing.T) {
	_ = NewByteCollectionCreator.Default()
	_ = NewByteCollectionCreator.UsingCap(5)
	_ = NewByteCollectionCreator.UsingItems(byte(1))
}

func TestNewAnyCollectionCreator(t *testing.T) {
	_ = NewAnyCollectionCreator.Default()
	_ = NewAnyCollectionCreator.UsingCap(5)
	_ = NewAnyCollectionCreator.UsingItems("a", 1)
}

func TestNewGenericCollectionCreator(t *testing.T) {
	_ = NewGenericCollectionCreator.Default()
	_ = NewGenericCollectionCreator.UsingCap(5)
}

// === TypedSimpleRequest/Result ===

func TestTypedSimpleRequest(t *testing.T) {
	tsr := NewTypedSimpleRequest[string]("hello", true, "")
	if tsr.Request() != "hello" { t.Fatal("expected hello") }
	if !tsr.IsValid() { t.Fatal("expected valid") }
	tsr2 := NewTypedSimpleRequestValid[int](42)
	if tsr2.Data() != 42 { t.Fatal("expected 42") }
	tsr3 := InvalidTypedSimpleRequest[string]("err")
	if tsr3.IsValid() { t.Fatal("expected invalid") }
	tsr4 := InvalidTypedSimpleRequestNoMessage[int]()
	if tsr4.IsValid() { t.Fatal("expected invalid") }
}

func TestTypedSimpleResult(t *testing.T) {
	tsr := NewTypedSimpleResult[string]("hello", true, "")
	if tsr.Result() != "hello" { t.Fatal("expected hello") }
	tsr2 := NewTypedSimpleResultValid[int](42)
	if tsr2.Data() != 42 { t.Fatal("expected 42") }
	tsr3 := InvalidTypedSimpleResult[string]("err")
	if tsr3.IsValid() { t.Fatal("expected invalid") }
	tsr4 := InvalidTypedSimpleResultNoMessage[int]()
	if tsr4.IsValid() { t.Fatal("expected invalid") }
}
