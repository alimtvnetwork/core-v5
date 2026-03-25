package enumimpltests

import (
	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var dynamicMapSimpleDiffCaseV1TestCases = []coretestcases.CaseV1{
	{
		Title: "Dynamic map simple diff [someKey2] mismatch verify",
		ArrangeInput: args.Map{
			"left": map[string]any{
				"someKey":  1,
				"someKey2": 2,
				"someKey3": 3,
			},
			"right": map[string]any{
				"someKey":  1,
				"someKey2": 4,
				"someKey3": 3,
			},
			"checker": enumimpl.DefaultDiffCheckerImpl,
		},
		ExpectedInput: []string{
			"Dynamic map simple diff [someKey2] mismatch verify",
			"",
			"Difference Between Map:",
			"",
			"{",
			"- Left Map - Has Diff from Right Map:",
			"",
			"  {",
			"  ",
			"    \"someKey2\":4",
			"  ",
			"  }",
			"",
			"- Right Map - Has Diff from Left Map:",
			"",
			"  {",
			"  ",
			"    \"someKey2\":2",
			"  ",
			"  }}",
		},
	},
	{
		Title: "Dynamic map simple diff [someKey2], [someKey4] mismatch verify",
		ArrangeInput: args.Map{
			"left": map[string]any{
				"someKey":  1,
				"someKey2": 2,
				"someKey3": 3,
			},
			"right": map[string]any{
				"someKey":  1,
				"someKey4": 4,
				"someKey3": 3,
			},
			"checker": enumimpl.DefaultDiffCheckerImpl,
		},
		ExpectedInput: []string{
			"Dynamic map simple diff [someKey2], [someKey4] mismatch verify",
			"",
			"Difference Between Map:",
			"",
			"{",
			"- Left Map - Has Diff from Right Map:",
			"",
			"  {",
			"  ",
			"    \"someKey4\":4",
			"  ",
			"  }",
			"",
			"- Right Map - Has Diff from Left Map:",
			"",
			"  {",
			"  ",
			"    \"someKey2\":2",
			"  ",
			"  }}",
		},
	},
	{
		Title: "Dynamic map simple diff all match - no diff",
		ArrangeInput: args.Map{
			"left": map[string]any{
				"someKey":  1,
				"someKey2": 2,
				"someKey4": 4,
				"someKey3": 3,
			},
			"right": map[string]any{
				"someKey":  1,
				"someKey2": 2,
				"someKey4": 4,
				"someKey3": 3,
			},
			"checker": enumimpl.DefaultDiffCheckerImpl,
		},
		ExpectedInput: []string{
			"",
		},
	},
}
