package coredynamictests

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

type leftRightTestCase struct {
	Case coretestcases.CaseV1
	LR   *coredynamic.LeftRight
}

// ==========================================
// IsEmpty
// ==========================================

var leftRightIsEmptyTestCases = []leftRightTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsEmpty true on nil receiver",
			ExpectedInput: "true",
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsEmpty true when both nil",
			ExpectedInput: "true",
		},
		LR: &coredynamic.LeftRight{Left: nil, Right: nil},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsEmpty false when has left",
			ExpectedInput: "false",
		},
		LR: &coredynamic.LeftRight{Left: "hello", Right: nil},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsEmpty false when has right",
			ExpectedInput: "false",
		},
		LR: &coredynamic.LeftRight{Left: nil, Right: 42},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsEmpty false when both set",
			ExpectedInput: "false",
		},
		LR: &coredynamic.LeftRight{Left: "a", Right: "b"},
	},
}

// ==========================================
// HasLeft / HasRight
// ==========================================

var leftRightHasLeftTestCases = []leftRightTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "HasLeft false on nil receiver",
			ExpectedInput: "false",
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasLeft true when present",
			ExpectedInput: "true",
		},
		LR: &coredynamic.LeftRight{Left: "hello"},
	},
}

var leftRightHasRightTestCases = []leftRightTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "HasRight false on nil receiver",
			ExpectedInput: "false",
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasRight true when present",
			ExpectedInput: "true",
		},
		LR: &coredynamic.LeftRight{Right: 42},
	},
}

// ==========================================
// IsLeftEmpty / IsRightEmpty
// ==========================================

var leftRightIsLeftEmptyTestCases = []leftRightTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsLeftEmpty true on nil receiver",
			ExpectedInput: "true",
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsLeftEmpty true when nil left",
			ExpectedInput: "true",
		},
		LR: &coredynamic.LeftRight{Left: nil, Right: "x"},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsLeftEmpty false when non-nil left",
			ExpectedInput: "false",
		},
		LR: &coredynamic.LeftRight{Left: "x"},
	},
}

var leftRightIsRightEmptyTestCases = []leftRightTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsRightEmpty true on nil receiver",
			ExpectedInput: "true",
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsRightEmpty false when non-nil right",
			ExpectedInput: "false",
		},
		LR: &coredynamic.LeftRight{Right: "y"},
	},
}

// ==========================================
// DeserializeLeft / DeserializeRight nil safety
// ==========================================

var leftRightDeserializeLeftTestCases = []leftRightTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "DeserializeLeft nil on nil receiver",
			ExpectedInput: "true",
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "DeserializeLeft valid data returns non-nil no error",
			ExpectedInput: args.Map{
				"isNil":    false,
				"hasError": false,
			},
		},
		LR: &coredynamic.LeftRight{Left: map[string]string{"key": "val"}},
	},
}

var leftRightDeserializeRightTestCases = []leftRightTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "DeserializeRight nil on nil receiver",
			ExpectedInput: "true",
		},
		LR: nil,
	},
}

// ==========================================
// TypeStatus nil safety
// ==========================================

var leftRightTypeStatusTestCases = []leftRightTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "TypeStatus on nil receiver returns both-null status",
			ExpectedInput: args.Map{
				"isSame":             "true",
				"isLeftUnknownNull":  "true",
				"isRightUnknownNull": "true",
			},
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "TypeStatus same value types",
			ExpectedInput: args.Map{
				"isSame":             "true",
				"isLeftUnknownNull":  "false",
				"isRightUnknownNull": "false",
			},
		},
		LR: &coredynamic.LeftRight{Left: 1, Right: 2},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "TypeStatus different value types",
			ExpectedInput: args.Map{
				"isSame":             "false",
				"isLeftUnknownNull":  "false",
				"isRightUnknownNull": "false",
			},
		},
		LR: &coredynamic.LeftRight{Left: "hello", Right: 42},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "TypeStatus mixed pointer and value of same underlying type",
			ExpectedInput: args.Map{
				"isSame":             "false",
				"isLeftUnknownNull":  "false",
				"isRightUnknownNull": "false",
				"isLeftPointer":      "false",
				"isRightPointer":     "true",
			},
		},
		LR: func() *coredynamic.LeftRight {
			v := 10
			return &coredynamic.LeftRight{Left: 5, Right: &v}
		}(),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "TypeStatus both pointer types same",
			ExpectedInput: args.Map{
				"isSame":             "true",
				"isLeftUnknownNull":  "false",
				"isRightUnknownNull": "false",
				"isLeftPointer":      "true",
				"isRightPointer":     "true",
			},
		},
		LR: func() *coredynamic.LeftRight {
			a, b := 1, 2
			return &coredynamic.LeftRight{Left: &a, Right: &b}
		}(),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "TypeStatus left nil right non-nil",
			ExpectedInput: args.Map{
				"isSame":             "false",
				"isLeftUnknownNull":  "true",
				"isRightUnknownNull": "false",
			},
		},
		LR: &coredynamic.LeftRight{Left: nil, Right: "value"},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "TypeStatus left non-nil right nil",
			ExpectedInput: args.Map{
				"isSame":             "false",
				"isLeftUnknownNull":  "false",
				"isRightUnknownNull": "true",
			},
		},
		LR: &coredynamic.LeftRight{Left: "value", Right: nil},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "TypeStatus both pointer types different",
			ExpectedInput: args.Map{
				"isSame":             "false",
				"isLeftUnknownNull":  "false",
				"isRightUnknownNull": "false",
				"isLeftPointer":      "true",
				"isRightPointer":     "true",
			},
		},
		LR: func() *coredynamic.LeftRight {
			i := 1
			s := "x"
			return &coredynamic.LeftRight{Left: &i, Right: &s}
		}(),
	},
}
