package corestr

import (
	"testing"
)

// Tests using unexported field 'next' — must remain in source package.

func Test_LinkedList_AppendChainOfNodes_C18(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	chain := &LinkedListNode{Element: "b", next: &LinkedListNode{Element: "c"}}
	ll.AppendChainOfNodes(chain)
	if ll.Length() != 3 {
		t.Fatalf("expected 3, got %d", ll.Length())
	}
}

func Test_LinkedList_AppendChainOfNodes_Empty_C18(t *testing.T) {
	ll := New.LinkedList.Create()
	chain := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	ll.AppendChainOfNodes(chain)
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedListNode_HasNext_C18(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	if !node.HasNext() {
		t.Fatal("expected has next")
	}
	if node.Next().HasNext() {
		t.Fatal("expected no next")
	}
}

func Test_LinkedListNode_EndOfChain_C18(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b", next: &LinkedListNode{Element: "c"}}}
	end, length := node.EndOfChain()
	if end.Element != "c" || length != 3 {
		t.Fatal("expected c, 3")
	}
}

func Test_LinkedListNode_Clone_C18(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	cloned := node.Clone()
	if cloned.Element != "a" || cloned.HasNext() {
		t.Fatal("expected a without next")
	}
}

func Test_LinkedListNode_List_C18(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	list := node.List()
	if len(list) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedListNode_Join_C18(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	if node.Join(",") != "a,b" {
		t.Fatal("expected a,b")
	}
}

func Test_LinkedListNode_IsEqual_C18(t *testing.T) {
	n1 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	n2 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	if !n1.IsEqual(n2) {
		t.Fatal("expected equal")
	}
}

func Test_LinkedListNode_IsChainEqual_C18(t *testing.T) {
	n1 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	n2 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	if !n1.IsChainEqual(n2, true) {
		t.Fatal("expected chain equal")
	}
}

func Test_LinkedListNode_CreateLinkedList_C18(t *testing.T) {
	n := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	ll := n.CreateLinkedList()
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LinkedListNode_LoopEndOfChain_C18(t *testing.T) {
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

func Test_LinkedListNode_LoopEndOfChain_Break_C18(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	end, length := node.LoopEndOfChain(func(arg *LinkedListProcessorParameter) bool {
		return true
	})
	if end.Element != "a" || length != 1 {
		t.Fatal("expected a, 1")
	}
}

func Test_LinkedListNode_LoopEndOfChain_BreakSecond_C18(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b", next: &LinkedListNode{Element: "c"}}}
	end, length := node.LoopEndOfChain(func(arg *LinkedListProcessorParameter) bool {
		return arg.Index == 1
	})
	if end.Element != "b" || length != 2 {
		t.Fatal("expected b, 2")
	}
}
