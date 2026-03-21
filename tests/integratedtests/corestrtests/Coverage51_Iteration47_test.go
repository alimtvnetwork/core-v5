package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ═══════════════════════════════════════════════════════════════
// LinkedListNode
// ═══════════════════════════════════════════════════════════════

func Test_Cov51_LinkedListNode_HasNext_False(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	tc := coretestcases.CaseV1{Name: "Node HasNext false", Expected: false, Actual: ll.Head().HasNext(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_HasNext_True(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "Node HasNext true", Expected: true, Actual: ll.Head().HasNext(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_Next(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "Node Next", Expected: "b", Actual: ll.Head().Next().Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_EndOfChain(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b").Add("c")
	end, length := ll.Head().EndOfChain()
	tc := coretestcases.CaseV1{Name: "Node EndOfChain elem", Expected: "c", Actual: end.Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "Node EndOfChain len", Expected: 3, Actual: length, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_Clone(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("x")
	clone := ll.Head().Clone()
	tc := coretestcases.CaseV1{Name: "Node Clone", Expected: "x", Actual: clone.Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "Node Clone no next", Expected: false, Actual: clone.HasNext(), Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_IsEqual_Same(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	tc := coretestcases.CaseV1{Name: "Node IsEqual same", Expected: true, Actual: ll.Head().IsEqual(ll.Head()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_IsEqual_BothNil(t *testing.T) {
	var n *corestr.LinkedListNode
	tc := coretestcases.CaseV1{Name: "Node IsEqual both nil", Expected: true, Actual: n.IsEqual(nil), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_IsEqualValue(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("hello")
	tc := coretestcases.CaseV1{Name: "Node IsEqualValue", Expected: true, Actual: ll.Head().IsEqualValue("hello"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_IsEqualValueSensitive_CI(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("Hello")
	tc := coretestcases.CaseV1{Name: "Node IsEqualValueSensitive CI", Expected: true, Actual: ll.Head().IsEqualValueSensitive("hello", false), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_String(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("test")
	tc := coretestcases.CaseV1{Name: "Node String", Expected: "test", Actual: ll.Head().String(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_List(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "Node List", Expected: 2, Actual: len(ll.Head().List()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_Join(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "Node Join", Expected: "a,b", Actual: ll.Head().Join(","), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_CreateLinkedList(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("x").Add("y")
	newLL := ll.Head().CreateLinkedList()
	tc := coretestcases.CaseV1{Name: "Node CreateLinkedList", Expected: 2, Actual: newLL.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_IsChainEqual_True(t *testing.T) {
	ll1 := corestr.Empty.LinkedList()
	ll1.Add("a").Add("b")
	ll2 := corestr.Empty.LinkedList()
	ll2.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "Node IsChainEqual true", Expected: true, Actual: ll1.Head().IsChainEqual(ll2.Head(), true), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedListNode_IsChainEqual_CaseInsensitive(t *testing.T) {
	ll1 := corestr.Empty.LinkedList()
	ll1.Add("A").Add("B")
	ll2 := corestr.Empty.LinkedList()
	ll2.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "Node IsChainEqual CI", Expected: true, Actual: ll1.Head().IsChainEqual(ll2.Head(), false), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// LinkedList
// ═══════════════════════════════════════════════════════════════

func Test_Cov51_LinkedList_Add(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	tc := coretestcases.CaseV1{Name: "LL Add", Expected: 1, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_Head(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("first")
	tc := coretestcases.CaseV1{Name: "LL Head", Expected: "first", Actual: ll.Head().Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_Tail(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "LL Tail", Expected: "b", Actual: ll.Tail().Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_IsEmpty_True(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	tc := coretestcases.CaseV1{Name: "LL IsEmpty true", Expected: true, Actual: ll.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_IsEmpty_False(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	tc := coretestcases.CaseV1{Name: "LL IsEmpty false", Expected: false, Actual: ll.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_HasItems(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	tc := coretestcases.CaseV1{Name: "LL HasItems", Expected: true, Actual: ll.HasItems(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_IsEmptyLock(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	tc := coretestcases.CaseV1{Name: "LL IsEmptyLock", Expected: true, Actual: ll.IsEmptyLock(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_LengthLock(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	tc := coretestcases.CaseV1{Name: "LL LengthLock", Expected: 1, Actual: ll.LengthLock(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddLock(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddLock("x")
	tc := coretestcases.CaseV1{Name: "LL AddLock", Expected: 1, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddFront(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("b").AddFront("a")
	tc := coretestcases.CaseV1{Name: "LL AddFront", Expected: "a", Actual: ll.Head().Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddFront_Empty(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddFront("a")
	tc := coretestcases.CaseV1{Name: "LL AddFront empty", Expected: "a", Actual: ll.Head().Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_PushFront(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("b").PushFront("a")
	tc := coretestcases.CaseV1{Name: "LL PushFront", Expected: "a", Actual: ll.Head().Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_PushBack(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.PushBack("a")
	tc := coretestcases.CaseV1{Name: "LL PushBack", Expected: 1, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_Push(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Push("a")
	tc := coretestcases.CaseV1{Name: "LL Push", Expected: 1, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddNonEmpty(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddNonEmpty("a").AddNonEmpty("")
	tc := coretestcases.CaseV1{Name: "LL AddNonEmpty", Expected: 1, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddNonEmptyWhitespace(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddNonEmptyWhitespace("a").AddNonEmptyWhitespace("  ")
	tc := coretestcases.CaseV1{Name: "LL AddNonEmptyWhitespace", Expected: 1, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddIf_True(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddIf(true, "a")
	tc := coretestcases.CaseV1{Name: "LL AddIf true", Expected: 1, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddIf_False(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddIf(false, "a")
	tc := coretestcases.CaseV1{Name: "LL AddIf false", Expected: 0, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddsIf_True(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddsIf(true, "a", "b")
	tc := coretestcases.CaseV1{Name: "LL AddsIf true", Expected: 2, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddsIf_False(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddsIf(false, "a", "b")
	tc := coretestcases.CaseV1{Name: "LL AddsIf false", Expected: 0, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddFunc(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddFunc(func() string { return "x" })
	tc := coretestcases.CaseV1{Name: "LL AddFunc", Expected: "x", Actual: ll.Head().Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddFuncErr_NoErr(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddFuncErr(func() (string, error) { return "ok", nil }, func(err error) {})
	tc := coretestcases.CaseV1{Name: "LL AddFuncErr no err", Expected: "ok", Actual: ll.Head().Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_Adds(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Adds("a", "b", "c")
	tc := coretestcases.CaseV1{Name: "LL Adds", Expected: 3, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_Adds_Empty(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Adds()
	tc := coretestcases.CaseV1{Name: "LL Adds empty", Expected: 0, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddStrings(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddStrings([]string{"x", "y"})
	tc := coretestcases.CaseV1{Name: "LL AddStrings", Expected: 2, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddsLock(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddsLock("a", "b")
	tc := coretestcases.CaseV1{Name: "LL AddsLock", Expected: 2, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddItemsMap(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddItemsMap(map[string]bool{"a": true, "b": false})
	tc := coretestcases.CaseV1{Name: "LL AddItemsMap", Expected: 1, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddCollection(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	col := corestr.New.Collection.Strings("a", "b")
	ll.AddCollection(col)
	tc := coretestcases.CaseV1{Name: "LL AddCollection", Expected: 2, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AddCollection_Nil(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.AddCollection(nil)
	tc := coretestcases.CaseV1{Name: "LL AddCollection nil", Expected: 0, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AppendNode(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	node := &corestr.LinkedListNode{Element: "b"}
	ll.AppendNode(node)
	tc := coretestcases.CaseV1{Name: "LL AppendNode", Expected: 2, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AppendNode_Empty(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	node := &corestr.LinkedListNode{Element: "a"}
	ll.AppendNode(node)
	tc := coretestcases.CaseV1{Name: "LL AppendNode empty", Expected: "a", Actual: ll.Head().Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_IsEquals_Same(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "LL IsEquals same", Expected: true, Actual: ll.IsEquals(ll), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_IsEquals_Equal(t *testing.T) {
	ll1 := corestr.Empty.LinkedList()
	ll1.Add("a").Add("b")
	ll2 := corestr.Empty.LinkedList()
	ll2.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "LL IsEquals equal", Expected: true, Actual: ll1.IsEquals(ll2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_IsEquals_DiffLen(t *testing.T) {
	ll1 := corestr.Empty.LinkedList()
	ll1.Add("a")
	ll2 := corestr.Empty.LinkedList()
	ll2.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "LL IsEquals diff len", Expected: false, Actual: ll1.IsEquals(ll2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_IsEqualsWithSensitive_CI(t *testing.T) {
	ll1 := corestr.Empty.LinkedList()
	ll1.Add("A")
	ll2 := corestr.Empty.LinkedList()
	ll2.Add("a")
	tc := coretestcases.CaseV1{Name: "LL IsEqualsWithSensitive CI", Expected: true, Actual: ll1.IsEqualsWithSensitive(ll2, false), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_List(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "LL List", Expected: 2, Actual: len(ll.List()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_List_Empty(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	tc := coretestcases.CaseV1{Name: "LL List empty", Expected: 0, Actual: len(ll.List()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_ListLock(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	tc := coretestcases.CaseV1{Name: "LL ListLock", Expected: 1, Actual: len(ll.ListLock()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_String(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	tc := coretestcases.CaseV1{Name: "LL String", Expected: true, Actual: len(ll.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_String_Empty(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	tc := coretestcases.CaseV1{Name: "LL String empty", Expected: true, Actual: len(ll.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_Join(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "LL Join", Expected: "a,b", Actual: ll.Join(","), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_JoinLock(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	tc := coretestcases.CaseV1{Name: "LL JoinLock", Expected: "a,b", Actual: ll.JoinLock(","), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_ToCollection(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	col := ll.ToCollection(0)
	tc := coretestcases.CaseV1{Name: "LL ToCollection", Expected: 2, Actual: col.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_ToCollection_Empty(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	col := ll.ToCollection(5)
	tc := coretestcases.CaseV1{Name: "LL ToCollection empty", Expected: 0, Actual: col.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_SafeIndexAt(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	node := ll.SafeIndexAt(1)
	tc := coretestcases.CaseV1{Name: "LL SafeIndexAt", Expected: "b", Actual: node.Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_SafeIndexAt_OOB(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	node := ll.SafeIndexAt(0)
	tc := coretestcases.CaseV1{Name: "LL SafeIndexAt oob", Expected: true, Actual: node == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_SafeIndexAt_Negative(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	node := ll.SafeIndexAt(-1)
	tc := coretestcases.CaseV1{Name: "LL SafeIndexAt neg", Expected: true, Actual: node == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_SafePointerIndexAt(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	ptr := ll.SafePointerIndexAt(0)
	tc := coretestcases.CaseV1{Name: "LL SafePointerIndexAt", Expected: "a", Actual: *ptr, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_SafePointerIndexAt_Nil(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ptr := ll.SafePointerIndexAt(0)
	tc := coretestcases.CaseV1{Name: "LL SafePointerIndexAt nil", Expected: true, Actual: ptr == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_SafePointerIndexAtUsingDefault(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	val := ll.SafePointerIndexAtUsingDefault(0, "def")
	tc := coretestcases.CaseV1{Name: "LL SafePointerIndexAtDefault", Expected: "def", Actual: val, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_SafeIndexAtLock(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	node := ll.SafeIndexAtLock(0)
	tc := coretestcases.CaseV1{Name: "LL SafeIndexAtLock", Expected: "a", Actual: node.Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_Clear(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	ll.Clear()
	tc := coretestcases.CaseV1{Name: "LL Clear", Expected: 0, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_RemoveAll(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	ll.RemoveAll()
	tc := coretestcases.CaseV1{Name: "LL RemoveAll", Expected: true, Actual: ll.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_JsonModel(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	tc := coretestcases.CaseV1{Name: "LL JsonModel", Expected: 1, Actual: len(ll.JsonModel()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_MarshalJSON(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	data, err := ll.MarshalJSON()
	tc := coretestcases.CaseV1{Name: "LL MarshalJSON", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_UnmarshalJSON(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	err := ll.UnmarshalJSON([]byte(`["x","y"]`))
	tc := coretestcases.CaseV1{Name: "LL UnmarshalJSON", Expected: true, Actual: err == nil && ll.Length() == 2, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_Loop(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	count := 0
	ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
		count++
		return false
	})
	tc := coretestcases.CaseV1{Name: "LL Loop", Expected: 2, Actual: count, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_Loop_Break(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b").Add("c")
	count := 0
	ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
		count++
		return true
	})
	tc := coretestcases.CaseV1{Name: "LL Loop break", Expected: 1, Actual: count, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_Filter(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b").Add("c")
	result := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
		return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: arg.Node.Element == "b", IsBreak: false}
	})
	tc := coretestcases.CaseV1{Name: "LL Filter", Expected: 1, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_GetAllLinkedNodes(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	nodes := ll.GetAllLinkedNodes()
	tc := coretestcases.CaseV1{Name: "LL GetAllLinkedNodes", Expected: 2, Actual: len(nodes), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_GetNextNodes(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b").Add("c")
	nodes := ll.GetNextNodes(2)
	tc := coretestcases.CaseV1{Name: "LL GetNextNodes", Expected: 2, Actual: len(nodes), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_InsertAt(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("c")
	ll.InsertAt(1, "b")
	tc := coretestcases.CaseV1{Name: "LL InsertAt", Expected: 3, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_InsertAt_Front(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("b")
	ll.InsertAt(0, "a")
	tc := coretestcases.CaseV1{Name: "LL InsertAt front", Expected: "a", Actual: ll.Head().Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_RemoveNodeByElementValue(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b").Add("c")
	ll.RemoveNodeByElementValue("b", true, false)
	tc := coretestcases.CaseV1{Name: "LL RemoveByElem", Expected: 2, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_RemoveNodeByElementValue_First(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b")
	ll.RemoveNodeByElementValue("a", true, false)
	tc := coretestcases.CaseV1{Name: "LL RemoveByElem first", Expected: "b", Actual: ll.Head().Element, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_RemoveNodeByIndex(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a").Add("b").Add("c")
	ll.RemoveNodeByIndex(1)
	tc := coretestcases.CaseV1{Name: "LL RemoveByIndex", Expected: 2, Actual: ll.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_StringLock(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	tc := coretestcases.CaseV1{Name: "LL StringLock", Expected: true, Actual: len(ll.StringLock()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_Joins(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	result := ll.Joins(",", "b", "c")
	tc := coretestcases.CaseV1{Name: "LL Joins", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_JsonModelAny(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	tc := coretestcases.CaseV1{Name: "LL JsonModelAny", Expected: true, Actual: ll.JsonModelAny() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_Json(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	j := ll.Json()
	tc := coretestcases.CaseV1{Name: "LL Json", Expected: true, Actual: j.HasSafeItems(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_LinkedList_AsJsonMarshaller(t *testing.T) {
	ll := corestr.Empty.LinkedList()
	ll.Add("a")
	m := ll.AsJsonMarshaller()
	tc := coretestcases.CaseV1{Name: "LL AsJsonMarshaller", Expected: true, Actual: m != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// CharCollectionDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov51_CharCollectionDataModel_NewUsing(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("hello")
	dm := corestr.NewCharCollectionMapDataModelUsing(ccm)
	tc := coretestcases.CaseV1{Name: "CharCollDM NewUsing", Expected: true, Actual: dm != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov51_CharCollectionDataModel_NewUsingDataModel(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("hello")
	dm := corestr.NewCharCollectionMapDataModelUsing(ccm)
	restored := corestr.NewCharCollectionMapUsingDataModel(dm)
	tc := coretestcases.CaseV1{Name: "CharCollDM restored", Expected: true, Actual: restored.Has("hello"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}
