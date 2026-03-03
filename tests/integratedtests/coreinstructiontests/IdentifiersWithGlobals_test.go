package coreinstructiontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

var identifiersWithGlobalsTestCases = []coretestcases.CaseV1{
	// Length
	{Title: "Length - empty returns 0", ArrangeInput: args.Map{"case": "length-empty"}, ExpectedInput: []string{"0"}},
	{Title: "Length - 3 items returns 3", ArrangeInput: args.Map{"case": "length-3"}, ExpectedInput: []string{"3"}},
	{Title: "Length - nil receiver returns 0", ArrangeInput: args.Map{"case": "length-nil"}, ExpectedInput: []string{"0"}},
	// GetById
	{Title: "GetById - found returns item", ArrangeInput: args.Map{"case": "getbyid-found"}, ExpectedInput: []string{"true", "beta", "true"}},
	{Title: "GetById - missing returns nil", ArrangeInput: args.Map{"case": "getbyid-missing"}, ExpectedInput: []string{"true"}},
	{Title: "GetById - empty id returns nil", ArrangeInput: args.Map{"case": "getbyid-empty"}, ExpectedInput: []string{"true"}},
	// Clone
	{Title: "Clone - independence", ArrangeInput: args.Map{"case": "clone-independence"}, ExpectedInput: []string{"2", "3"}},
	{Title: "Clone - empty clones to empty", ArrangeInput: args.Map{"case": "clone-empty"}, ExpectedInput: []string{"true", "0"}},
	{Title: "Clone - preserves values", ArrangeInput: args.Map{"case": "clone-preserves"}, ExpectedInput: []string{"true", "id-1", "false"}},
	// Add
	{Title: "Add - single item", ArrangeInput: args.Map{"case": "add-single"}, ExpectedInput: []string{"1", "true", "true"}},
	{Title: "Add - empty id ignored", ArrangeInput: args.Map{"case": "add-empty"}, ExpectedInput: []string{"0"}},
	{Title: "Add - multiple accumulate", ArrangeInput: args.Map{"case": "add-multiple"}, ExpectedInput: []string{"3", "true", "false"}},
	// IsEmpty / HasAnyItem
	{Title: "IsEmpty - empty true", ArrangeInput: args.Map{"case": "isempty-true"}, ExpectedInput: []string{"true", "false"}},
	{Title: "IsEmpty - non-empty false", ArrangeInput: args.Map{"case": "isempty-false"}, ExpectedInput: []string{"false", "true"}},
	// IndexOf
	{Title: "IndexOf - found returns correct index", ArrangeInput: args.Map{"case": "indexof-found"}, ExpectedInput: []string{"0", "1", "2"}},
	{Title: "IndexOf - missing returns -1", ArrangeInput: args.Map{"case": "indexof-missing"}, ExpectedInput: []string{"-1"}},
	{Title: "IndexOf - empty string returns -1", ArrangeInput: args.Map{"case": "indexof-empty"}, ExpectedInput: []string{"-1"}},
	{Title: "IndexOf - empty collection returns -1", ArrangeInput: args.Map{"case": "indexof-emptycol"}, ExpectedInput: []string{"-1"}},
	// Adds
	{Title: "Adds - batch add all items", ArrangeInput: args.Map{"case": "adds-batch"}, ExpectedInput: []string{"3", "true", "true", "true"}},
	{Title: "Adds - empty ids no add", ArrangeInput: args.Map{"case": "adds-empty"}, ExpectedInput: []string{"0"}},
	// NewIdentifiersWithGlobals edge
	{Title: "New - no ids creates empty", ArrangeInput: args.Map{"case": "new-noids"}, ExpectedInput: []string{"true", "0"}},
}

func Test_IdentifiersWithGlobals_Verification(t *testing.T) {
	for caseIndex, tc := range identifiersWithGlobalsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		caseType := input["case"].(string)

		var actLines []string

		// Act
		switch caseType {
		case "length-empty":
			ids := coreinstruction.EmptyIdentifiersWithGlobals()
			actLines = []string{fmt.Sprintf("%v", ids.Length())}
		case "length-3":
			ids := coreinstruction.NewIdentifiersWithGlobals(true, "a", "b", "c")
			actLines = []string{fmt.Sprintf("%v", ids.Length())}
		case "length-nil":
			var nilIds *coreinstruction.IdentifiersWithGlobals
			actLines = []string{fmt.Sprintf("%v", nilIds.Length())}
		case "getbyid-found":
			ids := coreinstruction.NewIdentifiersWithGlobals(true, "alpha", "beta")
			found := ids.GetById("beta")
			actLines = []string{fmt.Sprintf("%v", found != nil), found.Id, fmt.Sprintf("%v", found.IsGlobal)}
		case "getbyid-missing":
			ids := coreinstruction.NewIdentifiersWithGlobals(false, "alpha")
			actLines = []string{fmt.Sprintf("%v", ids.GetById("missing") == nil)}
		case "getbyid-empty":
			ids := coreinstruction.NewIdentifiersWithGlobals(true, "alpha")
			actLines = []string{fmt.Sprintf("%v", ids.GetById("") == nil)}
		case "clone-independence":
			orig := coreinstruction.NewIdentifiersWithGlobals(true, "x", "y")
			cloned := orig.Clone()
			cloned.Add(false, "z")
			actLines = []string{fmt.Sprintf("%v", orig.Length()), fmt.Sprintf("%v", cloned.Length())}
		case "clone-empty":
			orig := coreinstruction.EmptyIdentifiersWithGlobals()
			cloned := orig.Clone()
			actLines = []string{fmt.Sprintf("%v", cloned != nil), fmt.Sprintf("%v", cloned.Length())}
		case "clone-preserves":
			orig := coreinstruction.NewIdentifiersWithGlobals(false, "id-1", "id-2")
			cloned := orig.Clone()
			item := cloned.GetById("id-1")
			actLines = []string{fmt.Sprintf("%v", item != nil), item.Id, fmt.Sprintf("%v", item.IsGlobal)}
		case "add-single":
			ids := coreinstruction.EmptyIdentifiersWithGlobals()
			ids.Add(true, "new-id")
			found := ids.GetById("new-id")
			actLines = []string{fmt.Sprintf("%v", ids.Length()), fmt.Sprintf("%v", found != nil), fmt.Sprintf("%v", found.IsGlobal)}
		case "add-empty":
			ids := coreinstruction.EmptyIdentifiersWithGlobals()
			ids.Add(true, "")
			actLines = []string{fmt.Sprintf("%v", ids.Length())}
		case "add-multiple":
			ids := coreinstruction.NewIdentifiersWithGlobals(false, "existing")
			ids.Add(true, "second")
			ids.Add(false, "third")
			actLines = []string{fmt.Sprintf("%v", ids.Length()), fmt.Sprintf("%v", ids.GetById("second").IsGlobal), fmt.Sprintf("%v", ids.GetById("third").IsGlobal)}
		case "isempty-true":
			ids := coreinstruction.EmptyIdentifiersWithGlobals()
			actLines = []string{fmt.Sprintf("%v", ids.IsEmpty()), fmt.Sprintf("%v", ids.HasAnyItem())}
		case "isempty-false":
			ids := coreinstruction.NewIdentifiersWithGlobals(true, "item")
			actLines = []string{fmt.Sprintf("%v", ids.IsEmpty()), fmt.Sprintf("%v", ids.HasAnyItem())}
		case "indexof-found":
			ids := coreinstruction.NewIdentifiersWithGlobals(true, "a", "b", "c")
			actLines = []string{fmt.Sprintf("%v", ids.IndexOf("a")), fmt.Sprintf("%v", ids.IndexOf("b")), fmt.Sprintf("%v", ids.IndexOf("c"))}
		case "indexof-missing":
			ids := coreinstruction.NewIdentifiersWithGlobals(false, "x")
			actLines = []string{fmt.Sprintf("%v", ids.IndexOf("missing"))}
		case "indexof-empty":
			ids := coreinstruction.NewIdentifiersWithGlobals(true, "a")
			actLines = []string{fmt.Sprintf("%v", ids.IndexOf(""))}
		case "indexof-emptycol":
			ids := coreinstruction.EmptyIdentifiersWithGlobals()
			actLines = []string{fmt.Sprintf("%v", ids.IndexOf("any"))}
		case "adds-batch":
			ids := coreinstruction.EmptyIdentifiersWithGlobals()
			ids.Adds(true, "one", "two", "three")
			actLines = []string{
				fmt.Sprintf("%v", ids.Length()),
				fmt.Sprintf("%v", ids.GetById("one") != nil),
				fmt.Sprintf("%v", ids.GetById("two") != nil),
				fmt.Sprintf("%v", ids.GetById("three") != nil),
			}
		case "adds-empty":
			ids := coreinstruction.EmptyIdentifiersWithGlobals()
			ids.Adds(true)
			actLines = []string{fmt.Sprintf("%v", ids.Length())}
		case "new-noids":
			ids := coreinstruction.NewIdentifiersWithGlobals(true)
			actLines = []string{fmt.Sprintf("%v", ids != nil), fmt.Sprintf("%v", ids.Length())}
		}

		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}
