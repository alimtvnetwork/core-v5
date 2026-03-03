package coredynamictests

import (
	"fmt"
	"reflect"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// Test: Constructors
// ==========================================

func Test_Dynamic_Constructors_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicConstructorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)

		var actLines []string

		if input.Has("constructorRef") {
			ref := input["constructorRef"]

			switch fn := ref.(type) {
			case func(any) coredynamic.Dynamic:
				d := fn(input["inputData"])
				actLines = []string{
					fmt.Sprintf("%v", d.IsValid()),
					fmt.Sprintf("%v", d.Value()),
				}
			case func(any, bool) coredynamic.Dynamic:
				isValid := input["isValid"].(bool)
				d := fn(nil, isValid)
				actLines = []string{
					fmt.Sprintf("%v", d.IsValid()),
					fmt.Sprintf("%v", d.IsInvalid()),
				}
			case func() coredynamic.Dynamic:
				d := fn()
				actLines = []string{
					fmt.Sprintf("%v", d.IsValid()),
					fmt.Sprintf("%v", d.IsNull()),
				}
			case func() *coredynamic.Dynamic:
				d := fn()
				actLines = []string{
					fmt.Sprintf("%v", d != nil),
					fmt.Sprintf("%v", d.IsValid()),
					fmt.Sprintf("%v", d.IsNull()),
				}
			case func(any, bool) *coredynamic.Dynamic:
				isValid := input["isValid"].(bool)
				d := fn(input["inputData"], isValid)
				actLines = []string{
					fmt.Sprintf("%v", d != nil),
					fmt.Sprintf("%v", d.IsValid()),
					fmt.Sprintf("%v", d.Value()),
				}
			default:
				errcore.HandleErrMessage(
					fmt.Sprintf("unknown constructor ref type: %T", ref),
				)
			}
		}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: Clone
// ==========================================

func Test_Dynamic_Clone(t *testing.T) {
	tc := dynamicCloneTestCases[0]
	inputData, _ := tc.ArrangeInput.(args.Map).GetAsString("inputData")
	original := refNewDynamicValid(inputData)
	cloned := original.Clone()

	actLines := []string{
		fmt.Sprintf("%v", cloned.Value()),
		fmt.Sprintf("%v", cloned.IsValid()),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Dynamic_ClonePtr_NilReceiver(t *testing.T) {
	tc := dynamicCloneTestCases[1]
	var d *coredynamic.Dynamic

	actLines := []string{fmt.Sprintf("%v", d.ClonePtr() == nil)}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Dynamic_ClonePtr_Valid(t *testing.T) {
	tc := dynamicCloneTestCases[2]
	inputData, _ := tc.ArrangeInput.(args.Map).GetAsString("inputData")
	original := refNewDynamicPtr(inputData, true)
	cloned := original.ClonePtr()

	actLines := []string{
		fmt.Sprintf("%v", cloned != nil),
		fmt.Sprintf("%v", cloned.Value()),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Dynamic_NonPtr(t *testing.T) {
	tc := dynamicCloneTestCases[3]
	inputData, _ := tc.ArrangeInput.(args.Map).GetAsString("inputData")
	d := refNewDynamicValid(inputData)

	actLines := []string{fmt.Sprintf("%v", d.NonPtr().Value())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: Type Checks
// ==========================================

func Test_Dynamic_TypeChecks_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicTypeCheckTestCases {
		input := tc.ArrangeInput.(args.Map)
		var actLines []string

		scenario, hasScenario := input.GetAsString("scenario")

		if hasScenario {
			actLines = typeCheckSpecialScenario(scenario, input)
		} else {
			checkRef := input["checkRef"].(DynamicBoolMethodRef)
			d := createDynamicFromInput(input)
			actLines = []string{fmt.Sprintf("%v", checkRef.Call(d))}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

func typeCheckSpecialScenario(scenario string, input args.Map) []string {
	switch scenario {
	case "dataValue":
		inputData := input["inputData"]
		d := refNewDynamicValid(inputData)
		return []string{
			fmt.Sprintf("%v", d.Data()),
			fmt.Sprintf("%v", d.Data() == d.Value()),
		}
	case "stringNonEmpty":
		inputData, _ := input.GetAsString("inputData")
		d := refNewDynamicValid(inputData)
		return []string{fmt.Sprintf("%v", d.String() != "")}
	case "pointer":
		val := 42
		d := refNewDynamicValid(&val)
		checkRef := input["checkRef"].(DynamicBoolMethodRef)
		return []string{fmt.Sprintf("%v", checkRef.Call(&d))}
	default:
		errcore.HandleErrMessage("unknown typeCheck scenario: " + scenario)
		return nil
	}
}

// ==========================================
// Test: IsStruct
// ==========================================

func Test_Dynamic_IsStruct_Verification(t *testing.T) {
	type sample struct{ Name string }

	for caseIndex, tc := range dynamicIsStructTestCases {
		input := tc.ArrangeInput.(args.Map)
		checkRef := input["checkRef"].(DynamicBoolMethodRef)
		scenario, hasScenario := input.GetAsString("scenario")

		var d coredynamic.Dynamic
		if hasScenario && scenario == "struct" {
			d = refNewDynamicValid(sample{Name: "test"})
		} else {
			d = refNewDynamicValid(input["inputData"])
		}

		actLines := []string{fmt.Sprintf("%v", checkRef.Call(&d))}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: Length
// ==========================================

func Test_Dynamic_Length_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicLengthTestCases {
		input := tc.ArrangeInput.(args.Map)
		d := createDynamicFromInput(input)

		actLines := []string{fmt.Sprintf("%d", d.Length())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueInt
// ==========================================

func Test_Dynamic_ValueInt_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueIntTestCases {
		input := tc.ArrangeInput.(args.Map)
		d := refNewDynamicValid(input["inputData"])

		actLines := []string{fmt.Sprintf("%d", d.ValueInt())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueBool
// ==========================================

func Test_Dynamic_ValueBool_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueBoolTestCases {
		input := tc.ArrangeInput.(args.Map)
		d := refNewDynamicValid(input["inputData"])

		actLines := []string{fmt.Sprintf("%v", d.ValueBool())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueString
// ==========================================

func Test_Dynamic_ValueString_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueStringTestCases {
		input := tc.ArrangeInput.(args.Map)
		d := createDynamicFromInput(input)

		var actLines []string
		result := d.ValueString()
		inputData := input["inputData"]
		_, isString := inputData.(string)

		if inputData == nil {
			actLines = []string{fmt.Sprintf("%v", result == "")}
		} else if isString && inputData.(string) == "hello" {
			actLines = []string{result}
		} else {
			actLines = []string{fmt.Sprintf("%v", result != "")}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueStrings
// ==========================================

func Test_Dynamic_ValueStrings_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueStringsTestCases {
		input := tc.ArrangeInput.(args.Map)
		d := refNewDynamicValid(input["inputData"])

		var actLines []string
		result := d.ValueStrings()

		if result != nil {
			actLines = result
		} else {
			actLines = []string{fmt.Sprintf("%v", result == nil)}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueUInt
// ==========================================

func Test_Dynamic_ValueUInt_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueUIntTestCases {
		input := tc.ArrangeInput.(args.Map)
		d := refNewDynamicValid(input["inputData"])

		actLines := []string{fmt.Sprintf("%d", d.ValueUInt())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueInt64
// ==========================================

func Test_Dynamic_ValueInt64_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueInt64TestCases {
		input := tc.ArrangeInput.(args.Map)
		d := refNewDynamicValid(input["inputData"])

		actLines := []string{fmt.Sprintf("%d", d.ValueInt64())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: Bytes
// ==========================================

func Test_Dynamic_Bytes_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicBytesTestCases {
		input := tc.ArrangeInput.(args.Map)
		nilReceiver := getBoolDefault(input, "nilReceiver")

		var actLines []string

		if nilReceiver {
			var d *coredynamic.Dynamic
			raw, ok := d.Bytes()
			actLines = []string{
				fmt.Sprintf("%v", raw == nil),
				fmt.Sprintf("%v", ok),
			}
		} else {
			d := refNewDynamicValid(input["inputData"])
			raw, ok := d.Bytes()
			if ok {
				actLines = []string{
					fmt.Sprintf("%v", ok),
					string(raw),
				}
			} else {
				actLines = []string{fmt.Sprintf("%v", ok)}
			}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: IntDefault
// ==========================================

func Test_Dynamic_IntDefault_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicIntDefaultTestCases {
		input := tc.ArrangeInput.(args.Map)
		defaultVal := input.GetAsIntDefault("defaultValue", 0)
		d := createDynamicFromInput(input)

		val, ok := d.IntDefault(defaultVal)
		actLines := []string{
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", val),
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueNullErr
// ==========================================

func Test_Dynamic_ValueNullErr_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueNullErrTestCases {
		input := tc.ArrangeInput.(args.Map)
		nilReceiver := getBoolDefault(input, "nilReceiver")

		var actLines []string

		if nilReceiver {
			var d *coredynamic.Dynamic
			actLines = []string{fmt.Sprintf("%v", d.ValueNullErr() != nil)}
		} else {
			d := createDynamicFromInput(input)
			actLines = []string{fmt.Sprintf("%v", d.ValueNullErr() != nil)}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: Reflect — ReflectKind
// ==========================================

func Test_Dynamic_ReflectKind(t *testing.T) {
	tc := dynamicReflectTestCases[0]
	input := tc.ArrangeInput.(args.Map)
	d := refNewDynamicValid(input["inputData"])

	actLines := []string{d.ReflectKind().String()}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: Reflect — IsReflectKind Match
// ==========================================

func Test_Dynamic_IsReflectKindMatch(t *testing.T) {
	tc := dynamicReflectTestCases[1]
	input := tc.ArrangeInput.(args.Map)
	d := refNewDynamicValid(input["inputData"])

	actLines := []string{fmt.Sprintf("%v", d.IsReflectKind(reflect.String))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: Reflect — IsReflectKind Mismatch
// ==========================================

func Test_Dynamic_IsReflectKindMismatch(t *testing.T) {
	tc := dynamicReflectTestCases[2]
	input := tc.ArrangeInput.(args.Map)
	d := refNewDynamicValid(input["inputData"])

	actLines := []string{fmt.Sprintf("%v", d.IsReflectKind(reflect.Int))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: Reflect — ReflectTypeName
// ==========================================

func Test_Dynamic_ReflectTypeName(t *testing.T) {
	tc := dynamicReflectTestCases[3]
	input := tc.ArrangeInput.(args.Map)
	d := refNewDynamicValid(input["inputData"])

	actLines := []string{fmt.Sprintf("%v", d.ReflectTypeName() != "")}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: Reflect — ReflectType
// ==========================================

func Test_Dynamic_ReflectType(t *testing.T) {
	tc := dynamicReflectTestCases[4]
	input := tc.ArrangeInput.(args.Map)
	d := refNewDynamicValid(input["inputData"])

	actLines := []string{fmt.Sprintf("%v", d.ReflectType() == reflect.TypeOf(input["inputData"]))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: Reflect — IsReflectTypeOf
// ==========================================

func Test_Dynamic_IsReflectTypeOf(t *testing.T) {
	tc := dynamicReflectTestCases[5]
	input := tc.ArrangeInput.(args.Map)
	d := refNewDynamicValid(input["inputData"])

	actLines := []string{fmt.Sprintf("%v", d.IsReflectTypeOf(reflect.TypeOf("")))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: ReflectValue (cached)
// ==========================================

func Test_Dynamic_ReflectValue_Verification(t *testing.T) {
	d := refNewDynamicPtr(42, true)

	rv1 := d.ReflectValue()
	rv2 := d.ReflectValue()

	actLines := []string{
		fmt.Sprintf("%v", rv1 == rv2),
		fmt.Sprintf("%d", rv1.Int()),
	}
	expected := []string{"true", "42"}

	errcore.AssertDiffOnMismatch(t, 0, "ReflectValue returns cached reflect.Value", actLines, expected)
}

// ==========================================
// Test: Loop — Iterate
// ==========================================

func Test_Dynamic_Loop_Iterate(t *testing.T) {
	tc := dynamicLoopTestCases[0]
	d := refNewDynamicValid([]string{"a", "b", "c"})
	collected := make([]string, 0, 3)
	called := d.Loop(func(index int, item any) bool {
		collected = append(collected, item.(string))
		return false
	})

	actLines := append([]string{fmt.Sprintf("%v", called)}, collected...)

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: Loop — Invalid
// ==========================================

func Test_Dynamic_Loop_Invalid(t *testing.T) {
	tc := dynamicLoopTestCases[1]
	d := refInvalidDynamicPtr()
	called := d.Loop(func(index int, item any) bool { return false })

	actLines := []string{fmt.Sprintf("%v", called)}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: Loop — Break
// ==========================================

func Test_Dynamic_Loop_Break(t *testing.T) {
	tc := dynamicLoopTestCases[2]
	d := refNewDynamicValid([]int{1, 2, 3, 4})
	count := 0
	d.Loop(func(index int, item any) bool {
		count++
		return index == 1
	})

	actLines := []string{fmt.Sprintf("%d", count)}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: ItemAccess — ItemUsingIndex
// ==========================================

func Test_Dynamic_ItemUsingIndex(t *testing.T) {
	tc := dynamicItemAccessTestCases[0]
	d := refNewDynamicValid([]string{"a", "b"})

	actLines := []string{
		fmt.Sprintf("%v", d.ItemUsingIndex(0)),
		fmt.Sprintf("%v", d.ItemUsingIndex(1)),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: ItemAccess — ItemUsingKey
// ==========================================

func Test_Dynamic_ItemUsingKey(t *testing.T) {
	tc := dynamicItemAccessTestCases[1]
	d := refNewDynamicValid(map[string]int{"k": 42})

	actLines := []string{fmt.Sprintf("%v", d.ItemUsingKey("k"))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: IsStructStringNullOrEmpty
// ==========================================

func Test_Dynamic_StructStringNullOrEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicStructStringNullOrEmptyTestCases {
		input := tc.ArrangeInput.(args.Map)
		checkRef := input["checkRef"].(DynamicBoolMethodRef)
		d := createDynamicFromInput(input)

		actLines := []string{fmt.Sprintf("%v", checkRef.Call(d))}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// =============================================================================
// Helpers
// =============================================================================

func createDynamicFromInput(input args.Map) *coredynamic.Dynamic {
	inputData := input["inputData"]

	if isValid, has := input["isValid"]; has {
		d := refNewDynamic(inputData, isValid.(bool))
		return &d
	}

	d := refNewDynamicValid(inputData)
	return &d
}

func getBoolDefault(input args.Map, key string) bool {
	v, ok := input[key]
	if !ok {
		return false
	}

	b, isBool := v.(bool)
	if !isBool {
		return false
	}

	return b
}
