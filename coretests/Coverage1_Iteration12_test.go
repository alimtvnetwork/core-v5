package coretests

import (
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// AnyToBytes — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I12_AnyToBytes_Bytes(t *testing.T) {
	input := []byte("hello")
	out := AnyToBytes(input)
	if string(out) != "hello" {
		t.Fatal("expected bytes pass-through")
	}
}

func Test_I12_AnyToBytes_NilBytes(t *testing.T) {
	var input []byte
	out := AnyToBytes(input)
	if out != nil {
		t.Fatal("expected nil for nil bytes")
	}
}

func Test_I12_AnyToBytes_String(t *testing.T) {
	out := AnyToBytes("world")
	if string(out) != "world" {
		t.Fatal("expected string->bytes conversion")
	}
}

func Test_I12_AnyToBytes_OtherType(t *testing.T) {
	out := AnyToBytes(map[string]int{"a": 1})
	if len(out) == 0 {
		t.Fatal("expected json bytes")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// DraftType — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I12_DraftType_PtrOrNonPtr(t *testing.T) {
	d := &DraftType{SampleString1: "a", SampleInteger: 1}
	p := d.PtrOrNonPtr(true)
	if p == nil {
		t.Fatal("expected ptr")
	}
	np := d.PtrOrNonPtr(false)
	if _, ok := np.(DraftType); !ok {
		t.Fatal("expected non-ptr DraftType")
	}
	var nilD *DraftType
	if nilD.PtrOrNonPtr(true) != nil {
		t.Fatal("nil should return nil")
	}
}

func Test_I12_DraftType_ClonePtr_Nil(t *testing.T) {
	var d *DraftType
	if d.ClonePtr() != nil {
		t.Fatal("nil clone should be nil")
	}
}

func Test_I12_DraftType_IsEqual_Branches(t *testing.T) {
	d1 := &DraftType{SampleString1: "a", SampleString2: "b", SampleInteger: 1, Lines: []string{"x"}, RawBytes: []byte("r")}
	d2 := d1.ClonePtr()
	if !d1.IsEqualAll(d2) {
		t.Fatal("expected equal")
	}
	// diff SampleString2
	d3 := d1.ClonePtr()
	d3.SampleString2 = "c"
	if d1.IsEqual(false, d3) {
		t.Fatal("expected not equal on SampleString2")
	}
	// diff SampleInteger
	d4 := d1.ClonePtr()
	d4.SampleInteger = 99
	if d1.IsEqual(false, d4) {
		t.Fatal("expected not equal on SampleInteger")
	}
	// diff RawBytes
	d5 := d1.ClonePtr()
	d5.RawBytes = []byte("different")
	if d1.IsEqual(false, d5) {
		t.Fatal("expected not equal on RawBytes")
	}
	// diff Lines
	d6 := d1.ClonePtr()
	d6.Lines = []string{"y"}
	if d1.IsEqual(false, d6) {
		t.Fatal("expected not equal on Lines")
	}
	// nil vs nil
	var nilD1, nilD2 *DraftType
	if !nilD1.IsEqual(false, nilD2) {
		t.Fatal("nil==nil should be true")
	}
	// nil vs non-nil
	if nilD1.IsEqual(false, d1) {
		t.Fatal("nil!=non-nil should be false")
	}
	// same ptr
	if !d1.IsEqual(false, d1) {
		t.Fatal("same ptr should be equal")
	}
}

func Test_I12_DraftType_VerifyNotEqual(t *testing.T) {
	d1 := &DraftType{SampleString1: "a"}
	d2 := &DraftType{SampleString1: "b"}
	msg := d1.VerifyAllNotEqualMessage(d2)
	if msg == "" {
		t.Fatal("expected not-equal message")
	}
	err := d1.VerifyAllNotEqualErr(d2)
	if err == nil {
		t.Fatal("expected not-equal error")
	}
	err2 := d1.VerifyNotEqualExcludingInnerFieldsErr(d2)
	if err2 == nil {
		t.Fatal("expected not-equal excluding-inner error")
	}
	// equal case
	d3 := d1.ClonePtr()
	if d1.VerifyAllNotEqualErr(d3) != nil {
		t.Fatal("expected nil for equal drafts")
	}
}

func Test_I12_DraftType_JsonAndSetters(t *testing.T) {
	d := DraftType{SampleString1: "x"}
	s := d.JsonString()
	if s == "" {
		t.Fatal("expected json string")
	}
	b := d.JsonBytes()
	if len(b) == 0 {
		t.Fatal("expected json bytes")
	}
	b2 := d.JsonBytesPtr()
	if len(b2) == 0 {
		t.Fatal("expected json bytes ptr")
	}
	d.SetF2Integer(42)
	if d.F2Integer() != 42 {
		t.Fatal("expected f2=42")
	}
	if d.F1String() != "" {
		t.Fatal("expected empty f1")
	}
	np := d.NonPtr()
	_ = np
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleTestCase — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I12_SimpleTestCase_Titles(t *testing.T) {
	tc := SimpleTestCase{Title: "test-title"}
	if tc.CaseTitle() != "test-title" {
		t.Fatal("expected case title")
	}
	if tc.FormTitle(0) == "" {
		t.Fatal("expected form title")
	}
	if tc.CustomTitle(0, "custom") == "" {
		t.Fatal("expected custom title")
	}
}

func Test_I12_SimpleTestCase_ArrangeAndExpected(t *testing.T) {
	tc := SimpleTestCase{
		Title:         "tc",
		ArrangeInput:  "arrange-val",
		ExpectedInput: "expected-val",
	}
	if tc.ArrangeString() == "" {
		t.Fatal("expected arrange string")
	}
	if tc.Input() != "arrange-val" {
		t.Fatal("expected input")
	}
	if tc.Expected() != "expected-val" {
		t.Fatal("expected expected")
	}
	if tc.ExpectedString() == "" {
		t.Fatal("expected expected string")
	}
	tc.SetActual("actual-val")
	if tc.ActualString() == "" {
		t.Fatal("expected actual string")
	}
	if tc.String(0) == "" {
		t.Fatal("expected string repr")
	}
	if tc.LinesString(0) == "" {
		t.Fatal("expected lines string")
	}
	_ = tc.AsSimpleTestCaseWrapper()
}

// ══════════════════════════════════════════════════════════════════════════════
// messagePrinter — exercise branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I12_MessagePrinter_FailedExpected(t *testing.T) {
	p := printMessage{}
	// not failed — no output
	p.FailedExpected(false, "when", "actual", "expected", 0)
	// failed — exercises print
	p.FailedExpected(true, "when", "actual", "expected", 1)
}

func Test_I12_MessagePrinter_NameValueAndValue(t *testing.T) {
	p := printMessage{}
	p.NameValue("header", map[string]int{"a": 1})
	p.Value("header2", []int{1, 2, 3})
}

// ══════════════════════════════════════════════════════════════════════════════
// SkipOnUnix — exercise
// ══════════════════════════════════════════════════════════════════════════════

func Test_I12_SkipOnUnix(t *testing.T) {
	// Just exercise the function; on Unix it will skip, on Windows it won't.
	// We wrap in a sub-test so the parent doesn't get skipped.
	t.Run("sub", func(st *testing.T) {
		SkipOnUnix(st)
	})
}
