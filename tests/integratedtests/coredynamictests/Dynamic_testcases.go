package coredynamictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
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
	Title: "NewDynamicValid creates valid Dynamic",
	ExpectedInput: args.Map{
		"isValid":   true,
		"dataValue": "hello",
	},
}

var dynamicConstructorNewDynamicInvalidTestCase = coretestcases.CaseV1{
	Title: "NewDynamic with isValid=false creates invalid Dynamic",
	ExpectedInput: args.Map{
		"isValid": false,
		"isNull":  true,
	},
}

var dynamicConstructorInvalidDynamicTestCase = coretestcases.CaseV1{
	Title: "InvalidDynamic creates invalid nil Dynamic",
	ExpectedInput: args.Map{
		"isValid": false,
		"isNull":  true,
	},
}

var dynamicConstructorInvalidDynamicPtrTestCase = coretestcases.CaseV1{
	Title: "InvalidDynamicPtr creates invalid nil Dynamic pointer",
	ExpectedInput: args.Map{
		"isNotNilPtr": true,
		"isValid":     false,
		"isNull":      true,
	},
}

var dynamicConstructorNewDynamicPtrTestCase = coretestcases.CaseV1{
	Title: "NewDynamicPtr creates pointer Dynamic",
	ExpectedInput: args.Map{
		"isNotNilPtr": true,
		"isValid":     true,
		"dataValue":   "42",
	},
}

// ==========================================================================
// Clone
// ==========================================================================

var dynamicCloneTestCase = coretestcases.CaseV1{
	Title: "Clone creates independent copy",
	ExpectedInput: args.Map{
		"clonedValue":   "data",
		"isIndependent": true,
	},
}

// Note: ClonePtr, Bytes, ValueNullErr nil receiver test cases migrated to
// Dynamic_NilReceiver_testcases.go using CaseNilSafe pattern.

var dynamicClonePtrValidTestCase = coretestcases.CaseV1{
	Title: "ClonePtr creates independent pointer copy",
	ExpectedInput: args.Map{
		"isNotNilPtr": true,
		"clonedValue": "data",
	},
}

var dynamicNonPtrTestCase = coretestcases.CaseV1{
	Title:         "NonPtr returns value copy",
	ExpectedInput: "x", // valueCopy
}

// ==========================================================================
// Type Checks — Special scenarios (split into individual test cases)
// ==========================================================================

var dynamicDataValueEqualityTestCase = coretestcases.CaseV1{
	Title: "Data and Value return same inner data",
	ExpectedInput: args.Map{
		"dataValue":       "99",
		"dataEqualsValue": true,
	},
}

var dynamicStringNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "String returns non-empty for valid",
	ExpectedInput: "true", // isNonEmpty
}

var dynamicIsPointerTestCase = coretestcases.CaseV1{
	Title:         "IsPointer true for pointer data",
	ExpectedInput: "true", // isPointer
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
		ExpectedInput: "true",
	},
	{
		Title: "IsNull false for non-nil data",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsNull,
			InputData: "x",
			IsValid:   true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsStringType true for string",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsStringType,
			InputData: "text",
			IsValid:   true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsStringType false for int",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsStringType,
			InputData: 42,
			IsValid:   true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsNumber true for int",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsNumber,
			InputData: 42,
			IsValid:   true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsNumber false for string",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsNumber,
			InputData: "x",
			IsValid:   true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsPrimitive true for int",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsPrimitive,
			InputData: 10,
			IsValid:   true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsFunc true for function",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsFunc,
			InputData: func() {},
			IsValid:   true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsSliceOrArray true for slice",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsSliceOrArray,
			InputData: []int{1, 2, 3},
			IsValid:   true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsSliceOrArray false for string",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsSliceOrArray,
			InputData: "x",
			IsValid:   true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsSliceOrArrayOrMap true for map",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsSliceOrArrayOrMap,
			InputData: map[string]int{"a": 1},
			IsValid:   true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsMap true for map",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsMap,
			InputData: map[string]int{"x": 1},
			IsValid:   true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsValueType true for non-pointer",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsValueType,
			InputData: 42,
			IsValid:   true,
		},
		ExpectedInput: "true",
	},
}

// ==========================================================================
// IsStruct
// ==========================================================================

var dynamicIsStructTrueTestCase = coretestcases.CaseV1{
	Title:         "IsStruct true for struct",
	ExpectedInput: "true",
}

var dynamicIsStructFalseTestCase = coretestcases.CaseV1{
	Title:         "IsStruct false for int",
	ExpectedInput: "false",
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
		ExpectedInput: "3",
	},
	{
		Title: "Length returns 0 for nil data",
		ArrangeInput: dynamicInputMap{
			InputData: nil,
			IsValid:   false,
		},
		ExpectedInput: "0",
	},
	{
		Title: "Length returns map length",
		ArrangeInput: dynamicInputMap{
			InputData: map[string]int{"a": 1, "b": 2},
			IsValid:   true,
		},
		ExpectedInput: "2",
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
		ExpectedInput: "42",
	},
	{
		Title: "ValueInt returns -1 for non-int",
		ArrangeInput: dynamicInputMap{
			InputData: "not-int",
			IsValid:   true,
		},
		ExpectedInput: "-1",
	},
}

var dynamicValueBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueBool returns true",
		ArrangeInput: dynamicInputMap{
			InputData: true,
			IsValid:   true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "ValueBool returns false for non-bool",
		ArrangeInput: dynamicInputMap{
			InputData: "x",
			IsValid:   true,
		},
		ExpectedInput: "false",
	},
}

// ==========================================================================
// ValueString
// ==========================================================================

var dynamicValueStringDirectTestCase = coretestcases.CaseV1{
	Title:         "ValueString returns string directly",
	ExpectedInput: "hello",
}

var dynamicValueStringNonStringTestCase = coretestcases.CaseV1{
	Title:         "ValueString formats non-string as non-empty",
	ExpectedInput: "true", // isNonEmpty
}

var dynamicValueStringNilTestCase = coretestcases.CaseV1{
	Title:         "ValueString returns empty for nil data",
	ExpectedInput: "true", // isEmpty
}

// ==========================================================================
// ValueStrings
// ==========================================================================

var dynamicValueStringsSliceTestCase = coretestcases.CaseV1{
	Title: "ValueStrings returns []string",
	ExpectedInput: args.Map{
		"item0": "a",
		"item1": "b",
	},
}

var dynamicValueStringsNonSliceTestCase = coretestcases.CaseV1{
	Title:         "ValueStrings returns nil for non-[]string",
	ExpectedInput: "true", // isNil
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
		ExpectedInput: "10",
	},
}

var dynamicValueInt64TestCases = []coretestcases.CaseV1{
	{
		Title: "ValueInt64 returns int64 value",
		ArrangeInput: dynamicInputMap{
			InputData: int64(999),
			IsValid:   true,
		},
		ExpectedInput: "999",
	},
}

// ==========================================================================
// Bytes
// ==========================================================================

var dynamicBytesValidTestCase = coretestcases.CaseV1{
	Title: "Bytes returns []byte",
	ExpectedInput: args.Map{
		"hasBytes": true,
		"content":  "raw",
	},
}

var dynamicBytesNonBytesTestCase = coretestcases.CaseV1{
	Title:         "Bytes returns false for non-bytes",
	ExpectedInput: "false", // hasBytes
}

// Note: Bytes nil receiver test case migrated to Dynamic_NilReceiver_testcases.go.

// ==========================================================================
// IntDefault
// ==========================================================================

var dynamicIntDefaultValidTestCase = coretestcases.CaseV1{
	Title: "IntDefault parses int value",
	ExpectedInput: args.Map{
		"isValid":  true,
		"intValue": 42,
	},
}

var dynamicIntDefaultNilTestCase = coretestcases.CaseV1{
	Title: "IntDefault returns default on nil data",
	ExpectedInput: args.Map{
		"isValid":      false,
		"defaultValue": 99,
	},
}

// ==========================================================================
// ValueNullErr
// ==========================================================================

// Note: ValueNullErr nil receiver test case migrated to Dynamic_NilReceiver_testcases.go.

var dynamicValueNullErrNullDataTestCase = coretestcases.CaseV1{
	Title:         "ValueNullErr returns error on null data",
	ExpectedInput: "true", // hasError
}

var dynamicValueNullErrValidTestCase = coretestcases.CaseV1{
	Title:         "ValueNullErr returns nil for valid data",
	ExpectedInput: "false", // hasError
}

// ==========================================================================
// Reflect
// ==========================================================================

var dynamicReflectKindStringTestCase = coretestcases.CaseV1{
	Title:         "ReflectKind returns String for string",
	ExpectedInput: "string",
}

var dynamicReflectKindIntTestCase = coretestcases.CaseV1{
	Title:         "ReflectKind returns Int for int",
	ExpectedInput: "int",
}

var dynamicIsReflectKindMatchTestCase = coretestcases.CaseV1{
	Title:         "IsReflectKind matches correctly",
	ExpectedInput: "true",
}

var dynamicIsReflectKindMismatchTestCase = coretestcases.CaseV1{
	Title:         "IsReflectKind returns false on mismatch",
	ExpectedInput: "false",
}

var dynamicReflectTypeNameTestCase = coretestcases.CaseV1{
	Title:         "ReflectTypeName returns non-empty",
	ExpectedInput: "true", // isNonEmpty
}

var dynamicReflectTypeTestCase = coretestcases.CaseV1{
	Title:         "ReflectType returns correct type",
	ExpectedInput: "true", // isCorrectType
}

var dynamicIsReflectTypeOfTestCase = coretestcases.CaseV1{
	Title:         "IsReflectTypeOf matches type",
	ExpectedInput: "true",
}

var dynamicReflectValueCachedTestCase = coretestcases.CaseV1{
	Title: "ReflectValue returns cached reflect.Value",
	ExpectedInput: args.Map{
		"isCached":       true,
		"extractedValue": 42,
	},
}

// ==========================================================================
// Loop
// ==========================================================================

var dynamicLoopIterateTestCase = coretestcases.CaseV1{
	Title: "Loop iterates slice items",
	ExpectedInput: args.Map{
		"didLoop": true,
		"item0":   "a",
		"item1":   "b",
		"item2":   "c",
	},
}

var dynamicLoopInvalidTestCase = coretestcases.CaseV1{
	Title:         "Loop returns false for invalid",
	ExpectedInput: "false", // didLoop
}

var dynamicLoopBreakTestCase = coretestcases.CaseV1{
	Title:         "Loop respects break",
	ExpectedInput: "2", // iterationCount
}

// ==========================================================================
// ItemAccess
// ==========================================================================

var dynamicItemUsingIndexTestCase = coretestcases.CaseV1{
	Title: "ItemUsingIndex returns correct element",
	ExpectedInput: args.Map{
		"item0": "a",
		"item1": "b",
	},
}

var dynamicItemUsingKeyTestCase = coretestcases.CaseV1{
	Title:         "ItemUsingKey returns map value",
	ExpectedInput: "42", // mapValue
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
		ExpectedInput: "true",
	},
	{
		Title: "IsStructStringNullOrEmpty false for non-empty",
		ArrangeInput: dynamicBoolCheckInput{
			CheckRef:  refIsStructStringNullOrEmpty,
			InputData: "text",
			IsValid:   true,
		},
		ExpectedInput: "false",
	},
}
