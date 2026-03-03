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
		_ = input["constructorRef"] // ensures ref key exists

		var actLines []string

		if input.Has("constructorRef") {
			ref := input["constructorRef"]

			switch fn := ref.(type) {
			case func(any) coredynamic.Dynamic:
				// refNewDynamicValid
				d := fn(input["inputData"])
				actLines = []string{
					fmt.Sprintf("%v", d.IsValid()),
					fmt.Sprintf("%v", d.Value()),
				}
			case func(any, bool) coredynamic.Dynamic:
				// refNewDynamic
				isValid := input["isValid"].(bool)
				d := fn(nil, isValid)
				actLines = []string{
					fmt.Sprintf("%v", d.IsValid()),
					fmt.Sprintf("%v", d.IsInvalid()),
				}
			case func() coredynamic.Dynamic:
				// refInvalidDynamic
				d := fn()
				actLines = []string{
					fmt.Sprintf("%v", d.IsValid()),
					fmt.Sprintf("%v", d.IsNull()),
				}
			case func() *coredynamic.Dynamic:
				// refInvalidDynamicPtr
				d := fn()
				actLines = []string{
					fmt.Sprintf("%v", d != nil),
					fmt.Sprintf("%v", d.IsValid()),
					fmt.Sprintf("%v", d.IsNull()),
				}
			case func(any, bool) *coredynamic.Dynamic:
				// refNewDynamicPtr
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

func Test_Dynamic_Clone_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicCloneTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		scenario, _ := input.GetAsString("scenario")
		nilReceiver := getBoolDefault(input, "nilReceiver")

		var actLines []string

		switch scenario {
		case "clone":
			inputData, _ := input.GetAsString("inputData")
			original := refNewDynamicValid(inputData)
			cloned := original.Clone()
			actLines = []string{
				fmt.Sprintf("%v", cloned.Value()),
				fmt.Sprintf("%v", cloned.IsValid()),
			}
		case "clonePtr":
			if nilReceiver {
				var d *coredynamic.Dynamic
				actLines = []string{fmt.Sprintf("%v", d.ClonePtr() == nil)}
			} else {
				inputData, _ := input.GetAsString("inputData")
				original := refNewDynamicPtr(inputData, true)
				cloned := original.ClonePtr()
				actLines = []string{
					fmt.Sprintf("%v", cloned != nil),
					fmt.Sprintf("%v", cloned.Value()),
				}
			}
		case "nonPtr":
			inputData, _ := input.GetAsString("inputData")
			d := refNewDynamicValid(inputData)
			actLines = []string{fmt.Sprintf("%v", d.NonPtr().Value())}
		default:
			errcore.HandleErrMessage("unknown clone scenario: " + scenario)
		}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: Type Checks
//
// Uses DynamicBoolMethodRef stored in "checkRef" key.
// The method reference is compile-time safe — renaming the
// method on *Dynamic causes a build error.
// ==========================================

func Test_Dynamic_TypeChecks_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicTypeCheckTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)

		var actLines []string

		scenario, hasScenario := input.GetAsString("scenario")

		if hasScenario {
			// Special cases that need custom setup
			actLines = typeCheckSpecialScenario(scenario, input)
		} else {
			// Standard bool method ref check
			checkRef := input["checkRef"].(DynamicBoolMethodRef)
			d := createDynamicFromInput(input)
			actLines = []string{fmt.Sprintf("%v", checkRef.Call(d))}
		}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// typeCheckSpecialScenario handles type check cases that need
// custom Dynamic creation (e.g., pointer wrapping, Data==Value).
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
		// Create a pointer value to verify IsPointer
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
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		checkRef := input["checkRef"].(DynamicBoolMethodRef)
		scenario, hasScenario := input.GetAsString("scenario")

		var d coredynamic.Dynamic
		if hasScenario && scenario == "struct" {
			d = refNewDynamicValid(sample{Name: "test"})
		} else {
			d = refNewDynamicValid(input["inputData"])
		}

		// Act
		actLines := []string{fmt.Sprintf("%v", checkRef.Call(&d))}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: Length
// ==========================================

func Test_Dynamic_Length_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicLengthTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		d := createDynamicFromInput(input)

		// Act
		actLines := []string{fmt.Sprintf("%d", d.Length())}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueInt
// ==========================================

func Test_Dynamic_ValueInt_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueIntTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		d := refNewDynamicValid(input["inputData"])

		// Act
		actLines := []string{fmt.Sprintf("%d", d.ValueInt())}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueBool
// ==========================================

func Test_Dynamic_ValueBool_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueBoolTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		d := refNewDynamicValid(input["inputData"])

		// Act
		actLines := []string{fmt.Sprintf("%v", d.ValueBool())}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueString
// ==========================================

func Test_Dynamic_ValueString_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		d := createDynamicFromInput(input)

		// Act
		var actLines []string
		result := d.ValueString()

		inputData := input["inputData"]
		_, isString := inputData.(string)

		switch {
		case inputData == nil:
			// nil data → check that result is empty
			actLines = []string{fmt.Sprintf("%v", result == "")}
		case isString && inputData.(string) == "hello":
			// string data → return exact value
			actLines = []string{result}
		default:
			// non-string data → check that result is non-empty
			actLines = []string{fmt.Sprintf("%v", result != "")}
		}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueStrings
// ==========================================

func Test_Dynamic_ValueStrings_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueStringsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		d := refNewDynamicValid(input["inputData"])

		// Act
		var actLines []string
		result := d.ValueStrings()

		if result != nil {
			actLines = result
		} else {
			// nil result → report as "true" (ValueStrings() == nil)
			actLines = []string{fmt.Sprintf("%v", result == nil)}
		}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueUInt
// ==========================================

func Test_Dynamic_ValueUInt_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueUIntTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		d := refNewDynamicValid(input["inputData"])

		// Act
		actLines := []string{fmt.Sprintf("%d", d.ValueUInt())}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueInt64
// ==========================================

func Test_Dynamic_ValueInt64_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueInt64TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		d := refNewDynamicValid(input["inputData"])

		// Act
		actLines := []string{fmt.Sprintf("%d", d.ValueInt64())}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: Bytes
// ==========================================

func Test_Dynamic_Bytes_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicBytesTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		nilReceiver := getBoolDefault(input, "nilReceiver")

		var actLines []string

		if nilReceiver {
			// Test nil *Dynamic receiver safety
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

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: IntDefault
// ==========================================

func Test_Dynamic_IntDefault_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicIntDefaultTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		defaultVal := input.GetAsIntDefault("defaultValue", 0)
		d := createDynamicFromInput(input)

		// Act
		val, ok := d.IntDefault(defaultVal)
		actLines := []string{
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", val),
		}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ValueNullErr
// ==========================================

func Test_Dynamic_ValueNullErr_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueNullErrTestCases {
		// Arrange
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

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: Reflect
// ==========================================

func Test_Dynamic_Reflect_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicReflectTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		scenario, _ := input.GetAsString("scenario")
		d := refNewDynamicValid(input["inputData"])

		var actLines []string

		switch scenario {
		case "reflectKind":
			actLines = []string{d.ReflectKind().String()}
		case "isReflectKindMatch":
			// Checks that string data matches reflect.String
			actLines = []string{fmt.Sprintf("%v", d.IsReflectKind(reflect.String))}
		case "isReflectKindMismatch":
			// Checks that string data does NOT match reflect.Int
			actLines = []string{fmt.Sprintf("%v", d.IsReflectKind(reflect.Int))}
		case "reflectTypeName":
			actLines = []string{fmt.Sprintf("%v", d.ReflectTypeName() != "")}
		case "reflectType":
			actLines = []string{fmt.Sprintf("%v", d.ReflectType() == reflect.TypeOf(input["inputData"]))}
		case "isReflectTypeOf":
			actLines = []string{fmt.Sprintf("%v", d.IsReflectTypeOf(reflect.TypeOf("")))}
		default:
			errcore.HandleErrMessage("unknown reflect scenario: " + scenario)
		}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: ReflectValue (cached)
// ==========================================

func Test_Dynamic_ReflectValue_Verification(t *testing.T) {
	// Arrange
	d := refNewDynamicPtr(42, true)

	// Act
	rv1 := d.ReflectValue()
	rv2 := d.ReflectValue()

	actLines := []string{
		fmt.Sprintf("%v", rv1 == rv2),
		fmt.Sprintf("%d", rv1.Int()),
	}
	expected := []string{"true", "42"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "ReflectValue returns cached reflect.Value", actLines, expected)
}

// ==========================================
// Test: Loop
// ==========================================

func Test_Dynamic_Loop_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicLoopTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		scenario, _ := input.GetAsString("scenario")

		var actLines []string

		switch scenario {
		case "iterate":
			d := refNewDynamicValid([]string{"a", "b", "c"})
			collected := make([]string, 0, 3)
			called := d.Loop(func(index int, item any) bool {
				collected = append(collected, item.(string))
				return false
			})
			actLines = append([]string{fmt.Sprintf("%v", called)}, collected...)
		case "invalid":
			d := refInvalidDynamicPtr()
			called := d.Loop(func(index int, item any) bool { return false })
			actLines = []string{fmt.Sprintf("%v", called)}
		case "break":
			d := refNewDynamicValid([]int{1, 2, 3, 4})
			count := 0
			d.Loop(func(index int, item any) bool {
				count++
				return index == 1 // break after second item
			})
			actLines = []string{fmt.Sprintf("%d", count)}
		default:
			errcore.HandleErrMessage("unknown loop scenario: " + scenario)
		}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: Item Access
// ==========================================

func Test_Dynamic_ItemAccess_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicItemAccessTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		scenario, _ := input.GetAsString("scenario")

		var actLines []string

		switch scenario {
		case "itemUsingIndex":
			d := refNewDynamicValid([]string{"a", "b"})
			actLines = []string{
				fmt.Sprintf("%v", d.ItemUsingIndex(0)),
				fmt.Sprintf("%v", d.ItemUsingIndex(1)),
			}
		case "itemUsingKey":
			d := refNewDynamicValid(map[string]int{"k": 42})
			actLines = []string{fmt.Sprintf("%v", d.ItemUsingKey("k"))}
		default:
			errcore.HandleErrMessage("unknown itemAccess scenario: " + scenario)
		}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================
// Test: IsStructStringNullOrEmpty
// ==========================================

func Test_Dynamic_StructStringNullOrEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicStructStringNullOrEmptyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		checkRef := input["checkRef"].(DynamicBoolMethodRef)
		d := createDynamicFromInput(input)

		// Act
		actLines := []string{fmt.Sprintf("%v", checkRef.Call(d))}

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// =============================================================================
// Helpers
// =============================================================================

// createDynamicFromInput creates a *Dynamic from the standard input keys.
//   - "inputData": the actual Go value to wrap
//   - "isValid": if present, uses NewDynamic(data, isValid); otherwise NewDynamicValid(data)
//
// Returns a pointer to allow nil-safe method calls.
func createDynamicFromInput(input args.Map) *coredynamic.Dynamic {
	inputData := input["inputData"]

	if isValid, has := input["isValid"]; has {
		d := refNewDynamic(inputData, isValid.(bool))
		return &d
	}

	d := refNewDynamicValid(inputData)
	return &d
}

// getBoolDefault returns the bool value for the given key, or false if missing.
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
