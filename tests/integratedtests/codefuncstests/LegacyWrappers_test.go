package codefuncstests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/corefuncs"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// IsSuccessFuncWrapper — Exec
// =============================================================================

func Test_IsSuccessFuncWrapper_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range isSuccessExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		actionResult, _ := input.GetAsBool("actionResult")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.IsSuccessFuncWrapper{
			Name:   name,
			Action: func() bool { return actionResult },
		}

		// Act
		result := wrapper.Exec()
		actual := args.Map{
			"result": result,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// IsSuccessFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

func Test_IsSuccessFuncWrapper_AsActionReturnsErrorFunc_Verification(t *testing.T) {
	for caseIndex, tc := range isSuccessAsActionReturnsErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		actionResult, _ := input.GetAsBool("actionResult")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.IsSuccessFuncWrapper{
			Name:   name,
			Action: func() bool { return actionResult },
		}

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc()
		err := errFunc()
		actual := args.Map{
			"hasError": err != nil,
		}
		if err != nil {
			actual["containsName"] = strings.Contains(err.Error(), name)
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// NamedActionFuncWrapper — Exec
// =============================================================================

func Test_NamedActionFuncWrapper_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range namedActionExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		tracker := &namedActionTracker{}
		wrapper := corefuncs.NamedActionFuncWrapper{
			Name:   name,
			Action: tracker.Action,
		}

		// Act
		wrapper.Exec()
		actual := args.Map{
			"calledWith": tracker.CalledWith,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// NamedActionFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

func Test_NamedActionFuncWrapper_AsActionReturnsErrorFunc_Verification(t *testing.T) {
	for caseIndex, tc := range namedActionAsActionReturnsErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		tracker := &namedActionTracker{}
		wrapper := corefuncs.NamedActionFuncWrapper{
			Name:   name,
			Action: tracker.Action,
		}

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc()
		err := errFunc()
		actual := args.Map{
			"hasError":   err != nil,
			"calledWith": tracker.CalledWith,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// ActionReturnsErrorFuncWrapper — Exec
// =============================================================================

func Test_ActionReturnsErrorFuncWrapper_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range actionReturnsErrorExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		hasActionErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.ActionReturnsErrorFuncWrapper{
			Name: name,
			Action: func() error {
				if hasActionErr {
					return errTest
				}

				return nil
			},
		}

		// Act
		err := wrapper.Exec()
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// ActionReturnsErrorFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

func Test_ActionReturnsErrorFuncWrapper_AsActionReturnsErrorFunc_Verification(t *testing.T) {
	for caseIndex, tc := range actionReturnsErrorAsActionReturnsErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		hasActionErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.ActionReturnsErrorFuncWrapper{
			Name: name,
			Action: func() error {
				if hasActionErr {
					return errTest
				}

				return nil
			},
		}

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc()
		err := errFunc()
		actual := args.Map{
			"hasError": err != nil,
		}
		if err != nil {
			actual["containsName"] = strings.Contains(err.Error(), name)
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// InOutErrFuncWrapper — Exec
// =============================================================================

func Test_InOutErrFuncWrapper_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range inOutErrExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsString("input")
		hasActionErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.InOutErrFuncWrapper{
			Name: name,
			Action: func(in any) (any, error) {
				if hasActionErr {
					return nil, errTest
				}

				return strings.ToUpper(in.(string)), nil
			},
		}

		// Act
		output, err := wrapper.Exec(inputVal)
		actual := args.Map{
			"hasError": err != nil,
		}
		if output != nil {
			actual["output"] = output.(string)
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// InOutErrFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

func Test_InOutErrFuncWrapper_AsActionReturnsErrorFunc_Verification(t *testing.T) {
	for caseIndex, tc := range inOutErrAsActionReturnsErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsString("input")
		hasActionErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.InOutErrFuncWrapper{
			Name: name,
			Action: func(in any) (any, error) {
				if hasActionErr {
					return nil, errTest
				}

				return in, nil
			},
		}

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc(inputVal)
		err := errFunc()
		actual := args.Map{
			"hasError": err != nil,
		}
		if err != nil {
			actual["containsName"] = strings.Contains(err.Error(), name)
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// ResultDelegatingFuncWrapper — Exec
// =============================================================================

func Test_ResultDelegatingFuncWrapper_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range resultDelegatingExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		hasActionErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.ResultDelegatingFuncWrapper{
			Name: name,
			Action: func(target any) error {
				if hasActionErr {
					return errTest
				}

				return nil
			},
		}

		// Act
		err := wrapper.Exec("target-ptr")
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// ResultDelegatingFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

func Test_ResultDelegatingFuncWrapper_AsActionReturnsErrorFunc_Verification(t *testing.T) {
	for caseIndex, tc := range resultDelegatingAsActionReturnsErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		hasActionErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.ResultDelegatingFuncWrapper{
			Name: name,
			Action: func(target any) error {
				if hasActionErr {
					return errTest
				}

				return nil
			},
		}

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc("target-ptr")
		err := errFunc()
		actual := args.Map{
			"hasError": err != nil,
		}
		if err != nil {
			actual["containsName"] = strings.Contains(err.Error(), name)
		} else {
			actual["containsName"] = false
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
