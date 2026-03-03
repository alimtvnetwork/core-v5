package regexnewtests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/regexnew"
)

func Test_CreateMust_Verification(t *testing.T) {
	for caseIndex, testCase := range createMustTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		regex := regexnew.CreateMust(pattern)
		isNotNil := fmt.Sprintf("%v", regex != nil)
		isMatch := fmt.Sprintf("%v", regex.MatchString(compareInput))

		actLines := []string{isNotNil, isMatch}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_CreateMust_PanicsOnInvalidPattern(t *testing.T) {
	// Arrange
	pattern := "[invalid"

	// Act & Assert
	defer func() {
		r := recover()
		if r == nil {
			t.Error("CreateMust should panic on invalid pattern")
		}
	}()

	regexnew.CreateMust(pattern)
}

func Test_CreateMustLockIf_Verification(t *testing.T) {
	for caseIndex, testCase := range createMustLockIfTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")
		isLockRaw := testCase.GetExtra("isLock")
		isLock := isLockRaw == true

		// Act
		regex := regexnew.CreateMustLockIf(isLock, pattern)
		isNotNil := fmt.Sprintf("%v", regex != nil)
		isMatch := fmt.Sprintf("%v", regex.MatchString(compareInput))

		actLines := []string{isNotNil, isMatch}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_CreateMustLockIf_PanicsOnInvalidPattern(t *testing.T) {
	// Arrange
	pattern := "[invalid"

	// Act & Assert — with lock
	defer func() {
		r := recover()
		if r == nil {
			t.Error("CreateMustLockIf should panic on invalid pattern")
		}
	}()

	regexnew.CreateMustLockIf(true, pattern)
}

func Test_CreateLockIf_Verification(t *testing.T) {
	for caseIndex, testCase := range createLockIfTestCases {
		// Arrange
		pattern := testCase.Input
		isLockRaw := testCase.GetExtra("isLock")
		isLock := isLockRaw == true

		// Act
		regex, err := regexnew.CreateLockIf(isLock, pattern)
		isNotNil := fmt.Sprintf("%v", regex != nil)
		hasError := fmt.Sprintf("%v", err != nil)

		actLines := []string{isNotNil, hasError}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_CreateApplicableLock_Verification(t *testing.T) {
	for caseIndex, testCase := range createApplicableLockTestCases {
		// Arrange
		pattern := testCase.Input

		// Act
		regex, err, isApplicable := regexnew.CreateApplicableLock(pattern)
		isNotNil := fmt.Sprintf("%v", regex != nil)
		hasError := fmt.Sprintf("%v", err != nil)
		applicable := fmt.Sprintf("%v", isApplicable)

		actLines := []string{isNotNil, hasError, applicable}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_NewMustLock_Verification(t *testing.T) {
	for caseIndex, testCase := range newMustLockTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		regex := regexnew.NewMustLock(pattern)
		isNotNil := fmt.Sprintf("%v", regex != nil)
		isMatch := fmt.Sprintf("%v", regex.MatchString(compareInput))

		actLines := []string{isNotNil, isMatch}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_NewMustLock_PanicsOnInvalidPattern(t *testing.T) {
	// Arrange
	pattern := "[invalid"

	// Act & Assert
	defer func() {
		r := recover()
		if r == nil {
			t.Error("NewMustLock should panic on invalid pattern")
		}
	}()

	regexnew.NewMustLock(pattern)
}

func Test_MatchUsingFuncErrorLock_Verification(t *testing.T) {
	matchFunc := regexnew.RegexValidationFunc(
		func(re *regexp.Regexp, term string) bool {
			return re.MatchString(term)
		},
	)

	for caseIndex, testCase := range matchUsingFuncErrorLockTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")

		// Act
		err := regexnew.MatchUsingFuncErrorLock(
			pattern,
			compareInput,
			matchFunc,
		)
		isNoError := fmt.Sprintf("%v", err == nil)

		actLines := []string{isNoError}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}

func Test_MatchUsingCustomizeErrorFuncLock_Verification(t *testing.T) {
	matchFunc := regexnew.RegexValidationFunc(
		func(re *regexp.Regexp, term string) bool {
			return re.MatchString(term)
		},
	)

	customErrFunc := regexnew.CustomizeErr(
		func(regexPattern, matchLookingTerm string, err error, re *regexp.Regexp) error {
			return fmt.Errorf("CUSTOM: pattern %s failed on %s", regexPattern, matchLookingTerm)
		},
	)

	for caseIndex, testCase := range matchUsingCustomizeErrorFuncLockTestCases {
		// Arrange
		pattern := testCase.Input
		compareInput, _ := testCase.GetExtraAsString("compareInput")
		customizer, _ := testCase.GetExtraAsString("customizer")

		// Act
		var errFunc regexnew.CustomizeErr
		if customizer == "custom" {
			errFunc = customErrFunc
		}

		err := regexnew.MatchUsingCustomizeErrorFuncLock(
			pattern,
			compareInput,
			matchFunc,
			errFunc,
		)

		isNoError := fmt.Sprintf("%v", err == nil)

		var actLines []string
		actLines = append(actLines, isNoError)

		if err != nil {
			isCustomError := fmt.Sprintf("%v", strings.HasPrefix(err.Error(), "CUSTOM:"))
			actLines = append(actLines, isCustomError)
		}

		// Assert
		testCase.ShouldBeEqualUsingExpected(t, caseIndex, actLines)
	}
}
