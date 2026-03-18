package corestr

import (
	"encoding/json"
	"sync"
	"testing"
)

// ── Creators ──────────────────────────────────────────────

func Test_LinkedCollections_NewCreate(t *testing.T) {
	lc := New.LinkedCollection.Create()
	if lc == nil || lc.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedCollections_NewEmpty(t *testing.T) {
	lc := New.LinkedCollection.Empty()
	if lc.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_LinkedCollections_NewStrings(t *testing.T) {
	lc := New.LinkedCollection.Strings("a", "b")
	if lc.Length() != 1 {
		t.Fatal("expected 1 collection node")
	}
}

func Test_LinkedCollections_NewStrings_Empty(t *testing.T) {
	lc := New.LinkedCollection.Strings()
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedCollections_NewPointerStringsPtr(t *testing.T) {
	s1, s2 := "a", "b"
	items := []*string{&s1, &s2}
	lc := New.LinkedCollection.PointerStringsPtr(&items)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_NewPointerStringsPtr_Nil(t *testing.T) {
	lc := New.LinkedCollection.PointerStringsPtr(nil)
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedCollections_NewUsingCollections(t *testing.T) {
	c1 := New.Collection.Strings([]string{"a"})
	c2 := New.Collection.Strings([]string{"b"})
	lc := New.LinkedCollection.UsingCollections(c1, c2)
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_NewUsingCollections_Nil(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(nil)
	if lc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LinkedCollections_EmptyLinkedCollections(t *testing.T) {
	lc := Empty.LinkedCollections()
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── Head / Tail / First / Last ────────────────────────────

func Test_LinkedCollections_HeadTail(t *testing.T) {
	c1 := New.Collection.Strings([]string{"a"})
	c2 := New.Collection.Strings([]string{"b"})
	lc := New.LinkedCollection.UsingCollections(c1, c2)
	if lc.Head() == nil || lc.Tail() == nil {
		t.Fatal("expected non-nil")
	}
	if lc.First().Length() != 1 {
		t.Fatal("expected 1")
	}
	if lc.Last().Length() != 1 {
		t.Fatal("expected 1")
	}
	if lc.Single().Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_FirstOrDefault(t *testing.T) {
	lc := Empty.LinkedCollections()
	col := lc.FirstOrDefault()
	if col == nil {
		t.Fatal("expected non-nil default")
	}
}

func Test_LinkedCollections_LastOrDefault(t *testing.T) {
	lc := Empty.LinkedCollections()
	col := lc.LastOrDefault()
	if col == nil {
		t.Fatal("expected non-nil default")
	}
}

func Test_LinkedCollections_FirstOrDefault_HasItems(t *testing.T) {
	c1 := New.Collection.Strings([]string{"x"})
	lc := New.LinkedCollection.UsingCollections(c1)
	if lc.FirstOrDefault().Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ── Length / LengthLock / AllIndividualItemsLength ─────────

func Test_LinkedCollections_LengthLock(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	if lc.LengthLock() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AllIndividualItemsLength(t *testing.T) {
	c1 := New.Collection.Strings([]string{"a", "b"})
	c2 := New.Collection.Strings([]string{"c"})
	lc := New.LinkedCollection.UsingCollections(c1, c2)
	if lc.AllIndividualItemsLength() != 3 {
		t.Fatal("expected 3")
	}
}

// ── IsEmpty / HasItems / IsEmptyLock ──────────────────────

func Test_LinkedCollections_IsEmpty(t *testing.T) {
	lc := Empty.LinkedCollections()
	if !lc.IsEmpty() || lc.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedCollections_IsEmptyLock(t *testing.T) {
	lc := Empty.LinkedCollections()
	if !lc.IsEmptyLock() {
		t.Fatal("expected empty")
	}
}

// ── Add / AddLock / AddStrings / AddStringsLock ───────────

func Test_LinkedCollections_Add(t *testing.T) {
	lc := New.LinkedCollection.Create()
	col := New.Collection.Strings([]string{"a"})
	lc.Add(col)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
	lc.Add(New.Collection.Strings([]string{"b"}))
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_AddLock(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddLock(New.Collection.Strings([]string{"a"}))
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AddStrings(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddStrings("a", "b")
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AddStrings_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddStrings()
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedCollections_AddStringsLock(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddStringsLock("a")
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AddStringsLock_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddStringsLock()
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AddFront / PushFront / PushBack / Push ────────────────

func Test_LinkedCollections_AddFront(t *testing.T) {
	c1 := New.Collection.Strings([]string{"b"})
	lc := New.LinkedCollection.UsingCollections(c1)
	c0 := New.Collection.Strings([]string{"a"})
	lc.AddFront(c0)
	if lc.First().List()[0] != "a" {
		t.Fatal("expected a at front")
	}
}

func Test_LinkedCollections_AddFront_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddFront(New.Collection.Strings([]string{"a"}))
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AddFrontLock(t *testing.T) {
	lc := New.LinkedCollection.Strings("b")
	lc.AddFrontLock(New.Collection.Strings([]string{"a"}))
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_PushFront(t *testing.T) {
	lc := New.LinkedCollection.Strings("b")
	lc.PushFront(New.Collection.Strings([]string{"a"}))
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_PushBack(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.PushBack(New.Collection.Strings([]string{"a"}))
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_PushBackLock(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.PushBackLock(New.Collection.Strings([]string{"a"}))
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_Push(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.Push(New.Collection.Strings([]string{"a"}))
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ── AppendNode / AppendChainOfNodes ───────────────────────

func Test_LinkedCollections_AppendNode(t *testing.T) {
	lc := New.LinkedCollection.Create()
	node := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	lc.AppendNode(node)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
	lc.AppendNode(&LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})})
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_AddBackNode(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddBackNode(&LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})})
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AppendChainOfNodes(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	chain := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"b"}),
		next: &LinkedCollectionNode{
			Element: New.Collection.Strings([]string{"c"}),
		},
	}
	lc.AppendChainOfNodes(chain)
	if lc.Length() != 3 {
		t.Fatalf("expected 3, got %d", lc.Length())
	}
}

func Test_LinkedCollections_AppendChainOfNodes_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	chain := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	lc.AppendChainOfNodes(chain)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ── InsertAt ──────────────────────────────────────────────

func Test_LinkedCollections_InsertAt_Front(t *testing.T) {
	lc := New.LinkedCollection.Strings("b")
	lc.InsertAt(0, New.Collection.Strings([]string{"a"}))
	if lc.First().List()[0] != "a" {
		t.Fatal("expected a at front")
	}
}

func Test_LinkedCollections_InsertAt_Middle(t *testing.T) {
	c1 := New.Collection.Strings([]string{"a"})
	c3 := New.Collection.Strings([]string{"c"})
	lc := New.LinkedCollection.UsingCollections(c1, c3)
	c2 := New.Collection.Strings([]string{"b"})
	lc.InsertAt(1, c2)
	if lc.Length() != 3 {
		t.Fatal("expected 3")
	}
}

// ── AttachWithNode ────────────────────────────────────────

func Test_LinkedCollections_AttachWithNode_NilCurrent(t *testing.T) {
	lc := New.LinkedCollection.Create()
	err := lc.AttachWithNode(nil, &LinkedCollectionNode{Element: New.Collection.Strings([]string{"x"})})
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_LinkedCollections_AttachWithNode_NonNilNext(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	err := lc.AttachWithNode(lc.Head(), &LinkedCollectionNode{Element: New.Collection.Strings([]string{"x"})})
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_LinkedCollections_AttachWithNode_Success(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	node := lc.Tail()
	err := lc.AttachWithNode(node, &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})})
	if err != nil {
		t.Fatal(err)
	}
}

// ── AddAnother ────────────────────────────────────────────

func Test_LinkedCollections_AddAnother(t *testing.T) {
	lc1 := New.LinkedCollection.Strings("a")
	lc2 := New.LinkedCollection.Strings("b")
	lc1.AddAnother(lc2)
	if lc1.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_AddAnother_Nil(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	lc.AddAnother(nil)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ── AddCollection / AddCollectionsPtr / AddCollections ─────

func Test_LinkedCollections_AddCollection(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddCollection(New.Collection.Strings([]string{"a"}))
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AddCollection_Nil(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddCollection(nil)
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedCollections_AddCollectionsPtr(t *testing.T) {
	lc := New.LinkedCollection.Create()
	cols := []*Collection{New.Collection.Strings([]string{"a"})}
	lc.AddCollectionsPtr(cols)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AddCollections(t *testing.T) {
	lc := New.LinkedCollection.Create()
	cols := []*Collection{New.Collection.Strings([]string{"a"})}
	lc.AddCollections(cols)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AddCollections_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddCollections([]*Collection{})
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AppendCollections / AppendCollectionsPointers ──────────

func Test_LinkedCollections_AppendCollections(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AppendCollections(true, New.Collection.Strings([]string{"a"}), nil)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AppendCollections_NilInput(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AppendCollections(true, nil)
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedCollections_AppendCollectionsPointers(t *testing.T) {
	lc := New.LinkedCollection.Create()
	cols := []*Collection{New.Collection.Strings([]string{"a"}), nil}
	lc.AppendCollectionsPointers(true, &cols)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AppendCollectionsPointersLock(t *testing.T) {
	lc := New.LinkedCollection.Create()
	cols := []*Collection{New.Collection.Strings([]string{"a"})}
	lc.AppendCollectionsPointersLock(true, &cols)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ── Loop ──────────────────────────────────────────────────

func Test_LinkedCollections_Loop(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	count := 0
	lc.Loop(func(arg *LinkedCollectionProcessorParameter) bool {
		count++
		return false
	})
	if count != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_Loop_Break(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	count := 0
	lc.Loop(func(arg *LinkedCollectionProcessorParameter) bool {
		count++
		return true
	})
	if count != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_Loop_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.Loop(func(arg *LinkedCollectionProcessorParameter) bool {
		t.Fatal("should not be called")
		return false
	})
}

// ── Filter / FilterAsCollection / FilterAsCollections ─────

func Test_LinkedCollections_Filter(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	nodes := lc.Filter(func(arg *LinkedCollectionFilterParameter) *LinkedCollectionFilterResult {
		return &LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
	})
	if len(nodes) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_Filter_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	nodes := lc.Filter(func(arg *LinkedCollectionFilterParameter) *LinkedCollectionFilterResult {
		return &LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
	})
	if len(nodes) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_LinkedCollections_FilterAsCollection(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a", "b"}),
		New.Collection.Strings([]string{"c"}),
	)
	col := lc.FilterAsCollection(func(arg *LinkedCollectionFilterParameter) *LinkedCollectionFilterResult {
		return &LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
	}, 0)
	if col.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_LinkedCollections_FilterAsCollection_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	col := lc.FilterAsCollection(func(arg *LinkedCollectionFilterParameter) *LinkedCollectionFilterResult {
		return &LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
	}, 0)
	if col.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedCollections_FilterAsCollections(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	cols := lc.FilterAsCollections(func(arg *LinkedCollectionFilterParameter) *LinkedCollectionFilterResult {
		return &LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
	})
	if len(cols) != 2 {
		t.Fatal("expected 2")
	}
}

// ── GetNextNodes / GetAllLinkedNodes ───────────────────────

func Test_LinkedCollections_GetNextNodes(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
		New.Collection.Strings([]string{"c"}),
	)
	nodes := lc.GetNextNodes(2)
	if len(nodes) != 2 {
		t.Fatalf("expected 2, got %d", len(nodes))
	}
}

func Test_LinkedCollections_GetAllLinkedNodes(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	nodes := lc.GetAllLinkedNodes()
	if len(nodes) != 2 {
		t.Fatal("expected 2")
	}
}

// ── RemoveNodeByIndex ─────────────────────────────────────

func Test_LinkedCollections_RemoveNodeByIndex_First(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	lc.RemoveNodeByIndex(0)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_RemoveNodeByIndex_Last(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	lc.RemoveNodeByIndex(1)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_RemoveNodeByIndex_Middle(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
		New.Collection.Strings([]string{"c"}),
	)
	lc.RemoveNodeByIndex(1)
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

// ── RemoveNodeByIndexes ───────────────────────────────────

func Test_LinkedCollections_RemoveNodeByIndexes(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
		New.Collection.Strings([]string{"c"}),
	)
	lc.RemoveNodeByIndexes(true, 0, 2)
	if lc.Length() != 1 {
		t.Fatalf("expected 1, got %d", lc.Length())
	}
}

func Test_LinkedCollections_RemoveNodeByIndexes_Empty(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	lc.RemoveNodeByIndexes(true)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ── RemoveNode ────────────────────────────────────────────

func Test_LinkedCollections_RemoveNode_First(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	lc.RemoveNode(lc.Head())
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_RemoveNode_Middle(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
		New.Collection.Strings([]string{"c"}),
	)
	lc.RemoveNode(lc.Head().Next())
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

// ── AddAfterNode ──────────────────────────────────────────

func Test_LinkedCollections_AddAfterNode(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	lc.AddAfterNode(lc.Head(), New.Collection.Strings([]string{"b"}))
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

// ── ConcatNew ─────────────────────────────────────────────

func Test_LinkedCollections_ConcatNew(t *testing.T) {
	lc1 := New.LinkedCollection.Strings("a")
	lc2 := New.LinkedCollection.Strings("b")
	result := lc1.ConcatNew(false, lc2)
	if result.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_ConcatNew_EmptyClone(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	result := lc.ConcatNew(true)
	if result.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_ConcatNew_EmptyNoClone(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	result := lc.ConcatNew(false)
	if result != lc {
		t.Fatal("expected same ref")
	}
}

// ── IndexAt / SafeIndexAt / SafePointerIndexAt ────────────

func Test_LinkedCollections_IndexAt(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	node := lc.IndexAt(1)
	if node == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LinkedCollections_IndexAt_Zero(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	node := lc.IndexAt(0)
	if node == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LinkedCollections_IndexAt_Negative(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	node := lc.IndexAt(-1)
	if node != nil {
		t.Fatal("expected nil")
	}
}

func Test_LinkedCollections_SafeIndexAt(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	node := lc.SafeIndexAt(1)
	if node == nil {
		t.Fatal("expected non-nil")
	}
	node2 := lc.SafeIndexAt(5)
	if node2 != nil {
		t.Fatal("expected nil")
	}
	node3 := lc.SafeIndexAt(-1)
	if node3 != nil {
		t.Fatal("expected nil")
	}
}

func Test_LinkedCollections_SafePointerIndexAt(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
	)
	col := lc.SafePointerIndexAt(0)
	if col == nil {
		t.Fatal("expected non-nil")
	}
	col2 := lc.SafePointerIndexAt(5)
	if col2 != nil {
		t.Fatal("expected nil")
	}
}

// ── ToCollection / ToCollectionSimple / ToStrings ──────────

func Test_LinkedCollections_ToCollection(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a", "b"}),
		New.Collection.Strings([]string{"c"}),
	)
	col := lc.ToCollection(0)
	if col.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_LinkedCollections_ToCollection_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	col := lc.ToCollection(0)
	if col.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedCollections_ToCollectionSimple(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	col := lc.ToCollectionSimple()
	if col.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_ToStrings(t *testing.T) {
	lc := New.LinkedCollection.Strings("a", "b")
	strs := lc.ToStrings()
	if len(strs) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_ToStringsPtr(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	ptr := lc.ToStringsPtr()
	if ptr == nil || len(*ptr) != 1 {
		t.Fatal("expected 1")
	}
}

// ── ToCollectionsOfCollection ─────────────────────────────

func Test_LinkedCollections_ToCollectionsOfCollection(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	coc := lc.ToCollectionsOfCollection(0)
	if coc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LinkedCollections_ToCollectionsOfCollection_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	coc := lc.ToCollectionsOfCollection(0)
	if coc == nil {
		t.Fatal("expected non-nil empty")
	}
}

// ── ItemsOfItems / ItemsOfItemsCollection ─────────────────

func Test_LinkedCollections_ItemsOfItems(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a", "b"}),
		New.Collection.Strings([]string{"c"}),
	)
	items := lc.ItemsOfItems()
	if len(items) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_ItemsOfItems_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	items := lc.ItemsOfItems()
	if len(items) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_LinkedCollections_ItemsOfItemsCollection(t *testing.T) {
	lc := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
	)
	items := lc.ItemsOfItemsCollection()
	if len(items) != 1 {
		t.Fatal("expected 1")
	}
}

// ── SimpleSlice ───────────────────────────────────────────

func Test_LinkedCollections_SimpleSlice(t *testing.T) {
	lc := New.LinkedCollection.Strings("a", "b")
	ss := lc.SimpleSlice()
	if ss == nil {
		t.Fatal("expected non-nil")
	}
}

// ── AddStringsOfStrings ───────────────────────────────────

func Test_LinkedCollections_AddStringsOfStrings(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddStringsOfStrings(false, []string{"a"}, []string{"b"})
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_AddStringsOfStrings_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.AddStringsOfStrings(false)
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AddAsyncFuncItems ─────────────────────────────────────

func Test_LinkedCollections_AddAsyncFuncItems(t *testing.T) {
	lc := New.LinkedCollection.Create()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	lc.AddAsyncFuncItems(wg, false, func() []string {
		return []string{"a", "b"}
	})
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AddAsyncFuncItems_EmptyReturn(t *testing.T) {
	lc := New.LinkedCollection.Create()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	lc.AddAsyncFuncItems(wg, false, func() []string {
		return []string{}
	})
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedCollections_AddAsyncFuncItems_Nil(t *testing.T) {
	lc := New.LinkedCollection.Create()
	wg := &sync.WaitGroup{}
	lc.AddAsyncFuncItems(wg, false)
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AddAsyncFuncItemsPointer ──────────────────────────────

func Test_LinkedCollections_AddAsyncFuncItemsPointer(t *testing.T) {
	lc := New.LinkedCollection.Create()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	lc.AddAsyncFuncItemsPointer(wg, false, func() []string {
		return []string{"x"}
	})
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollections_AddAsyncFuncItemsPointer_Nil(t *testing.T) {
	lc := New.LinkedCollection.Create()
	wg := &sync.WaitGroup{}
	lc.AddAsyncFuncItemsPointer(wg, false)
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── String / StringLock / Join / Joins ─────────────────────

func Test_LinkedCollections_String(t *testing.T) {
	lc := New.LinkedCollection.Strings("a", "b")
	s := lc.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_LinkedCollections_String_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	s := lc.String()
	if s == "" {
		t.Fatal("expected NoElements string")
	}
}

func Test_LinkedCollections_StringLock(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	s := lc.StringLock()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_LinkedCollections_StringLock_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	s := lc.StringLock()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_LinkedCollections_Join(t *testing.T) {
	lc := New.LinkedCollection.Strings("a", "b")
	j := lc.Join(",")
	if j != "a,b" {
		t.Fatalf("expected a,b, got %s", j)
	}
}

func Test_LinkedCollections_Joins(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	j := lc.Joins(",", "b")
	if j != "a,b" {
		t.Fatalf("expected a,b, got %s", j)
	}
}

func Test_LinkedCollections_Joins_NilItems(t *testing.T) {
	lc := New.LinkedCollection.Create()
	j := lc.Joins(",", "a")
	if j != "a" {
		t.Fatal("expected a")
	}
}

// ── List / ListPtr ────────────────────────────────────────

func Test_LinkedCollections_List(t *testing.T) {
	lc := New.LinkedCollection.Strings("a", "b")
	list := lc.List()
	if len(list) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_List_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	list := lc.List()
	if len(list) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_LinkedCollections_ListPtr(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	ptr := lc.ListPtr()
	if ptr == nil || len(*ptr) != 1 {
		t.Fatal("expected 1")
	}
}

// ── IsEqualsPtr ───────────────────────────────────────────

func Test_LinkedCollections_IsEqualsPtr(t *testing.T) {
	lc1 := New.LinkedCollection.Strings("a", "b")
	lc2 := New.LinkedCollection.Strings("a", "b")
	if !lc1.IsEqualsPtr(lc2) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedCollections_IsEqualsPtr_Nil(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	if lc.IsEqualsPtr(nil) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedCollections_IsEqualsPtr_SameRef(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	if !lc.IsEqualsPtr(lc) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedCollections_IsEqualsPtr_BothEmpty(t *testing.T) {
	lc1 := New.LinkedCollection.Create()
	lc2 := New.LinkedCollection.Create()
	if !lc1.IsEqualsPtr(lc2) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedCollections_IsEqualsPtr_OneEmpty(t *testing.T) {
	lc1 := New.LinkedCollection.Strings("a")
	lc2 := New.LinkedCollection.Create()
	if lc1.IsEqualsPtr(lc2) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedCollections_IsEqualsPtr_DiffLength(t *testing.T) {
	lc1 := New.LinkedCollection.UsingCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Strings([]string{"b"}),
	)
	lc2 := New.LinkedCollection.Strings("a")
	if lc1.IsEqualsPtr(lc2) {
		t.Fatal("expected not equal")
	}
}

// ── GetCompareSummary ─────────────────────────────────────

func Test_LinkedCollections_GetCompareSummary(t *testing.T) {
	lc1 := New.LinkedCollection.Strings("a")
	lc2 := New.LinkedCollection.Strings("b")
	s := lc1.GetCompareSummary(lc2, "left", "right")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

// ── JSON ──────────────────────────────────────────────────

func Test_LinkedCollections_JsonModel(t *testing.T) {
	lc := New.LinkedCollection.Strings("a", "b")
	model := lc.JsonModel()
	if len(model) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollections_JsonModelAny(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	if lc.JsonModelAny() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LinkedCollections_MarshalUnmarshalJSON(t *testing.T) {
	lc := New.LinkedCollection.Strings("a", "b")
	data, err := json.Marshal(lc)
	if err != nil {
		t.Fatal(err)
	}
	lc2 := New.LinkedCollection.Create()
	err = json.Unmarshal(data, lc2)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_LinkedCollections_Json(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	r := lc.Json()
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_LinkedCollections_JsonPtr(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	r := lc.JsonPtr()
	if r == nil || r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_LinkedCollections_ParseInjectUsingJson(t *testing.T) {
	lc := New.LinkedCollection.Strings("a", "b")
	jr := lc.JsonPtr()
	lc2 := New.LinkedCollection.Create()
	_, err := lc2.ParseInjectUsingJson(jr)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_LinkedCollections_ParseInjectUsingJsonMust(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	jr := lc.JsonPtr()
	lc2 := New.LinkedCollection.Create()
	result := lc2.ParseInjectUsingJsonMust(jr)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LinkedCollections_JsonParseSelfInject(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	jr := lc.JsonPtr()
	lc2 := New.LinkedCollection.Create()
	err := lc2.JsonParseSelfInject(jr)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_LinkedCollections_AsJsonContractsBinder(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	if lc.AsJsonContractsBinder() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LinkedCollections_AsJsoner(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	if lc.AsJsoner() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LinkedCollections_AsJsonParseSelfInjector(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	if lc.AsJsonParseSelfInjector() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LinkedCollections_AsJsonMarshaller(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	if lc.AsJsonMarshaller() == nil {
		t.Fatal("expected non-nil")
	}
}

// ── RemoveAll / Clear ─────────────────────────────────────

func Test_LinkedCollections_RemoveAll(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	lc.RemoveAll()
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedCollections_Clear(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	lc.Clear()
	if lc.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_LinkedCollections_Clear_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	lc.Clear()
}

// ── LinkedCollectionNode ──────────────────────────────────

func Test_LinkedCollectionNode_IsEmpty(t *testing.T) {
	var node *LinkedCollectionNode
	if !node.IsEmpty() {
		t.Fatal("expected empty for nil node")
	}
	node2 := &LinkedCollectionNode{Element: nil}
	if !node2.IsEmpty() {
		t.Fatal("expected empty for nil element")
	}
	node3 := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	if node3.IsEmpty() {
		t.Fatal("expected not empty")
	}
}

func Test_LinkedCollectionNode_HasElement(t *testing.T) {
	node := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	if !node.HasElement() {
		t.Fatal("expected has element")
	}
}

func Test_LinkedCollectionNode_HasNext(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	if !node.HasNext() {
		t.Fatal("expected has next")
	}
}

func Test_LinkedCollectionNode_EndOfChain(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next: &LinkedCollectionNode{
			Element: New.Collection.Strings([]string{"b"}),
		},
	}
	end, length := node.EndOfChain()
	if length != 2 || end.Element.List()[0] != "b" {
		t.Fatal("expected b, 2")
	}
}

func Test_LinkedCollectionNode_Clone(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	cloned := node.Clone()
	if cloned.HasNext() {
		t.Fatal("expected no next")
	}
}

func Test_LinkedCollectionNode_List(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	list := node.List()
	if len(list) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollectionNode_ListPtr(t *testing.T) {
	node := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	ptr := node.ListPtr()
	if ptr == nil || len(*ptr) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollectionNode_Join(t *testing.T) {
	node := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a", "b"})}
	if node.Join(",") != "a,b" {
		t.Fatal("expected a,b")
	}
}

func Test_LinkedCollectionNode_String(t *testing.T) {
	node := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	s := node.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_LinkedCollectionNode_StringList(t *testing.T) {
	node := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	s := node.StringList("Header: ")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_LinkedCollectionNode_IsEqual(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	n1 := &LinkedCollectionNode{Element: c}
	n2 := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	if !n1.IsEqual(n2) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedCollectionNode_IsEqual_BothNil(t *testing.T) {
	var n1 *LinkedCollectionNode
	if !n1.IsEqual(nil) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedCollectionNode_IsEqual_OneNil(t *testing.T) {
	n1 := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	if n1.IsEqual(nil) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedCollectionNode_IsEqual_SameRef(t *testing.T) {
	n1 := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	if !n1.IsEqual(n1) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedCollectionNode_IsEqual_NilElements(t *testing.T) {
	n1 := &LinkedCollectionNode{Element: nil}
	n2 := &LinkedCollectionNode{Element: nil}
	if !n1.IsEqual(n2) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedCollectionNode_IsEqual_OneNilElement(t *testing.T) {
	n1 := &LinkedCollectionNode{Element: nil}
	n2 := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	if n1.IsEqual(n2) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedCollectionNode_IsEqual_SameElementRef(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	n1 := &LinkedCollectionNode{Element: c}
	n2 := &LinkedCollectionNode{Element: c}
	if !n1.IsEqual(n2) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedCollectionNode_IsChainEqual(t *testing.T) {
	n1 := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	n2 := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	if !n1.IsChainEqual(n2) {
		t.Fatal("expected chain equal")
	}
}

func Test_LinkedCollectionNode_IsChainEqual_BothNil(t *testing.T) {
	var n1 *LinkedCollectionNode
	if !n1.IsChainEqual(nil) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedCollectionNode_IsChainEqual_OneNil(t *testing.T) {
	n1 := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	if n1.IsChainEqual(nil) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedCollectionNode_IsChainEqual_SameRef(t *testing.T) {
	n1 := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	if !n1.IsChainEqual(n1) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedCollectionNode_IsEqualValue(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	n := &LinkedCollectionNode{Element: c}
	if !n.IsEqualValue(c) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedCollectionNode_IsEqualValue_BothNil(t *testing.T) {
	n := &LinkedCollectionNode{Element: nil}
	if !n.IsEqualValue(nil) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedCollectionNode_IsEqualValue_OneNil(t *testing.T) {
	n := &LinkedCollectionNode{Element: nil}
	if n.IsEqualValue(New.Collection.Strings([]string{"a"})) {
		t.Fatal("expected not equal")
	}
}

func Test_LinkedCollectionNode_CreateLinkedList(t *testing.T) {
	n := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	lc := n.CreateLinkedList()
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollectionNode_AddNext(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	node := lc.Head()
	newNode := node.AddNext(lc, New.Collection.Strings([]string{"b"}))
	if newNode == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LinkedCollectionNode_AddNextNode(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	node := lc.Head()
	newNode := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})}
	node.AddNextNode(lc, newNode)
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollectionNode_AddStringsToNode(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	lc.Head().AddStringsToNode(lc, true, []string{"b"}, false)
	if lc.Length() < 2 {
		t.Fatal("expected items added")
	}
}

func Test_LinkedCollectionNode_AddCollectionToNode(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	col := New.Collection.Strings([]string{"b"})
	lc.Head().AddCollectionToNode(lc, true, col)
	if lc.Length() < 2 {
		t.Fatal("expected items added")
	}
}

func Test_LinkedCollectionNode_LoopEndOfChain(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next: &LinkedCollectionNode{
			Element: New.Collection.Strings([]string{"b"}),
		},
	}
	count := 0
	end, length := node.LoopEndOfChain(func(arg *LinkedCollectionProcessorParameter) bool {
		count++
		return false
	})
	if length != 2 || count != 2 || end == nil {
		t.Fatal("expected 2, 2")
	}
}

func Test_LinkedCollectionNode_LoopEndOfChain_Break(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	end, length := node.LoopEndOfChain(func(arg *LinkedCollectionProcessorParameter) bool {
		return true
	})
	if length != 1 || end == nil {
		t.Fatal("expected 1")
	}
}

// ── NonChainedLinkedCollectionNodes ────────────────────────

func Test_NonChainedLinkedCollectionNodes_Basic(t *testing.T) {
	nc := &NonChainedLinkedCollectionNodes{
		items: []*LinkedCollectionNode{
			{Element: New.Collection.Strings([]string{"a"})},
			{Element: New.Collection.Strings([]string{"b"})},
		},
	}
	if nc.IsEmpty() {
		t.Fatal("expected not empty")
	}
	if nc.Length() != 2 {
		t.Fatal("expected 2")
	}
	if !nc.HasItems() {
		t.Fatal("expected has items")
	}
	if nc.First() == nil || nc.Last() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_NonChainedLinkedCollectionNodes_ApplyChaining(t *testing.T) {
	nc := &NonChainedLinkedCollectionNodes{
		items: []*LinkedCollectionNode{
			{Element: New.Collection.Strings([]string{"a"})},
			{Element: New.Collection.Strings([]string{"b"})},
		},
	}
	nc.ApplyChaining()
	if !nc.IsChainingApplied() {
		t.Fatal("expected chaining applied")
	}
}

func Test_NonChainedLinkedCollectionNodes_FirstOrDefault_Empty(t *testing.T) {
	nc := &NonChainedLinkedCollectionNodes{}
	if nc.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
}

func Test_NonChainedLinkedCollectionNodes_LastOrDefault_Empty(t *testing.T) {
	nc := &NonChainedLinkedCollectionNodes{}
	if nc.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
}
