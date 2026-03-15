package jsoninternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/jsoninternal"
)

// ── AnyTo — String (not pretty) ──

func Test_Cov4_AnyTo_String(t *testing.T) {
	// Act
	result, err := jsoninternal.Pretty.AnyTo.String(map[string]int{"a": 1})
	emptyResult, _ := jsoninternal.Pretty.AnyTo.String(nil)

	// Assert
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil, "empty": emptyResult}
	expected := args.Map{"notEmpty": true, "noErr": true, "empty": "null"}
	expected.ShouldBeEqual(t, 0, "AnyTo String", actual)
}

func Test_Cov4_AnyTo_SafeString(t *testing.T) {
	// Act
	result := jsoninternal.Pretty.AnyTo.SafeString(map[string]int{"a": 1})
	emptyResult := jsoninternal.Pretty.AnyTo.SafeString(nil)

	// Assert
	actual := args.Map{"notEmpty": result != "", "empty": emptyResult}
	expected := args.Map{"notEmpty": true, "empty": "null"}
	expected.ShouldBeEqual(t, 0, "AnyTo SafeString", actual)
}

// ── AnyTo — PrettyString ──

func Test_Cov4_AnyTo_PrettyString(t *testing.T) {
	// Act
	result, err := jsoninternal.Pretty.AnyTo.PrettyString("", map[string]int{"a": 1})
	emptyResult, _ := jsoninternal.Pretty.AnyTo.PrettyString("", nil)

	// Assert
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil, "empty": emptyResult}
	expected := args.Map{"notEmpty": true, "noErr": true, "empty": "null"}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyString", actual)
}

func Test_Cov4_AnyTo_SafePrettyString(t *testing.T) {
	// Act
	result := jsoninternal.Pretty.AnyTo.SafePrettyString("", map[string]int{"a": 1})
	emptyResult := jsoninternal.Pretty.AnyTo.SafePrettyString("", nil)

	// Assert
	actual := args.Map{"notEmpty": result != "", "empty": emptyResult}
	expected := args.Map{"notEmpty": true, "empty": "null"}
	expected.ShouldBeEqual(t, 0, "AnyTo SafePrettyString", actual)
}

func Test_Cov4_AnyTo_PrettyStringDefault(t *testing.T) {
	// Act
	result := jsoninternal.Pretty.AnyTo.PrettyStringDefault(map[string]int{"a": 1})
	emptyResult := jsoninternal.Pretty.AnyTo.PrettyStringDefault(nil)

	// Assert
	actual := args.Map{"notEmpty": result != "", "empty": emptyResult}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringDefault", actual)
}

// ── AnyTo — PrettyStringIndent ──

func Test_Cov4_AnyTo_PrettyStringIndent(t *testing.T) {
	// Act
	result, err := jsoninternal.Pretty.AnyTo.PrettyStringIndent("", "\t", map[string]int{"a": 1})
	emptyResult, _ := jsoninternal.Pretty.AnyTo.PrettyStringIndent("", "\t", nil)

	// Assert
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil, "empty": emptyResult}
	expected := args.Map{"notEmpty": true, "noErr": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringIndent", actual)
}

// ── Bytes — SafeDefault / Indent ──

func Test_Cov4_Bytes_SafeDefault(t *testing.T) {
	// Act
	result := jsoninternal.Pretty.Bytes.SafeDefault([]byte(`{"a":1}`))

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Bytes SafeDefault", actual)
}

func Test_Cov4_Bytes_Indent(t *testing.T) {
	// Act
	result, err := jsoninternal.Pretty.Bytes.Indent("", "\t", []byte(`{"a":1}`))

	// Assert
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Bytes Indent", actual)
}
