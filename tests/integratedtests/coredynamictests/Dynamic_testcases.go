package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Input types (replace args.Map for type safety and no branching)
// ==========================================================================

type dynamicInputMap struct {
	InputData any
	IsValid   bool
}

type dynamicBoolCheckInput struct {
	CheckRef  DynamicBoolMethodRef
	InputData any
	IsValid   bool
}

// ==========================================================================
// Constructors
// ==========================================================================

var dynamicConstructorNewDynamicValidTestCase = coretestcases.CaseV1{
	Title:         "NewDynamicValid creates valid Dynamic",
	ExpectedInput: []string{"true", "hello"},
}

var dynamicConstructorNewDynamicInvalidTestCase = coretestcases.CaseV1{
	Title:         "NewDynamic with isValid=false creates invalid Dynamic",
	ExpectedInput: []string{"false", "true"},
}

var dynamicConstructorInvalidDynamicTestCase = coretestcases.CaseV1{
	Title:         "InvalidDynamic creates invalid nil Dynamic",
	ExpectedInput: []string{"false", "true"},
}

var dynamicConstructorInvalidDynamicPtrTestCase = coretestcases.CaseV1{
	Title:         "InvalidDynamicPtr creates invalid nil Dynamic pointer",
	ExpectedInput: []string{"true", "false", "true"},
}

var dynamicConstructorNewDynamicPtrTestCase = coretestcases.CaseV1{
	Title:         "NewDynamicPtr creates pointer Dynamic",
	ExpectedInput: []string{"true", "true", "42"},
}

// ==========================================================================
// Clone
// ==========================================================================

var dynamicCloneTestCase = coretestcases.CaseV1{
	Title:         "Clone creates independent copy",
	ExpectedInput: []string{"data", "true"},
}

var dynamicClonePtrNilTestCase = coretestcases.CaseV1{
	Title:         "ClonePtr returns nil on nil receiver",
	ExpectedInput: []string{"true"},
}

var dynamicClonePtrValidTestCase = coretestcases.CaseV1{
	Title:         "ClonePtr creates independent pointer copy",
	ExpectedInput: []string{"true", "data"},
}

var dynamicNonPtrTestCase = coretestcases.CaseV1{
	Title:         "NonPtr returns value copy",
	ExpectedInput: []string{"x"},
}

// ==========================================================================
// Type Checks — Special scenarios (split into individual test cases)
// ==========================================================================

var dynamicDataValueEqualityTestCase = coretestcases.CaseV1{
	Title:         "Data and Value return same inner data",
	ExpectedInput: []string{"99", "true"},
}

var dynamicStringNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "String returns non-empty for valid",
	ExpectedInput: []string{"true"},
}

var dynamicIsPointerTestCase = coretestcases.CaseV1{
	Title:         "IsPointer true for pointer data",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// Type Checks — Uniform bool method ref checks
// ==========================================================================

var dynamicTypeCheckTestCases = []coretestcases.CaseV1{
	{
		Title: "IsNull true for nil data",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsNull,
			InputData: nil,
			IsValid:   true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsNull false for non-nil data",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsNull,
			InputData: "x",
			IsValid:   true,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsStringType true for string",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsStringType,
			InputData: "text",
			IsValid:   true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsStringType false for int",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsStringType,
			InputData: 42,
			IsValid:   true,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsNumber true for int",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsNumber,
			InputData: 42,
			IsValid:   true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsNumber false for string",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsNumber,
			InputData: "x",
			IsValid:   true,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsPrimitive true for int",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsPrimitive,
			InputData: 10,
			IsValid:   true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsFunc true for function",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsFunc,
			InputData: func() {},
			IsValid:   true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsSliceOrArray true for slice",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsSliceOrArray,
			InputData: []int{1, 2, 3},
			IsValid:   true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsSliceOrArray false for string",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsSliceOrArray,
			InputData: "x",
			IsValid:   true,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsSliceOrArrayOrMap true for map",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsSliceOrArrayOrMap,
			InputData: map[string]int{"a": 1},
			IsValid:   true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsMap true for map",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsMap,
			InputData: map[string]int{"x": 1},
			IsValid:   true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsValueType true for non-pointer",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsValueType,
			InputData: 42,
			IsValid:   true,
		},
		ExpectedInput: []string{"true"},
	},
}

// ==========================================================================
// IsStruct
// ==========================================================================

var dynamicIsStructTrueTestCase = coretestcases.CaseV1{
	Title:         "IsStruct true for struct",
	ExpectedInput: []string{"true"},
}

var dynamicIsStructFalseTestCase = coretestcases.CaseV1{
	Title:         "IsStruct false for int",
	ExpectedInput: []string{"false"},
}

// ==========================================================================
// Length
// ==========================================================================

var dynamicLengthTestCases = []coretestcases.CaseV1{
	{
		Title: "Length returns slice length",
		ArrangeInput: dynamicInputMap{
			InputData: []int{1, 2, 3},
			IsValid:   true,
		},
		ExpectedInput: []string{"3"},
	},
	{
		Title: "Length returns 0 for nil data",
		ArrangeInput: dynamicInputMap{
			InputData: nil,
			IsValid:   false,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "Length returns map length",
		ArrangeInput: dynamicInputMap{
			InputData: map[string]int{"a": 1, "b": 2},
			IsValid:   true,
		},
		ExpectedInput: []string{"2"},
	},
}

// ==========================================================================
// Value Extraction
// ==========================================================================

var dynamicValueIntTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueInt returns int value",
		ArrangeInput: dynamicInputMap{
			InputData: 42,
			IsValid:   true,
		},
		ExpectedInput: []string{"42"},
	},
	{
		Title: "ValueInt returns -1 for non-int",
		ArrangeInput: dynamicInputMap{
			InputData: "not-int",
			IsValid:   true,
		},
		ExpectedInput: []string{"-1"},
	},
}

var dynamicValueBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueBool returns true",
		ArrangeInput: dynamicInputMap{
			InputData: true,
			IsValid:   true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "ValueBool returns false for non-bool",
		ArrangeInput: dynamicInputMap{
			InputData: "x",
			IsValid:   true,
		},
		ExpectedInput: []string{"false"},
	},
}

// ==========================================================================
// ValueString
// ==========================================================================

var dynamicValueStringDirectTestCase = coretestcases.CaseV1{
	Title:         "ValueString returns string directly",
	ExpectedInput: []string{"hello"},
}

var dynamicValueStringNonStringTestCase = coretestcases.CaseV1{
	Title:         "ValueString formats non-string as non-empty",
	ExpectedInput: []string{"true"},
}

var dynamicValueStringNilTestCase = coretestcases.CaseV1{
	Title:         "ValueString returns empty for nil data",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// ValueStrings
// ==========================================================================

var dynamicValueStringsSliceTestCase = coretestcases.CaseV1{
	Title:         "ValueStrings returns []string",
	ExpectedInput: []string{"a", "b"},
}

var dynamicValueStringsNonSliceTestCase = coretestcases.CaseV1{
	Title:         "ValueStrings returns nil for non-[]string",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// ValueUInt / ValueInt64
// ==========================================================================

var dynamicValueUIntTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueUInt returns uint value",
		ArrangeInput: dynamicInputMap{
			InputData: uint(10),
			IsValid:   true,
		},
		ExpectedInput: []string{"10"},
	},
}

var dynamicValueInt64TestCases = []coretestcases.CaseV1{
	{
		Title: "ValueInt64 returns int64 value",
		ArrangeInput: dynamicInputMap{
			InputData: int64(999),
			IsValid:   true,
		},
		ExpectedInput: []string{"999"},
	},
}

// ==========================================================================
// Bytes
// ==========================================================================

var dynamicBytesValidTestCase = coretestcases.CaseV1{
	Title:         "Bytes returns []byte",
	ExpectedInput: []string{"true", "raw"},
}

var dynamicBytesNonBytesTestCase = coretestcases.CaseV1{
	Title:         "Bytes returns false for non-bytes",
	ExpectedInput: []string{"false"},
}

var dynamicBytesNilReceiverTestCase = coretestcases.CaseV1{
	Title:         "Bytes returns nil,false on nil receiver",
	ExpectedInput: []string{"true", "false"},
}

// ==========================================================================
// IntDefault
// ==========================================================================

var dynamicIntDefaultValidTestCase = coretestcases.CaseV1{
	Title:         "IntDefault parses int value",
	ExpectedInput: []string{"true", "42"},
}

var dynamicIntDefaultNilTestCase = coretestcases.CaseV1{
	Title:         "IntDefault returns default on nil data",
	ExpectedInput: []string{"false", "99"},
}

// ==========================================================================
// ValueNullErr
// ==========================================================================

var dynamicValueNullErrNilReceiverTestCase = coretestcases.CaseV1{
	Title:         "ValueNullErr returns error on nil receiver",
	ExpectedInput: []string{"true"},
}

var dynamicValueNullErrNullDataTestCase = coretestcases.CaseV1{
	Title:         "ValueNullErr returns error on null data",
	ExpectedInput: []string{"true"},
}

var dynamicValueNullErrValidTestCase = coretestcases.CaseV1{
	Title:         "ValueNullErr returns nil for valid data",
	ExpectedInput: []string{"false"},
}

// ==========================================================================
// Reflect
// ==========================================================================

var dynamicReflectKindStringTestCase = coretestcases.CaseV1{
	Title:         "ReflectKind returns String for string",
	ExpectedInput: []string{"string"},
}

var dynamicReflectKindIntTestCase = coretestcases.CaseV1{
	Title:         "ReflectKind returns Int for int",
	ExpectedInput: []string{"int"},
}

var dynamicIsReflectKindMatchTestCase = coretestcases.CaseV1{
	Title:         "IsReflectKind matches correctly",
	ExpectedInput: []string{"true"},
}

var dynamicIsReflectKindMismatchTestCase = coretestcases.CaseV1{
	Title:         "IsReflectKind returns false on mismatch",
	ExpectedInput: []string{"false"},
}

var dynamicReflectTypeNameTestCase = coretestcases.CaseV1{
	Title:         "ReflectTypeName returns non-empty",
	ExpectedInput: []string{"true"},
}

var dynamicReflectTypeTestCase = coretestcases.CaseV1{
	Title:         "ReflectType returns correct type",
	ExpectedInput: []string{"true"},
}

var dynamicIsReflectTypeOfTestCase = coretestcases.CaseV1{
	Title:         "IsReflectTypeOf matches type",
	ExpectedInput: []string{"true"},
}

var dynamicReflectValueCachedTestCase = coretestcases.CaseV1{
	Title:         "ReflectValue returns cached reflect.Value",
	ExpectedInput: []string{"true", "42"},
}

// ==========================================================================
// Loop
// ==========================================================================

var dynamicLoopIterateTestCase = coretestcases.CaseV1{
	Title:         "Loop iterates slice items",
	ExpectedInput: []string{"true", "a", "b", "c"},
}

var dynamicLoopInvalidTestCase = coretestcases.CaseV1{
	Title:         "Loop returns false for invalid",
	ExpectedInput: []string{"false"},
}

var dynamicLoopBreakTestCase = coretestcases.CaseV1{
	Title:         "Loop respects break",
	ExpectedInput: []string{"2"},
}

// ==========================================================================
// ItemAccess
// ==========================================================================

var dynamicItemUsingIndexTestCase = coretestcases.CaseV1{
	Title:         "ItemUsingIndex returns correct element",
	ExpectedInput: []string{"a", "b"},
}

var dynamicItemUsingKeyTestCase = coretestcases.CaseV1{
	Title:         "ItemUsingKey returns map value",
	ExpectedInput: []string{"42"},
}

// ==========================================================================
// IsStructStringNullOrEmpty
// ==========================================================================

var dynamicStructStringNullOrEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsStructStringNullOrEmpty true on nil data",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsStructStringNullOrEmpty,
			InputData: nil,
			IsValid:   true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsStructStringNullOrEmpty false for non-empty",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsStructStringNullOrEmpty,
			InputData: "text",
			IsValid:   true,
		},
		ExpectedInput: []string{"false"},
	},
}
