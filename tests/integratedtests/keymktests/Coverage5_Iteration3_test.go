package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/keymk"
)

// Test_Cov5_CompileSingleItem_WithBrackets covers
// keymk/KeyCompiler.go L142-145: compileSingleItem with IsUseBrackets=true.
func Test_Cov5_CompileSingleItem_WithBrackets(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Create(
		keymk.BracketJoinerOption,
		"root",
	)
	key.AppendChain("a", "b")

	// Act
	result := key.Compile()

	// Assert
	// With BracketJoinerOption, items are wrapped in [ ]
	coretests.ShouldNotBeEmptyString(t, result)
}

// Test_Cov5_CompileCompleteAdditional_Empty covers
// keymk/KeyCompiler.go L271-273: empty items in compileCompleteAdditional.
func Test_Cov5_CompileCompleteAdditional_Empty(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Create(
		keymk.JoinerOption,
		"root",
	)

	// Act — Compile with no additional items triggers compileCompleteAdditional with empty
	result := key.Compile()

	// Assert
	coretests.ShouldBeEqual(t, "root", result)
}

// Test_Cov5_CompileCompleteAdditionalStrings_Empty covers
// keymk/KeyCompiler.go L285-287: empty items in compileCompleteAdditionalStrings.
func Test_Cov5_CompileCompleteAdditionalStrings_Empty(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Create(
		keymk.JoinerOption,
		"root",
	)

	// Act
	result := key.CompileStrings()

	// Assert
	coretests.ShouldBeEqual(t, "root", result)
}

// Test_Cov5_AppendStringsWithBaseAnyItems_SkipEmpty covers
// keymk/appendStringsWithBaseAnyItems.go L13-14: skip empty entry.
func Test_Cov5_AppendStringsWithBaseAnyItems_SkipEmpty(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Create(
		keymk.JoinerOption,
		"root",
	)
	key.AppendChain("a", "", "b")

	// Act
	result := key.Compile()

	// Assert
	// IsSkipEmptyEntry=true, so empty string is skipped
	coretests.ShouldBeEqual(t, "root-a-b", result)
}
