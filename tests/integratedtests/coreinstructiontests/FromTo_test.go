package coreinstructiontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

var fromToTestCases = []coretestcases.CaseV1{
	{Title: "ClonePtr - copies From and To", ArrangeInput: args.Map{"case": "cloneptr-positive"}, ExpectedInput: []string{"true", "source", "destination"}},
	{Title: "ClonePtr - nil receiver returns nil", ArrangeInput: args.Map{"case": "cloneptr-nil"}, ExpectedInput: []string{"true"}},
	{Title: "Clone - copies values", ArrangeInput: args.Map{"case": "clone"}, ExpectedInput: []string{"a", "b"}},
	{Title: "IsNull - nil returns true", ArrangeInput: args.Map{"case": "isnull-true"}, ExpectedInput: []string{"true"}},
	{Title: "IsNull - non-nil returns false", ArrangeInput: args.Map{"case": "isnull-false"}, ExpectedInput: []string{"false"}},
	{Title: "IsFromEmpty - empty From returns true", ArrangeInput: args.Map{"case": "isfromempty-true"}, ExpectedInput: []string{"true"}},
	{Title: "IsFromEmpty - nil receiver returns true", ArrangeInput: args.Map{"case": "isfromempty-nil"}, ExpectedInput: []string{"true"}},
	{Title: "IsToEmpty - empty To returns true", ArrangeInput: args.Map{"case": "istoempty-true"}, ExpectedInput: []string{"true"}},
	{Title: "IsToEmpty - non-empty returns false", ArrangeInput: args.Map{"case": "istoempty-false"}, ExpectedInput: []string{"false"}},
	{Title: "String - contains From and To", ArrangeInput: args.Map{"case": "string"}, ExpectedInput: []string{"true", "true"}},
	{Title: "FromName/ToName return field values", ArrangeInput: args.Map{"case": "fromto-names"}, ExpectedInput: []string{"src", "dst"}},
	{Title: "SetFromName - updates From", ArrangeInput: args.Map{"case": "setfromname"}, ExpectedInput: []string{"new"}},
	{Title: "SetToName - updates To", ArrangeInput: args.Map{"case": "settoname"}, ExpectedInput: []string{"new"}},
	{Title: "SetFromName - nil receiver no panic", ArrangeInput: args.Map{"case": "setfromname-nil"}, ExpectedInput: []string{"true"}},
	{Title: "SourceDestination - maps From->Source To->Destination", ArrangeInput: args.Map{"case": "sourcedest"}, ExpectedInput: []string{"true", "src", "dst"}},
	{Title: "SourceDestination - nil returns nil", ArrangeInput: args.Map{"case": "sourcedest-nil"}, ExpectedInput: []string{"true"}},
	{Title: "Rename - maps From->Existing To->New", ArrangeInput: args.Map{"case": "rename"}, ExpectedInput: []string{"true", "old", "new"}},
	{Title: "Rename - nil returns nil", ArrangeInput: args.Map{"case": "rename-nil"}, ExpectedInput: []string{"true"}},
}

func Test_FromTo_Verification(t *testing.T) {
	for caseIndex, tc := range fromToTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		caseType := input["case"].(string)

		var actLines []string

		// Act
		switch caseType {
		case "cloneptr-positive":
			orig := &coreinstruction.FromTo{From: "source", To: "destination"}
			cloned := orig.ClonePtr()
			actLines = []string{fmt.Sprintf("%v", cloned != nil), cloned.From, cloned.To}
		case "cloneptr-nil":
			var nilFT *coreinstruction.FromTo
			actLines = []string{fmt.Sprintf("%v", nilFT.ClonePtr() == nil)}
		case "clone":
			orig := coreinstruction.FromTo{From: "a", To: "b"}
			c := orig.Clone()
			actLines = []string{c.From, c.To}
		case "isnull-true":
			var nilFT *coreinstruction.FromTo
			actLines = []string{fmt.Sprintf("%v", nilFT.IsNull())}
		case "isnull-false":
			ft := &coreinstruction.FromTo{From: "x", To: "y"}
			actLines = []string{fmt.Sprintf("%v", ft.IsNull())}
		case "isfromempty-true":
			ft := &coreinstruction.FromTo{From: "", To: "dest"}
			actLines = []string{fmt.Sprintf("%v", ft.IsFromEmpty())}
		case "isfromempty-nil":
			var nilFT *coreinstruction.FromTo
			actLines = []string{fmt.Sprintf("%v", nilFT.IsFromEmpty())}
		case "istoempty-true":
			ft := &coreinstruction.FromTo{From: "src", To: ""}
			actLines = []string{fmt.Sprintf("%v", ft.IsToEmpty())}
		case "istoempty-false":
			ft := &coreinstruction.FromTo{From: "src", To: "dest"}
			actLines = []string{fmt.Sprintf("%v", ft.IsToEmpty())}
		case "string":
			ft := coreinstruction.FromTo{From: "alpha", To: "beta"}
			s := ft.String()
			actLines = []string{
				fmt.Sprintf("%v", len(s) > 0 && contains(s, "alpha")),
				fmt.Sprintf("%v", contains(s, "beta")),
			}
		case "fromto-names":
			ft := coreinstruction.FromTo{From: "src", To: "dst"}
			actLines = []string{ft.FromName(), ft.ToName()}
		case "setfromname":
			ft := &coreinstruction.FromTo{From: "old", To: "t"}
			ft.SetFromName("new")
			actLines = []string{ft.From}
		case "settoname":
			ft := &coreinstruction.FromTo{From: "f", To: "old"}
			ft.SetToName("new")
			actLines = []string{ft.To}
		case "setfromname-nil":
			var nilFT *coreinstruction.FromTo
			didPanic := false
			func() {
				defer func() {
					if r := recover(); r != nil {
						didPanic = true
					}
				}()
				nilFT.SetFromName("x")
			}()
			actLines = []string{fmt.Sprintf("%v", !didPanic)}
		case "sourcedest":
			ft := &coreinstruction.FromTo{From: "src", To: "dst"}
			sd := ft.SourceDestination()
			actLines = []string{fmt.Sprintf("%v", sd != nil), sd.Source, sd.Destination}
		case "sourcedest-nil":
			var nilFT *coreinstruction.FromTo
			actLines = []string{fmt.Sprintf("%v", nilFT.SourceDestination() == nil)}
		case "rename":
			ft := &coreinstruction.FromTo{From: "old", To: "new"}
			rn := ft.Rename()
			actLines = []string{fmt.Sprintf("%v", rn != nil), rn.Existing, rn.New}
		case "rename-nil":
			var nilFT *coreinstruction.FromTo
			actLines = []string{fmt.Sprintf("%v", nilFT.Rename() == nil)}
		}

		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, expectedLines)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || findSubstr(s, substr))
}

func findSubstr(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
