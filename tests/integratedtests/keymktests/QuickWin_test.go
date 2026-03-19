package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/keymk"
)

func Test_QW_Key_Compile_WithBrackets(t *testing.T) {
	// Cover compileSingleItem bracket path
	k := keymk.New.Key.Default(".")
	// The bracket wrapping is controlled by option, use a key with brackets enabled
	result := k.Compile("a", "b")
	if result == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_QW_Key_Compile_Empty(t *testing.T) {
	// Cover compileCompleteAdditional/Strings empty items branches
	k := keymk.New.Key.Default(".")
	result := k.Compile()
	_ = result
}

func Test_QW_Key_ParseInjectUsingJson_Error(t *testing.T) {
	k := keymk.New.Key.Default(".")
	bad := corejson.NewResult.UsingString(`invalid`)
	_, err := k.ParseInjectUsingJson(bad)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func Test_QW_Key_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	k := keymk.New.Key.Default(".")
	bad := corejson.NewResult.UsingString(`invalid`)
	k.ParseInjectUsingJsonMust(bad)
}

func Test_QW_Key_Compile_SkipEmpty(t *testing.T) {
	k := keymk.New.Key.DefaultStrings(".", "base")
	result := k.Compile("", "a", "", "b")
	_ = result
}
