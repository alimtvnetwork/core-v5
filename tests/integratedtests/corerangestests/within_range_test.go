package corerangestests

import (
	"fmt"
	"strconv"
	"testing"
	
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/converters"
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
