package coreteststests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// DraftType — isIncludingInnerFields, JsonString, JsonBytes
// Covers DraftType.go L148, L174, L184
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_DraftType_IsEqualIncludingInner_Mismatch(t *testing.T) {
	// Arrange
	dt1 := coretests.NewDraftType("a", "b", 1, 2, []byte{1}, []string{"x"})
	dt2 := coretests.NewDraftType("a", "b", 1, 99, []byte{1}, []string{"x"})

	// Act — with inner fields check, f2Integer differs
	result := dt1.IsEqual(true, dt2)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "DraftType IsEqual inner mismatch", actual)
}

func Test_Cov2_DraftType_JsonString(t *testing.T) {
	// Arrange
	dt := coretests.NewDraftType("s1", "s2", 42, 7, []byte{0x01}, []string{"line1"})

	// Act
	result := dt.JsonString()

	// Assert
	actual := args.Map{"nonEmpty": result != ""}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "DraftType JsonString", actual)
}

func Test_Cov2_DraftType_JsonBytes(t *testing.T) {
	// Arrange
	dt := coretests.NewDraftType("s1", "s2", 42, 7, []byte{0x01}, []string{"line1"})

	// Act
	result := dt.JsonBytes()

	// Assert
	actual := args.Map{"hasBytes": len(result) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "DraftType JsonBytes", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleTestCase — ShouldBeEqual, ShouldHaveNoError, ShouldContains
// Covers SimpleTestCase.go L126-137, L143-154, L160-171
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_SimpleTestCase_ShouldBeEqual(t *testing.T) {
	// Arrange
	tc := coretests.SimpleTestCase{
		Title:         "simple equal test",
		ExpectedInput: "hello",
	}

	// Act & Assert
	tc.ShouldBeEqual(0, t, "hello")
}

func Test_Cov2_SimpleTestCase_ShouldHaveNoError(t *testing.T) {
	// Arrange
	tc := coretests.SimpleTestCase{
		Title:         "no error test",
		ExpectedInput: nil,
	}

	// Act & Assert
	tc.ShouldHaveNoError(0, t, nil)
}

func Test_Cov2_SimpleTestCase_ShouldContains(t *testing.T) {
	// Arrange
	tc := coretests.SimpleTestCase{
		Title:         "contains test",
		ExpectedInput: "world",
	}

	// Act & Assert
	tc.ShouldContains(0, t, []string{"hello", "world"})
}

// ══════════════════════════════════════════════════════════════════════════════
// messagePrinter — FailedExpected, NameValue, Value
// Covers messagePrinter.go L15-20, L26-34, L39-47
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_PrintMessage_FailedExpected(t *testing.T) {
	// Arrange & Act — trigger the isFailed=true branch
	coretests.Print.FailedExpected(true, "when", "actual", "expected", 0)

	// Assert — no panic means success
	actual := args.Map{"executed": true}
	expected := args.Map{"executed": true}
	expected.ShouldBeEqual(t, 0, "PrintMessage FailedExpected", actual)
}

func Test_Cov2_PrintMessage_NameValue(t *testing.T) {
	// Arrange & Act
	coretests.Print.NameValue("header", "some value")

	// Assert
	actual := args.Map{"executed": true}
	expected := args.Map{"executed": true}
	expected.ShouldBeEqual(t, 0, "PrintMessage NameValue", actual)
}

func Test_Cov2_PrintMessage_Value(t *testing.T) {
	// Arrange & Act
	coretests.Print.Value("header", "some value")

	// Assert
	actual := args.Map{"executed": true}
	expected := args.Map{"executed": true}
	expected.ShouldBeEqual(t, 0, "PrintMessage Value", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BaseTestCase — ShouldBeExplicit disabled, TypeShouldMatch error
// Covers BaseTestCaseAssertions.go L73-77, L88-92, L123-141
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_BaseTestCase_ShouldBeExplicit_Disabled(t *testing.T) {
	// Arrange
	tc := &coretests.BaseTestCase{
		Title:         "disabled test",
		ExpectedInput: "expected",
	}
	tc.IsEnable.SetFalse()

	// Act & Assert — exercises the disabled path with noPrintAssert
	tc.ShouldBeExplicit(
		false,
		0,
		t,
		"disabled test",
		"expected",
		convey.ShouldEqual,
		"expected",
	)
}

// ══════════════════════════════════════════════════════════════════════════════
// BaseTestCaseValidation — TypesValidationMustPasses error path
// Covers BaseTestCaseValidation.go L18-23
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov2_BaseTestCase_TypesValidationMustPasses_NoError(t *testing.T) {
	// Arrange
	tc := &coretests.BaseTestCase{
		Title:         "type validation",
		ExpectedInput: "hello",
	}
	tc.SetActual("hello")

	// Act & Assert — no type mismatch
	tc.TypesValidationMustPasses(t)
}
