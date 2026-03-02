package stringslicetests

import (
	"fmt"
	"testing"
	"unsafe"

	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// =============================================================================
// Tests: stringslice.CloneIf
// =============================================================================

func Test_StringSlice_CloneIf(t *testing.T) {
	for caseIndex, testCase := range cloneIfTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isCloneRaw, cloneErr := input.Get("isClone")
		errcore.HandleErrMessage("isClone", cloneErr)
		isClone := isCloneRaw.(bool)
		additionalCapRaw, capErr := input.Get("additionalCap")
		errcore.HandleErrMessage("additionalCap", capErr)
		additionalCap := additionalCapRaw.(int)
		isNilRaw, _ := input.Get("isNil")
		isNil, _ := isNilRaw.(bool)

		var inputSlice []string
		if !isNil {
			inputRaw, _ := input.Get("input")
			inputSlice = inputRaw.([]string)
		}

		// Act
		result := stringslice.CloneIf(isClone, additionalCap, inputSlice)

		actLines := []string{fmt.Sprintf("%d", len(result))}
		for _, v := range result {
			actLines = append(actLines, v)
		}

		// Check independence
		isSamePointer := false
		if len(inputSlice) > 0 && len(result) > 0 {
			isSamePointer = unsafe.Pointer(&inputSlice[0]) == unsafe.Pointer(&result[0])
		}
		isIndependent := isClone && !isSamePointer
		isReturned := !isClone && (isSamePointer || len(inputSlice) == 0)
		actLines = append(actLines, fmt.Sprintf("%v", isIndependent || isReturned))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// =============================================================================
// Tests: stringslice.AnyItemsCloneIf
// =============================================================================

func Test_StringSlice_AnyItemsCloneIf(t *testing.T) {
	for caseIndex, testCase := range anyItemsCloneIfTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isCloneRaw, cloneErr := input.Get("isClone")
		errcore.HandleErrMessage("isClone", cloneErr)
		isClone := isCloneRaw.(bool)
		additionalCapRaw, capErr := input.Get("additionalCap")
		errcore.HandleErrMessage("additionalCap", capErr)
		additionalCap := additionalCapRaw.(int)
		inputRaw, inputErr := input.Get("input")
		errcore.HandleErrMessage("input", inputErr)
		inputSlice := inputRaw.([]any)

		// Act
		result := stringslice.AnyItemsCloneIf(isClone, additionalCap, inputSlice)

		actLines := []string{fmt.Sprintf("%d", len(result))}
		for _, v := range result {
			actLines = append(actLines, fmt.Sprintf("%v", v))
		}

		// Check independence
		isSamePointer := false
		if len(inputSlice) > 0 && len(result) > 0 {
			isSamePointer = unsafe.Pointer(&inputSlice[0]) == unsafe.Pointer(&result[0])
		}
		isIndependent := isClone && !isSamePointer
		isReturned := !isClone && isSamePointer
		actLines = append(actLines, fmt.Sprintf("%v", isIndependent || isReturned))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
