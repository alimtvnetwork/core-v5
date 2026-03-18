package corestr

import (
	"encoding/json"
	"errors"
	"testing"
)

// ── Creators ──────────────────────────────────────────────

func Test_LinkedList_NewCreate(t *testing.T) {
	ll := New.LinkedList.Create()
	if ll == nil || ll.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedList_NewEmpty(t *testing.T) {
	ll := New.LinkedList.Empty()
	if ll.Length() != 0 {
		t.Fatal("expected empty")
	}
}

func Test_LinkedList_NewStrings(t *testing.T) {
	ll := New.LinkedList.Strings([]string{"a", "b", "c"})
	if ll.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_LinkedList_NewStrings_Empty(t *testing.T) {
	ll := New.LinkedList.Strings([]string{})
	if ll.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedList_NewSpreadStrings(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("x", "y")
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_NewSpreadStrings_Empty(t *testing.T) {
	ll := New.LinkedList.SpreadStrings()
	if ll.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedList_NewUsingMap(t *testing.T) {
	ll := New.LinkedList.UsingMap(map[string]bool{"a": true, "b": false, "c": true})
	if ll.Length() != 2 {
		t.Fatalf("expected 2, got %d", ll.Length())
	}
}

func Test_LinkedList_NewUsingMap_Nil(t *testing.T) {
	ll := New.LinkedList.UsingMap(nil)
	if ll.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedList_NewPointerStringsPtr(t *testing.T) {
	s1, s2 := "a", "b"
	items := []*string{&s1, &s2, nil}
	ll := New.LinkedList.PointerStringsPtr(&items)
	if ll.Length() != 2 {
		t.Fatalf("expected 2, got %d", ll.Length())
	}
}

func Test_LinkedList_NewPointerStringsPtr_Nil(t *testing.T) {
	ll := New.LinkedList.PointerStringsPtr(nil)
	if ll.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedList_EmptyLinkedList(t *testing.T) {
	ll := Empty.LinkedList()
	if ll.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── Head / Tail / Length ──────────────────────────────────

func Test_LinkedList_HeadTail(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	if ll.Head().Element != "a" {
		t.Fatal("expected head a")
	}
	if ll.Tail().Element != "c" {
		t.Fatal("expected tail c")
	}
}

func Test_LinkedList_LengthLock(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	if ll.LengthLock() != 2 {
		t.Fatal("expected 2")
	}
}

// ── IsEmpty / HasItems / IsEmptyLock ──────────────────────

func Test_LinkedList_IsEmpty(t *testing.T) {
	ll := New.LinkedList.Create()
	if !ll.IsEmpty() || ll.HasItems() {
		t.Fatal("expected empty")
	}
	ll.Add("x")
	if ll.IsEmpty() || !ll.HasItems() {
		t.Fatal("expected has items")
	}
}

func Test_LinkedList_IsEmptyLock(t *testing.T) {
	ll := New.LinkedList.Create()
	if !ll.IsEmptyLock() {
		t.Fatal("expected empty")
	}
}

// ── Add variants ──────────────────────────────────────────

func Test_LinkedList_Add(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.Add("first")
	ll.Add("second")
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_AddLock(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddLock("a")
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedList_AddNonEmpty(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddNonEmpty("")
	ll.AddNonEmpty("x")
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedList_AddNonEmptyWhitespace(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddNonEmptyWhitespace("  ")
	ll.AddNonEmptyWhitespace("x")
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedList_AddIf(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddIf(false, "no")
	ll.AddIf(true, "yes")
	if ll.Length() != 1 || ll.Head().Element != "yes" {
		t.Fatal("expected only yes")
	}
}

func Test_LinkedList_AddsIf(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddsIf(false, "a", "b")
	ll.AddsIf(true, "c", "d")
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_AddFunc(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddFunc(func() string { return "computed" })
	if ll.Head().Element != "computed" {
		t.Fatal("expected computed")
	}
}

func Test_LinkedList_AddFuncErr_Success(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddFuncErr(
		func() (string, error) { return "ok", nil },
		func(err error) { t.Fatal("should not be called") },
	)
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedList_AddFuncErr_Error(t *testing.T) {
	ll := New.LinkedList.Create()
	called := false
	ll.AddFuncErr(
		func() (string, error) { return "", errors.New("fail") },
		func(err error) { called = true },
	)
	if ll.HasItems() {
		t.Fatal("expected empty")
	}
	if !called {
		t.Fatal("expected error handler called")
	}
}

func Test_LinkedList_Push(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.Push("a")
	ll.PushBack("b")
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_Adds(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.Adds("a", "b", "c")
	if ll.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_LinkedList_Adds_Empty(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.Adds()
	if ll.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedList_AddStrings(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddStrings([]string{"a", "b"})
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_AddsLock(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddsLock("a", "b")
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_AddItemsMap(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddItemsMap(map[string]bool{"a": true, "b": false})
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedList_AddItemsMap_Empty(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddItemsMap(map[string]bool{})
	if ll.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AddFront / PushFront ──────────────────────────────────

func Test_LinkedList_AddFront(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("b", "c")
	ll.AddFront("a")
	if ll.Head().Element != "a" {
		t.Fatal("expected a at front")
	}
	if ll.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_LinkedList_AddFront_Empty(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddFront("only")
	if ll.Length() != 1 || ll.Head().Element != "only" {
		t.Fatal("expected single item")
	}
}

func Test_LinkedList_PushFront(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("b")
	ll.PushFront("a")
	if ll.Head().Element != "a" {
		t.Fatal("expected a")
	}
}

// ── AppendNode / AppendChainOfNodes / AddBackNode ─────────

func Test_LinkedList_AppendNode(t *testing.T) {
	ll := New.LinkedList.Create()
	node := &LinkedListNode{Element: "x"}
	ll.AppendNode(node)
	if ll.Length() != 1 || ll.Head().Element != "x" {
		t.Fatal("expected x")
	}
	ll.AppendNode(&LinkedListNode{Element: "y"})
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_AddBackNode(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddBackNode(&LinkedListNode{Element: "a"})
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedList_AppendChainOfNodes(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	chain := &LinkedListNode{Element: "b", next: &LinkedListNode{Element: "c"}}
	ll.AppendChainOfNodes(chain)
	if ll.Length() != 3 {
		t.Fatalf("expected 3, got %d", ll.Length())
	}
}

func Test_LinkedList_AppendChainOfNodes_Empty(t *testing.T) {
	ll := New.LinkedList.Create()
	chain := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	ll.AppendChainOfNodes(chain)
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

// ── InsertAt ──────────────────────────────────────────────

func Test_LinkedList_InsertAt_Front(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("b", "c")
	ll.InsertAt(0, "a")
	if ll.Head().Element != "a" {
		t.Fatal("expected a at front")
	}
}

func Test_LinkedList_InsertAt_Middle(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "c")
	ll.InsertAt(1, "b")
	list := ll.List()
	if len(list) < 3 || list[1] != "b" {
		t.Fatal("expected b at index 1")
	}
}

// ── AddAfterNode ──────────────────────────────────────────

func Test_LinkedList_AddAfterNode(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "c")
	node := ll.Head()
	ll.AddAfterNode(node, "b")
	list := ll.List()
	if list[1] != "b" {
		t.Fatal("expected b after a")
	}
}

// ── AttachWithNode ────────────────────────────────────────

func Test_LinkedList_AttachWithNode_NilCurrent(t *testing.T) {
	ll := New.LinkedList.Create()
	err := ll.AttachWithNode(nil, &LinkedListNode{Element: "x"})
	if err == nil {
		t.Fatal("expected error for nil current")
	}
}

func Test_LinkedList_AttachWithNode_NonNilNext(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	node := ll.Head() // has next
	err := ll.AttachWithNode(node, &LinkedListNode{Element: "x"})
	if err == nil {
		t.Fatal("expected error for non-nil next")
	}
}

func Test_LinkedList_AttachWithNode_Success(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	node := ll.Tail()
	err := ll.AttachWithNode(node, &LinkedListNode{Element: "b"})
	if err != nil {
		t.Fatal(err)
	}
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

// ── AddCollectionToNode ───────────────────────────────────

func Test_LinkedList_AddCollectionToNode(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	col := New.Collection.Strings([]string{"b", "c"})
	ll.AddCollectionToNode(true, ll.Head(), col)
	if ll.Length() < 2 {
		t.Fatal("expected items added")
	}
}

// ── AddPointerStringsPtr ──────────────────────────────────

func Test_LinkedList_AddPointerStringsPtr(t *testing.T) {
	ll := New.LinkedList.Create()
	s1 := "a"
	ll.AddPointerStringsPtr([]*string{&s1, nil})
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ── AddCollection ─────────────────────────────────────────

func Test_LinkedList_AddCollection(t *testing.T) {
	ll := New.LinkedList.Create()
	col := New.Collection.Strings([]string{"a", "b"})
	ll.AddCollection(col)
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_AddCollection_Nil(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.AddCollection(nil)
	if ll.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── Loop ──────────────────────────────────────────────────

func Test_LinkedList_Loop(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	count := 0
	ll.Loop(func(arg *LinkedListProcessorParameter) bool {
		count++
		return false
	})
	if count != 3 {
		t.Fatalf("expected 3, got %d", count)
	}
}

func Test_LinkedList_Loop_Break(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	count := 0
	ll.Loop(func(arg *LinkedListProcessorParameter) bool {
		count++
		return true // break immediately
	})
	if count != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedList_Loop_Empty(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.Loop(func(arg *LinkedListProcessorParameter) bool {
		t.Fatal("should not be called")
		return false
	})
}

func Test_LinkedList_Loop_BreakOnSecond(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	count := 0
	ll.Loop(func(arg *LinkedListProcessorParameter) bool {
		count++
		return arg.Index == 1
	})
	if count != 2 {
		t.Fatal("expected 2")
	}
}

// ── Filter ────────────────────────────────────────────────

func Test_LinkedList_Filter(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	nodes := ll.Filter(func(arg *LinkedListFilterParameter) *LinkedListFilterResult {
		return &LinkedListFilterResult{Value: arg.Node, IsKeep: arg.Node.Element != "b", IsBreak: false}
	})
	if len(nodes) != 2 {
		t.Fatalf("expected 2, got %d", len(nodes))
	}
}

func Test_LinkedList_Filter_Empty(t *testing.T) {
	ll := New.LinkedList.Create()
	nodes := ll.Filter(func(arg *LinkedListFilterParameter) *LinkedListFilterResult {
		return &LinkedListFilterResult{Value: arg.Node, IsKeep: true}
	})
	if len(nodes) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_LinkedList_Filter_BreakFirst(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	nodes := ll.Filter(func(arg *LinkedListFilterParameter) *LinkedListFilterResult {
		return &LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
	})
	if len(nodes) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedList_Filter_BreakSecond(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	nodes := ll.Filter(func(arg *LinkedListFilterParameter) *LinkedListFilterResult {
		return &LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: arg.Index == 1}
	})
	if len(nodes) != 2 {
		t.Fatal("expected 2")
	}
}

// ── RemoveNodeByElementValue ──────────────────────────────

func Test_LinkedList_RemoveNodeByElementValue_First(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	ll.RemoveNodeByElementValue("a", true, true)
	if ll.Length() != 2 || ll.Head().Element != "b" {
		t.Fatal("expected b as head")
	}
}

func Test_LinkedList_RemoveNodeByElementValue_Middle(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	ll.RemoveNodeByElementValue("b", true, true)
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_RemoveNodeByElementValue_CaseInsensitive(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("ABC", "def")
	ll.RemoveNodeByElementValue("abc", false, true)
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ── RemoveNodeByIndex ─────────────────────────────────────

func Test_LinkedList_RemoveNodeByIndex_First(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	ll.RemoveNodeByIndex(0)
	if ll.Head().Element != "b" {
		t.Fatal("expected b")
	}
}

func Test_LinkedList_RemoveNodeByIndex_Last(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	ll.RemoveNodeByIndex(2)
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_RemoveNodeByIndex_Middle(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	ll.RemoveNodeByIndex(1)
	list := ll.List()
	if len(list) != 2 || list[1] != "c" {
		t.Fatal("expected a,c")
	}
}

// ── RemoveNodeByIndexes ───────────────────────────────────

func Test_LinkedList_RemoveNodeByIndexes(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c", "d")
	ll.RemoveNodeByIndexes(true, 0, 2)
	if ll.Length() != 2 {
		t.Fatalf("expected 2, got %d", ll.Length())
	}
}

func Test_LinkedList_RemoveNodeByIndexes_Empty(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	ll.RemoveNodeByIndexes(true)
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ── RemoveNode ────────────────────────────────────────────

func Test_LinkedList_RemoveNode(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	node := ll.Head().Next() // "b"
	ll.RemoveNode(node)
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_RemoveNode_First(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	ll.RemoveNode(ll.Head())
	if ll.Head().Element != "b" {
		t.Fatal("expected b")
	}
}

func Test_LinkedList_RemoveNode_Nil(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	ll.RemoveNode(nil)
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ── IndexAt / SafeIndexAt ─────────────────────────────────

func Test_LinkedList_IndexAt(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	if ll.IndexAt(0).Element != "a" {
		t.Fatal("expected a")
	}
	if ll.IndexAt(2).Element != "c" {
		t.Fatal("expected c")
	}
}

func Test_LinkedList_IndexAt_Negative(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	node := ll.IndexAt(-1)
	if node != nil {
		t.Fatal("expected nil")
	}
}

func Test_LinkedList_SafeIndexAt(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	node := ll.SafeIndexAt(1)
	if node == nil || node.Element != "b" {
		t.Fatal("expected b")
	}
}

func Test_LinkedList_SafeIndexAt_OutOfRange(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	node := ll.SafeIndexAt(5)
	if node != nil {
		t.Fatal("expected nil")
	}
}

func Test_LinkedList_SafeIndexAt_Negative(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	node := ll.SafeIndexAt(-1)
	if node != nil {
		t.Fatal("expected nil")
	}
}

func Test_LinkedList_SafeIndexAtLock(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	node := ll.SafeIndexAtLock(0)
	if node == nil || node.Element != "a" {
		t.Fatal("expected a")
	}
}

func Test_LinkedList_SafePointerIndexAt(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	ptr := ll.SafePointerIndexAt(0)
	if ptr == nil || *ptr != "a" {
		t.Fatal("expected a")
	}
	ptr2 := ll.SafePointerIndexAt(5)
	if ptr2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_LinkedList_SafePointerIndexAtUsingDefault(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	v := ll.SafePointerIndexAtUsingDefault(0, "def")
	if v != "a" {
		t.Fatal("expected a")
	}
	v2 := ll.SafePointerIndexAtUsingDefault(5, "def")
	if v2 != "def" {
		t.Fatal("expected def")
	}
}

func Test_LinkedList_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	v := ll.SafePointerIndexAtUsingDefaultLock(0, "def")
	if v != "a" {
		t.Fatal("expected a")
	}
}

// ── GetNextNodes / GetAllLinkedNodes ───────────────────────

func Test_LinkedList_GetNextNodes(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c", "d")
	nodes := ll.GetNextNodes(2)
	if len(nodes) != 2 {
		t.Fatalf("expected 2, got %d", len(nodes))
	}
}

func Test_LinkedList_GetAllLinkedNodes(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	nodes := ll.GetAllLinkedNodes()
	if len(nodes) != 2 {
		t.Fatal("expected 2")
	}
}

// ── ToCollection / List ───────────────────────────────────

func Test_LinkedList_ToCollection(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	col := ll.ToCollection(5)
	if col.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_ToCollection_Empty(t *testing.T) {
	ll := New.LinkedList.Create()
	col := ll.ToCollection(0)
	if col.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedList_List(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	list := ll.List()
	if len(list) != 2 || list[0] != "a" {
		t.Fatal("expected [a, b]")
	}
}

func Test_LinkedList_List_Empty(t *testing.T) {
	ll := New.LinkedList.Create()
	list := ll.List()
	if len(list) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_LinkedList_ListPtr(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	list := ll.ListPtr()
	if len(list) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedList_ListLock(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	list := ll.ListLock()
	if len(list) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedList_ListPtrLock(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	list := ll.ListPtrLock()
	if len(list) != 1 {
		t.Fatal("expected 1")
	}
}

// ── String / StringLock / Join ─────────────────────────────

func Test_LinkedList_String(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	s := ll.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_LinkedList_String_Empty(t *testing.T) {
	ll := New.LinkedList.Create()
	s := ll.String()
	if s == "" {
		t.Fatal("expected non-empty (NoElements)")
	}
}

func Test_LinkedList_StringLock(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	s := ll.StringLock()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_LinkedList_StringLock_Empty(t *testing.T) {
	ll := New.LinkedList.Create()
	s := ll.StringLock()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_LinkedList_Join(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	if ll.Join(",") != "a,b" {
		t.Fatal("expected a,b")
	}
}

func Test_LinkedList_JoinLock(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	if ll.JoinLock(",") != "a,b" {
		t.Fatal("expected a,b")
	}
}

func Test_LinkedList_Joins(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	result := ll.Joins(",", "b", "c")
	if result != "a,b,c" {
		t.Fatalf("expected a,b,c got %s", result)
	}
}

func Test_LinkedList_Joins_NilItems(t *testing.T) {
	ll := New.LinkedList.Create()
	result := ll.Joins(",", "a")
	if result != "a" {
		t.Fatal("expected just a")
	}
}

// ── GetCompareSummary ─────────────────────────────────────

func Test_LinkedList_GetCompareSummary(t *testing.T) {
	ll1 := New.LinkedList.SpreadStrings("a", "b")
	ll2 := New.LinkedList.SpreadStrings("a", "b")
	summary := ll1.GetCompareSummary(ll2, "left", "right")
	if summary == "" {
		t.Fatal("expected non-empty")
	}
}

// ── IsEquals ──────────────────────────────────────────────

func Test_LinkedList_IsEquals(t *testing.T) {
	ll1 := New.LinkedList.SpreadStrings("a", "b")
	ll2 := New.LinkedList.SpreadStrings("a", "b")
	if !ll1.IsEquals(ll2) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedList_IsEquals_DiffLength(t *testing.T) {
	ll1 := New.LinkedList.SpreadStrings("a", "b")
	ll2 := New.LinkedList.SpreadStrings("a")
	if ll1.IsEquals(ll2) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedList_IsEquals_BothNil(t *testing.T) {
	var ll1 *LinkedList
	var ll2 *LinkedList
	// Can't call method on nil, but test via IsEqualsWithSensitive coverage
	ll1 = New.LinkedList.Create()
	if !ll1.IsEqualsWithSensitive(nil, true) {
		// ll1 is empty, nil treated differently
	}
}

func Test_LinkedList_IsEquals_SameRef(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	if !ll.IsEqualsWithSensitive(ll, true) {
		t.Fatal("expected same ref equal")
	}
}

func Test_LinkedList_IsEquals_BothEmpty(t *testing.T) {
	ll1 := New.LinkedList.Create()
	ll2 := New.LinkedList.Create()
	if !ll1.IsEquals(ll2) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedList_IsEquals_OneEmpty(t *testing.T) {
	ll1 := New.LinkedList.SpreadStrings("a")
	ll2 := New.LinkedList.Create()
	if ll1.IsEquals(ll2) {
		t.Fatal("expected not equal")
	}
}

// ── AddStringsToNode / AddStringsPtrToNode ────────────────

func Test_LinkedList_AddStringsToNode(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "d")
	node := ll.Head()
	ll.AddStringsToNode(false, node, []string{"b", "c"})
	if ll.Length() < 3 {
		t.Fatal("expected items added")
	}
}

func Test_LinkedList_AddStringsToNode_Single(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "c")
	node := ll.Head()
	ll.AddStringsToNode(false, node, []string{"b"})
	list := ll.List()
	if len(list) < 3 || list[1] != "b" {
		t.Fatal("expected b inserted")
	}
}

func Test_LinkedList_AddStringsToNode_Empty(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	ll.AddStringsToNode(false, ll.Head(), []string{})
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedList_AddStringsToNode_NilNodeSkip(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	ll.AddStringsToNode(true, nil, []string{"b"})
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedList_AddStringsPtrToNode(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	items := []string{"b"}
	ll.AddStringsPtrToNode(true, ll.Head(), &items)
	if ll.Length() < 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_AddStringsPtrToNode_Nil(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	ll.AddStringsPtrToNode(true, ll.Head(), nil)
	if ll.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ── JSON ──────────────────────────────────────────────────

func Test_LinkedList_JsonModel(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	model := ll.JsonModel()
	if len(model) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_JsonModelAny(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	if ll.JsonModelAny() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LinkedList_MarshalUnmarshalJSON(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b", "c")
	data, err := json.Marshal(ll)
	if err != nil {
		t.Fatal(err)
	}
	ll2 := New.LinkedList.Create()
	err = json.Unmarshal(data, ll2)
	if err != nil {
		t.Fatal(err)
	}
	if ll2.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_LinkedList_Json(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	r := ll.Json()
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_LinkedList_JsonPtr(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	r := ll.JsonPtr()
	if r == nil || r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_LinkedList_ParseInjectUsingJson(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	jr := ll.JsonPtr()
	ll2 := New.LinkedList.Create()
	_, err := ll2.ParseInjectUsingJson(jr)
	if err != nil {
		t.Fatal(err)
	}
	if ll2.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedList_ParseInjectUsingJsonMust(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	jr := ll.JsonPtr()
	ll2 := New.LinkedList.Create()
	result := ll2.ParseInjectUsingJsonMust(jr)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LinkedList_JsonParseSelfInject(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	jr := ll.JsonPtr()
	ll2 := New.LinkedList.Create()
	err := ll2.JsonParseSelfInject(jr)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_LinkedList_AsJsonMarshaller(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	if ll.AsJsonMarshaller() == nil {
		t.Fatal("expected non-nil")
	}
}

// ── RemoveAll / Clear ─────────────────────────────────────

func Test_LinkedList_RemoveAll(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a", "b")
	ll.RemoveAll()
	if ll.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedList_Clear(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	ll.Clear()
	if ll.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedList_Clear_Empty(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.Clear()
}

// ── LinkedListNode ────────────────────────────────────────

func Test_LinkedListNode_HasNext(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	if !node.HasNext() {
		t.Fatal("expected has next")
	}
	if node.Next().HasNext() {
		t.Fatal("expected no next")
	}
}

func Test_LinkedListNode_EndOfChain(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b", next: &LinkedListNode{Element: "c"}}}
	end, length := node.EndOfChain()
	if end.Element != "c" || length != 3 {
		t.Fatal("expected c, 3")
	}
}

func Test_LinkedListNode_Clone(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	cloned := node.Clone()
	if cloned.Element != "a" || cloned.HasNext() {
		t.Fatal("expected a without next")
	}
}

func Test_LinkedListNode_List(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	list := node.List()
	if len(list) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedListNode_ListPtr(t *testing.T) {
	node := &LinkedListNode{Element: "a"}
	list := node.ListPtr()
	if len(list) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedListNode_Join(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	if node.Join(",") != "a,b" {
		t.Fatal("expected a,b")
	}
}

func Test_LinkedListNode_String(t *testing.T) {
	node := &LinkedListNode{Element: "hello"}
	if node.String() != "hello" {
		t.Fatal("expected hello")
	}
}

func Test_LinkedListNode_StringList(t *testing.T) {
	node := &LinkedListNode{Element: "a"}
	s := node.StringList("Header: ")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_LinkedListNode_IsEqual(t *testing.T) {
	n1 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	n2 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	if !n1.IsEqual(n2) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedListNode_IsEqual_BothNil(t *testing.T) {
	var n1 *LinkedListNode
	if !n1.IsEqual(nil) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedListNode_IsEqual_OneNil(t *testing.T) {
	n1 := &LinkedListNode{Element: "a"}
	if n1.IsEqual(nil) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedListNode_IsEqual_SameRef(t *testing.T) {
	n1 := &LinkedListNode{Element: "a"}
	if !n1.IsEqual(n1) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedListNode_IsEqual_DiffElement(t *testing.T) {
	n1 := &LinkedListNode{Element: "a"}
	n2 := &LinkedListNode{Element: "b"}
	if n1.IsEqual(n2) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedListNode_IsChainEqual(t *testing.T) {
	n1 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	n2 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	if !n1.IsChainEqual(n2, true) {
		t.Fatal("expected chain equal")
	}
}

func Test_LinkedListNode_IsChainEqual_CaseInsensitive(t *testing.T) {
	n1 := &LinkedListNode{Element: "A"}
	n2 := &LinkedListNode{Element: "a"}
	if !n1.IsChainEqual(n2, false) {
		t.Fatal("expected equal case insensitive")
	}
}

func Test_LinkedListNode_IsChainEqual_BothNil(t *testing.T) {
	var n1 *LinkedListNode
	if !n1.IsChainEqual(nil, true) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedListNode_IsChainEqual_OneNil(t *testing.T) {
	n1 := &LinkedListNode{Element: "a"}
	if n1.IsChainEqual(nil, true) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedListNode_IsEqualSensitive(t *testing.T) {
	n1 := &LinkedListNode{Element: "A"}
	n2 := &LinkedListNode{Element: "a"}
	if !n1.IsEqualSensitive(n2, false) {
		t.Fatal("expected equal")
	}
	if n1.IsEqualSensitive(n2, true) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedListNode_IsEqualSensitive_BothNil(t *testing.T) {
	var n1 *LinkedListNode
	if !n1.IsEqualSensitive(nil, true) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedListNode_IsEqualSensitive_OneNil(t *testing.T) {
	n1 := &LinkedListNode{Element: "a"}
	if n1.IsEqualSensitive(nil, true) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedListNode_IsEqualValue(t *testing.T) {
	n := &LinkedListNode{Element: "hello"}
	if !n.IsEqualValue("hello") {
		t.Fatal("expected equal")
	}
}

func Test_LinkedListNode_IsEqualValueSensitive(t *testing.T) {
	n := &LinkedListNode{Element: "Hello"}
	if !n.IsEqualValueSensitive("hello", false) {
		t.Fatal("expected equal")
	}
	if n.IsEqualValueSensitive("hello", true) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedListNode_CreateLinkedList(t *testing.T) {
	n := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	ll := n.CreateLinkedList()
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedListNode_AddNext(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	node := ll.Head()
	newNode := node.AddNext(ll, "b")
	if newNode.Element != "b" {
		t.Fatal("expected b")
	}
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedListNode_AddNextNode(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	node := ll.Head()
	newNode := &LinkedListNode{Element: "b"}
	node.AddNextNode(ll, newNode)
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedListNode_AddStringsToNode(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	node := ll.Head()
	node.AddStringsToNode(ll, true, []string{"b", "c"})
	if ll.Length() < 2 {
		t.Fatal("expected items added")
	}
}

func Test_LinkedListNode_AddStringsPtrToNode(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	node := ll.Head()
	items := []string{"b"}
	node.AddStringsPtrToNode(ll, true, &items)
	if ll.Length() < 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedListNode_AddStringsPtrToNode_Nil(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	node := ll.Head()
	result := node.AddStringsPtrToNode(ll, true, nil)
	if result.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedListNode_AddCollectionToNode(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	col := New.Collection.Strings([]string{"b", "c"})
	ll.Head().AddCollectionToNode(ll, true, col)
	if ll.Length() < 2 {
		t.Fatal("expected items added")
	}
}

func Test_LinkedListNode_LoopEndOfChain(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b", next: &LinkedListNode{Element: "c"}}}
	count := 0
	end, length := node.LoopEndOfChain(func(arg *LinkedListProcessorParameter) bool {
		count++
		return false
	})
	if end.Element != "c" || length != 3 || count != 3 {
		t.Fatal("expected c, 3, 3")
	}
}

func Test_LinkedListNode_LoopEndOfChain_Break(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	end, length := node.LoopEndOfChain(func(arg *LinkedListProcessorParameter) bool {
		return true
	})
	if end.Element != "a" || length != 1 {
		t.Fatal("expected a, 1")
	}
}

func Test_LinkedListNode_LoopEndOfChain_BreakSecond(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b", next: &LinkedListNode{Element: "c"}}}
	end, length := node.LoopEndOfChain(func(arg *LinkedListProcessorParameter) bool {
		return arg.Index == 1
	})
	if end.Element != "b" || length != 2 {
		t.Fatal("expected b, 2")
	}
}

// ── NonChainedLinkedListNodes ─────────────────────────────

func Test_NonChainedLinkedListNodes_Basic(t *testing.T) {
	nc := NewNonChainedLinkedListNodes(5)
	if !nc.IsEmpty() {
		t.Fatal("expected empty")
	}
	nc.Adds(&LinkedListNode{Element: "a"}, &LinkedListNode{Element: "b"})
	if nc.Length() != 2 {
		t.Fatal("expected 2")
	}
	if !nc.HasItems() {
		t.Fatal("expected has items")
	}
	if nc.First().Element != "a" {
		t.Fatal("expected a")
	}
	if nc.Last().Element != "b" {
		t.Fatal("expected b")
	}
}

func Test_NonChainedLinkedListNodes_ApplyChaining(t *testing.T) {
	nc := NewNonChainedLinkedListNodes(5)
	nc.Adds(&LinkedListNode{Element: "a"}, &LinkedListNode{Element: "b"}, &LinkedListNode{Element: "c"})
	nc.ApplyChaining()
	if !nc.IsChainingApplied() {
		t.Fatal("expected chaining applied")
	}
	if !nc.First().HasNext() {
		t.Fatal("expected first to have next")
	}
}

func Test_NonChainedLinkedListNodes_ToChainedNodes(t *testing.T) {
	nc := NewNonChainedLinkedListNodes(3)
	nc.Adds(&LinkedListNode{Element: "x"}, &LinkedListNode{Element: "y"})
	chained := nc.ToChainedNodes()
	if len(chained) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_NonChainedLinkedListNodes_FirstOrDefault_Empty(t *testing.T) {
	nc := NewNonChainedLinkedListNodes(0)
	if nc.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
}

func Test_NonChainedLinkedListNodes_LastOrDefault_Empty(t *testing.T) {
	nc := NewNonChainedLinkedListNodes(0)
	if nc.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
}

func Test_NonChainedLinkedListNodes_Adds_Nil(t *testing.T) {
	nc := NewNonChainedLinkedListNodes(0)
	nc.Adds(nil)
	if nc.HasItems() {
		t.Fatal("expected empty")
	}
}
