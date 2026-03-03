package coredynamictests

import (
	"fmt"
	"reflect"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// Test: Constructors
// ==========================================

func Test_Dynamic_Constructors_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicConstructorTestCases {
		input := tc.ArrangeInput.(args.Map)
		constructor, _ := input.GetAsString("constructor")

		var actLines []string

		switch constructor {
		case "NewDynamicValid":
			value, _ := input.GetAsString("value")
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{
				fmt.Sprintf("%v", d.IsValid()),
				fmt.Sprintf("%v", d.Value()),
			}
		case "NewDynamic":
			isValid := input["isValid"].(bool)
			d := coredynamic.NewDynamic(nil, isValid)
			actLines = []string{
				fmt.Sprintf("%v", d.IsValid()),
				fmt.Sprintf("%v", d.IsInvalid()),
			}
		case "InvalidDynamic":
			d := coredynamic.InvalidDynamic()
			actLines = []string{
				fmt.Sprintf("%v", d.IsValid()),
				fmt.Sprintf("%v", d.IsNull()),
			}
		case "InvalidDynamicPtr":
			d := coredynamic.InvalidDynamicPtr()
			actLines = []string{
				fmt.Sprintf("%v", d != nil),
				fmt.Sprintf("%v", d.IsValid()),
				fmt.Sprintf("%v", d.IsNull()),
			}
		case "NewDynamicPtr":
			value := input["value"]
			isValid := input["isValid"].(bool)
			d := coredynamic.NewDynamicPtr(value, isValid)
			actLines = []string{
				fmt.Sprintf("%v", d != nil),
				fmt.Sprintf("%v", d.IsValid()),
				fmt.Sprintf("%v", d.Value()),
			}
		default:
			errcore.HandleErrMessage("unknown constructor: " + constructor)
		}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Clone
// ==========================================

func Test_Dynamic_Clone_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicCloneTestCases {
		input := tc.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")
		receiver, _ := input.GetAsString("receiver")

		var actLines []string

		switch method {
		case "Clone":
			value, _ := input.GetAsString("value")
			original := coredynamic.NewDynamicValid(value)
			cloned := original.Clone()
			actLines = []string{
				fmt.Sprintf("%v", cloned.Value()),
				fmt.Sprintf("%v", cloned.IsValid()),
			}
		case "ClonePtr":
			if receiver == "nil" {
				var d *coredynamic.Dynamic
				actLines = []string{fmt.Sprintf("%v", d.ClonePtr() == nil)}
			} else {
				value, _ := input.GetAsString("value")
				original := coredynamic.NewDynamicPtr(value, true)
				cloned := original.ClonePtr()
				actLines = []string{
					fmt.Sprintf("%v", cloned != nil),
					fmt.Sprintf("%v", cloned.Value()),
				}
			}
		case "NonPtr":
			value, _ := input.GetAsString("value")
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{fmt.Sprintf("%v", d.NonPtr().Value())}
		default:
			errcore.HandleErrMessage("unknown method: " + method)
		}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Type Checks
// ==========================================

func Test_Dynamic_TypeChecks_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicTypeCheckTestCases {
		input := tc.ArrangeInput.(args.Map)
		check, _ := input.GetAsString("check")

		var actLines []string

		switch check {
		case "DataValue":
			value := input["value"]
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{
				fmt.Sprintf("%v", d.Data()),
				fmt.Sprintf("%v", d.Data() == d.Value()),
			}
		case "IsNull":
			isValid, hasIsValid := input["isValid"]
			if hasIsValid {
				d := coredynamic.NewDynamic(nil, isValid.(bool))
				actLines = []string{fmt.Sprintf("%v", d.IsNull())}
			} else {
				value, _ := input.GetAsString("value")
				d := coredynamic.NewDynamicValid(value)
				actLines = []string{fmt.Sprintf("%v", d.IsNull())}
			}
		case "String":
			value, _ := input.GetAsString("value")
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{fmt.Sprintf("%v", d.String() != "")}
		case "IsStringType":
			value := input["value"]
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{fmt.Sprintf("%v", d.IsStringType())}
		case "IsNumber":
			value := input["value"]
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{fmt.Sprintf("%v", d.IsNumber())}
		case "IsPrimitive":
			value := input["value"]
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{fmt.Sprintf("%v", d.IsPrimitive())}
		case "IsFunc":
			d := coredynamic.NewDynamicValid(func() {})
			actLines = []string{fmt.Sprintf("%v", d.IsFunc())}
		case "IsSliceOrArray":
			value, _ := input.GetAsString("value")
			if value == "slice_int" {
				d := coredynamic.NewDynamicValid([]int{1, 2, 3})
				actLines = []string{fmt.Sprintf("%v", d.IsSliceOrArray())}
			} else {
				d := coredynamic.NewDynamicValid(value)
				actLines = []string{fmt.Sprintf("%v", d.IsSliceOrArray())}
			}
		case "IsSliceOrArrayOrMap":
			d := coredynamic.NewDynamicValid(map[string]int{"a": 1})
			actLines = []string{fmt.Sprintf("%v", d.IsSliceOrArrayOrMap())}
		case "IsMap":
			d := coredynamic.NewDynamicValid(map[string]int{"x": 1})
			actLines = []string{fmt.Sprintf("%v", d.IsMap())}
		case "IsPointer":
			val := 42
			d := coredynamic.NewDynamicValid(&val)
			actLines = []string{fmt.Sprintf("%v", d.IsPointer())}
		case "IsValueType":
			value := input["value"]
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{fmt.Sprintf("%v", d.IsValueType())}
		default:
			errcore.HandleErrMessage("unknown check: " + check)
		}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IsStruct
// ==========================================

func Test_Dynamic_IsStruct_Verification(t *testing.T) {
	type sample struct{ Name string }

	for caseIndex, tc := range dynamicIsStructTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"]

		var d coredynamic.Dynamic
		if value == "struct" {
			d = coredynamic.NewDynamicValid(sample{Name: "test"})
		} else {
			d = coredynamic.NewDynamicValid(value)
		}

		actLines := []string{fmt.Sprintf("%v", d.IsStruct())}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Length
// ==========================================

func Test_Dynamic_Length_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicLengthTestCases {
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsString("value")

		var d coredynamic.Dynamic
		switch value {
		case "slice_int":
			d = coredynamic.NewDynamicValid([]int{1, 2, 3})
		case "nil":
			d = coredynamic.NewDynamic(nil, false)
		case "map_2":
			d = coredynamic.NewDynamicValid(map[string]int{"a": 1, "b": 2})
		default:
			errcore.HandleErrMessage("unknown value: " + value)
		}

		actLines := []string{fmt.Sprintf("%d", d.Length())}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ValueInt
// ==========================================

func Test_Dynamic_ValueInt_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueIntTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"]
		d := coredynamic.NewDynamicValid(value)

		actLines := []string{fmt.Sprintf("%d", d.ValueInt())}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ValueBool
// ==========================================

func Test_Dynamic_ValueBool_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueBoolTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"]
		d := coredynamic.NewDynamicValid(value)

		actLines := []string{fmt.Sprintf("%v", d.ValueBool())}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ValueString
// ==========================================

func Test_Dynamic_ValueString_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueStringTestCases {
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsString("value")

		var actLines []string

		switch {
		case value == "nil":
			d := coredynamic.NewDynamic(nil, true)
			actLines = []string{fmt.Sprintf("%v", d.ValueString() == "")}
		case value == "hello":
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{d.ValueString()}
		default:
			d := coredynamic.NewDynamicValid(input["value"])
			actLines = []string{fmt.Sprintf("%v", d.ValueString() != "")}
		}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ValueStrings
// ==========================================

func Test_Dynamic_ValueStrings_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueStringsTestCases {
		input := tc.ArrangeInput.(args.Map)
		value, isStr := input.GetAsString("value")

		var actLines []string

		if isStr && value == "strings_ab" {
			d := coredynamic.NewDynamicValid([]string{"a", "b"})
			result := d.ValueStrings()
			actLines = result
		} else {
			d := coredynamic.NewDynamicValid(input["value"])
			actLines = []string{fmt.Sprintf("%v", d.ValueStrings() == nil)}
		}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ValueUInt
// ==========================================

func Test_Dynamic_ValueUInt_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueUIntTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"]
		d := coredynamic.NewDynamicValid(value)

		actLines := []string{fmt.Sprintf("%d", d.ValueUInt())}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ValueInt64
// ==========================================

func Test_Dynamic_ValueInt64_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueInt64TestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"]
		d := coredynamic.NewDynamicValid(value)

		actLines := []string{fmt.Sprintf("%d", d.ValueInt64())}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Bytes
// ==========================================

func Test_Dynamic_Bytes_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicBytesTestCases {
		input := tc.ArrangeInput.(args.Map)
		receiver, _ := input.GetAsString("receiver")
		value, _ := input.GetAsString("value")

		var actLines []string

		if receiver == "nil" {
			var d *coredynamic.Dynamic
			raw, ok := d.Bytes()
			actLines = []string{
				fmt.Sprintf("%v", raw == nil),
				fmt.Sprintf("%v", ok),
			}
		} else if value == "bytes_raw" {
			d := coredynamic.NewDynamicValid([]byte("raw"))
			raw, ok := d.Bytes()
			actLines = []string{
				fmt.Sprintf("%v", ok),
				string(raw),
			}
		} else {
			d := coredynamic.NewDynamicValid(value)
			_, ok := d.Bytes()
			actLines = []string{fmt.Sprintf("%v", ok)}
		}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IntDefault
// ==========================================

func Test_Dynamic_IntDefault_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicIntDefaultTestCases {
		input := tc.ArrangeInput.(args.Map)
		defaultVal := input.GetAsIntDefault("defaultValue", 0)

		var actLines []string

		valueStr, isStr := input.GetAsString("value")
		if isStr && valueStr == "nil" {
			d := coredynamic.NewDynamic(nil, true)
			val, ok := d.IntDefault(defaultVal)
			actLines = []string{
				fmt.Sprintf("%v", ok),
				fmt.Sprintf("%d", val),
			}
		} else {
			value := input["value"]
			d := coredynamic.NewDynamicValid(value)
			val, ok := d.IntDefault(defaultVal)
			actLines = []string{
				fmt.Sprintf("%v", ok),
				fmt.Sprintf("%d", val),
			}
		}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ValueNullErr
// ==========================================

func Test_Dynamic_ValueNullErr_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueNullErrTestCases {
		input := tc.ArrangeInput.(args.Map)
		receiver, _ := input.GetAsString("receiver")
		value, _ := input.GetAsString("value")

		var actLines []string

		if receiver == "nil" {
			var d *coredynamic.Dynamic
			actLines = []string{fmt.Sprintf("%v", d.ValueNullErr() != nil)}
		} else if value == "nil" {
			d := coredynamic.NewDynamic(nil, true)
			actLines = []string{fmt.Sprintf("%v", d.ValueNullErr() != nil)}
		} else {
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{fmt.Sprintf("%v", d.ValueNullErr() != nil)}
		}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Reflect
// ==========================================

func Test_Dynamic_Reflect_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicReflectTestCases {
		input := tc.ArrangeInput.(args.Map)
		check, _ := input.GetAsString("check")

		var actLines []string

		switch check {
		case "ReflectKind":
			value := input["value"]
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{d.ReflectKind().String()}
		case "IsReflectKind":
			value, _ := input.GetAsString("value")
			d := coredynamic.NewDynamicValid(value)
			match := input["match"].(bool)
			if match {
				actLines = []string{fmt.Sprintf("%v", d.IsReflectKind(reflect.String))}
			} else {
				actLines = []string{fmt.Sprintf("%v", d.IsReflectKind(reflect.Int))}
			}
		case "ReflectTypeName":
			value, _ := input.GetAsString("value")
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{fmt.Sprintf("%v", d.ReflectTypeName() != "")}
		case "ReflectType":
			value := input["value"]
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{fmt.Sprintf("%v", d.ReflectType() == reflect.TypeOf(value))}
		case "IsReflectTypeOf":
			value, _ := input.GetAsString("value")
			d := coredynamic.NewDynamicValid(value)
			actLines = []string{fmt.Sprintf("%v", d.IsReflectTypeOf(reflect.TypeOf("")))}
		default:
			errcore.HandleErrMessage("unknown check: " + check)
		}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ReflectValue (cached)
// ==========================================

func Test_Dynamic_ReflectValue_Verification(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	rv1 := d.ReflectValue()
	rv2 := d.ReflectValue()

	actLines := []string{
		fmt.Sprintf("%v", rv1 == rv2),
		fmt.Sprintf("%d", rv1.Int()),
	}
	expected := []string{"true", "42"}

	errcore.PrintLineDiff(0, "ReflectValue returns cached reflect.Value", actLines, expected)

	tc := coretestcases.CaseV1{
		Title:         "ReflectValue returns cached reflect.Value",
		ExpectedInput: expected,
	}
	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================
// Test: Loop
// ==========================================

func Test_Dynamic_Loop_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicLoopTestCases {
		input := tc.ArrangeInput.(args.Map)
		scenario, _ := input.GetAsString("scenario")

		var actLines []string

		switch scenario {
		case "iterate":
			d := coredynamic.NewDynamicValid([]string{"a", "b", "c"})
			collected := make([]string, 0, 3)
			called := d.Loop(func(index int, item any) bool {
				collected = append(collected, item.(string))
				return false
			})
			actLines = append([]string{fmt.Sprintf("%v", called)}, collected...)
		case "invalid":
			d := coredynamic.InvalidDynamicPtr()
			called := d.Loop(func(index int, item any) bool { return false })
			actLines = []string{fmt.Sprintf("%v", called)}
		case "break":
			d := coredynamic.NewDynamicValid([]int{1, 2, 3, 4})
			count := 0
			d.Loop(func(index int, item any) bool {
				count++
				return index == 1
			})
			actLines = []string{fmt.Sprintf("%d", count)}
		default:
			errcore.HandleErrMessage("unknown scenario: " + scenario)
		}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Item Access
// ==========================================

func Test_Dynamic_ItemAccess_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicItemAccessTestCases {
		input := tc.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")

		var actLines []string

		switch method {
		case "ItemUsingIndex":
			d := coredynamic.NewDynamicValid([]string{"a", "b"})
			actLines = []string{
				fmt.Sprintf("%v", d.ItemUsingIndex(0)),
				fmt.Sprintf("%v", d.ItemUsingIndex(1)),
			}
		case "ItemUsingKey":
			d := coredynamic.NewDynamicValid(map[string]int{"k": 42})
			actLines = []string{fmt.Sprintf("%v", d.ItemUsingKey("k"))}
		default:
			errcore.HandleErrMessage("unknown method: " + method)
		}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IsStructStringNullOrEmpty
// ==========================================

func Test_Dynamic_StructStringNullOrEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicStructStringNullOrEmptyTestCases {
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsString("value")

		var d coredynamic.Dynamic
		if value == "nil" {
			d = coredynamic.NewDynamic(nil, true)
		} else {
			d = coredynamic.NewDynamicValid(value)
		}

		actLines := []string{fmt.Sprintf("%v", d.IsStructStringNullOrEmpty())}

		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, tc.ExpectedInput)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
