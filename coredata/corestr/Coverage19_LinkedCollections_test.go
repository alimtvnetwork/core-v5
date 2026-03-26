package corestr

import (
	"testing"
)

// Tests using unexported fields (next, items) — must remain in source package.

func Test_LinkedCollections_AppendChainOfNodes_C19(t *testing.T) {
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

func Test_LinkedCollections_AppendChainOfNodes_Empty_C19(t *testing.T) {
	lc := New.LinkedCollection.Create()
	chain := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	lc.AppendChainOfNodes(chain)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_LinkedCollectionNode_HasNext_C19(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	if !node.HasNext() {
		t.Fatal("expected has next")
	}
}

func Test_LinkedCollectionNode_EndOfChain_C19(t *testing.T) {
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

func Test_LinkedCollectionNode_Clone_C19(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	cloned := node.Clone()
	if cloned.HasNext() {
		t.Fatal("expected no next")
	}
}

func Test_LinkedCollectionNode_List_C19(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	list := node.List()
	if len(list) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollectionNode_IsChainEqual_C19(t *testing.T) {
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

func Test_LinkedCollectionNode_CreateLinkedList_C19(t *testing.T) {
	n := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	lc := n.CreateLinkedList()
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedCollectionNode_LoopEndOfChain_C19(t *testing.T) {
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

func Test_LinkedCollectionNode_LoopEndOfChain_Break_C19(t *testing.T) {
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

func Test_NonChainedLinkedCollectionNodes_Basic_C19(t *testing.T) {
	nc := &NonChainedLinkedCollectionNodes{
		items: []*LinkedCollectionNode{
			{Element: New.Collection.Strings([]string{"a"})},
			{Element: New.Collection.Strings([]string{"b"})},
		},
	}
	if nc.IsEmpty() || nc.Length() != 2 || !nc.HasItems() {
		t.Fatal("expected 2 items")
	}
}

func Test_NonChainedLinkedCollectionNodes_ApplyChaining_C19(t *testing.T) {
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
