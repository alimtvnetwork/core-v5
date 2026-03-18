package corestr

import (
	"testing"
)

func TestLinkedList_Basic(t *testing.T) {
	ll := New.LinkedList.Create()
	if !ll.IsEmpty() || ll.HasItems() { t.Fatal("expected empty") }
	if ll.Length() != 0 { t.Fatal("expected 0") }
	if ll.IsEmptyLock() != true { t.Fatal("expected empty") }
}

func TestLinkedList_Add(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.Add("a").Add("b").Add("c")
	if ll.Length() != 3 { t.Fatal("expected 3") }
	if ll.Head().Element != "a" || ll.Tail().Element != "c" { t.Fatal("unexpected") }
	ll.AddLock("d")
	ll.AddFront("z")
	if ll.Head().Element != "z" { t.Fatal("expected z") }
	ll.AddNonEmpty("")
	ll.AddNonEmpty("e")
	ll.AddNonEmptyWhitespace("   ")
	ll.AddNonEmptyWhitespace("f")
	ll.AddIf(false, "skip")
	ll.AddIf(true, "g")
	ll.AddsIf(false, "x")
	ll.AddsIf(true, "h", "i")
	ll.AddFunc(func() string { return "j" })
	ll.AddFuncErr(func() (string, error) { return "k", nil }, func(e error) {})
	ll.Push("l")
	ll.PushFront("m")
	ll.PushBack("n")
}

func TestLinkedList_AddStrings(t *testing.T) {
	ll := New.LinkedList.Create()
	ll.Adds("a", "b")
	ll.Adds()
	ll.AddStrings([]string{"c"})
	ll.AddStrings(nil)
	ll.AddsLock("d")
	ll.AddCollection(New.Collection.Strings([]string{"e"}))
	ll.AddCollection(nil)
	ll.AddItemsMap(map[string]bool{"f": true, "g": false})
	ll.AddItemsMap(nil)
}

func TestLinkedList_List(t *testing.T) {
	ll := New.LinkedList.Strings([]string{"a", "b", "c"})
	list := ll.List()
	if len(list) != 3 { t.Fatal("expected 3") }
	_ = ll.ListPtr()
	_ = ll.ListLock()
	_ = ll.ListPtrLock()
	_ = ll.String()
	_ = ll.StringLock()
	_ = ll.Join(",")
	_ = ll.JoinLock(",")
}

func TestLinkedList_ToCollection(t *testing.T) {
	ll := New.LinkedList.Strings([]string{"a", "b"})
	c := ll.ToCollection(0)
	if c.Length() != 2 { t.Fatal("expected 2") }
	empty := New.LinkedList.Create()
	c2 := empty.ToCollection(0)
	if c2.Length() != 0 { t.Fatal("expected 0") }
}

func TestLinkedList_Loop(t *testing.T) {
	ll := New.LinkedList.Strings([]string{"a", "b", "c"})
	count := 0
	ll.Loop(func(arg *LinkedListProcessorParameter) bool {
		count++
		return false
	})
	if count != 3 { t.Fatal("expected 3") }
}

func TestLinkedList_Filter(t *testing.T) {
	ll := New.LinkedList.Strings([]string{"a", "b", "c"})
	nodes := ll.Filter(func(arg *LinkedListFilterParameter) *LinkedListFilterResult {
		return &LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
	})
	if len(nodes) != 3 { t.Fatal("expected 3") }
}

func TestLinkedList_IndexAt(t *testing.T) {
	ll := New.LinkedList.Strings([]string{"a", "b", "c"})
	node := ll.SafeIndexAt(1)
	if node == nil || node.Element != "b" { t.Fatal("unexpected") }
	if ll.SafeIndexAt(-1) != nil { t.Fatal("expected nil") }
	if ll.SafeIndexAt(99) != nil { t.Fatal("expected nil") }
	_ = ll.SafePointerIndexAt(1)
	_ = ll.SafePointerIndexAt(-1)
	_ = ll.SafePointerIndexAtUsingDefault(1, "def")
	_ = ll.SafePointerIndexAtUsingDefault(-1, "def")
	_ = ll.SafePointerIndexAtUsingDefaultLock(0, "def")
	_ = ll.SafeIndexAtLock(0)
}

func TestLinkedList_GetNextNodes(t *testing.T) {
	ll := New.LinkedList.Strings([]string{"a", "b", "c"})
	nodes := ll.GetNextNodes(2)
	if len(nodes) != 2 { t.Fatal("expected 2") }
	all := ll.GetAllLinkedNodes()
	if len(all) != 3 { t.Fatal("expected 3") }
}

func TestLinkedList_IsEquals(t *testing.T) {
	ll1 := New.LinkedList.Strings([]string{"a", "b"})
	ll2 := New.LinkedList.Strings([]string{"a", "b"})
	if !ll1.IsEquals(ll2) { t.Fatal("expected equal") }
	if !ll1.IsEqualsWithSensitive(ll2, false) { t.Fatal("expected equal") }
}

func TestLinkedList_RemoveNodeByIndex(t *testing.T) {
	ll := New.LinkedList.Strings([]string{"a", "b", "c"})
	ll.RemoveNodeByIndex(0)
	if ll.Length() != 2 { t.Fatal("expected 2") }
}

func TestLinkedList_Clear(t *testing.T) {
	ll := New.LinkedList.Strings([]string{"a", "b"})
	ll.Clear()
	if ll.Length() != 0 { t.Fatal("expected 0") }
	ll.RemoveAll()
}

func TestLinkedList_JsonAndMarshal(t *testing.T) {
	ll := New.LinkedList.Strings([]string{"a"})
	_ = ll.JsonModel()
	_ = ll.JsonModelAny()
	_, _ = ll.MarshalJSON()
	_ = ll.AsJsonMarshaller()
}

func TestLinkedList_Joins(t *testing.T) {
	ll := New.LinkedList.Strings([]string{"a", "b"})
	s := ll.Joins(",", "c")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestLinkedList_AppendNode(t *testing.T) {
	ll := New.LinkedList.Create()
	node := &LinkedListNode{Element: "a"}
	ll.AppendNode(node)
	if ll.Length() != 1 { t.Fatal("expected 1") }
	ll.AppendNode(&LinkedListNode{Element: "b"})
}

func TestLinkedList_AppendChainOfNodes(t *testing.T) {
	ll := New.LinkedList.Create()
	n1 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	ll.AppendChainOfNodes(n1)
	if ll.Length() != 2 { t.Fatal("expected 2") }
}

func TestLinkedListNode_Methods(t *testing.T) {
	n := &LinkedListNode{Element: "a"}
	if n.HasNext() { t.Fatal("expected no next") }
	if n.String() != "a" { t.Fatal("expected a") }
	if !n.IsEqualValue("a") { t.Fatal("expected true") }
	if !n.IsEqualValueSensitive("a", true) { t.Fatal("expected true") }
	if !n.IsEqualValueSensitive("A", false) { t.Fatal("expected true") }
	c := n.Clone()
	if c.Element != "a" { t.Fatal("expected a") }
	_ = n.List()
	_ = n.ListPtr()
	_ = n.Join(",")
	_ = n.StringList("header: ")
	_ = n.CreateLinkedList()
}

func TestLinkedListNode_EndOfChain(t *testing.T) {
	n1 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	end, length := n1.EndOfChain()
	if end.Element != "b" || length != 2 { t.Fatal("unexpected") }
}

func TestLinkedListNode_IsEqual(t *testing.T) {
	n1 := &LinkedListNode{Element: "a"}
	n2 := &LinkedListNode{Element: "a"}
	if !n1.IsEqual(n2) { t.Fatal("expected equal") }
	if !n1.IsChainEqual(n2, true) { t.Fatal("expected equal") }
	if !n1.IsEqualSensitive(n2, true) { t.Fatal("expected equal") }
}
