package coredynamictests

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
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
			ExpectedInput: []string{"true"},
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsEmpty true when both nil",
			ExpectedInput: []string{"true"},
		},
		LR: &coredynamic.LeftRight{Left: nil, Right: nil},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsEmpty false when has left",
			ExpectedInput: []string{"false"},
		},
		LR: &coredynamic.LeftRight{Left: "hello", Right: nil},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsEmpty false when has right",
			ExpectedInput: []string{"false"},
		},
		LR: &coredynamic.LeftRight{Left: nil, Right: 42},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsEmpty false when both set",
			ExpectedInput: []string{"false"},
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
			ExpectedInput: []string{"false"},
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasLeft true when present",
			ExpectedInput: []string{"true"},
		},
		LR: &coredynamic.LeftRight{Left: "hello"},
	},
}

var leftRightHasRightTestCases = []leftRightTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "HasRight false on nil receiver",
			ExpectedInput: []string{"false"},
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasRight true when present",
			ExpectedInput: []string{"true"},
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
			ExpectedInput: []string{"true"},
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsLeftEmpty true when nil left",
			ExpectedInput: []string{"true"},
		},
		LR: &coredynamic.LeftRight{Left: nil, Right: "x"},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsLeftEmpty false when non-nil left",
			ExpectedInput: []string{"false"},
		},
		LR: &coredynamic.LeftRight{Left: "x"},
	},
}

var leftRightIsRightEmptyTestCases = []leftRightTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsRightEmpty true on nil receiver",
			ExpectedInput: []string{"true"},
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsRightEmpty false when non-nil right",
			ExpectedInput: []string{"false"},
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
			ExpectedInput: []string{"true"},
		},
		LR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "DeserializeLeft valid data returns non-nil no error",
			ExpectedInput: []string{"false", "false"},
		},
		LR: &coredynamic.LeftRight{Left: map[string]string{"key": "val"}},
	},
}

var leftRightDeserializeRightTestCases = []leftRightTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "DeserializeRight nil on nil receiver",
			ExpectedInput: []string{"true"},
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
			Title:         "TypeStatus not nil on nil receiver",
			ExpectedInput: []string{"true"},
		},
		LR: nil,
	},
}
