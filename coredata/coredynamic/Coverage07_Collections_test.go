package coredynamic

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// === Collection[T] ===

func TestCollection_Basic(t *testing.T) {
	c := NewCollection[string](5)
	if c.Length() != 0 { t.Fatal("expected 0") }
	if !c.IsEmpty() { t.Fatal("expected empty") }
	c.Add("a").Add("b").Add("c")
	if c.Length() != 3 { t.Fatal("expected 3") }
	if c.Count() != 3 { t.Fatal("expected 3") }
	if !c.HasAnyItem() { t.Fatal("expected items") }
	if c.At(0) != "a" { t.Fatal("expected a") }
	if c.First() != "a" { t.Fatal("expected a") }
	if c.Last() != "c" { t.Fatal("expected c") }
	if c.LastIndex() != 2 { t.Fatal("expected 2") }
	if !c.HasIndex(1) { t.Fatal("expected true") }
	if c.HasIndex(-1) { t.Fatal("expected false") }
}

func TestCollection_NilReceiver(t *testing.T) {
	var c *Collection[string]
	if c.Length() != 0 { t.Fatal("expected 0") }
	if !c.IsEmpty() { t.Fatal("expected empty") }
	if len(c.Items()) != 0 { t.Fatal("expected empty") }
}

func TestCollection_EmptyAndFrom(t *testing.T) {
	ec := EmptyCollection[int]()
	if ec.Length() != 0 { t.Fatal("expected 0") }
	fc := CollectionFrom[int](nil)
	if fc.Length() != 0 { t.Fatal("expected 0") }
	fc2 := CollectionFrom([]int{1, 2})
	if fc2.Length() != 2 { t.Fatal("expected 2") }
	cc := CollectionClone([]int{1, 2, 3})
	if cc.Length() != 3 { t.Fatal("expected 3") }
}

func TestCollection_FirstLastOrDefault(t *testing.T) {
	c := EmptyCollection[string]()
	f, ok := c.FirstOrDefault()
	if f != nil || ok { t.Fatal("expected nil") }
	l, ok2 := c.LastOrDefault()
	if l != nil || ok2 { t.Fatal("expected nil") }
	c.Add("x")
	f2, ok3 := c.FirstOrDefault()
	if f2 == nil || !ok3 { t.Fatal("expected x") }
	l2, ok4 := c.LastOrDefault()
	if l2 == nil || !ok4 { t.Fatal("expected x") }
}

func TestCollection_SkipTakeLimit(t *testing.T) {
	c := NewCollection[int](5)
	c.AddMany(1, 2, 3, 4, 5)
	if len(c.Skip(2)) != 3 { t.Fatal("expected 3") }
	if len(c.Take(3)) != 3 { t.Fatal("expected 3") }
	if len(c.Limit(3)) != 3 { t.Fatal("expected 3") }
	sc := c.SkipCollection(2)
	if sc.Length() != 3 { t.Fatal("expected 3") }
	tc := c.TakeCollection(3)
	if tc.Length() != 3 { t.Fatal("expected 3") }
	lc := c.LimitCollection(3)
	if lc.Length() != 3 { t.Fatal("expected 3") }
	slc := c.SafeLimitCollection(10)
	if slc.Length() != 5 { t.Fatal("expected 5") }
}

func TestCollection_AddNonNil(t *testing.T) {
	c := NewCollection[string](2)
	c.AddNonNil(nil)
	s := "hello"
	c.AddNonNil(&s)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_RemoveAt(t *testing.T) {
	c := NewCollection[string](3)
	c.AddMany("a", "b", "c")
	if !c.RemoveAt(1) { t.Fatal("expected success") }
	if c.RemoveAt(10) { t.Fatal("expected failure") }
	if c.RemoveAt(-1) { t.Fatal("expected failure") }
}

func TestCollection_ClearDispose(t *testing.T) {
	c := NewCollection[int](3)
	c.AddMany(1, 2, 3)
	c.Clear()
	if c.Length() != 0 { t.Fatal("expected 0") }
	c.AddMany(1, 2)
	c.Dispose()
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_Loop(t *testing.T) {
	c := NewCollection[int](3)
	c.AddMany(1, 2, 3)
	count := 0
	c.Loop(func(index int, item int) bool {
		count++
		return false
	})
	if count != 3 { t.Fatal("expected 3") }
	ec := EmptyCollection[int]()
	ec.Loop(func(index int, item int) bool { return false })
}

func TestCollection_Loop_Break(t *testing.T) {
	c := NewCollection[int](3)
	c.AddMany(1, 2, 3)
	count := 0
	c.Loop(func(index int, item int) bool {
		count++
		return true
	})
	if count != 1 { t.Fatal("expected 1") }
}

func TestCollection_LoopAsync(t *testing.T) {
	c := NewCollection[int](3)
	c.AddMany(1, 2, 3)
	c.LoopAsync(func(index int, item int) {})
	ec := EmptyCollection[int]()
	ec.LoopAsync(func(index int, item int) {})
}

func TestCollection_Filter(t *testing.T) {
	c := NewCollection[int](5)
	c.AddMany(1, 2, 3, 4, 5)
	filtered := c.Filter(func(i int) bool { return i > 3 })
	if filtered.Length() != 2 { t.Fatal("expected 2") }
	ec := EmptyCollection[int]()
	if ec.Filter(func(i int) bool { return true }).Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_Paging(t *testing.T) {
	c := NewCollection[int](10)
	for i := 0; i < 10; i++ { c.Add(i) }
	if c.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if c.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	if c.GetPagesSize(-1) != 0 { t.Fatal("expected 0") }
	paged := c.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4") }
	single := c.GetSinglePageCollection(3, 1)
	if single.Length() != 3 { t.Fatal("expected 3") }
	small := NewCollection[int](1)
	small.Add(1)
	if len(small.GetPagedCollection(5)) != 1 { t.Fatal("expected 1") }
	if small.GetSinglePageCollection(5, 1).Length() != 1 { t.Fatal("expected 1") }
	_ = c.GetPagingInfo(3, 1)
}

func TestCollection_JSON(t *testing.T) {
	c := NewCollection[string](2)
	c.AddMany("a", "b")
	b, err := c.MarshalJSON()
	if err != nil { t.Fatal("unexpected") }
	c2 := EmptyCollection[string]()
	err2 := c2.UnmarshalJSON(b)
	if err2 != nil { t.Fatal("unexpected") }
	s, err3 := c.JsonString()
	if err3 != nil || s == "" { t.Fatal("unexpected") }
	_ = c.JsonStringMust()
}

func TestCollection_Strings(t *testing.T) {
	c := NewCollection[int](2)
	c.AddMany(1, 2)
	strs := c.Strings()
	if len(strs) != 2 { t.Fatal("expected 2") }
	_ = c.String()
}

// === CollectionMethods ===

func TestCollection_AddIf(t *testing.T) {
	c := NewCollection[string](2)
	c.AddIf(true, "a")
	c.AddIf(false, "b")
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddManyIf(t *testing.T) {
	c := NewCollection[string](3)
	c.AddManyIf(true, "a", "b")
	c.AddManyIf(false, "c")
	c.AddManyIf(true)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_AddCollection(t *testing.T) {
	c := NewCollection[string](3)
	c.Add("a")
	c2 := NewCollection[string](2)
	c2.AddMany("b", "c")
	c.AddCollection(c2)
	if c.Length() != 3 { t.Fatal("expected 3") }
	c.AddCollection(nil)
	c.AddCollection(EmptyCollection[string]())
}

func TestCollection_AddCollections(t *testing.T) {
	c := NewCollection[int](5)
	c.Add(1)
	c2 := CollectionFrom([]int{2, 3})
	c3 := CollectionFrom([]int{4})
	c.AddCollections(nil, c2, c3, EmptyCollection[int]())
	if c.Length() != 4 { t.Fatal("expected 4") }
}

func TestCollection_ConcatNew(t *testing.T) {
	c := NewCollection[string](2)
	c.AddMany("a", "b")
	c2 := c.ConcatNew("c", "d")
	if c2.Length() != 4 { t.Fatal("expected 4") }
	if c.Length() != 2 { t.Fatal("original should be unchanged") }
}

func TestCollection_Clone_Method(t *testing.T) {
	c := NewCollection[string](2)
	c.AddMany("a", "b")
	cc := c.Clone()
	if cc.Length() != 2 { t.Fatal("expected 2") }
	var nc *Collection[string]
	ncc := nc.Clone()
	if ncc.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_Capacity(t *testing.T) {
	c := NewCollection[int](10)
	if c.Capacity() != 10 { t.Fatal("expected 10") }
	var nc *Collection[int]
	if nc.Capacity() != 0 { t.Fatal("expected 0") }
}

func TestCollection_AddCapacity(t *testing.T) {
	c := NewCollection[int](2)
	c.AddCapacity(0)
	c.AddCapacity(10)
	if c.Capacity() < 12 { t.Fatal("expected >= 12") }
}

func TestCollection_Resize(t *testing.T) {
	c := NewCollection[int](2)
	c.AddMany(1, 2)
	c.Resize(10)
	if c.Capacity() < 10 { t.Fatal("expected >= 10") }
	c.Resize(5) // smaller, should not shrink
}

func TestCollection_Reverse(t *testing.T) {
	c := NewCollection[int](3)
	c.AddMany(1, 2, 3)
	c.Reverse()
	if c.At(0) != 3 { t.Fatal("expected 3") }
	if c.At(2) != 1 { t.Fatal("expected 1") }
	ec := EmptyCollection[int]()
	ec.Reverse()
	sc := NewCollection[int](1)
	sc.Add(1)
	sc.Reverse()
}

func TestCollection_InsertAt(t *testing.T) {
	c := NewCollection[int](3)
	c.AddMany(1, 3)
	c.InsertAt(1, 2)
	if c.At(1) != 2 { t.Fatal("expected 2") }
	c.InsertAt(0) // empty insert
}

func TestCollection_IndexOfFunc(t *testing.T) {
	c := NewCollection[string](3)
	c.AddMany("a", "b", "c")
	idx := c.IndexOfFunc(func(s string) bool { return s == "b" })
	if idx != 1 { t.Fatal("expected 1") }
	idx2 := c.IndexOfFunc(func(s string) bool { return s == "z" })
	if idx2 != -1 { t.Fatal("expected -1") }
}

func TestCollection_ContainsFunc(t *testing.T) {
	c := NewCollection[int](3)
	c.AddMany(1, 2, 3)
	if !c.ContainsFunc(func(i int) bool { return i == 2 }) { t.Fatal("expected true") }
	if c.ContainsFunc(func(i int) bool { return i == 9 }) { t.Fatal("expected false") }
}

func TestCollection_SafeAt(t *testing.T) {
	c := NewCollection[int](2)
	c.AddMany(1, 2)
	if c.SafeAt(0) != 1 { t.Fatal("expected 1") }
	if c.SafeAt(10) != 0 { t.Fatal("expected 0") }
	if c.SafeAt(-1) != 0 { t.Fatal("expected 0") }
}

func TestCollection_SprintItems(t *testing.T) {
	c := NewCollection[int](2)
	c.AddMany(1, 2)
	strs := c.SprintItems("%d")
	if len(strs) != 2 { t.Fatal("expected 2") }
}

// === DynamicCollection ===

func TestDynamicCollection_Basic(t *testing.T) {
	dc := NewDynamicCollection(5)
	if dc.Length() != 0 { t.Fatal("expected 0") }
	if !dc.IsEmpty() { t.Fatal("expected empty") }
	dc.Add(NewDynamic("a", true))
	dc.Add(NewDynamic("b", true))
	if dc.Length() != 2 { t.Fatal("expected 2") }
	if dc.Count() != 2 { t.Fatal("expected 2") }
	if !dc.HasAnyItem() { t.Fatal("expected items") }
	_ = dc.At(0)
	_ = dc.First()
	_ = dc.Last()
	_ = dc.FirstDynamic()
	_ = dc.LastDynamic()
	_ = dc.FirstOrDefault()
	_ = dc.LastOrDefault()
	_ = dc.FirstOrDefaultDynamic()
	_ = dc.LastOrDefaultDynamic()
	if dc.LastIndex() != 1 { t.Fatal("expected 1") }
	if !dc.HasIndex(0) { t.Fatal("expected true") }
}

func TestDynamicCollection_NilReceiver(t *testing.T) {
	var dc *DynamicCollection
	if dc.Length() != 0 { t.Fatal("expected 0") }
	if !dc.IsEmpty() { t.Fatal("expected empty") }
	if len(dc.Items()) != 0 { t.Fatal("expected empty") }
}

func TestDynamicCollection_Empty(t *testing.T) {
	dc := EmptyDynamicCollection()
	if dc.FirstOrDefault() != nil { t.Fatal("expected nil") }
	if dc.LastOrDefault() != nil { t.Fatal("expected nil") }
}

func TestDynamicCollection_SkipTakeLimit(t *testing.T) {
	dc := NewDynamicCollection(5)
	for i := 0; i < 5; i++ { dc.Add(NewDynamic(i, true)) }
	if len(dc.Skip(2)) != 3 { t.Fatal("expected 3") }
	if len(dc.Take(3)) != 3 { t.Fatal("expected 3") }
	if len(dc.Limit(3)) != 3 { t.Fatal("expected 3") }
	_ = dc.SkipDynamic(1)
	_ = dc.TakeDynamic(2)
	_ = dc.LimitDynamic(2)
	_ = dc.SkipCollection(2)
	_ = dc.TakeCollection(3)
	_ = dc.LimitCollection(3)
	_ = dc.SafeLimitCollection(10)
}

func TestDynamicCollection_AddMethods(t *testing.T) {
	dc := NewDynamicCollection(10)
	dc.AddAny("x", true)
	dc.AddAnyNonNull(nil, false)
	dc.AddAnyNonNull("y", true)
	dc.AddAnyMany(nil)
	dc.AddAnyMany("a", "b")
	dp := NewDynamicPtr("z", true)
	dc.AddPtr(nil)
	dc.AddPtr(dp)
	dc.AddManyPtr(nil, dp)
	dc.AddAnySliceFromSingleItem(true, nil)
	dc.AddAnySliceFromSingleItem(true, []string{"q"})
}

func TestDynamicCollection_AddAnyWithTypeValidation(t *testing.T) {
	dc := NewDynamicCollection(2)
	err := dc.AddAnyWithTypeValidation(false, reflect.TypeOf(""), "hello")
	if err != nil { t.Fatal("unexpected") }
	err2 := dc.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(""), "a")
	if err2 != nil { t.Fatal("unexpected") }
	err3 := dc.AddAnyItemsWithTypeValidation(true, false, reflect.TypeOf(0), "a")
	if err3 == nil { t.Fatal("expected error") }
	err4 := dc.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(""))
	if err4 != nil { t.Fatal("unexpected") }
}

func TestDynamicCollection_RemoveAt(t *testing.T) {
	dc := NewDynamicCollection(3)
	dc.Add(NewDynamic("a", true)).Add(NewDynamic("b", true))
	if !dc.RemoveAt(0) { t.Fatal("expected success") }
	if dc.RemoveAt(10) { t.Fatal("expected failure") }
}

func TestDynamicCollection_ListStrings(t *testing.T) {
	dc := NewDynamicCollection(2)
	dc.Add(NewDynamic("a", true))
	strs := dc.ListStrings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	_ = dc.ListStringsPtr()
}

func TestDynamicCollection_AnyItems(t *testing.T) {
	dc := EmptyDynamicCollection()
	if len(dc.AnyItems()) != 0 { t.Fatal("expected empty") }
	dc.Add(NewDynamic("x", true))
	if len(dc.AnyItems()) != 1 { t.Fatal("expected 1") }
	_ = dc.AnyItemsCollection()
}

func TestDynamicCollection_Loop(t *testing.T) {
	dc := NewDynamicCollection(2)
	dc.Add(NewDynamic("a", true)).Add(NewDynamic("b", true))
	count := 0
	dc.Loop(func(index int, d *Dynamic) bool { count++; return false })
	if count != 2 { t.Fatal("expected 2") }
	ec := EmptyDynamicCollection()
	ec.Loop(func(index int, d *Dynamic) bool { return false })
}

func TestDynamicCollection_JSON(t *testing.T) {
	dc := NewDynamicCollection(1)
	dc.Add(NewDynamic("x", true))
	s, _ := dc.JsonString()
	if s == "" { t.Fatal("expected non-empty") }
	_ = dc.JsonStringMust()
	_, _ = dc.MarshalJSON()
	_ = dc.JsonModel()
	_ = dc.JsonModelAny()
	_ = dc.Json()
	_ = dc.JsonPtr()
	_ = dc.Strings()
	_ = dc.String()
}

func TestDynamicCollection_JsonResults(t *testing.T) {
	dc := EmptyDynamicCollection()
	rc := dc.JsonResultsCollection()
	if rc.Length() != 0 { t.Fatal("expected 0") }
	dc.Add(NewDynamic("x", true))
	_ = dc.JsonResultsCollection()
	_ = dc.JsonResultsPtrCollection()
}

func TestDynamicCollection_Paging(t *testing.T) {
	dc := NewDynamicCollection(10)
	for i := 0; i < 10; i++ { dc.Add(NewDynamic(i, true)) }
	if dc.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if dc.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	paged := dc.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4") }
	single := dc.GetSinglePageCollection(3, 1)
	if single.Length() != 3 { t.Fatal("expected 3") }
	small := NewDynamicCollection(1)
	small.Add(NewDynamic("x", true))
	if len(small.GetPagedCollection(5)) != 1 { t.Fatal("expected 1") }
	if small.GetSinglePageCollection(5, 1).Length() != 1 { t.Fatal("expected 1") }
}

func TestDynamicCollection_ParseInjectUsingJson(t *testing.T) {
	dc := NewDynamicCollection(1)
	dc.Add(NewDynamic("x", true))
	jr := dc.Json()
	dc2 := EmptyDynamicCollection()
	_, _ = dc2.ParseInjectUsingJson(&jr)
	_ = dc2.JsonParseSelfInject(&jr)
}

func TestDynamicCollection_UnmarshalJSON(t *testing.T) {
	dc := EmptyDynamicCollection()
	_ = dc.UnmarshalJSON([]byte(`{"Items":[]}`))
}

// === CollectionTypes ===

func TestCollectionTypes_Factories(t *testing.T) {
	_ = NewStringCollection(5)
	_ = EmptyStringCollection()
	_ = NewIntCollection(5)
	_ = EmptyIntCollection()
	_ = NewInt64Collection(5)
	_ = NewByteCollection(5)
	_ = NewBoolCollection(5)
	_ = NewFloat64Collection(5)
	_ = NewAnyMapCollection(5)
	_ = NewStringMapCollection(5)
}

// === DynamicStatus ===

func TestDynamicStatus(t *testing.T) {
	ds := InvalidDynamicStatus("test")
	if ds.IsValid() { t.Fatal("expected invalid") }
	ds2 := InvalidDynamicStatusNoMessage()
	if ds2.IsValid() { t.Fatal("expected invalid") }
	c := ds.Clone()
	if c.IsValid() { t.Fatal("expected invalid") }
	cp := ds.ClonePtr()
	if cp == nil { t.Fatal("expected non-nil") }
	var nilDS *DynamicStatus
	if nilDS.ClonePtr() != nil { t.Fatal("expected nil") }
}

// === ValueStatus ===

func TestValueStatus(t *testing.T) {
	vs := InvalidValueStatus("test")
	if vs.IsValid { t.Fatal("expected invalid") }
	vs2 := InvalidValueStatusNoMessage()
	if vs2.IsValid { t.Fatal("expected invalid") }
}

// === BytesConverter ===

func TestBytesConverter_AllMethods(t *testing.T) {
	bc := NewBytesConverter([]byte(`"hello"`))
	var s string
	err := bc.Deserialize(&s)
	if err != nil { t.Fatal("unexpected") }
	bc.DeserializeMust(&s)
	s2, err2 := bc.ToString()
	if err2 != nil || s2 != "hello" { t.Fatal("unexpected") }
	_ = bc.ToStringMust()
	_ = bc.SafeCastString()
	cs, err3 := bc.CastString()
	if err3 != nil || cs == "" { t.Fatal("unexpected") }

	bcE := NewBytesConverter(nil)
	if bcE.SafeCastString() != "" { t.Fatal("expected empty") }
	_, err4 := bcE.CastString()
	if err4 == nil { t.Fatal("expected error") }
}

func TestBytesConverter_TypeConversions(t *testing.T) {
	bc := NewBytesConverter([]byte(`true`))
	b, err := bc.ToBool()
	if err != nil || !b { t.Fatal("unexpected") }
	_ = bc.ToBoolMust()

	bc2 := NewBytesConverter([]byte(`["a","b"]`))
	strs, err2 := bc2.ToStrings()
	if err2 != nil || len(strs) != 2 { t.Fatal("unexpected") }
	_ = bc2.ToStringsMust()

	bc3 := NewBytesConverter([]byte(`42`))
	i64, err3 := bc3.ToInt64()
	if err3 != nil || i64 != 42 { t.Fatal("unexpected") }
	_ = bc3.ToInt64Must()
}

func TestBytesConverter_ComplexTypes(t *testing.T) {
	bc := NewBytesConverter([]byte(`{"a":"b"}`))
	hm, err := bc.ToHashmap()
	if err != nil || hm == nil { t.Fatal("unexpected") }
	_ = bc.ToHashmapMust()

	_, _ = bc.ToCollection()
	_, _ = bc.ToSimpleSlice()
	_, _ = bc.ToKeyValCollection()
	_, _ = bc.ToAnyCollection()
	_, _ = bc.ToMapAnyItems()
	_, _ = bc.ToDynamicCollection()
	_, _ = bc.ToJsonResultCollection()
	_, _ = bc.ToJsonMapResults()
	_, _ = bc.ToBytesCollection()
}

func TestBytesConverter_FromJsonResult(t *testing.T) {
	jr := corejson.NewResult.Any("hello")
	bc, err := NewBytesConverterUsingJsonResult(&jr)
	if err != nil || bc == nil { t.Fatal("unexpected") }
	var emptyJR *corejson.Result
	_, err2 := NewBytesConverterUsingJsonResult(emptyJR)
	if err2 == nil { t.Fatal("expected error") }
}

// === CastTo/CastedResult ===

func TestCastTo_Basic(t *testing.T) {
	r := CastTo(false, "hello", reflect.TypeOf(""))
	if r.HasError() { t.Fatal("unexpected error") }
	if r.IsInvalid() { t.Fatal("expected valid") }
	if !r.IsNotNull() { t.Fatal("expected not null") }
	if !r.IsNotPointer() { t.Fatal("expected not pointer") }
	if r.IsNotMatchingAcceptedType() { t.Fatal("expected matching") }
	if !r.IsSourceKind(reflect.String) { t.Fatal("expected string") }
	if r.HasAnyIssues() { t.Fatal("expected no issues") }
}

func TestCastTo_NoMatch(t *testing.T) {
	r := CastTo(false, "hello", reflect.TypeOf(0))
	if !r.HasError() { t.Fatal("expected error") }
}

func TestCastedResult_NilReceiver(t *testing.T) {
	var cr *CastedResult
	if !cr.IsInvalid() { t.Fatal("expected invalid") }
	if cr.IsNotNull() { t.Fatal("expected false") }
	if cr.IsNotPointer() { t.Fatal("expected false") }
	if cr.IsNotMatchingAcceptedType() { t.Fatal("expected false") }
	if cr.IsSourceKind(reflect.String) { t.Fatal("expected false") }
	if cr.HasError() { t.Fatal("expected false") }
}
