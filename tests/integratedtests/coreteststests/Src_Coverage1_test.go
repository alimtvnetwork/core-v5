package coreteststests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Src_AnyToBytes_Verification(t *testing.T) {
	for caseIndex, tc := range srcAnyToBytesTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal := input["input"]
		inputType := input["type"].(string)

		// Act
		var result []byte
		switch inputType {
		case "bytes":
			result = coretests.AnyToBytes(inputVal.([]byte))
		case "nilBytes":
			var nilBytes []byte
			result = coretests.AnyToBytes(nilBytes)
		case "string":
			result = coretests.AnyToBytes(inputVal.(string))
		case "other":
			result = coretests.AnyToBytes(inputVal)
		}

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["result"]; has {
			actual["result"] = string(result)
		}
		if _, has := expected["isNil"]; has {
			actual["isNil"] = result == nil
		}
		if _, has := expected["nonEmpty"]; has {
			actual["nonEmpty"] = len(result) > 0
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Src_DraftType_PtrOrNonPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcDraftTypePtrOrNonPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		d := &coretests.DraftType{
			SampleString1: input["string1"].(string),
			SampleInteger: input["integer"].(int),
		}
		asPtr := input["asPtr"].(bool)

		// Act
		result := d.PtrOrNonPtr(asPtr)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["isNotNil"]; has {
			actual["isNotNil"] = result != nil
		}
		if _, has := expected["isDraftType"]; has {
			_, ok := result.(coretests.DraftType)
			actual["isDraftType"] = ok
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Src_DraftType_PtrOrNonPtr_NilReceiver(t *testing.T) {
	// Arrange
	var nilD *coretests.DraftType

	// Act
	result := nilD.PtrOrNonPtr(true)

	// Assert
	if result != nil {
		t.Fatal("nil receiver should return nil")
	}
}

func Test_Src_DraftType_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var d *coretests.DraftType

	// Act
	result := d.ClonePtr()

	// Assert
	tc := srcDraftTypeClonePtrNilTestCase
	tc.ShouldBeEqualMap(t, 0, args.Map{
		"isNil": result == nil,
	})
}

func Test_Src_DraftType_IsEqual_Verification(t *testing.T) {
	base := &coretests.DraftType{
		SampleString1: "a",
		SampleString2: "b",
		SampleInteger: 1,
		Lines:         []string{"x"},
		RawBytes:      []byte("r"),
	}

	for caseIndex, tc := range srcDraftTypeIsEqualTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		scenario := input["scenario"].(string)

		var result bool
		switch scenario {
		case "equal":
			// Act
			d2 := base.ClonePtr()
			result = base.IsEqualAll(d2)
		case "diffString2":
			d3 := base.ClonePtr()
			d3.SampleString2 = "c"
			result = base.IsEqual(false, d3)
		case "diffInteger":
			d4 := base.ClonePtr()
			d4.SampleInteger = 99
			result = base.IsEqual(false, d4)
		case "diffRawBytes":
			d5 := base.ClonePtr()
			d5.RawBytes = []byte("different")
			result = base.IsEqual(false, d5)
		case "diffLines":
			d6 := base.ClonePtr()
			d6.Lines = []string{"y"}
			result = base.IsEqual(false, d6)
		case "bothNil":
			var n1, n2 *coretests.DraftType
			result = n1.IsEqual(false, n2)
		case "nilVsNonNil":
			var n1 *coretests.DraftType
			result = n1.IsEqual(false, base)
		case "samePtr":
			result = base.IsEqual(false, base)
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_Src_DraftType_VerifyNotEqual(t *testing.T) {
	// Arrange
	d1 := &coretests.DraftType{SampleString1: "a", Lines: []string{}, RawBytes: []byte{}}
	d2 := &coretests.DraftType{SampleString1: "b", Lines: []string{}, RawBytes: []byte{}}

	// Act
	msg := d1.VerifyAllNotEqualMessage(d2)
	err := d1.VerifyAllNotEqualErr(d2)
	err2 := d1.VerifyNotEqualExcludingInnerFieldsErr(d2)

	// Assert
	convey.Convey("VerifyAllNotEqualMessage returns non-empty -- different drafts", t, func() {
		convey.So(msg, should.NotBeEmpty)
	})
	convey.Convey("VerifyAllNotEqualErr returns error -- different drafts", t, func() {
		convey.So(err, should.NotBeNil)
	})
	convey.Convey("VerifyNotEqualExcludingInnerFieldsErr returns error -- different drafts", t, func() {
		convey.So(err2, should.NotBeNil)
	})

	// Arrange (equal case)
	d3 := d1.ClonePtr()

	// Act
	err3 := d1.VerifyAllNotEqualErr(d3)

	// Assert
	convey.Convey("VerifyAllNotEqualErr returns nil -- equal drafts", t, func() {
		convey.So(err3, should.BeNil)
	})
}

func Test_Src_DraftType_JsonAndSetters(t *testing.T) {
	// Arrange
	d := coretests.DraftType{SampleString1: "x"}

	// Act
	s := d.JsonString()
	b := d.JsonBytes()
	b2 := d.JsonBytesPtr()

	// Assert
	if s == "" {
		t.Fatal("expected json string")
	}
	if len(b) == 0 {
		t.Fatal("expected json bytes")
	}
	if len(b2) == 0 {
		t.Fatal("expected json bytes ptr")
	}

	// Arrange + Act (setters)
	d.SetF2Integer(42)

	// Assert
	if d.F2Integer() != 42 {
		t.Fatal("expected f2=42")
	}
	if d.F1String() != "" {
		t.Fatal("expected empty f1")
	}
	_ = d.NonPtr()
}

func Test_Src_SimpleTestCase_Titles_Verification(t *testing.T) {
	// Arrange
	tc := srcSimpleTestCaseTitlesTestCase
	input := tc.ArrangeInput.(args.Map)
	title := input["title"].(string)

	stc := coretests.SimpleTestCase{Title: title}

	// Act
	caseTitle := stc.CaseTitle()
	formTitle := stc.FormTitle(0)
	customTitle := stc.CustomTitle(0, "custom")

	// Assert
	tc.ShouldBeEqualMap(t, 0, args.Map{
		"caseTitle":           caseTitle,
		"formTitleNotEmpty":   formTitle != "",
		"customTitleNotEmpty": customTitle != "",
	})
}

func Test_Src_SimpleTestCase_ArrangeAndExpected(t *testing.T) {
	// Arrange
	stc := coretests.SimpleTestCase{
		Title:         "tc",
		ArrangeInput:  "arrange-val",
		ExpectedInput: "expected-val",
	}

	// Act
	arrangeStr := stc.ArrangeString()
	inputVal := stc.Input()
	expectedVal := stc.Expected()
	expectedStr := stc.ExpectedString()

	// Assert
	if arrangeStr == "" {
		t.Fatal("expected arrange string")
	}
	if inputVal != "arrange-val" {
		t.Fatal("expected input")
	}
	if expectedVal != "expected-val" {
		t.Fatal("expected expected")
	}
	if expectedStr == "" {
		t.Fatal("expected expected string")
	}

	// Act (setters)
	stc.SetActual("actual-val")
	actualStr := stc.ActualString()
	str := stc.String(0)
	linesStr := stc.LinesString(0)

	// Assert
	if actualStr == "" {
		t.Fatal("expected actual string")
	}
	if str == "" {
		t.Fatal("expected string repr")
	}
	if linesStr == "" {
		t.Fatal("expected lines string")
	}
	_ = stc.AsSimpleTestCaseWrapper()
}
