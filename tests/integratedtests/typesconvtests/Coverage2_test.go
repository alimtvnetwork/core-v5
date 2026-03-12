package typesconvtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/typesconv"
)

func Test_Cov2_IntPtrToSimple_NonNil(t *testing.T) {
	v := 42
	actual := args.Map{"result": typesconv.IntPtrToSimple(&v)}
	expected := args.Map{"result": 42}
	expected.ShouldBeEqual(t, 0, "IntPtrToSimple_NonNil", actual)
}

func Test_Cov2_IntPtrToSimpleDef_NonNil(t *testing.T) {
	v := 10
	actual := args.Map{"result": typesconv.IntPtrToSimpleDef(&v, 99)}
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "IntPtrToSimpleDef_NonNil", actual)
}

func Test_Cov2_IntPtrToDefPtr_NonNil(t *testing.T) {
	v := 10
	r := typesconv.IntPtrToDefPtr(&v, 99)
	actual := args.Map{"result": *r}
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "IntPtrToDefPtr_NonNil", actual)
}

func Test_Cov2_IntPtrDefValFunc_NonNil(t *testing.T) {
	v := 10
	r := typesconv.IntPtrDefValFunc(&v, func() int { return 99 })
	actual := args.Map{"result": *r}
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "IntPtrDefValFunc_NonNil", actual)
}

func Test_Cov2_BytePtrToSimple_NonNil(t *testing.T) {
	v := byte(5)
	actual := args.Map{"result": int(typesconv.BytePtrToSimple(&v))}
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "BytePtrToSimple_NonNil", actual)
}

func Test_Cov2_BytePtrToSimpleDef_NonNil(t *testing.T) {
	v := byte(5)
	actual := args.Map{"result": int(typesconv.BytePtrToSimpleDef(&v, 9))}
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "BytePtrToSimpleDef_NonNil", actual)
}

func Test_Cov2_BytePtrToDefPtr_NonNil(t *testing.T) {
	v := byte(5)
	r := typesconv.BytePtrToDefPtr(&v, 9)
	actual := args.Map{"result": int(*r)}
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "BytePtrToDefPtr_NonNil", actual)
}

func Test_Cov2_BytePtrDefValFunc_NonNil(t *testing.T) {
	v := byte(5)
	r := typesconv.BytePtrDefValFunc(&v, func() byte { return 9 })
	actual := args.Map{"result": int(*r)}
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "BytePtrDefValFunc_NonNil", actual)
}

func Test_Cov2_FloatPtrToSimple_NonNil(t *testing.T) {
	v := float32(3.14)
	actual := args.Map{"result": fmt.Sprintf("%.2f", typesconv.FloatPtrToSimple(&v))}
	expected := args.Map{"result": "3.14"}
	expected.ShouldBeEqual(t, 0, "FloatPtrToSimple_NonNil", actual)
}

func Test_Cov2_FloatPtrToSimpleDef_NonNil(t *testing.T) {
	v := float32(3.14)
	actual := args.Map{"result": fmt.Sprintf("%.2f", typesconv.FloatPtrToSimpleDef(&v, 9.9))}
	expected := args.Map{"result": "3.14"}
	expected.ShouldBeEqual(t, 0, "FloatPtrToSimpleDef_NonNil", actual)
}

func Test_Cov2_FloatPtrToDefPtr_NonNil(t *testing.T) {
	v := float32(3.14)
	r := typesconv.FloatPtrToDefPtr(&v, 9.9)
	actual := args.Map{"result": fmt.Sprintf("%.2f", *r)}
	expected := args.Map{"result": "3.14"}
	expected.ShouldBeEqual(t, 0, "FloatPtrToDefPtr_NonNil", actual)
}

func Test_Cov2_FloatPtrDefValFunc_NonNil(t *testing.T) {
	v := float32(3.14)
	r := typesconv.FloatPtrDefValFunc(&v, func() float32 { return 9.9 })
	actual := args.Map{"result": fmt.Sprintf("%.2f", *r)}
	expected := args.Map{"result": "3.14"}
	expected.ShouldBeEqual(t, 0, "FloatPtrDefValFunc_NonNil", actual)
}

func Test_StringPtrToSimpleDef_NonNil_Cov2(t *testing.T) {
	v := "hello"
	actual := args.Map{"result": typesconv.StringPtrToSimpleDef(&v, "fb")}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "StringPtrToSimpleDef_NonNil", actual)
}

func Test_StringPtrToDefPtr_NonNil_Cov2(t *testing.T) {
	v := "hello"
	r := typesconv.StringPtrToDefPtr(&v, "fb")
	actual := args.Map{"result": *r}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "StringPtrToDefPtr_NonNil", actual)
}

func Test_StringPtrDefValFunc_NonNil_Cov2(t *testing.T) {
	v := "hello"
	r := typesconv.StringPtrDefValFunc(&v, func() string { return "fb" })
	actual := args.Map{"result": *r}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "StringPtrDefValFunc_NonNil", actual)
}

func Test_StringToBool_YES_Cov2(t *testing.T) {
	actual := args.Map{"result": typesconv.StringToBool("YES")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StringToBool_YES", actual)
}

func Test_StringToBool_No_Cov2(t *testing.T) {
	actual := args.Map{"result": typesconv.StringToBool("No")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringToBool_No", actual)
}

func Test_StringToBool_NO_Cov2(t *testing.T) {
	actual := args.Map{"result": typesconv.StringToBool("NO")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringToBool_NO", actual)
}

func Test_StringToBool_false_Cov2(t *testing.T) {
	actual := args.Map{"result": typesconv.StringToBool("false")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringToBool_false", actual)
}

func Test_StringPointerToBool_Empty_Cov2(t *testing.T) {
	s := ""
	actual := args.Map{"result": typesconv.StringPointerToBool(&s)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringPointerToBool_Empty", actual)
}

func Test_StringPointerToBoolPtr_Empty_Cov2(t *testing.T) {
	s := ""
	r := typesconv.StringPointerToBoolPtr(&s)
	actual := args.Map{"result": *r}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringPointerToBoolPtr_Empty", actual)
}

func Test_StringToBoolPtr_NonEmpty_Cov2(t *testing.T) {
	r := typesconv.StringToBoolPtr("yes")
	actual := args.Map{"result": *r}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StringToBoolPtr_NonEmpty", actual)
}
