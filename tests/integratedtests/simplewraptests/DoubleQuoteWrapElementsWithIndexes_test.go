package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

func Test_DoubleQuoteWrapElementsWithIndexes_Should_Return_Slice_WithDoubleQuoteWithIndexes(t *testing.T) {
	// Arrange
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
		"\"some-elem[0]\"",
		"\"alim-elem[1]\"",
		"\"\"has-quote\"[2]\"",
		"\"[3]\"",
		"\"\"[4]\"",
		"\"\"first[5]\"",
		"\"last\"[6]\"",
		"\"'[7]\"",
		"\"simple[8]\"",
	}

	// Act
	actual := simplewrap.
		DoubleQuoteWrapElementsWithIndexes(
			testCases...,
		)

	// Assert
	convey.Convey(
		"Wrap strings with double quote with indexes - "+
			"doesn't verify existing double quote, "+
			"and possible duplicate double quote possible", t, func() {
			convey.So(actual, should.Equal, expectation)
		},
	)
}
