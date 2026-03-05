package converterstests

import (
	"fmt"
	"testing"
	"unsafe"

	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// =============================================================================
// Tests: IntegersWithDefaults
// =============================================================================

func Test_StringsTo_IntegersWithDefaults(t *testing.T) {
	for caseIndex, testCase := range integersWithDefaultsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputRaw, rawFound := input.Get("input")
		if !rawFound {
			errcore.HandleErrMessage("input is required")
		}
		inputSlice := inputRaw.([]string)
		defaultInt, defFound := input.GetAsInt("defaultInt")
		if !defFound {
			errcore.HandleErrMessage("defaultInt is required")
		}

		// Act
		result := converters.StringsTo.IntegersWithDefaults(defaultInt, inputSlice...)
		actLines := []string{fmt.Sprintf("%d", result.Length())}

		for _, v := range result.Values {
			actLines = append(actLines, fmt.Sprintf("%d", v))
		}

		actLines = append(actLines, fmt.Sprintf("%v", result.HasError()))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// =============================================================================
// Tests: BytesWithDefaults
// =============================================================================

func Test_StringsTo_BytesWithDefaults(t *testing.T) {
	for caseIndex, testCase := range bytesWithDefaultsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputRaw, rawFound := input.Get("input")
		if !rawFound {
			errcore.HandleErrMessage("input is required")
		}
		inputSlice := inputRaw.([]string)
		defaultByteRaw, defFound := input.Get("defaultByte")
		if !defFound {
			errcore.HandleErrMessage("defaultByte is required")
		}
		defaultByte := defaultByteRaw.(byte)

		// Act
		result := converters.StringsTo.BytesWithDefaults(defaultByte, inputSlice...)
		actLines := []string{fmt.Sprintf("%d", result.Length())}

		for _, v := range result.Values {
			actLines = append(actLines, fmt.Sprintf("%d", v))
		}

		actLines = append(actLines, fmt.Sprintf("%v", result.HasError()))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// =============================================================================
// Tests: CloneIf
// =============================================================================

func Test_StringsTo_CloneIf(t *testing.T) {
	for caseIndex, testCase := range cloneIfTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputRaw, rawFound := input.Get("input")
		if !rawFound {
			errcore.HandleErrMessage("input is required")
		}
		inputSlice := inputRaw.([]string)
		isCloneRaw, cloneFound := input.Get("isClone")
		if !cloneFound {
			errcore.HandleErrMessage("isClone is required")
		}
		isClone := isCloneRaw.(bool)

		// Act
		result := converters.StringsTo.CloneIf(isClone, inputSlice...)
		actLines := []string{fmt.Sprintf("%d", len(result))}

		for _, v := range result {
			actLines = append(actLines, v)
		}

		// Check if it's a different slice (clone independence)
		isSamePointer := false
		if len(inputSlice) > 0 && len(result) > 0 {
			isSamePointer = unsafe.Pointer(&inputSlice[0]) == unsafe.Pointer(&result[0])
		}

		// When isClone=true and len>0, should NOT be same pointer
		// When isClone=false, SHOULD be same pointer
		isIndependent := isClone && !isSamePointer && len(inputSlice) > 0
		isReturned := !isClone && isSamePointer
		actLines = append(actLines, fmt.Sprintf("%v", isIndependent || isReturned || len(inputSlice) == 0))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// =============================================================================
// Tests: PtrOfPtrToPtrStrings
// =============================================================================

func Test_StringsTo_PtrOfPtrToPtrStrings(t *testing.T) {
	for caseIndex, testCase := range ptrOfPtrToPtrStringsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)

		isNilRaw, _ := input.Get("isNil")
		isNil, _ := isNilRaw.(bool)
		isNilInnerRaw, _ := input.Get("isNilInner")
		isNilInner, _ := isNilInnerRaw.(bool)

		var result *[]string

		if isNil {
			// Act — nil outer pointer
			result = converters.StringsTo.PtrOfPtrToPtrStrings(nil)
		} else if isNilInner {
			// Act — nil inner pointer
			var nilInner []*string
			result = converters.StringsTo.PtrOfPtrToPtrStrings(&nilInner)
		} else {
			inputRaw, _ := input.Get("input")
			inputSlice := inputRaw.([]string)
			hasNilRaw, _ := input.Get("hasNil")
			hasNil, _ := hasNilRaw.(bool)

			// Build []*string
			ptrSlice := make([]*string, 0, len(inputSlice)+1)
			for i := range inputSlice {
				ptrSlice = append(ptrSlice, &inputSlice[i])
			}

			if hasNil {
				ptrSlice = append(ptrSlice, nil)
			}

			// Act
			result = converters.StringsTo.PtrOfPtrToPtrStrings(&ptrSlice)
		}

		actLines := []string{fmt.Sprintf("%d", len(*result))}
		for _, v := range *result {
			actLines = append(actLines, v)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// =============================================================================
// Tests: PtrOfPtrToMapStringBool
// =============================================================================

func Test_StringsTo_PtrOfPtrToMapStringBool(t *testing.T) {
	for caseIndex, testCase := range ptrOfPtrToMapStringBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilRaw, _ := input.Get("isNil")
		isNil, _ := isNilRaw.(bool)

		var result map[string]bool

		if isNil {
			// Act — nil input
			result = converters.StringsTo.PtrOfPtrToMapStringBool(nil)
		} else {
			inputRaw, _ := input.Get("input")
			inputSlice := inputRaw.([]string)
			hasNilRaw, _ := input.Get("hasNil")
			hasNil, _ := hasNilRaw.(bool)

			ptrSlice := make([]*string, 0, len(inputSlice)+1)
			for i := range inputSlice {
				ptrSlice = append(ptrSlice, &inputSlice[i])
			}

			if hasNil {
				ptrSlice = append(ptrSlice, nil)
			}

			// Act
			result = converters.StringsTo.PtrOfPtrToMapStringBool(&ptrSlice)
		}

		actLines := []string{fmt.Sprintf("%d", len(result))}
		for _, v := range result {
			actLines = append(actLines, fmt.Sprintf("%v", v))
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
