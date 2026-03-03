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

// --- Clone (value) ---

func Test_FromTo_Clone_CopiesValues(t *testing.T) {
	// Arrange
	original := coreinstruction.FromTo{From: "a", To: "b"}

	// Act
	cloned := original.Clone()

	// Assert
	convey.Convey("Clone - value clone copies From and To", t, func() {
		convey.So(cloned.From, should.Equal, "a")
		convey.So(cloned.To, should.Equal, "b")
	})
}

// --- IsNull ---

func Test_FromTo_IsNull_True(t *testing.T) {
	// Arrange
	var nilFT *coreinstruction.FromTo

	// Assert
	convey.Convey("IsNull - nil receiver returns true", t, func() {
		convey.So(nilFT.IsNull(), should.BeTrue)
	})
}

func Test_FromTo_IsNull_False(t *testing.T) {
	// Arrange
	ft := &coreinstruction.FromTo{From: "x", To: "y"}

	// Assert
	convey.Convey("IsNull - non-nil receiver returns false", t, func() {
		convey.So(ft.IsNull(), should.BeFalse)
	})
}

// --- IsFromEmpty / IsToEmpty ---

func Test_FromTo_IsFromEmpty_True(t *testing.T) {
	// Arrange
	ft := &coreinstruction.FromTo{From: "", To: "dest"}

	// Assert
	convey.Convey("IsFromEmpty - empty From returns true", t, func() {
		convey.So(ft.IsFromEmpty(), should.BeTrue)
	})
}

func Test_FromTo_IsFromEmpty_NilReceiver(t *testing.T) {
	// Arrange
	var nilFT *coreinstruction.FromTo

	// Assert
	convey.Convey("IsFromEmpty - nil receiver returns true", t, func() {
		convey.So(nilFT.IsFromEmpty(), should.BeTrue)
	})
}

func Test_FromTo_IsToEmpty_True(t *testing.T) {
	// Arrange
	ft := &coreinstruction.FromTo{From: "src", To: ""}

	// Assert
	convey.Convey("IsToEmpty - empty To returns true", t, func() {
		convey.So(ft.IsToEmpty(), should.BeTrue)
	})
}

func Test_FromTo_IsToEmpty_False(t *testing.T) {
	// Arrange
	ft := &coreinstruction.FromTo{From: "src", To: "dest"}

	// Assert
	convey.Convey("IsToEmpty - non-empty To returns false", t, func() {
		convey.So(ft.IsToEmpty(), should.BeFalse)
	})
}

// --- String ---

func Test_FromTo_String(t *testing.T) {
	// Arrange
	ft := coreinstruction.FromTo{From: "alpha", To: "beta"}

	// Act
	result := ft.String()

	// Assert
	convey.Convey("String - should contain From and To values", t, func() {
		convey.So(result, should.ContainSubstring, "alpha")
		convey.So(result, should.ContainSubstring, "beta")
	})
}

// --- FromName / ToName ---

func Test_FromTo_FromName_ToName(t *testing.T) {
	// Arrange
	ft := coreinstruction.FromTo{From: "src", To: "dst"}

	// Assert
	convey.Convey("FromName/ToName return field values", t, func() {
		convey.So(ft.FromName(), should.Equal, "src")
		convey.So(ft.ToName(), should.Equal, "dst")
	})
}

// --- SetFromName / SetToName ---

func Test_FromTo_SetFromName(t *testing.T) {
	// Arrange
	ft := &coreinstruction.FromTo{From: "old", To: "t"}

	// Act
	ft.SetFromName("new")

	// Assert
	convey.Convey("SetFromName - updates From field", t, func() {
		convey.So(ft.From, should.Equal, "new")
	})
}

func Test_FromTo_SetToName(t *testing.T) {
	// Arrange
	ft := &coreinstruction.FromTo{From: "f", To: "old"}

	// Act
	ft.SetToName("new")

	// Assert
	convey.Convey("SetToName - updates To field", t, func() {
		convey.So(ft.To, should.Equal, "new")
	})
}

func Test_FromTo_SetFromName_NilReceiver(t *testing.T) {
	// Arrange
	var nilFT *coreinstruction.FromTo

	// Act & Assert (should not panic)
	convey.Convey("SetFromName - nil receiver should not panic", t, func() {
		convey.So(func() { nilFT.SetFromName("x") }, should.NotPanic)
	})
}

// --- SourceDestination ---

func Test_FromTo_SourceDestination(t *testing.T) {
	// Arrange
	ft := &coreinstruction.FromTo{From: "src", To: "dst"}

	// Act
	sd := ft.SourceDestination()

	// Assert
	convey.Convey("SourceDestination - maps From->Source, To->Destination", t, func() {
		convey.So(sd, should.NotBeNil)
		convey.So(sd.Source, should.Equal, "src")
		convey.So(sd.Destination, should.Equal, "dst")
	})
}

func Test_FromTo_SourceDestination_NilReceiver(t *testing.T) {
	// Arrange
	var nilFT *coreinstruction.FromTo

	// Assert
	convey.Convey("SourceDestination - nil receiver returns nil", t, func() {
		convey.So(nilFT.SourceDestination(), should.BeNil)
	})
}

// --- Rename ---

func Test_FromTo_Rename(t *testing.T) {
	// Arrange
	ft := &coreinstruction.FromTo{From: "old", To: "new"}

	// Act
	rn := ft.Rename()

	// Assert
	convey.Convey("Rename - maps From->Existing, To->New", t, func() {
		convey.So(rn, should.NotBeNil)
		convey.So(rn.Existing, should.Equal, "old")
		convey.So(rn.New, should.Equal, "new")
	})
}

func Test_FromTo_Rename_NilReceiver(t *testing.T) {
	// Arrange
	var nilFT *coreinstruction.FromTo

	// Assert
	convey.Convey("Rename - nil receiver returns nil", t, func() {
		convey.So(nilFT.Rename(), should.BeNil)
	})
}
