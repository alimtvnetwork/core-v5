package codefuncstests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/corefuncs"
	"gitlab.com/auk-go/core/coretests/args"
)

func sampleFunc() {}

// =============================================================================
// GetFuncName
// =============================================================================

func Test_GetFuncName_Verification(t *testing.T) {
	for caseIndex, tc := range getFuncNameTestCases {
		// Act
		name := corefuncs.GetFuncName(sampleFunc)
		isNotEmpty := fmt.Sprintf("%v", name != "")

		// Assert
		tc.ShouldBeEqual(t, caseIndex, isNotEmpty)
	}
}

// =============================================================================
// GetFuncFullName
// =============================================================================

func Test_GetFuncFullName_Verification(t *testing.T) {
	for caseIndex, tc := range getFuncFullNameTestCases {
		// Act
		fullName := corefuncs.GetFuncFullName(sampleFunc)
		actual := args.Map{
			"isNotEmpty":      fullName != "",
			"containsPackage": strings.Contains(fullName, "codefuncstests"),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// GetFunc
// =============================================================================

func Test_GetFunc_Verification(t *testing.T) {
	for caseIndex, tc := range getFuncTestCases {
		// Act
		f := corefuncs.GetFunc(sampleFunc)
		actual := args.Map{
			"isNotNil": f != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// newCreator — factory methods
// =============================================================================

func Test_NewCreator_Verification(t *testing.T) {
	for caseIndex, tc := range newCreatorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")
		name, _ := input.GetAsString("name")

		var actual args.Map

		// Act
		switch method {
		case "ActionErr":
			w := corefuncs.New.ActionErr(name, func() error { return nil })
			err := w.Exec()
			actual = args.Map{
				"hasError": err != nil,
			}
		case "IsSuccess":
			w := corefuncs.New.IsSuccess(name, func() bool { return true })
			actual = args.Map{
				"result": w.Exec(),
			}
		case "NamedAction":
			tracker := &namedActionTracker{}
			w := corefuncs.New.NamedAction(name, tracker.Action)
			w.Exec()
			actual = args.Map{
				"calledWith": tracker.CalledWith,
			}
		case "LegacyInOutErr":
			w := corefuncs.New.LegacyInOutErr(name, func(in any) (any, error) {
				return "processed", nil
			})
			output, err := w.Exec("input")
			actual = args.Map{
				"output":   output.(string),
				"hasError": err != nil,
			}
		case "LegacyResultDelegating":
			w := corefuncs.New.LegacyResultDelegating(name, func(_ any) error {
				return nil
			})
			err := w.Exec("target")
			actual = args.Map{
				"hasError": err != nil,
			}
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
