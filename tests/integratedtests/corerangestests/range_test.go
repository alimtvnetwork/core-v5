package corerangestests

import (
	"fmt"
	"strconv"
	"testing"
	
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coredata/corerange"
)

func Test_RangeInt_Valid_CheckWithInRange_Verification(t *testing.T) {
	// Arrange
	validCases := []int{
		5, 13, 5, 10, 25,
	}
	toString := converters.AnyToValueString(validCases)
	
	// Act, Assert
	title := toString + " -- all these are valid for (range) : " + someRange.String()
	convey.Convey(title, t, func() {
		for _, v := range validCases {
			isInRange := someRange.IsValidPlusWithinRange(v)
			
			if !isInRange {
				fmt.Println("Should be valid but invalid for : " + strconv.Itoa(v))
			}
			
			convey.So(
				isInRange,
				should.BeTrue)
		}
	})
}

func Test_RangeInt_Invalid_CheckWithInRange_Verification(t *testing.T) {
	// Arrange
	invalidCases := []int{
		265, 311, 4, 26, 100,
	}
	toString := converters.AnyToValueString(invalidCases)
	
	// Act, Assert
	title := toString + " -- all these are valid for (range) : " + someRange.String()
	convey.Convey(title, t, func() {
		for _, v := range invalidCases {
			isInRange := someRange.IsValidPlusWithinRange(v)
			
			if isInRange {
				fmt.Println("Should be Invalid but valid for : " + strconv.Itoa(v))
			}
			
			convey.So(
				isInRange,
				should.BeFalse)
		}
	})
}

func Test_Int_Ranges_ValidCases(t *testing.T) {
	for _, testCase := range validIntRangeTestCases {
		// Arrange
		arrangeInputs := testCase.Arrange()
		first := arrangeInputs[0]
		rest := arrangeInputs[1:]
		
		// Act
		actualRanges := first.CreateRanges(rest...)
		
		// Assert
		convey.Convey(testCase.Title, t, func() {
			convey.So(
				actualRanges,
				should.Equal,
				testCase.ExpectedInput)
		})
		
		convey.Convey(testCase.Title+" - type verify", t, func() {
			convey.So(
				testCase.TypeValidationError(),
				should.BeNil)
		})
	}
}

func Test_Int_ExceptRanges_Verify(t *testing.T) {
	// Arrange
	arrangeInput := corerange.MinMaxInt{
		Min: 1,
		Max: 15,
	}
	
	// Act
	actualRanges := arrangeInput.RangesExcept(
		3, 4, 5)
	
	// Assert
	convey.Convey("Ranges 1-15, RangesExcept(3, 4, 5), should not contain 3,4,5", t, func() {
		convey.So(actualRanges, should.Equal, []int{
			1, 2, 6,
			7, 8, 9,
			10, 11,
			12, 13,
			14, 15,
		})
	})
}

func Test_Int8_Ranges_ValidCases(t *testing.T) {
	for _, testCase := range validInt8RangeTestCases {
		// Arrange
		arrangeInputs := testCase.ArrangeInput.([]corerange.MinMaxInt8)
		first := arrangeInputs[0]
		rest := arrangeInputs[1:]
		
		// Act
		actualRanges := first.CreateRanges(rest...)
		
		// Assert
		convey.Convey(testCase.Title, t, func() {
			convey.So(
				actualRanges,
				should.Equal,
				testCase.ExpectedInput)
		})
		
		convey.Convey(testCase.Title+" - type verify", t, func() {
			convey.So(
				testCase.TypeValidationError(),
				should.BeNil)
		})
	}
}
