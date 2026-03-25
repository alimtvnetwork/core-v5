package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// =============================================================================
// New.Default
// =============================================================================

func Test_NewCreator_Default_Verification(t *testing.T) {
	for caseIndex, testCase := range newDefaultTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		regex, err := regexnew.New.Default(pattern)

		actual := args.Map{
			params.regexNotNil: regex != nil,
			params.hasError:    err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// New.DefaultLockIf
// =============================================================================

func Test_NewCreator_DefaultLockIf_Verification(t *testing.T) {
	for caseIndex, testCase := range newDefaultLockIfTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		isLock, _ := input.GetAsBool(params.isLock)

		// Act
		regex, err := regexnew.New.DefaultLockIf(isLock, pattern)

		actual := args.Map{
			params.regexNotNil: regex != nil,
			params.hasError:    err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// New.DefaultApplicableLock
// =============================================================================

func Test_NewCreator_DefaultApplicableLock_Verification(t *testing.T) {
	for caseIndex, testCase := range newDefaultApplicableLockTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		regex, err, isApplicable := regexnew.New.DefaultApplicableLock(pattern)

		actual := args.Map{
			params.regexNotNil:  regex != nil,
			params.hasError:     err != nil,
			params.isApplicable: isApplicable,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// New.LazyRegex.NewLockIf
// =============================================================================

func Test_NewLazyRegexCreator_NewLockIf_Verification(t *testing.T) {
	for caseIndex, testCase := range newLazyRegexNewLockIfTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)
		isLock, _ := input.GetAsBool(params.isLock)

		// Act
		lazy := regexnew.New.LazyRegex.NewLockIf(isLock, pattern)

		actual := args.Map{
			params.isDefined:    lazy.IsDefined(),
			params.isApplicable: lazy.IsApplicable(),
			params.isMatch:      lazy.IsMatch(compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// New.LazyRegex.AllPatternsMap
// =============================================================================

func Test_NewLazyRegexCreator_AllPatternsMap_Verification(t *testing.T) {
	for caseIndex, testCase := range allPatternsMapTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// ensure at least one pattern exists in the map
		regexnew.New.LazyLock(pattern)

		// Act
		allPatterns := regexnew.New.LazyRegex.AllPatternsMap()

		actual := args.Map{
			params.isNotEmpty: len(allPatterns) > 0,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
