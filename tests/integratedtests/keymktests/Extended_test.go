package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/keymk"
)

// TestKey_Default verifies default key creation.
func TestKey_Default(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "child1", "child2")

	// Act
	compiled := key.Compile()

	// Assert
	if compiled != "root-child1-child2" {
		t.Errorf("expected 'root-child1-child2', got '%s'", compiled)
	}
}

// TestKey_DefaultMain verifies main-only key.
func TestKey_DefaultMain(t *testing.T) {
	key := keymk.NewKey.DefaultMain("main")
	if key.Compile() != "main" {
		t.Errorf("expected 'main', got '%s'", key.Compile())
	}
}

// TestKey_DefaultStrings verifies string-based construction.
func TestKey_DefaultStrings(t *testing.T) {
	key := keymk.NewKey.DefaultStrings("root", "a", "b")
	if key.Compile() != "root-a-b" {
		t.Errorf("expected 'root-a-b', got '%s'", key.Compile())
	}
}

// TestKey_CurlyBraces verifies curly brace option.
func TestKey_CurlyBraces(t *testing.T) {
	key := keymk.NewKey.Curly("root", "a")
	compiled := key.Compile()
	if compiled != "{root}-{a}" {
		t.Errorf("expected '{root}-{a}', got '%s'", compiled)
	}
}

// TestKey_SquareBrackets verifies square bracket option.
func TestKey_SquareBrackets(t *testing.T) {
	key := keymk.NewKey.SquareBrackets("root", "a")
	compiled := key.Compile()
	if compiled != "[root]-[a]" {
		t.Errorf("expected '[root]-[a]', got '%s'", compiled)
	}
}

// TestKey_Parenthesis verifies parenthesis option.
func TestKey_Parenthesis(t *testing.T) {
	key := keymk.NewKey.Parenthesis("root", "a")
	compiled := key.Compile()
	if compiled != "(root)-(a)" {
		t.Errorf("expected '(root)-(a)', got '%s'", compiled)
	}
}

// TestKey_AppendChain verifies append chain.
func TestKey_AppendChain(t *testing.T) {
	key := keymk.NewKey.Default("root")
	key.AppendChain("a", "b")
	if key.Compile() != "root-a-b" {
		t.Errorf("expected 'root-a-b', got '%s'", key.Compile())
	}
}

// TestKey_AppendChainStrings verifies string append.
func TestKey_AppendChainStrings(t *testing.T) {
	key := keymk.NewKey.Default("root")
	key.AppendChainStrings("x", "y")
	if key.Compile() != "root-x-y" {
		t.Errorf("expected 'root-x-y', got '%s'", key.Compile())
	}
}

// TestKey_SkipEmpty verifies empty entry skipping.
func TestKey_SkipEmpty(t *testing.T) {
	key := keymk.NewKey.Default("root", "", "b")
	compiled := key.Compile()
	if compiled != "root-b" {
		t.Errorf("expected 'root-b', got '%s'", compiled)
	}
}

// TestKey_Length verifies chain length.
func TestKey_Length(t *testing.T) {
	key := keymk.NewKey.Default("root", "a", "b")
	if key.Length() != 2 {
		t.Errorf("expected 2, got %d", key.Length())
	}
}

// TestKey_NilLength verifies nil key length.
func TestKey_NilLength(t *testing.T) {
	var key *keymk.Key
	if key.Length() != 0 {
		t.Error("nil length should be 0")
	}
}

// TestKey_IsEmpty verifies empty check.
func TestKey_IsEmpty(t *testing.T) {
	key := keymk.NewKey.Default("")
	if !key.IsEmpty() {
		t.Error("empty main with no chains should be empty")
	}
}

// TestKey_MainName verifies main name getter.
func TestKey_MainName(t *testing.T) {
	key := keymk.NewKey.Default("myroot")
	if key.MainName() != "myroot" {
		t.Error("expected 'myroot'")
	}
}

// TestKey_KeyChains verifies chain getter.
func TestKey_KeyChains(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	chains := key.KeyChains()
	if len(chains) != 1 {
		t.Errorf("expected 1 chain, got %d", len(chains))
	}
}

// TestKey_NilKeyChains verifies nil chains.
func TestKey_NilKeyChains(t *testing.T) {
	var key *keymk.Key
	if key.KeyChains() != nil {
		t.Error("nil key should return nil chains")
	}
}

// TestKey_AllRawItems verifies all raw items.
func TestKey_AllRawItems(t *testing.T) {
	key := keymk.NewKey.Default("root", "a", "b")
	items := key.AllRawItems()
	if len(items) != 3 {
		t.Errorf("expected 3, got %d", len(items))
	}
}

// TestKey_NilAllRawItems verifies nil raw items.
func TestKey_NilAllRawItems(t *testing.T) {
	var key *keymk.Key
	if key.AllRawItems() != nil {
		t.Error("nil should return nil")
	}
}

// TestKey_HasInChains verifies chain search.
func TestKey_HasInChains(t *testing.T) {
	key := keymk.NewKey.Default("root", "a", "b")
	if !key.HasInChains("a") {
		t.Error("should have 'a'")
	}
	if key.HasInChains("c") {
		t.Error("should not have 'c'")
	}
}

// TestKey_NilHasInChains verifies nil chain search.
func TestKey_NilHasInChains(t *testing.T) {
	var key *keymk.Key
	if key.HasInChains("x") {
		t.Error("nil should return false")
	}
}

// TestKey_Finalized verifies finalization.
func TestKey_Finalized(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	key.Finalized("b")
	if !key.IsComplete() {
		t.Error("should be complete")
	}
	if key.CompiledChain() != "root-a-b" {
		t.Errorf("expected 'root-a-b', got '%s'", key.CompiledChain())
	}
}

// TestKey_CompiledChain_NotComplete verifies incomplete returns empty.
func TestKey_CompiledChain_NotComplete(t *testing.T) {
	key := keymk.NewKey.Default("root")
	if key.CompiledChain() != "" {
		t.Error("incomplete should return empty")
	}
}

// TestKey_ClonePtr verifies clone.
func TestKey_ClonePtr(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	cloned := key.ClonePtr()
	if cloned.Compile() != key.Compile() {
		t.Error("clone should match original")
	}
}

// TestKey_NilClonePtr verifies nil clone.
func TestKey_NilClonePtr(t *testing.T) {
	var key *keymk.Key
	if key.ClonePtr() != nil {
		t.Error("nil clone should be nil")
	}
}

// TestKey_String verifies String method.
func TestKey_String(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	if key.String() != "root-a" {
		t.Errorf("expected 'root-a', got '%s'", key.String())
	}
}

// TestKey_Name verifies Name method.
func TestKey_Name(t *testing.T) {
	key := keymk.NewKey.Default("root")
	if key.Name() != "root" {
		t.Error("expected 'root'")
	}
}

// TestKey_KeyCompiled verifies KeyCompiled.
func TestKey_KeyCompiled(t *testing.T) {
	key := keymk.NewKey.Default("r", "a")
	if key.KeyCompiled() != "r-a" {
		t.Errorf("expected 'r-a', got '%s'", key.KeyCompiled())
	}
}

// TestKey_CompileStrings verifies CompileStrings.
func TestKey_CompileStrings(t *testing.T) {
	key := keymk.NewKey.Default("root")
	r := key.CompileStrings("x", "y")
	if r != "root-x-y" {
		t.Errorf("expected 'root-x-y', got '%s'", r)
	}
}

// TestKey_CompileDefault verifies CompileDefault.
func TestKey_CompileDefault(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	if key.CompileDefault() != "root-a" {
		t.Errorf("expected 'root-a', got '%s'", key.CompileDefault())
	}
}

// TestKey_IntRange verifies IntRange.
func TestKey_IntRange(t *testing.T) {
	key := keymk.NewKey.Default("item")
	r := key.IntRange(0, 2)
	if len(r) != 3 {
		t.Errorf("expected 3 items, got %d", len(r))
	}
	if r[0] != "item-0" {
		t.Errorf("expected 'item-0', got '%s'", r[0])
	}
}

// TestKey_IntRangeEnding verifies IntRangeEnding.
func TestKey_IntRangeEnding(t *testing.T) {
	key := keymk.NewKey.Default("item")
	r := key.IntRangeEnding(1)
	if len(r) != 2 {
		t.Errorf("expected 2 items, got %d", len(r))
	}
}

// TestKey_AppendChainKeys verifies keys append.
func TestKey_AppendChainKeys(t *testing.T) {
	key1 := keymk.NewKey.Default("root")
	key2 := keymk.NewKey.Default("sub", "a")
	key1.AppendChainKeys(key2)
	if key1.Compile() != "root-sub-a" {
		t.Errorf("expected 'root-sub-a', got '%s'", key1.Compile())
	}
}

// TestKey_AppendChainKeys_Nil verifies nil key skipping.
func TestKey_AppendChainKeys_Nil(t *testing.T) {
	key := keymk.NewKey.Default("root")
	key.AppendChainKeys(nil)
	if key.Compile() != "root" {
		t.Errorf("expected 'root', got '%s'", key.Compile())
	}
}

// TestKey_ConcatNewUsingKeys verifies concat.
func TestKey_ConcatNewUsingKeys(t *testing.T) {
	key1 := keymk.NewKey.Default("root")
	key2 := keymk.NewKey.Default("sub")
	result := key1.ConcatNewUsingKeys(key2)
	if result.Compile() != "root-sub" {
		t.Errorf("expected 'root-sub', got '%s'", result.Compile())
	}
}

// TestKey_CompileKeys verifies compile with extra keys.
func TestKey_CompileKeys(t *testing.T) {
	key := keymk.NewKey.Default("root")
	sub := keymk.NewKey.Default("sub", "x")
	r := key.CompileKeys(sub)
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestKey_CompileKeys_Empty verifies empty compile keys.
func TestKey_CompileKeys_Empty(t *testing.T) {
	key := keymk.NewKey.Default("root")
	r := key.CompileKeys()
	if r != "root" {
		t.Errorf("expected 'root', got '%s'", r)
	}
}

// TestKey_JoinUsingJoiner verifies custom joiner.
func TestKey_JoinUsingJoiner(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	r := key.JoinUsingJoiner("/", "b")
	if r != "root/a/b" {
		t.Errorf("expected 'root/a/b', got '%s'", r)
	}
}

// TestKey_CompileReplaceCurlyKeyMap verifies curly replace.
func TestKey_CompileReplaceCurlyKeyMap(t *testing.T) {
	key := keymk.NewKey.Curly("root", "name")
	compiled := key.CompileReplaceCurlyKeyMap(map[string]string{
		"root": "myroot",
		"name": "myname",
	})
	if compiled == "" {
		t.Error("should return non-empty")
	}
}

// TestOption_ClonePtr verifies option clone.
func TestOption_ClonePtr(t *testing.T) {
	opt := keymk.JoinerOption.ClonePtr()
	if opt == nil {
		t.Error("clone should not be nil")
	}
	if opt.Joiner != keymk.JoinerOption.Joiner {
		t.Error("joiner should match")
	}
}

// TestOption_NilClonePtr verifies nil option clone.
func TestOption_NilClonePtr(t *testing.T) {
	var opt *keymk.Option
	if opt.ClonePtr() != nil {
		t.Error("nil clone should be nil")
	}
}

// TestOption_Clone verifies value clone.
func TestOption_Clone(t *testing.T) {
	opt := keymk.JoinerOption.Clone()
	if opt.Joiner != keymk.JoinerOption.Joiner {
		t.Error("joiner should match")
	}
}

// TestOption_IsAddEntryRegardlessOfEmptiness verifies option method.
func TestOption_IsAddEntryRegardlessOfEmptiness(t *testing.T) {
	opt := &keymk.Option{IsSkipEmptyEntry: false}
	if !opt.IsAddEntryRegardlessOfEmptiness() {
		t.Error("should be true when not skipping")
	}
	opt2 := &keymk.Option{IsSkipEmptyEntry: true}
	if opt2.IsAddEntryRegardlessOfEmptiness() {
		t.Error("should be false when skipping")
	}
	var nilOpt *keymk.Option
	if nilOpt.IsAddEntryRegardlessOfEmptiness() {
		t.Error("nil should return false")
	}
}

// TestKey_PathTemplate verifies path template.
func TestKey_PathTemplate(t *testing.T) {
	key := keymk.NewKey.PathTemplate("root", "sub")
	r := key.Compile()
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestKey_PathTemplateDefault verifies default path template.
func TestKey_PathTemplateDefault(t *testing.T) {
	key := keymk.NewKey.PathTemplateDefault("sub")
	if key.Compile() == "" {
		t.Error("should return non-empty")
	}
}

// TestKey_Strings verifies Strings method.
func TestKey_Strings(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	s := key.Strings()
	if len(s) != 2 {
		t.Errorf("expected 2, got %d", len(s))
	}
}
