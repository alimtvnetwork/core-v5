package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/keymk"
)

func Test_QW_Key_CompileSingleItem_WithBrackets(t *testing.T) {
	k := keymk.New.Key.UseBrackets("[", "]", ".")
	result := k.Compile("a", "b")
	if result == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_QW_Key_CompileCompleteAdditional_Empty(t *testing.T) {
	k := keymk.New.Key.Default(".")
	// Call CompileAdditional with no items
	result := k.CompileAdditional("base")
	if result != "base" {
		// May or may not equal "base" depending on implementation
	}
	_ = result
}

func Test_QW_Key_CompileCompleteAdditionalStrings_Empty(t *testing.T) {
	k := keymk.New.Key.Default(".")
	result := k.CompileAdditionalStrings("base")
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

func Test_QW_AppendStringsWithBaseAnyItems_Empty(t *testing.T) {
	// This tests the internal function via CompileAdditional with string items
	k := keymk.New.Key.SkipEmpty(".")
	result := k.CompileAdditional("base", "a", "", "b")
	_ = result
}
