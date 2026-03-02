package coreinstructiontests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coreinstruction"
)

// --- ClonePtr positive ---

func Test_FromTo_ClonePtr_Positive(t *testing.T) {
	// Arrange
	original := &coreinstruction.FromTo{
		From: "source",
		To:   "destination",
	}

	// Act
	cloned := original.ClonePtr()

	// Assert
	convey.Convey("ClonePtr - should copy From and To values", t, func() {
		convey.So(cloned, should.NotBeNil)
		convey.So(cloned.From, should.Equal, "source")
		convey.So(cloned.To, should.Equal, "destination")
		convey.So(cloned, should.NotPointTo, original)
	})
}

// --- ClonePtr nil receiver ---

func Test_FromTo_ClonePtr_NilReceiver(t *testing.T) {
	// Arrange
	var nilFromTo *coreinstruction.FromTo

	// Act
	cloned := nilFromTo.ClonePtr()

	// Assert
	convey.Convey("ClonePtr - nil receiver should return nil", t, func() {
		convey.So(cloned, should.BeNil)
	})
}
