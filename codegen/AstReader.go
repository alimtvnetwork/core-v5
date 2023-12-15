package codegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

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

	var fileErr error

	it.fileSet = fileSet
	it.node = node
	it.fullCode, fileErr = chmodhelper.
		SimpleFileWriter.
		FileReader.
		Read(it.filePath)

	combineErr := errcore.MergeErrors(err, fileErr)

	if err != nil {
		finalErr := errcore.ParsingFailed.MsgCsvRefError(
			combineErr.Error(),
			it.filePath,
		)

		it.parseErr = finalErr

		return node, finalErr
	}

	return node, err
}

func (it *AstReader) HasError() bool {
	return it != nil && it.parseErr != nil
}

func (it *AstReader) IsEmptyError() bool {
	return it == nil || it.parseErr == nil
}

func (it *AstReader) IsValid() bool {
	return it != nil && it.parseErr == nil
}

func (it *AstReader) IsInvalid() bool {
	return !it.IsValid()
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

func (it *AstReader) Substring(start, end int) (string, error) {
	if it.HasError() {
		return "", it.parseErr
	}

	return it.fullCode[start:end], nil
}
func (it *AstReader) SubstringByNode(n ast.Node) (string, error) {
	if it.HasError() {
		return "", it.parseErr
	}

	if n == nil {
		return "", errcore.FailedToParseType.ErrorNoRefs("node is nil")
	}

	start := n.Pos() - 1
	end := n.End() - 1

	return it.fullCode[start:end], nil
}

func (it *AstReader) NodesMap() (map[string]args.Map, error) {
	node, err := it.Initialize()

	if iserror.Defined(err) {
		return map[string]args.Map{}, err
	}

	// okay
	// Collect the struct types in this slice.
	curMap := make(map[string]args.Map, 30)
	var rawErrSlice errcore.RawErrCollection

	// Use the Inspect function to walk AST looking for struct
	// type nodes.
	ast.Inspect(
		node, func(n ast.Node) bool {
			if n == nil {
				return true
			}

			toString, subsErr := it.SubstringByNode(n)

			if subsErr != nil {
				rawErrSlice.Add(subsErr)

				return true
			}

			typeName := fmt.Sprintf("%T", n)
			typeName = strings.ReplaceAll(typeName, "->", "")

			m, isFound := curMap[typeName]

			if isFound {
				m[toString] = n
			} else {
				curMap[typeName] = make(map[string]interface{}, 100)
				curMap[typeName][toString] = n
			}

			return true
		},
	)

	return curMap, rawErrSlice.CompiledError()
}

func (it *AstReader) NestedNodesMap() (map[string]args.Map, error) {
	node, err := it.Initialize()

	if iserror.Defined(err) {
		return map[string]args.Map{}, err
	}

	// okay
	// Collect the struct types in this slice.
	curMap := make(map[string]args.Map, 30)
	var rawErrSlice errcore.RawErrCollection

	// Use the Inspect function to walk AST looking for struct
	// type nodes.
	ast.Inspect(
		node, func(n ast.Node) bool {
			if n == nil {
				return true
			}

			toString, subsErr := it.SubstringByNode(n)

			if subsErr != nil {
				rawErrSlice.Add(subsErr)

				return true
			}

			typeName := it.TypeName(n)

			m, isFound := curMap[typeName]

			if isFound {
				m[toString] = n
			} else {
				curMap[typeName] = make(map[string]interface{}, 100)
				curMap[typeName][toString] = n
			}

			return true
		},
	)

	return curMap, rawErrSlice.CompiledError()
}

func (it *AstReader) TypeName(n ast.Node) string {
	typeName := fmt.Sprintf("%T", n)

	if len(typeName) > 3 {
		return typeName[1 : len(typeName)-2]
	}

	return typeName
}

func (it *AstReader) NestedNodesMap() (map[string]args.Map, error) {

}
func (it *AstReader) StructTypes() ([]*ast.StructType, error) {
	node, err := it.Initialize()

	if iserror.Defined(err) {
		return nil, err
	}

	// okay
	var rawErrSlice errcore.RawErrCollection
	var structTypes []*ast.StructType

	// Use the Inspect function to walk AST looking for struct
	// type nodes.
	ast.Inspect(
		node, func(n ast.Node) bool {
			if x, isOkay := n.(*ast.StructType); isOkay {
				structTypes = append(structTypes, x)
			}

			return true
		},
	)

	return structTypes, rawErrSlice.CompiledError()
}
