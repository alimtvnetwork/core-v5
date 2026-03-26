package corestr

import (
	"testing"
)

// Tests that access unexported 'next' field — must remain in source package.

func TestLinkedList_AppendChainOfNodes_C08(t *testing.T) {
	ll := New.LinkedList.Create()
	n1 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	ll.AppendChainOfNodes(n1)
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestLinkedListNode_EndOfChain_C08(t *testing.T) {
	n1 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	end, length := n1.EndOfChain()
	if end.Element != "b" || length != 2 {
		t.Fatal("unexpected")
	}
}
