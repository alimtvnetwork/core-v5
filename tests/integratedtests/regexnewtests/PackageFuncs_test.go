package regexnewtests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// =============================================================================
// CreateApplicableLock
// =============================================================================

func Test_CreateApplicableLock(t *testing.T) {
	for caseIndex, tc := range createApplicableLockTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")

		// Act
		regex, err, isApplicable := regexnew.CreateApplicableLock(pattern)

		actual := args.Map{
			"regexNotNil":  regex != nil,
			"hasError":     err != nil,
			"isApplicable": isApplicable,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// CreateLockIf
// =============================================================================

func Test_CreateLockIf(t *testing.T) {
	for caseIndex, tc := range createLockIfTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isLock, _ := input.GetAsBool("isLock")
		pattern, _ := input.GetAsString("pattern")

		// Act
		regex, err := regexnew.CreateLockIf(isLock, pattern)

		actual := args.Map{
			"regexNotNil": regex != nil,
			"hasError":    err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// CreateMustLockIf
// =============================================================================

func Test_CreateMustLockIf(t *testing.T) {
	for caseIndex, tc := range createMustLockIfTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isLock, _ := input.GetAsBool("isLock")
		pattern, _ := input.GetAsString("pattern")

		var panicked bool
		var result *regexp.Regexp

		// Act
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()

			result = regexnew.CreateMustLockIf(isLock, pattern)
		}()

		actual := args.Map{
			"regexNotNil": result != nil,
			"panicked":    panicked,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// MatchUsingFuncErrorLock
// =============================================================================

func Test_MatchUsingFuncErrorLock(t *testing.T) {
	matchFunc := func(regex *regexp.Regexp, lookingTerm string) bool {
		return regex.MatchString(lookingTerm)
	}

	for caseIndex, tc := range matchUsingFuncErrorLockTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		comparing, _ := input.GetAsString("comparing")

		// Act
		err := regexnew.MatchUsingFuncErrorLock(pattern, comparing, matchFunc)

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// MatchUsingCustomizeErrorFuncLock
// =============================================================================

func Test_MatchUsingCustomizeErrorFuncLock(t *testing.T) {
	matchFunc := func(regex *regexp.Regexp, lookingTerm string) bool {
		return regex.MatchString(lookingTerm)
	}

	customErrFunc := func(
		regexPattern, matchLookingTerm string,
		err error,
		_ *regexp.Regexp,
	) error {
		return fmt.Errorf("CUSTOM: pattern=%s input=%s", regexPattern, matchLookingTerm)
	}

	for caseIndex, tc := range matchUsingCustomizeErrorFuncLockTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		comparing, _ := input.GetAsString("comparing")
		useCustomizer, _ := input.GetAsBool("useCustomizer")

		// Act
		var errResult error
		if useCustomizer {
			errResult = regexnew.MatchUsingCustomizeErrorFuncLock(
				pattern, comparing, matchFunc, customErrFunc,
			)
		} else {
			errResult = regexnew.MatchUsingCustomizeErrorFuncLock(
				pattern, comparing, matchFunc, nil,
			)
		}

		actual := args.Map{
			"hasError": errResult != nil,
		}

		if useCustomizer && errResult != nil {
			actual["isCustomError"] = strings.HasPrefix(errResult.Error(), "CUSTOM:")
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
