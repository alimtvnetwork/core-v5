package typesconvtests

import (
	"testing"

	"github.com/alimtvnetwork/core/typesconv"
)

// Cover non-nil paths for all Ptr functions and remaining StringToBool branches

func Test_IntPtrToSimple_NonNil_Cov2(t *testing.T) {
	v := 42
	if r := typesconv.IntPtrToSimple(&v); r != 42 {
		t.Error("expected 42")
	}
}

func Test_IntPtrToSimpleDef_NonNil_Cov2(t *testing.T) {
	v := 10
	if r := typesconv.IntPtrToSimpleDef(&v, 99); r != 10 {
		t.Error("expected 10")
	}
}

func Test_IntPtrToDefPtr_NonNil_Cov2(t *testing.T) {
	v := 10
	r := typesconv.IntPtrToDefPtr(&v, 99)
	if *r != 10 {
		t.Error("expected 10")
	}
}

func Test_IntPtrDefValFunc_NonNil_Cov2(t *testing.T) {
	v := 10
	r := typesconv.IntPtrDefValFunc(&v, func() int { return 99 })
	if *r != 10 {
		t.Error("expected 10")
	}
}

func Test_BytePtrToSimple_NonNil_Cov2(t *testing.T) {
	v := byte(5)
	if r := typesconv.BytePtrToSimple(&v); r != 5 {
		t.Error("expected 5")
	}
}

func Test_BytePtrToSimpleDef_NonNil_Cov2(t *testing.T) {
	v := byte(5)
	if r := typesconv.BytePtrToSimpleDef(&v, 9); r != 5 {
		t.Error("expected 5")
	}
}

func Test_BytePtrToDefPtr_NonNil_Cov2(t *testing.T) {
	v := byte(5)
	r := typesconv.BytePtrToDefPtr(&v, 9)
	if *r != 5 {
		t.Error("expected 5")
	}
}

func Test_BytePtrDefValFunc_NonNil_Cov2(t *testing.T) {
	v := byte(5)
	r := typesconv.BytePtrDefValFunc(&v, func() byte { return 9 })
	if *r != 5 {
		t.Error("expected 5")
	}
}

func Test_FloatPtrToSimple_NonNil_Cov2(t *testing.T) {
	v := float32(3.14)
	if r := typesconv.FloatPtrToSimple(&v); r != 3.14 {
		t.Errorf("expected 3.14, got %v", r)
	}
}

func Test_FloatPtrToSimpleDef_NonNil_Cov2(t *testing.T) {
	v := float32(3.14)
	if r := typesconv.FloatPtrToSimpleDef(&v, 9.9); r != 3.14 {
		t.Error("expected 3.14")
	}
}

func Test_FloatPtrToDefPtr_NonNil_Cov2(t *testing.T) {
	v := float32(3.14)
	r := typesconv.FloatPtrToDefPtr(&v, 9.9)
	if *r != 3.14 {
		t.Error("expected 3.14")
	}
}

func Test_FloatPtrDefValFunc_NonNil_Cov2(t *testing.T) {
	v := float32(3.14)
	r := typesconv.FloatPtrDefValFunc(&v, func() float32 { return 9.9 })
	if *r != 3.14 {
		t.Error("expected 3.14")
	}
}

func Test_StringPtrToSimpleDef_NonNil_Cov2(t *testing.T) {
	v := "hello"
	if r := typesconv.StringPtrToSimpleDef(&v, "fb"); r != "hello" {
		t.Error("expected hello")
	}
}

func Test_StringPtrToDefPtr_NonNil_Cov2(t *testing.T) {
	v := "hello"
	r := typesconv.StringPtrToDefPtr(&v, "fb")
	if *r != "hello" {
		t.Error("expected hello")
	}
}

func Test_StringPtrDefValFunc_NonNil_Cov2(t *testing.T) {
	v := "hello"
	r := typesconv.StringPtrDefValFunc(&v, func() string { return "fb" })
	if *r != "hello" {
		t.Error("expected hello")
	}
}

func Test_StringToBool_YES_Cov2(t *testing.T) {
	if !typesconv.StringToBool("YES") {
		t.Error("YES should be true")
	}
}

func Test_StringToBool_No_Cov2(t *testing.T) {
	if typesconv.StringToBool("No") {
		t.Error("No should be false")
	}
}

func Test_StringToBool_NO_Cov2(t *testing.T) {
	if typesconv.StringToBool("NO") {
		t.Error("NO should be false")
	}
}

func Test_StringToBool_false_Cov2(t *testing.T) {
	if typesconv.StringToBool("false") {
		t.Error("false should be false")
	}
}

func Test_StringPointerToBool_Empty_Cov2(t *testing.T) {
	s := ""
	if typesconv.StringPointerToBool(&s) {
		t.Error("empty should be false")
	}
}

func Test_StringPointerToBoolPtr_Empty_Cov2(t *testing.T) {
	s := ""
	r := typesconv.StringPointerToBoolPtr(&s)
	if *r != false {
		t.Error("expected false")
	}
}

func Test_StringToBoolPtr_NonEmpty_Cov2(t *testing.T) {
	r := typesconv.StringToBoolPtr("yes")
	if *r != true {
		t.Error("expected true")
	}
}
