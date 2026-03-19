package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ── LinkedList ──

func Test_C33_LL_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Empty()
	if !ll.IsEmpty() { t.Fatal("expected empty") }
	if ll.HasItems() { t.Fatal("expected no items") }
}

func Test_C33_LL_Add(t *testing.T) {
	ll := corestr.New.LinkedList.Empty()
	ll.Add("a")
	ll.Add("b")
	if ll.Length() != 2 { t.Fatal("expected 2") }
}

func Test_C33_LL_Head(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	if ll.Head().Element != "a" { t.Fatal("expected a") }
}

func Test_C33_LL_Tail(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	if ll.Tail().Element != "b" { t.Fatal("expected b") }
}

func Test_C33_LL_LengthLock(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.LengthLock()
}

func Test_C33_LL_IsEquals(t *testing.T) {
	a := corestr.New.LinkedList.Strings([]string{"a"})
	b := corestr.New.LinkedList.Strings([]string{"a"})
	if !a.IsEquals(b) { t.Fatal("expected true") }
}

func Test_C33_LL_IsEqualsWithSensitive(t *testing.T) {
	a := corestr.New.LinkedList.Strings([]string{"A"})
	b := corestr.New.LinkedList.Strings([]string{"a"})
	if a.IsEqualsWithSensitive(b, true) { t.Fatal("expected false") }
	if !a.IsEqualsWithSensitive(b, false) { t.Fatal("expected true") }
}

func Test_C33_LL_IsEquals_Nil(t *testing.T) {
	var a, b *corestr.LinkedList
	if !a.IsEqualsWithSensitive(b, true) { t.Fatal("expected true") }
}

func Test_C33_LL_AddBack(t *testing.T) {
	ll := corestr.New.LinkedList.Empty()
	ll.AddBack("a")
	ll.AddBack("b")
}

func Test_C33_LL_AddLock(t *testing.T) {
	ll := corestr.New.LinkedList.Empty()
	ll.AddLock("a")
}

func Test_C33_LL_AddBackLock(t *testing.T) {
	ll := corestr.New.LinkedList.Empty()
	ll.AddBackLock("a")
}

func Test_C33_LL_AddCollection(t *testing.T) {
	ll := corestr.New.LinkedList.Empty()
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	ll.AddCollection(c)
	ll.AddCollection(corestr.New.Collection.Empty())
}

func Test_C33_LL_AddStrings(t *testing.T) {
	ll := corestr.New.LinkedList.Empty()
	ll.AddStrings([]string{"a"})
	ll.AddStrings(nil)
}

func Test_C33_LL_Adds(t *testing.T) {
	ll := corestr.New.LinkedList.Empty()
	ll.Adds("a", "b")
}

func Test_C33_LL_AddAfterNode(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	ll.AddAfterNode(ll.Head(), "x")
}

func Test_C33_LL_AddBackNode(t *testing.T) {
	ll := corestr.New.LinkedList.Empty()
	ll.AddBack("a")
	// AddBackNode tested implicitly
}

func Test_C33_LL_Remove(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
	ll.Remove("b")
	ll.Remove("a")
	ll.Remove("c")
}

func Test_C33_LL_RemoveLock(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	ll.RemoveLock("a")
}

func Test_C33_LL_RemoveAt(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
	ll.RemoveAt(1)
	ll.RemoveAt(0)
}

func Test_C33_LL_List(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	r := ll.List()
	if len(r) != 2 { t.Fatal("expected 2") }
	ll2 := corestr.New.LinkedList.Empty()
	_ = ll2.List()
}

func Test_C33_LL_ListLock(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.ListLock()
}

func Test_C33_LL_Collection(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.Collection()
}

func Test_C33_LL_SimpleSlice(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.SimpleSlice()
}

func Test_C33_LL_Loop(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
		return false
	})
}

func Test_C33_LL_LoopLock(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	ll.LoopLock(func(arg *corestr.LinkedListProcessorParameter) bool {
		return false
	})
}

func Test_C33_LL_String(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.String()
	_ = corestr.New.LinkedList.Empty().String()
}

func Test_C33_LL_StringLock(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.StringLock()
}

func Test_C33_LL_SummaryString(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.SummaryString()
	_ = ll.SummaryStringLock()
}

func Test_C33_LL_Clear(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	ll.Clear()
}

func Test_C33_LL_ClearLock(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	ll.ClearLock()
}

func Test_C33_LL_Dispose(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	ll.Dispose()
}

func Test_C33_LL_DisposeLock(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	ll.DisposeLock()
}

func Test_C33_LL_JsonMethods(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.Json()
	_ = ll.JsonPtr()
	_ = ll.JsonModel()
	_ = ll.JsonModelAny()
	_, _ = ll.MarshalJSON()
	_ = ll.AsJsonContractsBinder()
	_ = ll.AsJsoner()
	_ = ll.AsJsonMarshaller()
	_ = ll.AsJsonParseSelfInjector()
}

func Test_C33_LL_Clone(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.Clone()
	_ = corestr.New.LinkedList.Empty().Clone()
}

func Test_C33_LL_CloneLock(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.CloneLock()
}

func Test_C33_LL_Has(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	if !ll.Has("a") { t.Fatal("expected true") }
}

func Test_C33_LL_HasLock(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.HasLock("a")
}

func Test_C33_LL_IndexOf(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	_ = ll.IndexOf("b")
	_ = ll.IndexOf("z")
}

func Test_C33_LL_Hashset(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.Hashset()
}

func Test_C33_LL_FirstOrDefault(t *testing.T) {
	ll := corestr.New.LinkedList.Empty()
	_ = ll.FirstOrDefault()
	ll.Add("a")
	_ = ll.FirstOrDefault()
}

func Test_C33_LL_LastOrDefault(t *testing.T) {
	ll := corestr.New.LinkedList.Empty()
	_ = ll.LastOrDefault()
	ll.Add("a")
	_ = ll.LastOrDefault()
}

func Test_C33_LL_Serialize(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_, _ = ll.Serialize()
}

func Test_C33_LL_Deserialize(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	var target []string
	_ = ll.Deserialize(&target)
}

func Test_C33_LL_AddCollectionToNode(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	ll.AddCollectionToNode(false, ll.Head(), corestr.New.Collection.Strings([]string{"b"}))
}

func Test_C33_LL_AddStringsPtrToNode(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	ll.AddStringsPtrToNode(false, ll.Head(), []string{"b"}, false)
}

// ── newLinkedListCreator ──

func Test_C33_NLLC_Empty(t *testing.T)   { _ = corestr.New.LinkedList.Empty() }
func Test_C33_NLLC_Create(t *testing.T)  { _ = corestr.New.LinkedList.Create([]string{"a"}) }
func Test_C33_NLLC_Strings(t *testing.T) { _ = corestr.New.LinkedList.Strings([]string{"a"}) }
func Test_C33_NLLC_SpreadStrings(t *testing.T) {
	_ = corestr.New.LinkedList.SpreadStrings("a", "b")
	_ = corestr.New.LinkedList.SpreadStrings()
}
func Test_C33_NLLC_PointerStringsPtr(t *testing.T) {
	s := "a"
	_ = corestr.New.LinkedList.PointerStringsPtr(&[]*string{&s})
	_ = corestr.New.LinkedList.PointerStringsPtr(nil)
}
func Test_C33_NLLC_StringsWithCap(t *testing.T) {
	_ = corestr.New.LinkedList.StringsWithCap([]string{"a"})
}

// ── LinkedCollections ──

func Test_C33_LC_Empty(t *testing.T) {
	lc := corestr.New.LinkedCollection.Empty()
	if !lc.IsEmpty() { t.Fatal("expected empty") }
}

func Test_C33_LC_Add(t *testing.T) {
	lc := corestr.New.LinkedCollection.Empty()
	lc.Add(corestr.New.Collection.Strings([]string{"a"}))
	if lc.Length() != 1 { t.Fatal("expected 1") }
}

func Test_C33_LC_AddLock(t *testing.T) {
	lc := corestr.New.LinkedCollection.Empty()
	lc.AddLock(corestr.New.Collection.Strings([]string{"a"}))
}

func Test_C33_LC_AddAsync(t *testing.T) {
	lc := corestr.New.LinkedCollection.Empty()
	lc.AddAsync(corestr.New.Collection.Strings([]string{"a"}))
}

func Test_C33_LC_AddAnother(t *testing.T) {
	lc := corestr.New.LinkedCollection.Empty()
	other := corestr.New.LinkedCollection.Strings([]string{"a"})
	lc.AddAnother(other)
	lc.AddAnother(corestr.New.LinkedCollection.Empty())
}

func Test_C33_LC_AddStrings(t *testing.T) {
	lc := corestr.New.LinkedCollection.Empty()
	lc.AddStrings(false, []string{"a"})
	lc.AddStrings(false, nil)
}

func Test_C33_LC_AddAfterNode(t *testing.T) {
	lc := corestr.New.LinkedCollection.Empty()
	lc.Add(corestr.New.Collection.Strings([]string{"a"}))
	lc.AddAfterNode(lc.Head(), corestr.New.Collection.Strings([]string{"x"}))
}

func Test_C33_LC_AddBack(t *testing.T) {
	lc := corestr.New.LinkedCollection.Empty()
	lc.AddBack(corestr.New.Collection.Strings([]string{"a"}))
}

func Test_C33_LC_AllIndividualItemsLength(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a", "b"})
	_ = lc.AllIndividualItemsLength()
}

func Test_C33_LC_First(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_ = lc.First()
}

func Test_C33_LC_Single(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_ = lc.Single()
}

func Test_C33_LC_Last(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_ = lc.Last()
}

func Test_C33_LC_FirstOrDefault(t *testing.T) {
	lc := corestr.New.LinkedCollection.Empty()
	_ = lc.FirstOrDefault()
	lc.Add(corestr.New.Collection.Strings([]string{"a"}))
	_ = lc.FirstOrDefault()
}

func Test_C33_LC_LastOrDefault(t *testing.T) {
	lc := corestr.New.LinkedCollection.Empty()
	_ = lc.LastOrDefault()
	lc.Add(corestr.New.Collection.Strings([]string{"a"}))
	_ = lc.LastOrDefault()
}

func Test_C33_LC_Remove(t *testing.T) {
	lc := corestr.New.LinkedCollection.Empty()
	lc.Add(corestr.New.Collection.Strings([]string{"a"}))
	lc.Add(corestr.New.Collection.Strings([]string{"b"}))
	lc.RemoveAt(0)
}

func Test_C33_LC_List(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_ = lc.List()
}

func Test_C33_LC_ListFlat(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_ = lc.ListFlat()
}

func Test_C33_LC_Loop(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
		return false
	})
}

func Test_C33_LC_String(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_ = lc.String()
	_ = corestr.New.LinkedCollection.Empty().String()
}

func Test_C33_LC_Clear(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	lc.Clear()
}

func Test_C33_LC_Dispose(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	lc.Dispose()
}

func Test_C33_LC_Clone(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_ = lc.Clone()
}

func Test_C33_LC_JsonMethods(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_ = lc.Json()
	_ = lc.JsonPtr()
	_ = lc.JsonModel()
	_ = lc.JsonModelAny()
	_, _ = lc.MarshalJSON()
	_ = lc.AsJsonContractsBinder()
	_ = lc.AsJsoner()
	_ = lc.AsJsonMarshaller()
	_ = lc.AsJsonParseSelfInjector()
}

func Test_C33_LC_Has(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_ = lc.Has("a")
}

func Test_C33_LC_Hashset(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_ = lc.Hashset()
}

func Test_C33_LC_Serialize(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_, _ = lc.Serialize()
}

func Test_C33_LC_AddCollectionToNode(t *testing.T) {
	lc := corestr.New.LinkedCollection.Empty()
	lc.Add(corestr.New.Collection.Strings([]string{"a"}))
	lc.AddCollectionToNode(false, lc.Head(), corestr.New.Collection.Strings([]string{"x"}))
}

// ── newLinkedListCollectionsCreator ──

func Test_C33_NLLCC_Empty(t *testing.T)  { _ = corestr.New.LinkedCollection.Empty() }
func Test_C33_NLLCC_Create(t *testing.T) { _ = corestr.New.LinkedCollection.Create([]string{"a"}) }
func Test_C33_NLLCC_Strings(t *testing.T) {
	_ = corestr.New.LinkedCollection.Strings([]string{"a"})
}
func Test_C33_NLLCC_PointerStringsPtr(t *testing.T) {
	s := "a"
	_ = corestr.New.LinkedCollection.PointerStringsPtr(&[]*string{&s})
	_ = corestr.New.LinkedCollection.PointerStringsPtr(nil)
}
func Test_C33_NLLCC_UsingCollections(t *testing.T) {
	_ = corestr.New.LinkedCollection.UsingCollections(
		corestr.New.Collection.Strings([]string{"a"}),
	)
	_ = corestr.New.LinkedCollection.UsingCollections()
}

// ── LinkedListNode ──

func Test_C33_LLN_HasNext(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	if !ll.Head().HasNext() { t.Fatal("expected true") }
}

func Test_C33_LLN_Next(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	_ = ll.Head().Next()
}

func Test_C33_LLN_EndOfChain(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	_, l := ll.Head().EndOfChain()
	if l != 2 { t.Fatal("expected 2") }
}

func Test_C33_LLN_LoopEndOfChain(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	_, _ = ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
		return false
	})
}

func Test_C33_LLN_Clone(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.Head().Clone()
}

func Test_C33_LLN_String(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	_ = ll.Head().String()
}

func Test_C33_LLN_AddNext(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	ll.Head().AddNext(ll, "x")
}

func Test_C33_LLN_AddNextNode(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	ll.Head().AddNextNode(ll, &corestr.LinkedListNode{Element: "x"})
}

func Test_C33_LLN_AddStringsToNode(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	ll.Head().AddStringsToNode(ll, false, []string{"b"}, false)
}

func Test_C33_LLN_AddStringsPtrToNode(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	ll.Head().AddStringsPtrToNode(ll, false, []string{"b"}, false)
}

func Test_C33_LLN_AddCollectionToNode(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a"})
	ll.Head().AddCollectionToNode(ll, false, corestr.New.Collection.Strings([]string{"b"}))
}

// ── LinkedCollectionNode ──

func Test_C33_LCN_IsEmpty(t *testing.T) {
	node := &corestr.LinkedCollectionNode{}
	_ = node.IsEmpty()
}

func Test_C33_LCN_HasElement(t *testing.T) {
	node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
	_ = node.HasElement()
}

func Test_C33_LCN_AddNext(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	lc.Head().AddNext(lc, corestr.New.Collection.Strings([]string{"x"}))
}

func Test_C33_LCN_AddNextNode(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	newNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"x"})}
	lc.Head().AddNextNode(lc, newNode)
}

func Test_C33_LCN_AddStringsToNode(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	lc.Head().AddStringsToNode(lc, false, []string{"b"}, false)
}

func Test_C33_LCN_AddCollectionToNode(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	lc.Head().AddCollectionToNode(lc, false, corestr.New.Collection.Strings([]string{"b"}))
}

func Test_C33_LCN_Clone(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_ = lc.Head().Clone()
}

func Test_C33_LCN_String(t *testing.T) {
	lc := corestr.New.LinkedCollection.Strings([]string{"a"})
	_ = lc.Head().String()
}
