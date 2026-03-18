package corestr

import (
	"testing"
)

// ── SimpleStringOnce ──

func TestSimpleStringOnce_Core(t *testing.T) {
	s := New.SimpleStringOnce.Init("hello")
	if s.Value() != "hello" || !s.IsInitialized() || !s.IsDefined() { t.Fatal("unexpected") }
	if s.IsUninitialized() || s.IsInvalid() { t.Fatal("unexpected") }
	if s.SafeValue() != "hello" { t.Fatal("unexpected") }
	_ = s.ValueBytes()
	_ = s.ValueBytesPtr()
	if s.IsEmpty() || s.IsWhitespace() { t.Fatal("unexpected") }
	if s.Trim() != "hello" { t.Fatal("unexpected") }
	if !s.HasValidNonEmpty() || !s.HasValidNonWhitespace() || !s.HasSafeNonEmpty() { t.Fatal("unexpected") }
	if !s.Is("hello") || s.Is("world") { t.Fatal("unexpected") }
	if !s.IsAnyOf("hello") || s.IsAnyOf("x") { t.Fatal("unexpected") }
	if !s.IsContains("hel") || !s.IsAnyContains("hel") { t.Fatal("unexpected") }
	if !s.IsEqualNonSensitive("HELLO") { t.Fatal("unexpected") }
}

func TestSimpleStringOnce_Set(t *testing.T) {
	s := New.SimpleStringOnce.Uninitialized("")
	err := s.SetOnUninitialized("val")
	if err != nil { t.Fatal(err) }
	err2 := s.SetOnUninitialized("val2")
	if err2 == nil { t.Fatal("expected error") }
}

func TestSimpleStringOnce_GetSetOnce(t *testing.T) {
	s := New.SimpleStringOnce.Uninitialized("")
	v := s.GetSetOnce("first")
	if v != "first" { t.Fatal("unexpected") }
	v2 := s.GetSetOnce("second")
	if v2 != "first" { t.Fatal("expected first") }
}

func TestSimpleStringOnce_GetOnce(t *testing.T) {
	s := New.SimpleStringOnce.Uninitialized("")
	v := s.GetOnce()
	if v != "" { t.Fatal("expected empty") }
}

func TestSimpleStringOnce_GetOnceFunc(t *testing.T) {
	s := New.SimpleStringOnce.Uninitialized("")
	v := s.GetOnceFunc(func() string { return "computed" })
	if v != "computed" { t.Fatal("unexpected") }
	v2 := s.GetOnceFunc(func() string { return "other" })
	if v2 != "computed" { t.Fatal("expected cached") }
}

func TestSimpleStringOnce_SetOnceIfUninitialized(t *testing.T) {
	s := New.SimpleStringOnce.Uninitialized("")
	ok := s.SetOnceIfUninitialized("val")
	if !ok { t.Fatal("expected set") }
	ok2 := s.SetOnceIfUninitialized("val2")
	if ok2 { t.Fatal("expected not set") }
}

func TestSimpleStringOnce_Invalidate(t *testing.T) {
	s := New.SimpleStringOnce.Init("hello")
	s.Invalidate()
	if s.IsInitialized() { t.Fatal("expected uninit") }
	s2 := New.SimpleStringOnce.Init("world")
	s2.Reset()
	if s2.IsInitialized() { t.Fatal("expected uninit") }
}

func TestSimpleStringOnce_Conversions(t *testing.T) {
	s := New.SimpleStringOnce.Init("42")
	if s.Int() != 42 { t.Fatal("expected 42") }
	if s.ValueInt(0) != 42 || s.ValueDefInt() != 42 { t.Fatal("unexpected") }
	if s.Byte() != 42 { t.Fatal("unexpected") }
	if s.ValueByte(0) != 42 || s.ValueDefByte() != 42 { t.Fatal("unexpected") }
	s2 := New.SimpleStringOnce.Init("3.14")
	if s2.ValueFloat64(0) == 0 || s2.ValueDefFloat64() == 0 { t.Fatal("unexpected") }
	s3 := New.SimpleStringOnce.Init("true")
	if !s3.Boolean(false) || !s3.BooleanDefault() || !s3.IsValueBool() { t.Fatal("unexpected") }
	s4 := New.SimpleStringOnce.Init("yes")
	if !s4.Boolean(false) { t.Fatal("expected true") }
	s5 := New.SimpleStringOnce.Init("abc")
	if s5.Int() != 0 { t.Fatal("expected 0") }
	_ = s5.Byte()
	_ = s5.Int16()
	_ = s5.Int32()
	_ = s5.ValueByte(0)
	_ = s5.IsSetter(false)
	_ = s5.IsSetter(true)
}

func TestSimpleStringOnce_WithinRange(t *testing.T) {
	s := New.SimpleStringOnce.Init("50")
	v, ok := s.WithinRange(true, 0, 100)
	if !ok || v != 50 { t.Fatal("unexpected") }
	v2, ok2 := s.WithinRangeDefault(0, 100)
	if !ok2 || v2 != 50 { t.Fatal("unexpected") }
	_, _ = s.Uint16()
	_, _ = s.Uint32()
}

func TestSimpleStringOnce_ConcatNew(t *testing.T) {
	s := New.SimpleStringOnce.Init("hello")
	c := s.ConcatNew(" world")
	if c.Value() != "hello world" { t.Fatal("unexpected") }
	c2 := s.ConcatNewUsingStrings(" ", "beautiful", "world")
	_ = c2
}

func TestSimpleStringOnce_Split(t *testing.T) {
	s := New.SimpleStringOnce.Init("a,b,c")
	_ = s.Split(",")
	_ = s.SplitNonEmpty(",")
	_ = s.SplitTrimNonWhitespace(",")
	l, r := s.SplitLeftRight(",")
	_ = l
	_ = r
	l2, r2 := s.SplitLeftRightTrim(",")
	_ = l2
	_ = r2
}

func TestSimpleStringOnce_Various(t *testing.T) {
	s := New.SimpleStringOnce.Init("hello")
	_ = s.LinesSimpleSlice()
	_ = s.SimpleSlice(",")
	_ = s.IsRegexMatches(nil)
	_ = s.RegexFindString(nil)
	_, _ = s.RegexFindAllStringsWithFlag(nil, -1)
	_ = s.RegexFindAllStrings(nil, -1)
	_ = s.NonPtr()
	_ = s.Ptr()
	_ = s.String()
	_ = s.StringPtr()
	_ = s.Clone()
	_ = s.ClonePtr()
	_ = s.CloneUsingNewVal("new")
	s.Dispose()
	var nilS *SimpleStringOnce
	if nilS.String() != "" { t.Fatal("expected empty") }
	if nilS.StringPtr() == nil { t.Fatal("expected non-nil") }
	if nilS.ClonePtr() != nil { t.Fatal("expected nil") }
}

func TestSimpleStringOnce_Json(t *testing.T) {
	s := New.SimpleStringOnce.Init("hello")
	_ = s.JsonModel()
	_ = s.JsonModelAny()
	_, _ = s.MarshalJSON()
	_, _ = s.Serialize()
	_ = s.AsJsoner()
	_ = s.AsJsonContractsBinder()
	_ = s.AsJsonParseSelfInjector()
	_ = s.AsJsonMarshaller()
}

// ── KeyValueCollection ──

func TestKeyValueCollection_Core(t *testing.T) {
	kv := New.KeyValues.Empty()
	if !kv.IsEmpty() || kv.HasAnyItem() { t.Fatal("expected empty") }
	kv.Add("k1", "v1").Add("k2", "v2")
	if kv.Length() != 2 || kv.Count() != 2 { t.Fatal("expected 2") }
	if kv.LastIndex() != 1 { t.Fatal("expected 1") }
	if !kv.HasIndex(0) || kv.HasIndex(5) { t.Fatal("unexpected") }
	_ = kv.First()
	_ = kv.Last()
	_ = kv.FirstOrDefault()
	_ = kv.LastOrDefault()
	_ = kv.Strings()
	_ = kv.String()
	_ = kv.AllKeys()
	_ = kv.AllKeysSorted()
	_ = kv.AllValues()
	_ = kv.Join(",")
	_ = kv.JoinKeys(",")
	_ = kv.JoinValues(",")
	_ = kv.Compile()
	if !kv.HasKey("k1") { t.Fatal("expected") }
	if !kv.IsContains("k1") { t.Fatal("expected") }
	v, found := kv.Get("k1")
	if !found || v != "v1" { t.Fatal("unexpected") }
}

func TestKeyValueCollection_AddVariants(t *testing.T) {
	kv := New.KeyValues.Empty()
	kv.AddIf(false, "skip", "val")
	kv.AddIf(true, "k", "v")
	kv.Adds(KeyValuePair{Key: "a", Value: "b"})
	kv.Adds()
	kv.AddMap(map[string]string{"c": "d"})
	kv.AddMap(nil)
	kv.AddHashsetMap(map[string]bool{"e": true})
	kv.AddHashsetMap(nil)
	kv.AddHashset(New.Hashset.StringsSpreadItems("f"))
	kv.AddHashset(nil)
	kv.AddsHashmap(New.Hashmap.UsingMap(map[string]string{"g": "h"}))
	kv.AddsHashmap(nil)
	kv.AddsHashmaps(New.Hashmap.UsingMap(map[string]string{"i": "j"}))
	kv.AddsHashmaps()
	kv.AddStringBySplit("=", "k=l")
	kv.AddStringBySplitTrim("=", " m = n ")
}

func TestKeyValueCollection_Find(t *testing.T) {
	kv := New.KeyValues.Empty()
	kv.Add("a", "1").Add("b", "2")
	r := kv.Find(func(i int, curr KeyValuePair) (KeyValuePair, bool, bool) {
		return curr, curr.Key == "a", false
	})
	if len(r) != 1 { t.Fatal("expected 1") }
}

func TestKeyValueCollection_Safe(t *testing.T) {
	kv := New.KeyValues.Empty()
	kv.Add("a", "1")
	if kv.SafeValueAt(0) != "1" || kv.SafeValueAt(99) != "" { t.Fatal("unexpected") }
	_ = kv.SafeValuesAtIndexes(0)
	_ = kv.StringsUsingFormat("%s=%s")
	_ = kv.Hashmap()
	_ = kv.Map()
}

func TestKeyValueCollection_Json(t *testing.T) {
	kv := New.KeyValues.Empty()
	kv.Add("a", "1")
	_ = kv.JsonModel()
	_ = kv.JsonModelAny()
	_, _ = kv.Serialize()
	_, _ = kv.MarshalJSON()
	_ = kv.SerializeMust()
	_ = kv.AsJsoner()
	_ = kv.AsJsonContractsBinder()
	_ = kv.AsJsonParseSelfInjector()
}

func TestKeyValueCollection_ClearDispose(t *testing.T) {
	kv := New.KeyValues.Empty()
	kv.Add("a", "1")
	kv.Clear()
	kv.Dispose()
	var nilKv *KeyValueCollection
	nilKv.Clear()
	nilKv.Dispose()
}

// ── NonChainedLinkedListNodes ──

func TestNonChainedLinkedListNodes(t *testing.T) {
	nc := NewNonChainedLinkedListNodes(5)
	if !nc.IsEmpty() || nc.HasItems() { t.Fatal("expected empty") }
	n1 := &LinkedListNode{Element: "a"}
	n2 := &LinkedListNode{Element: "b"}
	nc.Adds(n1, n2)
	if nc.Length() != 2 { t.Fatal("expected 2") }
	if nc.First().Element != "a" || nc.Last().Element != "b" { t.Fatal("unexpected") }
	_ = nc.FirstOrDefault()
	_ = nc.LastOrDefault()
	_ = nc.Items()
	if nc.IsChainingApplied() { t.Fatal("expected false") }
	nc.ApplyChaining()
	if !nc.IsChainingApplied() { t.Fatal("expected true") }
	_ = nc.ToChainedNodes()
}

// ── NonChainedLinkedCollectionNodes ──

func TestNonChainedLinkedCollectionNodes(t *testing.T) {
	nc := NewNonChainedLinkedCollectionNodes(5)
	if !nc.IsEmpty() || nc.HasItems() { t.Fatal("expected empty") }
	c1 := New.Collection.Strings([]string{"a"})
	c2 := New.Collection.Strings([]string{"b"})
	n1 := &LinkedCollectionNode{Element: c1}
	n2 := &LinkedCollectionNode{Element: c2}
	nc.Adds(n1, n2)
	if nc.Length() != 2 { t.Fatal("expected 2") }
	_ = nc.First()
	_ = nc.Last()
	_ = nc.FirstOrDefault()
	_ = nc.LastOrDefault()
	_ = nc.Items()
	nc.ApplyChaining()
	if !nc.IsChainingApplied() { t.Fatal("expected true") }
	_ = nc.ToChainedNodes()
}

// ── CollectionsOfCollection remaining ──

func TestCollectionsOfCollection_Methods(t *testing.T) {
	coc := New.CollectionsOfCollection.Empty()
	coc.Add(New.Collection.Strings([]string{"a", "b"}))
	coc.Add(New.Collection.Strings([]string{"c"}))
	if coc.Length() != 2 { t.Fatal("expected 2") }
	if coc.AllIndividualItemsLength() != 3 { t.Fatal("expected 3") }
	list := coc.List(0)
	if len(list) != 3 { t.Fatal("expected 3") }
	_ = coc.ToCollection()
	_ = coc.Items()
	_ = coc.String()
	_ = coc.JsonModel()
	_ = coc.JsonModelAny()
	_, _ = coc.MarshalJSON()
	_ = coc.AsJsoner()
	_ = coc.AsJsonContractsBinder()
	_ = coc.AsJsonParseSelfInjector()
	_ = coc.AsJsonMarshaller()
}

// ── HashsetsCollection remaining ──

func TestHashsetsCollection_Methods(t *testing.T) {
	hc := New.HashsetsCollection.Empty()
	h1 := New.Hashset.StringsSpreadItems("a")
	h2 := New.Hashset.StringsSpreadItems("b")
	hc.Add(h1)
	hc.AddNonNil(h2)
	hc.AddNonNil(nil)
	hc.AddNonEmpty(New.Hashset.Empty())
	hc.Adds(New.Hashset.StringsSpreadItems("c"))
	if hc.Length() < 3 { t.Fatal("expected >= 3") }
	_ = hc.LastIndex()
	_ = hc.List()
	_ = hc.ListPtr()
	_ = hc.ListDirectPtr()
	_ = hc.StringsList()
	_ = hc.String()
	_ = hc.Join(",")
	_ = hc.IsEqual(*hc)
	_ = hc.IsEqualPtr(hc)
	_ = hc.JsonModel()
	_ = hc.JsonModelAny()
	_, _ = hc.MarshalJSON()
	_, _ = hc.Serialize()
	_ = hc.AsJsoner()
	_ = hc.AsJsonContractsBinder()
	_ = hc.AsJsonParseSelfInjector()
	_ = hc.AsJsonMarshaller()
}

func TestHashsetsCollection_HasAll(t *testing.T) {
	hc := New.HashsetsCollection.Empty()
	h := New.Hashset.StringsSpreadItems("a", "b")
	hc.Add(h)
	if !hc.HasAll("a", "b") { t.Fatal("expected true") }
	if hc.HasAll("z") { t.Fatal("expected false") }
	empty := New.HashsetsCollection.Empty()
	if empty.HasAll("a") { t.Fatal("expected false") }
}

func TestHashsetsCollection_ConcatNew(t *testing.T) {
	hc := New.HashsetsCollection.Empty()
	hc.Add(New.Hashset.StringsSpreadItems("a"))
	hc2 := New.HashsetsCollection.Empty()
	hc2.Add(New.Hashset.StringsSpreadItems("b"))
	r := hc.ConcatNew(hc2)
	_ = r
	r2 := hc.ConcatNew()
	_ = r2
	hc.AddHashsetsCollection(hc2)
	hc.AddHashsetsCollection(nil)
}

// ── CharCollectionMap remaining ──

func TestCharCollectionMap_Methods(t *testing.T) {
	ccm := New.CharCollectionMap.Items([]string{"abc", "adef", "bcd"})
	if ccm.IsEmpty() { t.Fatal("expected non-empty") }
	if ccm.Length() == 0 { t.Fatal("expected > 0") }
	_ = ccm.AllLengthsSum()
	_ = ccm.AllLengthsSumLock()
	_ = ccm.LengthLock()
	_ = ccm.GetMap()
	_ = ccm.GetCopyMapLock()
	_ = ccm.List()
	_ = ccm.ListLock()
	_ = ccm.SortedListAsc()
	_ = ccm.String()
	_ = ccm.StringLock()
	_ = ccm.SummaryString()
	_ = ccm.SummaryStringLock()
	if ccm.GetChar("") != 0 { t.Fatal("expected 0") }
	if ccm.GetChar("a") != 'a' { t.Fatal("expected a") }
	_ = ccm.LengthOf('a')
	_ = ccm.LengthOfLock('a')
	_ = ccm.LengthOfCollectionFromFirstChar("abc")
	_ = ccm.Has("abc")
	_ = ccm.Has("zzz")
	_, _ = ccm.HasWithCollection("abc")
	_, _ = ccm.HasWithCollectionLock("abc")
	_ = ccm.GetCollection("a", true)
	_ = ccm.GetCollectionLock("a", false)
	_ = ccm.GetCollectionByChar('a')
	_ = ccm.HashsetByChar('a')
	_ = ccm.HashsetByCharLock('a')
	_ = ccm.HashsetByStringFirstChar("abc")
	_ = ccm.HashsetByStringFirstCharLock("abc")
	_ = ccm.HashsetsCollection()
	_ = ccm.HashsetsCollectionByChars('a')
	_ = ccm.HashsetsCollectionByStringFirstChar("abc")
	_ = ccm.IsEquals(ccm)
	_ = ccm.IsEqualsLock(ccm)
	_ = ccm.IsEqualsCaseSensitive(true, ccm)
	_ = ccm.IsEqualsCaseSensitiveLock(false, ccm)
}

func TestCharCollectionMap_Add(t *testing.T) {
	ccm := New.CharCollectionMap.Empty()
	ccm.Add("hello")
	ccm.AddLock("world")
	ccm.AddStrings("a", "b")
	ccm.AddStrings()
	ccm.AddCollectionItems(New.Collection.Strings([]string{"c"}))
	ccm.AddCollectionItems(nil)
}

func TestCharCollectionMap_ClearDispose(t *testing.T) {
	ccm := New.CharCollectionMap.Items([]string{"a"})
	ccm.Clear()
	ccm.Dispose()
	var nilCcm *CharCollectionMap
	nilCcm.Dispose()
}

// ── CharHashsetMap remaining ──

func TestCharHashsetMap_Methods(t *testing.T) {
	chm := New.CharHashsetMap.CapItems(20, 20, "abc", "adef", "bcd")
	if chm.IsEmpty() { t.Fatal("expected non-empty") }
	_ = chm.Length()
	_ = chm.LengthLock()
	_ = chm.AllLengthsSum()
	_ = chm.AllLengthsSumLock()
	_ = chm.GetMap()
	_ = chm.GetCopyMapLock()
	_ = chm.List()
	_ = chm.SortedListAsc()
	_ = chm.SortedListDsc()
	_ = chm.String()
	_ = chm.StringLock()
	_ = chm.SummaryString()
	_ = chm.SummaryStringLock()
	_ = chm.GetCharOf("")
	_ = chm.GetCharOf("a")
	_ = chm.LengthOf('a')
	_ = chm.LengthOfLock('a')
	_ = chm.LengthOfHashsetFromFirstChar("abc")
	_ = chm.Has("abc")
	_, _ = chm.HasWithHashset("abc")
	_, _ = chm.HasWithHashsetLock("abc")
	_ = chm.GetHashset("a", true)
	_ = chm.GetHashsetLock(true, "a")
	_ = chm.GetHashsetByChar('a')
	_ = chm.HashsetByChar('a')
	_ = chm.HashsetByCharLock('a')
	_ = chm.HashsetByStringFirstChar("abc")
	_ = chm.HashsetByStringFirstCharLock("abc")
	_ = chm.HashsetsCollection()
	_ = chm.HashsetsCollectionByChars('a')
	_ = chm.HashsetsCollectionByStringsFirstChar("abc")
	_ = chm.IsEquals(chm)
	_ = chm.IsEqualsLock(chm)
}

func TestCharHashsetMap_Add(t *testing.T) {
	chm := New.CharHashsetMap.Cap(20, 20)
	chm.Add("hello")
	chm.AddLock("world")
	chm.AddStrings("a", "b")
	chm.AddStrings()
	chm.AddStringsLock("c")
	chm.AddStringsLock()
	chm.AddCollectionItems(New.Collection.Strings([]string{"d"}))
	chm.AddCollectionItems(nil)
	chm.AddHashsetItems(New.Hashset.StringsSpreadItems("e"))
	chm.AddCharCollectionMapItems(New.CharCollectionMap.Items([]string{"f"}))
	chm.AddCharCollectionMapItems(nil)
}

func TestCharHashsetMap_ClearDispose(t *testing.T) {
	chm := New.CharHashsetMap.CapItems(20, 20, "a")
	chm.Clear()
	chm.RemoveAll()
}

// ── LinkedCollections remaining ──

func TestLinkedCollections_Basic(t *testing.T) {
	lc := New.LinkedCollection.Create()
	if !lc.IsEmpty() || lc.HasItems() { t.Fatal("expected empty") }
	c := New.Collection.Strings([]string{"a"})
	lc.Add(c)
	if lc.Length() != 1 { t.Fatal("expected 1") }
	_ = lc.Head()
	_ = lc.Tail()
	_ = lc.First()
	_ = lc.Last()
	_ = lc.FirstOrDefault()
	_ = lc.LastOrDefault()
	_ = lc.AllIndividualItemsLength()
	_ = lc.ToStrings()
	_ = lc.ToCollectionSimple()
	_ = lc.ToCollection(0)
	_ = lc.ToCollectionsOfCollection(0)
	_ = lc.ItemsOfItems()
	_ = lc.ItemsOfItemsCollection()
	_ = lc.SimpleSlice()
}

func TestLinkedCollections_AddVariants(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddStrings("a", "b")
	lc.AddStrings()
	lc.AddStringsLock("c")
	lc.AddStringsLock()
	lc.AddCollection(New.Collection.Strings([]string{"d"}))
	lc.AddCollection(nil)
	lc.AddLock(New.Collection.Strings([]string{"e"}))
	lc.Push(New.Collection.Strings([]string{"f"}))
	lc.PushBack(New.Collection.Strings([]string{"g"}))
	lc.PushBackLock(New.Collection.Strings([]string{"h"}))
	lc.PushFront(New.Collection.Strings([]string{"i"}))
	lc.AddFront(New.Collection.Strings([]string{"j"}))
	lc.AddFrontLock(New.Collection.Strings([]string{"k"}))
}

func TestLinkedCollections_Loop(t *testing.T) {
	lc := New.LinkedCollection.Strings("a", "b")
	count := 0
	lc.Loop(func(arg *LinkedCollectionProcessorParameter) bool {
		count++
		return false
	})
	if count == 0 { t.Fatal("expected > 0") }
}

func TestLinkedCollections_IsEquals(t *testing.T) {
	lc1 := New.LinkedCollection.Strings("a")
	lc2 := New.LinkedCollection.Strings("a")
	if !lc1.IsEqualsPtr(lc2) { t.Fatal("expected equal") }
}

// ── Collection remaining methods ──

func TestCollection_RemainingMethods(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	_ = c.First()
	_ = c.Last()
	_ = c.FirstOrDefault()
	_ = c.LastOrDefault()
	_ = c.IndexAt(0)
	_ = c.SafeIndexAtUsingLength("def", 3, 0)
	_ = c.List()
	_ = c.HasItems()
	_ = c.Reverse()
	_ = c.GetPagesSize(2)
	_ = c.Take(2)
	_ = c.Skip(0)
	_ = c.UniqueList()
	_ = c.UniqueListLock()
	_ = c.UniqueBoolMap()
	_ = c.UniqueBoolMapLock()
}

func TestCollection_Filter(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "bb", "ccc"})
	f := func(s string, i int) (string, bool, bool) { return s, len(s) > 1, false }
	r := c.Filter(f)
	if len(r) != 2 { t.Fatal("expected 2") }
	_ = c.FilteredCollection(f)
	_ = c.FilterLock(f)
	_ = c.FilteredCollectionLock(f)
}

func TestCollection_AppendAnys(t *testing.T) {
	c := New.Collection.Empty()
	c.AppendAnys("a", 42, nil)
	c.AppendAnys()
	c.AppendAnysLock("b")
	c.AppendAnysLock()
	c.AppendNonEmptyAnys("c", nil)
	c.AppendNonEmptyAnys()
	c.AddsNonEmpty("d", "", "e")
	c.AddsNonEmpty()
}

// ── SimpleSlice remaining ──

func TestSimpleSlice_Remaining(t *testing.T) {
	s := New.SimpleSlice.Lines("a", "b", "c")
	_ = s.Join(",")
	_ = s.JoinLine()
	_ = s.JoinLineEofLine()
	_ = s.JoinSpace()
	_ = s.JoinComma()
	_ = s.JoinCsv()
	_ = s.JoinCsvLine()
	_ = s.JoinWith(",")
	_ = s.JoinCsvString(",")
	_ = s.CsvStrings()
	_ = s.String()
	_ = s.Collection(false)
	_ = s.ToCollection(false)
	_ = s.NonPtr()
	_ = s.Ptr()
	_ = s.ToPtr()
	_ = s.ToNonPtr()
	_ = s.Sort()
	_ = s.Reverse()
	_ = s.Hashset()
	_ = s.EachItemSplitBy(",")
	_ = s.TranspileJoin(func(ss string) string { return ss }, ",")
	_ = s.PrependJoin(",", "z")
	_ = s.AppendJoin(",", "z")
	_ = s.ConcatNew("d")
	_ = s.ConcatNewStrings("d")
	_ = s.ConcatNewSimpleSlices(New.SimpleSlice.Lines("e"))
	s.PrependAppend([]string{"pre"}, []string{"post"})
	_ = s.JsonModel()
	_ = s.JsonModelAny()
	_, _ = s.MarshalJSON()
	_, _ = s.Serialize()
	_ = s.SafeStrings()
	_ = s.AsJsoner()
	_ = s.AsJsonContractsBinder()
	_ = s.AsJsonParseSelfInjector()
	_ = s.AsJsonMarshaller()
}

func TestSimpleSlice_IsEqual(t *testing.T) {
	s1 := New.SimpleSlice.Lines("a", "b")
	s2 := New.SimpleSlice.Lines("a", "b")
	if !s1.IsEqual(s2) { t.Fatal("expected equal") }
	if !s1.IsEqualLines([]string{"a", "b"}) { t.Fatal("expected equal") }
	if !s1.IsEqualUnorderedLines([]string{"b", "a"}) { t.Fatal("expected equal") }
	if !s1.IsEqualUnorderedLinesClone([]string{"b", "a"}) { t.Fatal("expected equal") }
	if !s1.IsDistinctEqual(s2) { t.Fatal("expected equal") }
	if !s1.IsDistinctEqualRaw("a", "b") { t.Fatal("expected equal") }
	if !s1.IsUnorderedEqual(true, s2) { t.Fatal("expected equal") }
	if !s1.IsUnorderedEqualRaw(true, "b", "a") { t.Fatal("expected equal") }
}

func TestSimpleSlice_DistinctDiff(t *testing.T) {
	s := New.SimpleSlice.Lines("a", "b")
	_ = s.DistinctDiffRaw("b", "c")
	_ = s.DistinctDiff(New.SimpleSlice.Lines("b", "c"))
	_ = s.AddedRemovedLinesDiff("b", "c")
}

func TestSimpleSlice_RemoveIndexes(t *testing.T) {
	s := New.SimpleSlice.Lines("a", "b", "c")
	newS, err := s.RemoveIndexes(1)
	if err != nil || newS.Length() != 2 { t.Fatal("unexpected") }
	empty := New.SimpleSlice.Empty()
	_, err2 := empty.RemoveIndexes(0)
	if err2 == nil { t.Fatal("expected error") }
}

func TestSimpleSlice_Clone(t *testing.T) {
	s := New.SimpleSlice.Lines("a", "b")
	_ = s.Clone(true)
	_ = s.ClonePtr(true)
	_ = s.DeepClone()
	_ = s.ShadowClone()
	var nilS *SimpleSlice
	if nilS.ClonePtr(true) != nil { t.Fatal("expected nil") }
}

func TestSimpleSlice_ClearDispose(t *testing.T) {
	s := New.SimpleSlice.Lines("a")
	s.Clear()
	s.Dispose()
	var nilS *SimpleSlice
	if nilS.Clear() != nil { t.Fatal("expected nil") }
}
