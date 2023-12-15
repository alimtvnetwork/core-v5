package codegen

import (
	"go/ast"
	"go/parser"
	"go/token"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/iserror"
	"golang.org/x/tools/go/packages"
)

type AstReader struct {
	filePath string
	src      interface{}
	node     *ast.File
	fullCode string
	parseErr error
	fileSet  *token.FileSet
	mode     parser.Mode
}

func (it *AstReader) Initialize() (*ast.File, error) {
	if it.fileSet != nil {
		return it.node, it.parseErr
	}

	fileSet := token.NewFileSet()

	node, err := parser.ParseFile(
		fileSet,
		it.filePath,
		it.src,
		it.mode,
	)

	it.fileSet = fileSet
	it.node = node
	it.fullCode = chmodhelper.SimpleFileWriter.FileWriter.

	if err != nil {
		finalErr := errcore.ParsingFailed.MsgCsvRefError(
			err.Error(),
			it.filePath,
		)

		it.parseErr = finalErr

		return node, finalErr
	}

	return node, err
}

func (it *AstReader) InitializeMust() *ast.File {
	node, err := it.Initialize()

	errcore.HandleErr(err)

	return node
}

func (it *AstReader) Config() *packages.Config {
	loadConfig := new(packages.Config)
	loadConfig.Mode = globalLoadMode
	loadConfig.Fset = token.NewFileSet()

	return loadConfig
}

func (it *AstReader) AllPackages() ([]*packages.Package, error) {
	loadConfig := it.Config()
	imports, loadErr := packages.Load(
		loadConfig,
		"syscall",
	)

	return imports, loadErr
}

func (it *AstReader) NodesMap() (args.Map, error) {
	node, err := it.Initialize()

	if iserror.Defined(err) {
		return args.Map{}, err
	}

	// okay
	// Collect the struct types in this slice.
	curMap := make(map[string]interface{}, 100)

	// Use the Inspect function to walk AST looking for struct
	// type nodes.
	ast.Inspect(
		node, func(n ast.Node) bool {
			start := n.Pos() - 1
			end := n.End() - 1

			curMap[n.]

return true
})
}
