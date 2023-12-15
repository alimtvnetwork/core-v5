package codegen

import (
	"go/ast"
	"go/parser"
	"go/token"

	"gitlab.com/auk-go/core/errcore"
	"golang.org/x/tools/go/packages"
)

type astReader struct {
	filePath string
	src      any
	node     *ast.File
	parseErr error
	fileSet  *token.FileSet
}

func (it *astReader) Initialize() (*ast.File, error) {
	if it.fileSet != nil {
		return it.node, it.parseErr
	}

	fileSet := token.NewFileSet()

	node, err := parser.ParseFile(
		fileSet,
		it.filePath,
		it.src,
		parser.AllErrors,
	)

	it.fileSet = fileSet
	it.node = node

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

func (it *astReader) InitializeMust() *ast.File {
	node, err := it.Initialize()

	errcore.HandleErr(err)

	return node
}

func (it *astReader) Config() *packages.Config {
	loadConfig := new(packages.Config)
	loadConfig.Mode = globalLoadMode
	loadConfig.Fset = token.NewFileSet()

	return loadConfig
}

func (it *astReader) AllPackages() ([]*packages.Package, error) {
	loadConfig := it.Config()
	imports, loadErr := packages.Load(
		loadConfig,
		"syscall",
	)

	return imports, loadErr
}
