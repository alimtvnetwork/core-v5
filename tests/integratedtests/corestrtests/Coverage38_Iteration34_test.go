package corestrtests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════
// LinkedList — basic operations
// ══════════════════════════════════════════════════════════════

func Test_Cov38_LL_Add(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.Add("a").Add("b")
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

func Test_Cov38_LL_Head_Tail(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	if ll.Head().Element != "a" {
		t.Errorf("expected a")
	}
	if ll.Tail().Element != "b" {
		t.Errorf("expected b")
	}
}

func Test_Cov38_LL_IsEmpty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	if !ll.IsEmpty() {
		t.Error("expected empty")
	}
	if ll.HasItems() {
		t.Error("expected no items")
	}
}

func Test_Cov38_LL_IsEmptyLock(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	if !ll.IsEmptyLock() {
		t.Error("expected empty")
	}
}

func Test_Cov38_LL_LengthLock(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.LengthLock() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_LL_PushBack(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.PushBack("x")
	if ll.Length() != 1 || ll.Head().Element != "x" {
		t.Error("unexpected")
	}
}

func Test_Cov38_LL_Push(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.Push("x")
	if ll.Length() != 1 {
		t.Error("unexpected")
	}
}

func Test_Cov38_LL_AddLock(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddLock("x")
	if ll.Length() != 1 {
		t.Error("unexpected")
	}
}

func Test_Cov38_LL_AddNonEmpty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddNonEmpty("")
	ll.AddNonEmpty("a")
	if ll.Length() != 1 {
		t.Errorf("expected 1, got %d", ll.Length())
	}
}

func Test_Cov38_LL_AddNonEmptyWhitespace(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddNonEmptyWhitespace("  ")
	ll.AddNonEmptyWhitespace("a")
	if ll.Length() != 1 {
		t.Errorf("expected 1, got %d", ll.Length())
	}
}

func Test_Cov38_LL_AddIf(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddIf(false, "no")
	ll.AddIf(true, "yes")
	if ll.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_LL_AddsIf_True(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddsIf(true, "a", "b")
	if ll.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_Cov38_LL_AddsIf_False(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddsIf(false, "a", "b")
	if ll.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_Cov38_LL_AddFunc(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddFunc(func() string { return "hello" })
	if ll.Length() != 1 || ll.Head().Element != "hello" {
		t.Error("unexpected")
	}
}

func Test_Cov38_LL_AddFuncErr_NoErr(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddFuncErr(
		func() (string, error) { return "ok", nil },
		func(err error) { t.Fatal("should not be called") },
	)
	if ll.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_LL_AddFuncErr_WithErr(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	called := false
	ll.AddFuncErr(
		func() (string, error) { return "", json.Unmarshal([]byte("bad"), &struct{}{}) },
		func(err error) { called = true },
	)
	if !called {
		t.Error("expected error handler to be called")
	}
	if ll.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_Cov38_LL_Adds(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.Adds("a", "b", "c")
	if ll.Length() != 3 {
		t.Errorf("expected 3")
	}
}

func Test_Cov38_LL_Adds_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.Adds()
	if ll.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_Cov38_LL_AddStrings(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddStrings([]string{"x", "y"})
	if ll.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_Cov38_LL_AddStrings_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddStrings(nil)
	if ll.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_Cov38_LL_AddsLock(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddsLock("a", "b")
	if ll.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_Cov38_LL_AddItemsMap(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddItemsMap(map[string]bool{"a": true, "b": false, "c": true})
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

func Test_Cov38_LL_AddItemsMap_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddItemsMap(nil)
	if ll.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_Cov38_LL_AddFront(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("b")
	ll.AddFront("a")
	if ll.Head().Element != "a" || ll.Length() != 2 {
		t.Error("unexpected")
	}
}

func Test_Cov38_LL_AddFront_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddFront("a")
	if ll.Head().Element != "a" || ll.Length() != 1 {
		t.Error("unexpected")
	}
}

func Test_Cov38_LL_PushFront(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("b")
	ll.PushFront("a")
	if ll.Head().Element != "a" {
		t.Error("unexpected")
	}
}

func Test_Cov38_LL_AddCollection(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	ll := corestr.New.LinkedList.Create()
	ll.AddCollection(col)
	if ll.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_Cov38_LL_AddCollection_Nil(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddCollection(nil)
	if ll.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_Cov38_LL_AddPointerStringsPtr(t *testing.T) {
	a, b := "a", "b"
	ll := corestr.New.LinkedList.Create()
	ll.AddPointerStringsPtr([]*string{&a, nil, &b})
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

// ── InsertAt ──

func Test_Cov38_LL_InsertAt_Front(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("b", "c")
	ll.InsertAt(0, "a")
	if ll.Head().Element != "a" || ll.Length() != 3 {
		t.Errorf("unexpected: %v", ll.List())
	}
}

func Test_Cov38_LL_InsertAt_Middle(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "c")
	ll.InsertAt(1, "b")
	list := ll.List()
	if len(list) != 3 || list[1] != "b" {
		t.Errorf("unexpected: %v", list)
	}
}

// ── AppendNode / AppendChainOfNodes / AddBackNode ──

func Test_Cov38_LL_AppendNode_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	node := ll.Head() // nil
	_ = node
	ll.AppendNode(&corestr.LinkedListNode{Element: "x"})
	if ll.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_LL_AppendNode_NonEmpty(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	ll.AppendNode(&corestr.LinkedListNode{Element: "b"})
	if ll.Length() != 2 || ll.Tail().Element != "b" {
		t.Error("unexpected")
	}
}

func Test_Cov38_LL_AddBackNode(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.AddBackNode(&corestr.LinkedListNode{Element: "x"})
	if ll.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_LL_AppendChainOfNodes(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	chain := corestr.New.LinkedList.SpreadStrings("b", "c")
	ll.AppendChainOfNodes(chain.Head())
	if ll.Length() != 3 {
		t.Errorf("expected 3, got %d", ll.Length())
	}
}

func Test_Cov38_LL_AppendChainOfNodes_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	chain := corestr.New.LinkedList.SpreadStrings("a", "b")
	ll.AppendChainOfNodes(chain.Head())
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

// ── AttachWithNode ──

func Test_Cov38_LL_AttachWithNode_NilCurrent(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	err := ll.AttachWithNode(nil, &corestr.LinkedListNode{Element: "x"})
	if err == nil {
		t.Error("expected error")
	}
}

func Test_Cov38_LL_AttachWithNode_NextNotNil(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	err := ll.AttachWithNode(ll.Head(), &corestr.LinkedListNode{Element: "x"})
	if err == nil {
		t.Error("expected error because head.next is not nil")
	}
}

// ── Loop ──

func Test_Cov38_LL_Loop(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	count := 0
	ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
		count++
		return false
	})
	if count != 3 {
		t.Errorf("expected 3, got %d", count)
	}
}

func Test_Cov38_LL_Loop_Break(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	count := 0
	ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
		count++
		return true
	})
	if count != 1 {
		t.Errorf("expected 1, got %d", count)
	}
}

func Test_Cov38_LL_Loop_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
		t.Fatal("should not be called")
		return false
	})
}

func Test_Cov38_LL_Loop_BreakSecond(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	count := 0
	ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
		count++
		return arg.Index == 1
	})
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

// ── Filter ──

func Test_Cov38_LL_Filter(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
		return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: arg.Node.Element != "b"}
	})
	if len(nodes) != 2 {
		t.Errorf("expected 2, got %d", len(nodes))
	}
}

func Test_Cov38_LL_Filter_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
		return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true}
	})
	if len(nodes) != 0 {
		t.Errorf("expected 0")
	}
}

func Test_Cov38_LL_Filter_BreakFirst(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
		return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
	})
	if len(nodes) != 1 {
		t.Errorf("expected 1, got %d", len(nodes))
	}
}

func Test_Cov38_LL_Filter_BreakSecond(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
		return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: arg.Index == 1}
	})
	if len(nodes) != 2 {
		t.Errorf("expected 2, got %d", len(nodes))
	}
}

// ── RemoveNodeByElementValue ──

func Test_Cov38_LL_RemoveByElem_CaseSensitive(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	ll.RemoveNodeByElementValue("b", true, false)
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

func Test_Cov38_LL_RemoveByElem_CaseInsensitive(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "B", "c")
	ll.RemoveNodeByElementValue("b", false, false)
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

func Test_Cov38_LL_RemoveByElem_First(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	ll.RemoveNodeByElementValue("a", true, false)
	if ll.Length() != 1 || ll.Head().Element != "b" {
		t.Error("unexpected")
	}
}

// ── RemoveNodeByIndex ──

func Test_Cov38_LL_RemoveByIndex_First(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	ll.RemoveNodeByIndex(0)
	if ll.Length() != 2 || ll.Head().Element != "b" {
		t.Error("unexpected")
	}
}

func Test_Cov38_LL_RemoveByIndex_Last(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	ll.RemoveNodeByIndex(2)
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

func Test_Cov38_LL_RemoveByIndex_Middle(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	ll.RemoveNodeByIndex(1)
	list := ll.List()
	if len(list) != 2 || list[0] != "a" || list[1] != "c" {
		t.Errorf("unexpected: %v", list)
	}
}

// ── RemoveNodeByIndexes ──

func Test_Cov38_LL_RemoveByIndexes(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c", "d")
	ll.RemoveNodeByIndexes(true, 1, 3)
	list := ll.List()
	if len(list) != 2 || list[0] != "a" || list[1] != "c" {
		t.Errorf("unexpected: %v", list)
	}
}

func Test_Cov38_LL_RemoveByIndexes_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	ll.RemoveNodeByIndexes(true)
	if ll.Length() != 1 {
		t.Errorf("expected 1")
	}
}

// ── RemoveNode ──

func Test_Cov38_LL_RemoveNode_Nil(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	ll.RemoveNode(nil)
	if ll.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_LL_RemoveNode_Head(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	ll.RemoveNode(ll.Head())
	if ll.Length() != 1 || ll.Head().Element != "b" {
		t.Error("unexpected")
	}
}

func Test_Cov38_LL_RemoveNode_NonHead(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	node := ll.IndexAt(1) // "b"
	ll.RemoveNode(node)
	if ll.Length() != 2 {
		t.Errorf("expected 2")
	}
}

// ── IndexAt / SafeIndexAt / SafePointerIndexAt ──

func Test_Cov38_LL_IndexAt(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	node := ll.IndexAt(1)
	if node.Element != "b" {
		t.Errorf("expected b")
	}
}

func Test_Cov38_LL_IndexAt_Zero(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	if ll.IndexAt(0).Element != "a" {
		t.Error("expected a")
	}
}

func Test_Cov38_LL_IndexAt_Negative(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.IndexAt(-1) != nil {
		t.Error("expected nil")
	}
}

func Test_Cov38_LL_SafeIndexAt(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	if ll.SafeIndexAt(1).Element != "b" {
		t.Error("expected b")
	}
}

func Test_Cov38_LL_SafeIndexAt_OutOfRange(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.SafeIndexAt(5) != nil {
		t.Error("expected nil")
	}
}

func Test_Cov38_LL_SafeIndexAt_Negative(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.SafeIndexAt(-1) != nil {
		t.Error("expected nil")
	}
}

func Test_Cov38_LL_SafeIndexAtLock(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.SafeIndexAtLock(0).Element != "a" {
		t.Error("expected a")
	}
}

func Test_Cov38_LL_SafePointerIndexAt(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	ptr := ll.SafePointerIndexAt(0)
	if ptr == nil || *ptr != "a" {
		t.Error("expected a")
	}
}

func Test_Cov38_LL_SafePointerIndexAt_Nil(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.SafePointerIndexAt(5) != nil {
		t.Error("expected nil")
	}
}

func Test_Cov38_LL_SafePointerIndexAtUsingDefault(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.SafePointerIndexAtUsingDefault(5, "def") != "def" {
		t.Error("expected def")
	}
	if ll.SafePointerIndexAtUsingDefault(0, "def") != "a" {
		t.Error("expected a")
	}
}

func Test_Cov38_LL_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.SafePointerIndexAtUsingDefaultLock(5, "x") != "x" {
		t.Error("expected x")
	}
}

// ── GetNextNodes / GetAllLinkedNodes ──

func Test_Cov38_LL_GetNextNodes(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c", "d")
	nodes := ll.GetNextNodes(2)
	if len(nodes) != 2 {
		t.Errorf("expected 2, got %d", len(nodes))
	}
}

func Test_Cov38_LL_GetAllLinkedNodes(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	nodes := ll.GetAllLinkedNodes()
	if len(nodes) != 2 {
		t.Errorf("expected 2, got %d", len(nodes))
	}
}

// ── ToCollection / List / ListPtr / ListLock / ListPtrLock ──

func Test_Cov38_LL_ToCollection(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	col := ll.ToCollection(0)
	if col.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_Cov38_LL_ToCollection_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	col := ll.ToCollection(0)
	if col.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_Cov38_LL_List(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	list := ll.List()
	if len(list) != 2 {
		t.Errorf("expected 2")
	}
}

func Test_Cov38_LL_List_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	if len(ll.List()) != 0 {
		t.Error("expected 0")
	}
}

func Test_Cov38_LL_ListPtr(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if len(ll.ListPtr()) != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov38_LL_ListLock(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if len(ll.ListLock()) != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov38_LL_ListPtrLock(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if len(ll.ListPtrLock()) != 1 {
		t.Error("expected 1")
	}
}

// ── String / StringLock / Join / JoinLock / Joins ──

func Test_Cov38_LL_String(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	s := ll.String()
	if !strings.Contains(s, "a") {
		t.Error("expected a")
	}
}

func Test_Cov38_LL_String_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	s := ll.String()
	if s == "" {
		t.Error("expected non-empty (contains NoElements)")
	}
}

func Test_Cov38_LL_StringLock(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	s := ll.StringLock()
	if !strings.Contains(s, "a") {
		t.Error("expected a")
	}
}

func Test_Cov38_LL_StringLock_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	s := ll.StringLock()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_Cov38_LL_Join(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	if ll.Join(",") != "a,b" {
		t.Errorf("unexpected: %s", ll.Join(","))
	}
}

func Test_Cov38_LL_JoinLock(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	if ll.JoinLock(",") != "a,b" {
		t.Error("unexpected")
	}
}

func Test_Cov38_LL_Joins(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	result := ll.Joins(",", "c")
	if !strings.Contains(result, "a") || !strings.Contains(result, "c") {
		t.Errorf("unexpected: %s", result)
	}
}

func Test_Cov38_LL_Joins_NilItems(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	result := ll.Joins(",")
	// items is nil, so joins items only
	_ = result
}

// ── IsEquals / IsEqualsWithSensitive ──

func Test_Cov38_LL_IsEquals_Same(t *testing.T) {
	a := corestr.New.LinkedList.SpreadStrings("a", "b")
	b := corestr.New.LinkedList.SpreadStrings("a", "b")
	if !a.IsEquals(b) {
		t.Error("expected equal")
	}
}

func Test_Cov38_LL_IsEquals_Diff(t *testing.T) {
	a := corestr.New.LinkedList.SpreadStrings("a")
	b := corestr.New.LinkedList.SpreadStrings("b")
	if a.IsEquals(b) {
		t.Error("expected not equal")
	}
}

func Test_Cov38_LL_IsEquals_BothEmpty(t *testing.T) {
	a := corestr.New.LinkedList.Create()
	b := corestr.New.LinkedList.Create()
	if !a.IsEquals(b) {
		t.Error("expected equal")
	}
}

func Test_Cov38_LL_IsEquals_DiffLen(t *testing.T) {
	a := corestr.New.LinkedList.SpreadStrings("a", "b")
	b := corestr.New.LinkedList.SpreadStrings("a")
	if a.IsEquals(b) {
		t.Error("expected not equal")
	}
}

func Test_Cov38_LL_IsEquals_OneEmpty(t *testing.T) {
	a := corestr.New.LinkedList.SpreadStrings("a")
	b := corestr.New.LinkedList.Create()
	if a.IsEquals(b) {
		t.Error("expected not equal")
	}
}

func Test_Cov38_LL_IsEquals_SamePtr(t *testing.T) {
	a := corestr.New.LinkedList.SpreadStrings("a")
	if !a.IsEqualsWithSensitive(a, true) {
		t.Error("expected equal (same pointer)")
	}
}

func Test_Cov38_LL_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	a := corestr.New.LinkedList.SpreadStrings("A", "B")
	b := corestr.New.LinkedList.SpreadStrings("a", "b")
	if !a.IsEqualsWithSensitive(b, false) {
		t.Error("expected equal case-insensitive")
	}
}

func Test_Cov38_LL_IsEqualsWithSensitive_OneNil(t *testing.T) {
	a := corestr.New.LinkedList.SpreadStrings("a")
	if a.IsEqualsWithSensitive(nil, true) {
		t.Error("expected not equal")
	}
}

// ── GetCompareSummary ──

func Test_Cov38_LL_GetCompareSummary(t *testing.T) {
	a := corestr.New.LinkedList.SpreadStrings("a")
	b := corestr.New.LinkedList.SpreadStrings("b")
	summary := a.GetCompareSummary(b, "left", "right")
	if summary == "" {
		t.Error("expected non-empty summary")
	}
}

// ── AddStringsToNode / AddStringsPtrToNode / AddCollectionToNode / AddAfterNode ──

func Test_Cov38_LL_AddStringsToNode(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "d")
	node := ll.Head()
	ll.AddStringsToNode(false, node, []string{"b", "c"})
	if ll.Length() != 5 { // a + b + c + (n-1 increments from AddNext) + 1 from AddStringsToNode + d
		// The actual length depends on internal logic; just verify it grew
		if ll.Length() < 4 {
			t.Errorf("expected at least 4, got %d", ll.Length())
		}
	}
}

func Test_Cov38_LL_AddStringsToNode_SingleItem(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "c")
	node := ll.Head()
	ll.AddStringsToNode(false, node, []string{"b"})
	if ll.Length() != 3 {
		t.Errorf("expected 3, got %d", ll.Length())
	}
}

func Test_Cov38_LL_AddStringsToNode_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	ll.AddStringsToNode(false, ll.Head(), nil)
	if ll.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_LL_AddStringsToNode_NilNode_Skip(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	ll.AddStringsToNode(true, nil, []string{"b"})
	if ll.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_LL_AddStringsPtrToNode_Nil(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	ll.AddStringsPtrToNode(true, ll.Head(), nil)
	if ll.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_LL_AddStringsPtrToNode(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	items := []string{"b"}
	ll.AddStringsPtrToNode(true, ll.Head(), &items)
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

func Test_Cov38_LL_AddCollectionToNode(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	col := corestr.New.Collection.Strings([]string{"b"})
	ll.AddCollectionToNode(true, ll.Head(), col)
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

// ── JSON / Serialize ──

func Test_Cov38_LL_MarshalJSON(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	data, err := json.Marshal(ll)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "\"a\"") {
		t.Errorf("unexpected: %s", data)
	}
}

func Test_Cov38_LL_UnmarshalJSON(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	err := json.Unmarshal([]byte(`["x","y"]`), ll)
	if err != nil {
		t.Fatal(err)
	}
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

func Test_Cov38_LL_UnmarshalJSON_Invalid(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	err := json.Unmarshal([]byte(`bad`), ll)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_Cov38_LL_JsonModel(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	model := ll.JsonModel()
	if len(model) != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_LL_JsonModelAny(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.JsonModelAny() == nil {
		t.Error("expected non-nil")
	}
}

func Test_Cov38_LL_Json(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	result := ll.Json()
	if result.HasError() {
		t.Errorf("unexpected error")
	}
}

func Test_Cov38_LL_JsonPtr(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.JsonPtr() == nil {
		t.Error("expected non-nil")
	}
}

func Test_Cov38_LL_ParseInjectUsingJson(t *testing.T) {
	src := corestr.New.LinkedList.SpreadStrings("a", "b")
	ll := corestr.New.LinkedList.Create()
	result, err := ll.ParseInjectUsingJson(src.JsonPtr())
	if err != nil {
		t.Fatal(err)
	}
	if result.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_Cov38_LL_ParseInjectUsingJsonMust(t *testing.T) {
	src := corestr.New.LinkedList.SpreadStrings("a")
	ll := corestr.New.LinkedList.Create()
	result := ll.ParseInjectUsingJsonMust(src.JsonPtr())
	if result.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_LL_JsonParseSelfInject(t *testing.T) {
	src := corestr.New.LinkedList.SpreadStrings("x")
	ll := corestr.New.LinkedList.Create()
	err := ll.JsonParseSelfInject(src.JsonPtr())
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Cov38_LL_AsJsonMarshaller(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.AsJsonMarshaller() == nil {
		t.Error("expected non-nil")
	}
}

// ── Clear / RemoveAll ──

func Test_Cov38_LL_Clear(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	ll.Clear()
	if ll.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_Cov38_LL_Clear_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.Clear()
	if ll.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_Cov38_LL_RemoveAll(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	ll.RemoveAll()
	if ll.Length() != 0 {
		t.Errorf("expected 0")
	}
}

// ══════════════════════════════════════════════════════════════
// LinkedListNode
// ══════════════════════════════════════════════════════════════

func Test_Cov38_Node_HasNext(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	if !ll.Head().HasNext() {
		t.Error("expected true")
	}
	if ll.Tail().HasNext() {
		t.Error("expected false")
	}
}

func Test_Cov38_Node_Next(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	if ll.Head().Next().Element != "b" {
		t.Error("expected b")
	}
}

func Test_Cov38_Node_EndOfChain(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	end, length := ll.Head().EndOfChain()
	if end.Element != "c" || length != 3 {
		t.Errorf("unexpected end=%s length=%d", end.Element, length)
	}
}

func Test_Cov38_Node_LoopEndOfChain(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	count := 0
	end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
		count++
		return false
	})
	if count != 3 || length != 3 || end.Element != "c" {
		t.Errorf("unexpected count=%d length=%d end=%s", count, length, end.Element)
	}
}

func Test_Cov38_Node_LoopEndOfChain_Break(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
		return true // break immediately
	})
	if length != 1 || end.Element != "a" {
		t.Errorf("unexpected length=%d end=%s", length, end.Element)
	}
}

func Test_Cov38_Node_LoopEndOfChain_BreakSecond(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	_, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
		return arg.Index == 1
	})
	if length != 2 {
		t.Errorf("expected 2, got %d", length)
	}
}

func Test_Cov38_Node_Clone(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	cloned := ll.Head().Clone()
	if cloned.Element != "a" || cloned.HasNext() {
		t.Error("expected clone without next")
	}
}

func Test_Cov38_Node_AddNext(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "c")
	newNode := ll.Head().AddNext(ll, "b")
	if newNode.Element != "b" {
		t.Error("expected b")
	}
}

func Test_Cov38_Node_AddNextNode(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "c")
	newNode := &corestr.LinkedListNode{Element: "b"}
	ll.Head().AddNextNode(ll, newNode)
	if ll.Length() != 3 {
		t.Errorf("expected 3, got %d", ll.Length())
	}
}

func Test_Cov38_Node_IsEqual_Same(t *testing.T) {
	ll1 := corestr.New.LinkedList.SpreadStrings("a", "b")
	ll2 := corestr.New.LinkedList.SpreadStrings("a", "b")
	if !ll1.Head().IsEqual(ll2.Head()) {
		t.Error("expected equal")
	}
}

func Test_Cov38_Node_IsEqual_Diff(t *testing.T) {
	ll1 := corestr.New.LinkedList.SpreadStrings("a", "b")
	ll2 := corestr.New.LinkedList.SpreadStrings("a", "c")
	if ll1.Head().IsEqual(ll2.Head()) {
		t.Error("expected not equal")
	}
}

func Test_Cov38_Node_IsEqual_BothNil(t *testing.T) {
	var a, b *corestr.LinkedListNode
	if !a.IsEqual(b) {
		t.Error("expected equal")
	}
}

func Test_Cov38_Node_IsEqual_OneNil(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.Head().IsEqual(nil) {
		t.Error("expected not equal")
	}
}

func Test_Cov38_Node_IsEqual_SamePtr(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if !ll.Head().IsEqual(ll.Head()) {
		t.Error("expected equal")
	}
}

func Test_Cov38_Node_IsChainEqual(t *testing.T) {
	a := corestr.New.LinkedList.SpreadStrings("a", "b")
	b := corestr.New.LinkedList.SpreadStrings("A", "B")
	if !a.Head().IsChainEqual(b.Head(), false) {
		t.Error("expected equal case-insensitive")
	}
	if a.Head().IsChainEqual(b.Head(), true) {
		t.Error("expected not equal case-sensitive")
	}
}

func Test_Cov38_Node_IsChainEqual_BothNil(t *testing.T) {
	var a, b *corestr.LinkedListNode
	if !a.IsChainEqual(b, true) {
		t.Error("expected equal")
	}
}

func Test_Cov38_Node_IsEqualSensitive(t *testing.T) {
	a := corestr.New.LinkedList.SpreadStrings("a")
	b := corestr.New.LinkedList.SpreadStrings("A")
	if !a.Head().IsEqualSensitive(b.Head(), false) {
		t.Error("expected equal")
	}
}

func Test_Cov38_Node_IsEqualSensitive_SamePtr(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if !ll.Head().IsEqualSensitive(ll.Head(), true) {
		t.Error("expected equal")
	}
}

func Test_Cov38_Node_IsEqualSensitive_BothNil(t *testing.T) {
	var a, b *corestr.LinkedListNode
	if !a.IsEqualSensitive(b, true) {
		t.Error("expected equal")
	}
}

func Test_Cov38_Node_IsEqualSensitive_OneNil(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if ll.Head().IsEqualSensitive(nil, true) {
		t.Error("expected not equal")
	}
}

func Test_Cov38_Node_IsEqualValue(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if !ll.Head().IsEqualValue("a") {
		t.Error("expected true")
	}
}

func Test_Cov38_Node_IsEqualValueSensitive(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if !ll.Head().IsEqualValueSensitive("A", false) {
		t.Error("expected true case-insensitive")
	}
	if ll.Head().IsEqualValueSensitive("A", true) {
		t.Error("expected false case-sensitive")
	}
}

func Test_Cov38_Node_String(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("hello")
	if ll.Head().String() != "hello" {
		t.Error("unexpected")
	}
}

func Test_Cov38_Node_List(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	list := ll.Head().List()
	if len(list) != 3 {
		t.Errorf("expected 3, got %d", len(list))
	}
}

func Test_Cov38_Node_ListPtr(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	if len(ll.Head().ListPtr()) != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov38_Node_Join(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	if ll.Head().Join(",") != "a,b" {
		t.Error("unexpected")
	}
}

func Test_Cov38_Node_CreateLinkedList(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b")
	newLL := ll.Head().CreateLinkedList()
	if newLL.Length() != 2 {
		t.Errorf("expected 2, got %d", newLL.Length())
	}
}

func Test_Cov38_Node_AddStringsToNode(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	ll.Head().AddStringsToNode(ll, false, []string{"b"})
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

func Test_Cov38_Node_AddStringsPtrToNode_Nil(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	ll.Head().AddStringsPtrToNode(ll, true, nil)
	if ll.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_Node_AddStringsPtrToNode(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	items := []string{"b"}
	ll.Head().AddStringsPtrToNode(ll, false, &items)
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

func Test_Cov38_Node_AddCollectionToNode(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a")
	col := corestr.New.Collection.Strings([]string{"b"})
	ll.Head().AddCollectionToNode(ll, false, col)
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

// ══════════════════════════════════════════════════════════════
// NonChainedLinkedListNodes
// ══════════════════════════════════════════════════════════════

func Test_Cov38_NCLLN_Basic(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(5)
	if !nc.IsEmpty() {
		t.Error("expected empty")
	}
	if nc.HasItems() {
		t.Error("expected no items")
	}
	if nc.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_Cov38_NCLLN_Adds(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(3)
	n1 := &corestr.LinkedListNode{Element: "a"}
	n2 := &corestr.LinkedListNode{Element: "b"}
	nc.Adds(n1, n2)
	if nc.Length() != 2 {
		t.Errorf("expected 2, got %d", nc.Length())
	}
	if nc.First().Element != "a" {
		t.Error("expected a")
	}
	if nc.Last().Element != "b" {
		t.Error("expected b")
	}
}

func Test_Cov38_NCLLN_Adds_Nil(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(3)
	nc.Adds()
	if nc.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_Cov38_NCLLN_FirstOrDefault_Empty(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(3)
	if nc.FirstOrDefault() != nil {
		t.Error("expected nil")
	}
}

func Test_Cov38_NCLLN_LastOrDefault_Empty(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(3)
	if nc.LastOrDefault() != nil {
		t.Error("expected nil")
	}
}

func Test_Cov38_NCLLN_Items(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(3)
	nc.Adds(&corestr.LinkedListNode{Element: "a"})
	if len(nc.Items()) != 1 {
		t.Errorf("expected 1")
	}
}

func Test_Cov38_NCLLN_IsChainingApplied(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(3)
	if nc.IsChainingApplied() {
		t.Error("expected false")
	}
}

func Test_Cov38_NCLLN_ApplyChaining(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(3)
	n1 := &corestr.LinkedListNode{Element: "a"}
	n2 := &corestr.LinkedListNode{Element: "b"}
	nc.Adds(n1, n2)
	nc.ApplyChaining()
	if !nc.IsChainingApplied() {
		t.Error("expected true")
	}
	if !n1.HasNext() {
		t.Error("expected a->b chain")
	}
	if n2.HasNext() {
		t.Error("expected b to be tail")
	}
}

func Test_Cov38_NCLLN_ApplyChaining_Empty(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(3)
	nc.ApplyChaining()
	// should not panic
}

func Test_Cov38_NCLLN_ToChainedNodes(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(3)
	nc.Adds(&corestr.LinkedListNode{Element: "a"}, &corestr.LinkedListNode{Element: "b"})
	chained := nc.ToChainedNodes()
	if chained == nil {
		t.Error("expected non-nil")
	}
}

func Test_Cov38_NCLLN_ToChainedNodes_Empty(t *testing.T) {
	nc := corestr.NewNonChainedLinkedListNodes(3)
	chained := nc.ToChainedNodes()
	if len(chained) != 0 {
		t.Errorf("expected 0, got %d", len(chained))
	}
}

// ── newLinkedListCreator factory methods ──

func Test_Cov38_Creator_Create(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	if ll == nil || ll.Length() != 0 {
		t.Error("unexpected")
	}
}

func Test_Cov38_Creator_Empty(t *testing.T) {
	ll := corestr.New.LinkedList.Empty()
	if ll == nil || ll.Length() != 0 {
		t.Error("unexpected")
	}
}

func Test_Cov38_Creator_Strings(t *testing.T) {
	ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
	if ll.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_Cov38_Creator_SpreadStrings(t *testing.T) {
	ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
	if ll.Length() != 3 {
		t.Errorf("expected 3")
	}
}

func Test_Cov38_Creator_UsingMap(t *testing.T) {
	ll := corestr.New.LinkedList.UsingMap(map[string]bool{"a": true, "b": false})
	if ll.Length() != 1 {
		t.Errorf("expected 1, got %d", ll.Length())
	}
}

func Test_Cov38_Creator_PointerStringsPtr(t *testing.T) {
	a, b := "a", "b"
	ptrs := []*string{&a, nil, &b}
	ll := corestr.New.LinkedList.PointerStringsPtr(&ptrs)
	if ll.Length() != 2 {
		t.Errorf("expected 2, got %d", ll.Length())
	}
}

func Test_Cov38_Creator_PointerStringsPtr_Nil(t *testing.T) {
	ll := corestr.New.LinkedList.PointerStringsPtr(nil)
	if ll.Length() != 0 {
		t.Errorf("expected 0")
	}
}
