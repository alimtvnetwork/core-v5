package corecmptests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var isStringsEqualPtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "Both nil returns true",
		ArrangeInput:  args.Map{"leftNil": true, "rightNil": true},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "Left nil right non-nil returns false",
		ArrangeInput:  args.Map{"leftNil": true, "rightNil": false, "right": []string{"a"}},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "Right nil left non-nil returns false",
		ArrangeInput:  args.Map{"leftNil": false, "rightNil": true, "left": []string{"a"}},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "Equal slices returns true",
		ArrangeInput:  args.Map{"leftNil": false, "rightNil": false, "left": []string{"a", "b"}, "right": []string{"a", "b"}},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "Different length returns false",
		ArrangeInput:  args.Map{"leftNil": false, "rightNil": false, "left": []string{"a"}, "right": []string{"a", "b"}},
		ExpectedInput: args.Map{"result": false},
	},
}

var timePtrTestCases = []coretestcases.CaseV1{
	{
		Title:         "Both nil returns Equal",
		ArrangeInput:  args.Map{"leftNil": true, "rightNil": true},
		ExpectedInput: args.Map{"isEqual": true},
	},
	{
		Title:         "Left nil returns NotEqual",
		ArrangeInput:  args.Map{"leftNil": true, "rightNil": false},
		ExpectedInput: args.Map{"isEqual": false},
	},
	{
		Title:         "Right nil returns NotEqual",
		ArrangeInput:  args.Map{"leftNil": false, "rightNil": true},
		ExpectedInput: args.Map{"isEqual": false},
	},
	{
		Title:         "Both same time returns Equal",
		ArrangeInput:  args.Map{"leftNil": false, "rightNil": false, "sameTime": true},
		ExpectedInput: args.Map{"isEqual": true},
	},
}
