package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

var linkedListTestCases = []coretestcases.CaseV1{
	// Constructors
	{Title: "EmptyLinkedList creates empty list", ArrangeInput: args.Map{"case": "empty"}, ExpectedInput: []string{"true", "0", "false"}},
	{Title: "LinkedListFrom creates from slice", ArrangeInput: args.Map{"case": "from"}, ExpectedInput: []string{"3", "a", "c"}},
	{Title: "LinkedListFrom empty slice", ArrangeInput: args.Map{"case": "from-empty"}, ExpectedInput: []string{"true"}},
	// Add
	{Title: "Add single sets head and tail", ArrangeInput: args.Map{"case": "add-single"}, ExpectedInput: []string{"1", "42", "42"}},
	{Title: "Add multiple appends to back", ArrangeInput: args.Map{"case": "add-multiple"}, ExpectedInput: []string{"1", "3", "3"}},
	{Title: "AddFront prepends", ArrangeInput: args.Map{"case": "addfront"}, ExpectedInput: []string{"1", "3", "3"}},
	{Title: "AddFront empty", ArrangeInput: args.Map{"case": "addfront-empty"}, ExpectedInput: []string{"first", "first", "1"}},
	{Title: "Adds multiple", ArrangeInput: args.Map{"case": "adds"}, ExpectedInput: []string{"3"}},
	{Title: "AddSlice appends", ArrangeInput: args.Map{"case": "addslice"}, ExpectedInput: []string{"2"}},
	{Title: "AddIf true adds", ArrangeInput: args.Map{"case": "addif-true"}, ExpectedInput: []string{"1"}},
	{Title: "AddIf false skips", ArrangeInput: args.Map{"case": "addif-false"}, ExpectedInput: []string{"true"}},
	{Title: "AddsIf false skips", ArrangeInput: args.Map{"case": "addsif-false"}, ExpectedInput: []string{"true"}},
	{Title: "AddFunc adds result", ArrangeInput: args.Map{"case": "addfunc"}, ExpectedInput: []string{"99"}},
	{Title: "Push aliases work", ArrangeInput: args.Map{"case": "push"}, ExpectedInput: []string{"3"}},
	// FirstOrDefault / LastOrDefault
	{Title: "FirstOrDefault empty returns zero", ArrangeInput: args.Map{"case": "firstdefault-empty"}, ExpectedInput: []string{"0"}},
	{Title: "LastOrDefault empty returns zero", ArrangeInput: args.Map{"case": "lastdefault-empty"}, ExpectedInput: []string{""}},
	{Title: "FirstOrDefault non-empty", ArrangeInput: args.Map{"case": "firstdefault"}, ExpectedInput: []string{"10"}},
	{Title: "LastOrDefault non-empty", ArrangeInput: args.Map{"case": "lastdefault"}, ExpectedInput: []string{"20"}},
	// Items / Collection / String
	{Title: "Items returns all elements", ArrangeInput: args.Map{"case": "items"}, ExpectedInput: []string{"3"}},
	{Title: "Items empty returns empty", ArrangeInput: args.Map{"case": "items-empty"}, ExpectedInput: []string{"0"}},
	{Title: "Collection converts", ArrangeInput: args.Map{"case": "collection"}, ExpectedInput: []string{"2"}},
	{Title: "String representation", ArrangeInput: args.Map{"case": "string"}, ExpectedInput: []string{"[1 2 3]"}},
	// IndexAt
	{Title: "IndexAt valid returns node", ArrangeInput: args.Map{"case": "indexat-valid"}, ExpectedInput: []string{"true", "b"}},
	{Title: "IndexAt first", ArrangeInput: args.Map{"case": "indexat-first"}, ExpectedInput: []string{"10"}},
	{Title: "IndexAt last", ArrangeInput: args.Map{"case": "indexat-last"}, ExpectedInput: []string{"30"}},
	{Title: "IndexAt out of bounds", ArrangeInput: args.Map{"case": "indexat-oob"}, ExpectedInput: []string{"true", "true"}},
	{Title: "IndexAt empty", ArrangeInput: args.Map{"case": "indexat-empty"}, ExpectedInput: []string{"true"}},
	// ForEach
	{Title: "ForEach visits all", ArrangeInput: args.Map{"case": "foreach"}, ExpectedInput: []string{"6"}},
	{Title: "ForEach empty noop", ArrangeInput: args.Map{"case": "foreach-empty"}, ExpectedInput: []string{"false"}},
	{Title: "ForEachBreak stops early", ArrangeInput: args.Map{"case": "foreachbreak"}, ExpectedInput: []string{"3"}},
	{Title: "ForEachBreak first element", ArrangeInput: args.Map{"case": "foreachbreak-first"}, ExpectedInput: []string{"1"}},
	// Head / Tail
	{Title: "Head/Tail nodes", ArrangeInput: args.Map{"case": "headtail"}, ExpectedInput: []string{"1", "3", "true", "false"}},
	{Title: "Node.Next traverses", ArrangeInput: args.Map{"case": "node-next"}, ExpectedInput: []string{"10", "20", "30", "false"}},
	// Lock variants
	{Title: "LengthLock", ArrangeInput: args.Map{"case": "lengthlock"}, ExpectedInput: []string{"2"}},
	{Title: "IsEmptyLock", ArrangeInput: args.Map{"case": "isemptylock"}, ExpectedInput: []string{"true"}},
	{Title: "AddLock", ArrangeInput: args.Map{"case": "addlock"}, ExpectedInput: []string{"2"}},
	// Nil receiver
	{Title: "IsEmpty nil receiver", ArrangeInput: args.Map{"case": "isempty-nil"}, ExpectedInput: []string{"true"}},
	// AppendNode
	{Title: "AppendNode appends", ArrangeInput: args.Map{"case": "appendnode"}, ExpectedInput: []string{"3", "3"}},
	{Title: "AppendNode empty", ArrangeInput: args.Map{"case": "appendnode-empty"}, ExpectedInput: []string{"1", "99"}},
}

func Test_LinkedList_Verification(t *testing.T) {
	for caseIndex, tc := range linkedListTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		caseType := input["case"].(string)

		var actLines []string

		// Act
		switch caseType {
		case "empty":
			ll := coregeneric.EmptyLinkedList[int]()
			actLines = []string{fmt.Sprintf("%v", ll.IsEmpty()), fmt.Sprintf("%v", ll.Length()), fmt.Sprintf("%v", ll.HasItems())}
		case "from":
			ll := coregeneric.LinkedListFrom([]string{"a", "b", "c"})
			actLines = []string{fmt.Sprintf("%v", ll.Length()), ll.First(), ll.Last()}
		case "from-empty":
			ll := coregeneric.LinkedListFrom([]int{})
			actLines = []string{fmt.Sprintf("%v", ll.IsEmpty())}
		case "add-single":
			ll := coregeneric.EmptyLinkedList[int]()
			ll.Add(42)
			actLines = []string{fmt.Sprintf("%v", ll.Length()), fmt.Sprintf("%v", ll.First()), fmt.Sprintf("%v", ll.Last())}
		case "add-multiple":
			ll := coregeneric.EmptyLinkedList[int]()
			ll.Add(1).Add(2).Add(3)
			actLines = []string{fmt.Sprintf("%v", ll.First()), fmt.Sprintf("%v", ll.Last()), fmt.Sprintf("%v", ll.Length())}
		case "addfront":
			ll := coregeneric.LinkedListFrom([]int{2, 3})
			ll.AddFront(1)
			actLines = []string{fmt.Sprintf("%v", ll.First()), fmt.Sprintf("%v", ll.Last()), fmt.Sprintf("%v", ll.Length())}
		case "addfront-empty":
			ll := coregeneric.EmptyLinkedList[string]()
			ll.AddFront("first")
			actLines = []string{ll.First(), ll.Last(), fmt.Sprintf("%v", ll.Length())}
		case "adds":
			ll := coregeneric.EmptyLinkedList[int]()
			ll.Adds(1, 2, 3)
			actLines = []string{fmt.Sprintf("%v", ll.Length())}
		case "addslice":
			ll := coregeneric.EmptyLinkedList[int]()
			ll.AddSlice([]int{10, 20})
			actLines = []string{fmt.Sprintf("%v", ll.Length())}
		case "addif-true":
			ll := coregeneric.EmptyLinkedList[int]()
			ll.AddIf(true, 5)
			actLines = []string{fmt.Sprintf("%v", ll.Length())}
		case "addif-false":
			ll := coregeneric.EmptyLinkedList[int]()
			ll.AddIf(false, 5)
			actLines = []string{fmt.Sprintf("%v", ll.IsEmpty())}
		case "addsif-false":
			ll := coregeneric.EmptyLinkedList[int]()
			ll.AddsIf(false, 1, 2, 3)
			actLines = []string{fmt.Sprintf("%v", ll.IsEmpty())}
		case "addfunc":
			ll := coregeneric.EmptyLinkedList[int]()
			ll.AddFunc(func() int { return 99 })
			actLines = []string{fmt.Sprintf("%v", ll.First())}
		case "push":
			ll := coregeneric.EmptyLinkedList[int]()
			ll.Push(1)
			ll.PushBack(2)
			ll.PushFront(0)
			actLines = []string{fmt.Sprintf("%v", ll.Length())}
		case "firstdefault-empty":
			ll := coregeneric.EmptyLinkedList[int]()
			actLines = []string{fmt.Sprintf("%v", ll.FirstOrDefault())}
		case "lastdefault-empty":
			ll := coregeneric.EmptyLinkedList[string]()
			actLines = []string{ll.LastOrDefault()}
		case "firstdefault":
			ll := coregeneric.LinkedListFrom([]int{10, 20})
			actLines = []string{fmt.Sprintf("%v", ll.FirstOrDefault())}
		case "lastdefault":
			ll := coregeneric.LinkedListFrom([]int{10, 20})
			actLines = []string{fmt.Sprintf("%v", ll.LastOrDefault())}
		case "items":
			ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
			actLines = []string{fmt.Sprintf("%v", len(ll.Items()))}
		case "items-empty":
			ll := coregeneric.EmptyLinkedList[int]()
			actLines = []string{fmt.Sprintf("%v", len(ll.Items()))}
		case "collection":
			ll := coregeneric.LinkedListFrom([]int{1, 2})
			actLines = []string{fmt.Sprintf("%v", ll.Collection().Length())}
		case "string":
			ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
			actLines = []string{ll.String()}
		case "indexat-valid":
			ll := coregeneric.LinkedListFrom([]string{"a", "b", "c"})
			node := ll.IndexAt(1)
			actLines = []string{fmt.Sprintf("%v", node != nil), node.Element}
		case "indexat-first":
			ll := coregeneric.LinkedListFrom([]int{10, 20})
			actLines = []string{fmt.Sprintf("%v", ll.IndexAt(0).Element)}
		case "indexat-last":
			ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
			actLines = []string{fmt.Sprintf("%v", ll.IndexAt(2).Element)}
		case "indexat-oob":
			ll := coregeneric.LinkedListFrom([]int{1, 2})
			actLines = []string{fmt.Sprintf("%v", ll.IndexAt(5) == nil), fmt.Sprintf("%v", ll.IndexAt(-1) == nil)}
		case "indexat-empty":
			ll := coregeneric.EmptyLinkedList[int]()
			actLines = []string{fmt.Sprintf("%v", ll.IndexAt(0) == nil)}
		case "foreach":
			ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
			sum := 0
			ll.ForEach(func(_ int, item int) { sum += item })
			actLines = []string{fmt.Sprintf("%v", sum)}
		case "foreach-empty":
			ll := coregeneric.EmptyLinkedList[int]()
			called := false
			ll.ForEach(func(_ int, _ int) { called = true })
			actLines = []string{fmt.Sprintf("%v", called)}
		case "foreachbreak":
			ll := coregeneric.LinkedListFrom([]int{1, 2, 3, 4, 5})
			count := 0
			ll.ForEachBreak(func(_ int, item int) bool { count++; return item == 3 })
			actLines = []string{fmt.Sprintf("%v", count)}
		case "foreachbreak-first":
			ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
			count := 0
			ll.ForEachBreak(func(_ int, _ int) bool { count++; return true })
			actLines = []string{fmt.Sprintf("%v", count)}
		case "headtail":
			ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
			actLines = []string{fmt.Sprintf("%v", ll.Head().Element), fmt.Sprintf("%v", ll.Tail().Element), fmt.Sprintf("%v", ll.Head().HasNext()), fmt.Sprintf("%v", ll.Tail().HasNext())}
		case "node-next":
			ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
			n := ll.Head()
			actLines = []string{fmt.Sprintf("%v", n.Element)}
			n = n.Next()
			actLines = append(actLines, fmt.Sprintf("%v", n.Element))
			n = n.Next()
			actLines = append(actLines, fmt.Sprintf("%v", n.Element), fmt.Sprintf("%v", n.HasNext()))
		case "lengthlock":
			ll := coregeneric.LinkedListFrom([]int{1, 2})
			actLines = []string{fmt.Sprintf("%v", ll.LengthLock())}
		case "isemptylock":
			ll := coregeneric.EmptyLinkedList[int]()
			actLines = []string{fmt.Sprintf("%v", ll.IsEmptyLock())}
		case "addlock":
			ll := coregeneric.EmptyLinkedList[int]()
			ll.AddLock(1)
			ll.AddLock(2)
			actLines = []string{fmt.Sprintf("%v", ll.Length())}
		case "isempty-nil":
			var ll *coregeneric.LinkedList[int]
			actLines = []string{fmt.Sprintf("%v", ll.IsEmpty())}
		case "appendnode":
			ll := coregeneric.LinkedListFrom([]int{1, 2})
			ll.AppendNode(&coregeneric.LinkedListNode[int]{Element: 3})
			actLines = []string{fmt.Sprintf("%v", ll.Length()), fmt.Sprintf("%v", ll.Last())}
		case "appendnode-empty":
			ll := coregeneric.EmptyLinkedList[int]()
			ll.AppendNode(&coregeneric.LinkedListNode[int]{Element: 99})
			actLines = []string{fmt.Sprintf("%v", ll.Length()), fmt.Sprintf("%v", ll.First())}
		}

		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}
