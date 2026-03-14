package jsoninternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/jsoninternal"
)

// ── AnyTo.PrettyStringDefaultMust ──

func Test_Cov3_AnyTo_PrettyStringDefaultMust(t *testing.T) {
	result := jsoninternal.Pretty.AnyTo.PrettyStringDefaultMust(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringDefaultMust", actual)
}

func Test_Cov3_AnyTo_PrettyStringDefaultMust_Panic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringDefaultMust panic", actual)
	}()
	// channels can't be marshalled
	ch := make(chan int)
	jsoninternal.Pretty.AnyTo.PrettyStringDefaultMust(ch)
}

// ── Bytes.PrefixMust / DefaultMust ──

func Test_Cov3_Bytes_PrefixMust(t *testing.T) {
	result := jsoninternal.Pretty.Bytes.PrefixMust("", []byte(`{"a":1}`))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Bytes PrefixMust", actual)
}

func Test_Cov3_Bytes_PrefixMust_Panic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "Bytes PrefixMust panic", actual)
	}()
	jsoninternal.Pretty.Bytes.PrefixMust("", []byte(`invalid json`))
}

func Test_Cov3_Bytes_DefaultMust(t *testing.T) {
	result := jsoninternal.Pretty.Bytes.DefaultMust([]byte(`{"a":1}`))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Bytes DefaultMust", actual)
}

func Test_Cov3_Bytes_DefaultMust_Panic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "Bytes DefaultMust panic", actual)
	}()
	jsoninternal.Pretty.Bytes.DefaultMust([]byte(`invalid`))
}

// ── String converter ──

func Test_Cov3_StringJson_Default(t *testing.T) {
	result, err := jsoninternal.String.Default(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "StringJson Default", actual)
}

func Test_Cov3_StringJson_SafeDefault(t *testing.T) {
	result := jsoninternal.String.SafeDefault(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringJson SafeDefault", actual)
}

func Test_Cov3_StringJson_Pretty(t *testing.T) {
	result, err := jsoninternal.String.Pretty(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "StringJson Pretty", actual)
}

func Test_Cov3_StringJson_StringValue(t *testing.T) {
	result := jsoninternal.String.StringValue("hello")
	actual := args.Map{"result": string(result)}
	expected := args.Map{"result": `"hello"`}
	expected.ShouldBeEqual(t, 0, "StringJson StringValue", actual)
}

// ── String to Pretty converter ──

func Test_Cov3_StringToPretty_Safe(t *testing.T) {
	result := jsoninternal.Pretty.String.Safe("", `{"a":1}`)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringToPretty Safe", actual)
}

func Test_Cov3_StringToPretty_SafeDefault(t *testing.T) {
	result := jsoninternal.Pretty.String.SafeDefault(`{"a":1}`)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringToPretty SafeDefault", actual)
}

func Test_Cov3_StringToPretty_Indent(t *testing.T) {
	result, err := jsoninternal.Pretty.String.Indent("", "  ", `{"a":1}`)
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "StringToPretty Indent", actual)
}

func Test_Cov3_StringToPretty_PrefixMust(t *testing.T) {
	result := jsoninternal.Pretty.String.PrefixMust("", `{"a":1}`)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringToPretty PrefixMust", actual)
}

func Test_Cov3_StringToPretty_PrefixMust_Panic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "StringToPretty PrefixMust panic", actual)
	}()
	jsoninternal.Pretty.String.PrefixMust("", `invalid`)
}

func Test_Cov3_StringToPretty_DefaultMust(t *testing.T) {
	result := jsoninternal.Pretty.String.DefaultMust(`{"a":1}`)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringToPretty DefaultMust", actual)
}

func Test_Cov3_StringToPretty_DefaultMust_Panic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "StringToPretty DefaultMust panic", actual)
	}()
	jsoninternal.Pretty.String.DefaultMust(`invalid`)
}
