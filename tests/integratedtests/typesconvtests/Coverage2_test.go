package typesconvtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/typesconv"
)

func Test_Cov2_IntPtrToSimple_NonNil(t *testing.T) {
	v := 42
	actual := args.Map{"result": typesconv.IntPtrToSimple(&v)}
	expected := args.Map{"result": 42}
	coretestcases.CaseV1{Title: "IntPtrToSimple_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_IntPtrToSimpleDef_NonNil(t *testing.T) {
	v := 10
	actual := args.Map{"result": typesconv.IntPtrToSimpleDef(&v, 99)}
	expected := args.Map{"result": 10}
	coretestcases.CaseV1{Title: "IntPtrToSimpleDef_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_IntPtrToDefPtr_NonNil(t *testing.T) {
	v := 10
	r := typesconv.IntPtrToDefPtr(&v, 99)
	actual := args.Map{"result": *r}
	expected := args.Map{"result": 10}
	coretestcases.CaseV1{Title: "IntPtrToDefPtr_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_IntPtrDefValFunc_NonNil(t *testing.T) {
	v := 10
	r := typesconv.IntPtrDefValFunc(&v, func() int { return 99 })
	actual := args.Map{"result": *r}
	expected := args.Map{"result": 10}
	coretestcases.CaseV1{Title: "IntPtrDefValFunc_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_BytePtrToSimple_NonNil(t *testing.T) {
	v := byte(5)
	actual := args.Map{"result": int(typesconv.BytePtrToSimple(&v))}
	expected := args.Map{"result": 5}
	coretestcases.CaseV1{Title: "BytePtrToSimple_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_BytePtrToSimpleDef_NonNil(t *testing.T) {
	v := byte(5)
	actual := args.Map{"result": int(typesconv.BytePtrToSimpleDef(&v, 9))}
	expected := args.Map{"result": 5}
	coretestcases.CaseV1{Title: "BytePtrToSimpleDef_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_BytePtrToDefPtr_NonNil(t *testing.T) {
	v := byte(5)
	r := typesconv.BytePtrToDefPtr(&v, 9)
	actual := args.Map{"result": int(*r)}
	expected := args.Map{"result": 5}
	coretestcases.CaseV1{Title: "BytePtrToDefPtr_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_BytePtrDefValFunc_NonNil(t *testing.T) {
	v := byte(5)
	r := typesconv.BytePtrDefValFunc(&v, func() byte { return 9 })
	actual := args.Map{"result": int(*r)}
	expected := args.Map{"result": 5}
	coretestcases.CaseV1{Title: "BytePtrDefValFunc_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_FloatPtrToSimple_NonNil(t *testing.T) {
	v := float32(3.14)
	actual := args.Map{"result": fmt.Sprintf("%.2f", typesconv.FloatPtrToSimple(&v))}
	expected := args.Map{"result": "3.14"}
	coretestcases.CaseV1{Title: "FloatPtrToSimple_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_FloatPtrToSimpleDef_NonNil(t *testing.T) {
	v := float32(3.14)
	actual := args.Map{"result": fmt.Sprintf("%.2f", typesconv.FloatPtrToSimpleDef(&v, 9.9))}
	expected := args.Map{"result": "3.14"}
	coretestcases.CaseV1{Title: "FloatPtrToSimpleDef_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_FloatPtrToDefPtr_NonNil(t *testing.T) {
	v := float32(3.14)
	r := typesconv.FloatPtrToDefPtr(&v, 9.9)
	actual := args.Map{"result": fmt.Sprintf("%.2f", *r)}
	expected := args.Map{"result": "3.14"}
	coretestcases.CaseV1{Title: "FloatPtrToDefPtr_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_FloatPtrDefValFunc_NonNil(t *testing.T) {
	v := float32(3.14)
	r := typesconv.FloatPtrDefValFunc(&v, func() float32 { return 9.9 })
	actual := args.Map{"result": fmt.Sprintf("%.2f", *r)}
	expected := args.Map{"result": "3.14"}
	coretestcases.CaseV1{Title: "FloatPtrDefValFunc_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_StringPtrToSimpleDef_NonNil(t *testing.T) {
	v := "hello"
	actual := args.Map{"result": typesconv.StringPtrToSimpleDef(&v, "fb")}
	expected := args.Map{"result": "hello"}
	coretestcases.CaseV1{Title: "StringPtrToSimpleDef_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_StringPtrToDefPtr_NonNil(t *testing.T) {
	v := "hello"
	r := typesconv.StringPtrToDefPtr(&v, "fb")
	actual := args.Map{"result": *r}
	expected := args.Map{"result": "hello"}
	coretestcases.CaseV1{Title: "StringPtrToDefPtr_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_StringPtrDefValFunc_NonNil(t *testing.T) {
	v := "hello"
	r := typesconv.StringPtrDefValFunc(&v, func() string { return "fb" })
	actual := args.Map{"result": *r}
	expected := args.Map{"result": "hello"}
	coretestcases.CaseV1{Title: "StringPtrDefValFunc_NonNil", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_StringToBool_YES(t *testing.T) {
	actual := args.Map{"result": typesconv.StringToBool("YES")}
	expected := args.Map{"result": true}
	coretestcases.CaseV1{Title: "StringToBool_YES", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_StringToBool_No(t *testing.T) {
	actual := args.Map{"result": typesconv.StringToBool("No")}
	expected := args.Map{"result": false}
	coretestcases.CaseV1{Title: "StringToBool_No", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_StringToBool_NO(t *testing.T) {
	actual := args.Map{"result": typesconv.StringToBool("NO")}
	expected := args.Map{"result": false}
	coretestcases.CaseV1{Title: "StringToBool_NO", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_StringToBool_false(t *testing.T) {
	actual := args.Map{"result": typesconv.StringToBool("false")}
	expected := args.Map{"result": false}
	coretestcases.CaseV1{Title: "StringToBool_false", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_StringPointerToBool_Empty(t *testing.T) {
	s := ""
	actual := args.Map{"result": typesconv.StringPointerToBool(&s)}
	expected := args.Map{"result": false}
	coretestcases.CaseV1{Title: "StringPointerToBool_Empty", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_StringPointerToBoolPtr_Empty(t *testing.T) {
	s := ""
	r := typesconv.StringPointerToBoolPtr(&s)
	actual := args.Map{"result": *r}
	expected := args.Map{"result": false}
	coretestcases.CaseV1{Title: "StringPointerToBoolPtr_Empty", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov2_StringToBoolPtr_NonEmpty(t *testing.T) {
	r := typesconv.StringToBoolPtr("yes")
	actual := args.Map{"result": *r}
	expected := args.Map{"result": true}
	coretestcases.CaseV1{Title: "StringToBoolPtr_NonEmpty", ExpectedInput: expected}.ShouldBeEqualMapFirst(t, actual)
}
