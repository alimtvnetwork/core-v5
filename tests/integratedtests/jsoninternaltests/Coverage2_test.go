package jsoninternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/jsoninternal"
)

// ── bytesToPrettyConvert ──

func Test_Cov2_Pretty_Bytes_Safe(t *testing.T) {
	input := []byte(`{"key":"value"}`)
	result := jsoninternal.Pretty.Bytes.Safe("", input)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.Safe returns correct value -- with args", actual)
}

func Test_Cov2_Pretty_Bytes_SafeDefault(t *testing.T) {
	input := []byte(`{"key":"value"}`)
	result := jsoninternal.Pretty.Bytes.SafeDefault(input)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.SafeDefault returns correct value -- with args", actual)
}

func Test_Cov2_Pretty_Bytes_Prefix(t *testing.T) {
	input := []byte(`{"key":"value"}`)
	result, err := jsoninternal.Pretty.Bytes.Prefix("  ", input)
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.Prefix returns correct value -- with args", actual)
}

func Test_Cov2_Pretty_Bytes_Indent(t *testing.T) {
	input := []byte(`{"key":"value"}`)
	result, err := jsoninternal.Pretty.Bytes.Indent("", "\t", input)
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.Indent returns correct value -- with args", actual)
}

func Test_Cov2_Pretty_Bytes_PrefixMust(t *testing.T) {
	input := []byte(`{"key":"value"}`)
	result := jsoninternal.Pretty.Bytes.PrefixMust("", input)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.PrefixMust returns correct value -- with args", actual)
}

func Test_Cov2_Pretty_Bytes_DefaultMust(t *testing.T) {
	input := []byte(`{"key":"value"}`)
	result := jsoninternal.Pretty.Bytes.DefaultMust(input)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.DefaultMust returns correct value -- with args", actual)
}

// ── stringToPrettyConvert ──

func Test_Cov2_Pretty_String_Safe(t *testing.T) {
	result := jsoninternal.Pretty.String.Safe("", `{"key":"value"}`)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.String.Safe returns correct value -- with args", actual)
}

func Test_Cov2_Pretty_String_SafeDefault(t *testing.T) {
	result := jsoninternal.Pretty.String.SafeDefault(`{"key":"value"}`)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.String.SafeDefault returns correct value -- with args", actual)
}

func Test_Cov2_Pretty_String_Prefix(t *testing.T) {
	result, err := jsoninternal.Pretty.String.Prefix("", `{"key":"value"}`)
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Pretty.String.Prefix returns correct value -- with args", actual)
}

func Test_Cov2_Pretty_String_Indent(t *testing.T) {
	result, err := jsoninternal.Pretty.String.Indent("", "\t", `{"key":"value"}`)
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Pretty.String.Indent returns correct value -- with args", actual)
}

func Test_Cov2_Pretty_String_PrefixMust(t *testing.T) {
	result := jsoninternal.Pretty.String.PrefixMust("", `{"key":"value"}`)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.String.PrefixMust returns correct value -- with args", actual)
}

func Test_Cov2_Pretty_String_DefaultMust(t *testing.T) {
	result := jsoninternal.Pretty.String.DefaultMust(`{"key":"value"}`)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.String.DefaultMust returns correct value -- with args", actual)
}

// ── anyToConvert ──

func Test_Cov2_AnyTo_SafeString(t *testing.T) {
	result := jsoninternal.Pretty.AnyTo.SafeString(map[string]string{"k": "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SafeString returns correct value -- with args", actual)
}

func Test_Cov2_AnyTo_String(t *testing.T) {
	result, err := jsoninternal.Pretty.AnyTo.String(map[string]string{"k": "v"})
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.String returns correct value -- with args", actual)
}

func Test_Cov2_AnyTo_PrettyString(t *testing.T) {
	result, err := jsoninternal.Pretty.AnyTo.PrettyString("", map[string]string{"k": "v"})
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyString returns correct value -- with args", actual)
}

func Test_Cov2_AnyTo_PrettyStringIndent(t *testing.T) {
	result, err := jsoninternal.Pretty.AnyTo.PrettyStringIndent("", "\t", map[string]string{"k": "v"})
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyStringIndent returns correct value -- with args", actual)
}

func Test_Cov2_AnyTo_SafePrettyString(t *testing.T) {
	result := jsoninternal.Pretty.AnyTo.SafePrettyString("", map[string]string{"k": "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SafePrettyString returns correct value -- with args", actual)
}

func Test_Cov2_AnyTo_PrettyStringDefault(t *testing.T) {
	result := jsoninternal.Pretty.AnyTo.PrettyStringDefault(map[string]string{"k": "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyStringDefault returns correct value -- with args", actual)
}

func Test_Cov2_AnyTo_PrettyStringDefaultMust(t *testing.T) {
	result := jsoninternal.Pretty.AnyTo.PrettyStringDefaultMust(map[string]string{"k": "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyStringDefaultMust returns correct value -- with args", actual)
}

// ── stringJsonConverter ──

func Test_Cov2_StringJson_SafeDefault(t *testing.T) {
	result := jsoninternal.String.SafeDefault(map[string]string{"k": "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringJson.SafeDefault returns correct value -- with args", actual)
}

func Test_Cov2_StringJson_Default(t *testing.T) {
	result, err := jsoninternal.String.Default(map[string]string{"k": "v"})
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "StringJson.Default returns correct value -- with args", actual)
}

func Test_Cov2_StringJson_Pretty(t *testing.T) {
	result, err := jsoninternal.String.Pretty(map[string]string{"k": "v"})
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "StringJson.Pretty returns correct value -- with args", actual)
}

func Test_Cov2_StringJson_StringValue(t *testing.T) {
	result := jsoninternal.String.StringValue("test")
	actual := args.Map{"notEmpty": len(result) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringJson.StringValue returns correct value -- with args", actual)
}
