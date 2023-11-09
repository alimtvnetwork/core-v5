package corerangestests

import (
	"testing"
	
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/simplewrap"
)

func Test_CurlyIf_SkipQuoteOnPresent_Should_Only_Have_SingleDoubleQuotation_NotDuplicates(t *testing.T) {
	testCases := []string{
		"some-elem",
		"alim-elem",
		"\"has-quote\"",
		"",
		"\"",
		"\"first",
		"last\"",
		"'",
		"simple",
	}
	expectation := []string{
		"\"some-elem\"",
		"\"alim-elem\"",
		"\"has-quote\"",
		"\"\"",
		"\"\"",
		"\"first\"",
		"\"last\"",
		"\"'\"",
		"\"simple\"",
	}
	actual := simplewrap.
		DoubleQuoteWrapElements(true,
			testCases...)
	
	// Assert
	convey.Convey("Wrap strings with double quote, if exists already then skip adding", t, func() {
		convey.So(actual, should.Equal, expectation)
	})
}

// func Test_When_DoubleQuoteWrapElements_SkipQuoteOnPresent_Disabled_Should_Have_DuplicateDoubleQuotations(t *testing.T) {
// 	testCases := []string{
// 		"some-elem",
// 		"alim-elem",
// 		"\"has-quote\"",
// 		"",
// 		"\"",
// 		"\"first",
// 		"last\"",
// 		"'",
// 		"simple",
// 	}
// 	expectation := []string{
// 		"\"some-elem\"",
// 		"\"alim-elem\"",
// 		"\"\"has-quote\"\"",
// 		"\"\"",
// 		"\"\"\"",
// 		"\"\"first\"",
// 		"\"last\"\"",
// 		"\"'\"",
// 		"\"simple\"",
// 	}
// 	
// 	// Act
// 	actual := simplewrap.
// 		DoubleQuoteWrapElements(
// 			false,
// 			testCases...)
// 	
// 	// Assert
// 	convey.Convey("Wrap strings with double quote, if exists already then skip adding", t, func() {
// 		convey.So(actual, should.Equal, expectation)
// 	})
// }
