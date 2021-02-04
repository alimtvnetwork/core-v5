package corestr

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type LinkedListNode struct {
	Element string
	next    *LinkedListNode
}

func (linkedListNode *LinkedListNode) HasNext() bool {
	return linkedListNode.next != nil
}

func (linkedListNode *LinkedListNode) Next() *LinkedListNode {
	return linkedListNode.next
}

func (linkedListNode *LinkedListNode) AddNext(
	linkedListForIncrement *LinkedList,
	item string,
) *LinkedListNode {
	newNode := &LinkedListNode{
		Element: item,
		next:    linkedListNode.Next(),
	}

	linkedListNode.next = newNode
	linkedListForIncrement.incrementLength()

	return newNode
}

func (linkedListNode *LinkedListNode) AddStringsPtrToNode(
	linkedListForIncrement *LinkedList,
	isSkipOnNull bool,
	items *[]string,
) *LinkedList {
	return linkedListForIncrement.AddStringsPtrToNode(
		isSkipOnNull,
		linkedListNode,
		items)
}

func (linkedListNode *LinkedListNode) AddCollectionToNode(
	linkedListForIncrement *LinkedList,
	isSkipOnNull bool,
	collection *Collection,
) *LinkedList {
	return linkedListForIncrement.AddStringsPtrToNode(
		isSkipOnNull,
		linkedListNode,
		collection.items)
}

func (linkedListNode *LinkedListNode) AddNextNode(
	linkedListForIncrement *LinkedList,
	nextNode *LinkedListNode,
) *LinkedListNode {
	nextNode.next = linkedListNode.Next()
	linkedListNode.next = nextNode
	linkedListForIncrement.incrementLength()

	return nextNode
}

func (linkedListNode *LinkedListNode) IsEqual(another *LinkedListNode) bool {
	if linkedListNode == nil && nil == another {
		return true
	}

	if linkedListNode == nil || nil == another {
		return false
	}

	if linkedListNode == another {
		return true
	}

	if another == nil && linkedListNode != nil {
		return false
	}

	//goland:noinspection GoNilness
	if linkedListNode.Element == another.Element {
		if !linkedListNode.HasNext() && !another.HasNext() {
			return false
		}

		if !linkedListNode.HasNext() || !another.HasNext() {
			return false
		}

		if linkedListNode.HasNext() && another.HasNext() {
			return linkedListNode.Next().Element == another.Element
		}

		if linkedListNode.HasNext() || another.HasNext() {
			return false
		}
	}

	return false
}

func (linkedListNode *LinkedListNode) IsEqualSensitive(another *LinkedListNode, isCaseSensitive bool) bool {
	if linkedListNode == another {
		return true
	}

	if another == nil && linkedListNode != nil {
		return false
	}

	//goland:noinspection GoNilness
	if linkedListNode.IsEqualValueSensitive(another.Element, isCaseSensitive) {
		return true
	}

	return false
}

func (linkedListNode *LinkedListNode) IsEqualValue(value string) bool {
	return linkedListNode.Element == value
}

func (linkedListNode *LinkedListNode) IsEqualValueSensitive(value string, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return value == linkedListNode.Element
	}

	return strings.EqualFold(linkedListNode.Element, value)
}

func (linkedListNode *LinkedListNode) String() string {
	return linkedListNode.Element
}

func (linkedListNode *LinkedListNode) ListPtr() *[]string {
	list := make([]string, 0, constants.ArbitraryCapacity100)

	node := linkedListNode
	list = append(list, node.Element)

	for node.HasNext() {
		node = node.Next()

		list = append(list, node.Element)
	}

	return &list
}

func (linkedListNode *LinkedListNode) Join(separator string) *string {
	list := linkedListNode.ListPtr()
	toString := strings.Join(*list, separator)

	return &toString
}

func (linkedListNode *LinkedListNode) StringListPtr(header string) *string {
	finalString := header +
		*linkedListNode.Join(commonJoiner)

	return &finalString
}

func (linkedListNode *LinkedListNode) Print(header string) {
	finalString := linkedListNode.StringListPtr(header)
	fmt.Println(finalString)
}
