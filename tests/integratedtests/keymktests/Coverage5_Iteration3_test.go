package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/keymk"
)

// Test_Cov5_CompileSingleItem_WithBrackets covers
// keymk/KeyCompiler.go L143-144: compileSingleItem with IsUseBrackets=true.
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
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[root]-[a]-[b]"}
	expected.ShouldBeEqual(t, 0, "compileSingleItem with brackets", actual)
}

// Test_Cov5_CompileCompleteAdditional_Empty covers
// keymk/KeyCompiler.go L271-273: empty items in compileCompleteAdditional.
func Test_Cov5_CompileCompleteAdditional_Empty(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Create(
		keymk.JoinerOption,
		"root",
	)

	// Act
	result := key.Compile()

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": "root"}
	expected.ShouldBeEqual(t, 0, "compileCompleteAdditional empty items", actual)
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
	actual := args.Map{"result": result}
	expected := args.Map{"result": "root"}
	expected.ShouldBeEqual(t, 0, "compileCompleteAdditionalStrings empty items", actual)
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
	actual := args.Map{"result": result}
	expected := args.Map{"result": "root-a-b"}
	expected.ShouldBeEqual(t, 0, "appendStringsWithBaseAnyItems skip empty", actual)
}
