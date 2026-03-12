package jsoninternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/jsoninternal"
)

func Test_Pretty_AnyTo_String(t *testing.T) {
	// Arrange
	item := map[string]string{"key": "value"}

	// Act
	str, err := jsoninternal.Pretty.AnyTo.String(item)
	safeStr := jsoninternal.Pretty.AnyTo.SafeString(item)
	nilStr := jsoninternal.Pretty.AnyTo.SafeString(nil)

	actual := args.Map{
		"noError":     err == nil,
		"strNotEmpty": str != "",
		"safeNotEmpty": safeStr != "",
		"nilEmpty":    nilStr,
	}
	expected := args.Map{
		"noError":     true,
		"strNotEmpty": true,
		"safeNotEmpty": true,
		"nilEmpty":    "",
	}
	expected.ShouldBeEqual(t, 0, "Pretty_AnyTo_String", actual)
}

func Test_Pretty_AnyTo_PrettyString(t *testing.T) {
	// Arrange
	item := map[string]string{"key": "value"}

	// Act
	pretty, err := jsoninternal.Pretty.AnyTo.PrettyString("", item)
	prettyIndent, err2 := jsoninternal.Pretty.AnyTo.PrettyStringIndent("", "  ", item)
	safePretty := jsoninternal.Pretty.AnyTo.SafePrettyString("", item)
	defaultPretty := jsoninternal.Pretty.AnyTo.PrettyStringDefault(item)
	nilPretty := jsoninternal.Pretty.AnyTo.PrettyStringDefault(nil)

	actual := args.Map{
		"noErr":            err == nil,
		"noErr2":           err2 == nil,
		"prettyNotEmpty":   pretty != "",
		"indentNotEmpty":   prettyIndent != "",
		"safeNotEmpty":     safePretty != "",
		"defaultNotEmpty":  defaultPretty != "",
		"nilEmpty":         nilPretty,
	}
	expected := args.Map{
		"noErr":            true,
		"noErr2":           true,
		"prettyNotEmpty":   true,
		"indentNotEmpty":   true,
		"safeNotEmpty":     true,
		"defaultNotEmpty":  true,
		"nilEmpty":         "",
	}
	expected.ShouldBeEqual(t, 0, "Pretty_AnyTo_PrettyString", actual)
}

func Test_Pretty_Bytes(t *testing.T) {
	// Arrange
	jsonBytes := []byte(`{"key":"value"}`)

	// Act
	safe := jsoninternal.Pretty.Bytes.Safe("", jsonBytes)
	safeDefault := jsoninternal.Pretty.Bytes.SafeDefault(jsonBytes)
	prefix, err := jsoninternal.Pretty.Bytes.Prefix("", jsonBytes)
	indent, err2 := jsoninternal.Pretty.Bytes.Indent("", "  ", jsonBytes)

	actual := args.Map{
		"safeNotEmpty":    safe != "",
		"safeDefNotEmpty": safeDefault != "",
		"noErr":           err == nil,
		"noErr2":          err2 == nil,
		"prefixNotEmpty":  prefix != "",
		"indentNotEmpty":  indent != "",
	}
	expected := args.Map{
		"safeNotEmpty":    true,
		"safeDefNotEmpty": true,
		"noErr":           true,
		"noErr2":          true,
		"prefixNotEmpty":  true,
		"indentNotEmpty":  true,
	}
	expected.ShouldBeEqual(t, 0, "Pretty_Bytes", actual)
}

func Test_Pretty_String(t *testing.T) {
	// Arrange
	jsonStr := `{"key":"value"}`

	// Act
	safe := jsoninternal.Pretty.String.Safe("", jsonStr)
	safeDefault := jsoninternal.Pretty.String.SafeDefault(jsonStr)
	prefix, err := jsoninternal.Pretty.String.Prefix("", jsonStr)
	indent, err2 := jsoninternal.Pretty.String.Indent("", "  ", jsonStr)

	actual := args.Map{
		"safeNotEmpty":    safe != "",
		"safeDefNotEmpty": safeDefault != "",
		"noErr":           err == nil,
		"noErr2":          err2 == nil,
		"prefixNotEmpty":  prefix != "",
		"indentNotEmpty":  indent != "",
	}
	expected := args.Map{
		"safeNotEmpty":    true,
		"safeDefNotEmpty": true,
		"noErr":           true,
		"noErr2":          true,
		"prefixNotEmpty":  true,
		"indentNotEmpty":  true,
	}
	expected.ShouldBeEqual(t, 0, "Pretty_String", actual)
}

func Test_StringJsonConverter(t *testing.T) {
	// Arrange
	item := map[string]string{"key": "value"}

	// Act
	defStr, err := jsoninternal.StringJson.Default(item)
	safeStr := jsoninternal.StringJson.SafeDefault(item)
	prettyStr, err2 := jsoninternal.StringJson.Pretty(item)
	strVal := jsoninternal.StringJson.StringValue("hello")

	actual := args.Map{
		"noErr":            err == nil,
		"noErr2":           err2 == nil,
		"defNotEmpty":      defStr != "",
		"safeNotEmpty":     safeStr != "",
		"prettyNotEmpty":   prettyStr != "",
		"strValNotEmpty":   len(strVal) > 0,
	}
	expected := args.Map{
		"noErr":            true,
		"noErr2":           true,
		"defNotEmpty":      true,
		"safeNotEmpty":     true,
		"prettyNotEmpty":   true,
		"strValNotEmpty":   true,
	}
	expected.ShouldBeEqual(t, 0, "StringJsonConverter", actual)
}
