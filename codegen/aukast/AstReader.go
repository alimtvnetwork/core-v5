package aukast

import (
	"go/ast"
	"go/parser"
	"go/token"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/isany"
	"golang.org/x/tools/go/packages"
)

type AstReader struct {
	filePath   string
	src        interface{}
	astFile    *ast.File
	fullCode   string
	fileSet    *token.FileSet
	mode       parser.Mode
	childNodes *AstCollection
}

func (it *AstReader) AstFile() *ast.File {
	if it.IsInvalid() {
		return nil
	}

	return it.astFile
}

func (it *AstReader) FullCode() (string, error) {
	if it == nil {
		return "", errcore.CannotBeNilType.ErrorRefOnly(it)
	}

	return it.fullCode, nil
}

func (it *AstReader) IsValid() bool {
	return it != nil && len(it.fullCode) > 0 || it.astFile == nil
}

func (it *AstReader) IsEmpty() bool {
	return it == nil || len(it.fullCode) == 0 || it.astFile == nil
}

func (it *AstReader) IsInvalid() bool {
	return !it.IsValid()
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
	return it.fullCode[start:end], nil
}
func (it *AstReader) SubstringByNode(n ast.Node) (string, error) {
	if n == nil {
		return "", errcore.FailedToParseType.ErrorNoRefs("astFile is nil")
	}

	start := n.Pos() - 1
	end := n.End() - 1

	return it.fullCode[start:end], nil
}

func (it *AstReader) NodesMap() (map[string]args.Map, error) {
	if it.IsInvalid() {
		return map[string]args.Map{}, it.invalidErr()
	}

	curMap := make(map[string]args.Map, 30)
	var rawErrSlice errcore.RawErrCollection

	// Use the Inspect function to walk AST looking for struct
	// type nodes.
	ast.Inspect(
		it.AstFile(), func(n ast.Node) bool {
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
				curMap[typeName] = map[string]interface{}{}
				curMap[typeName][toString] = n
			}

			return true
		},
	)

	return curMap, rawErrSlice.CompiledError()
}

func (it *AstReader) invalidErr() error {
	if it.IsInvalid() {
		return nil
	}

	return errcore.
		InvalidEmptyValueType.
		ErrorNoRefs(
			"invalid ast, either nil, full code empty or astFile is nil",
		)
}

func (it *AstReader) NestedNodesMap() (map[string]args.Map, error) {
	if it.IsInvalid() {
		return map[string]args.Map{}, it.invalidErr()
	}

	// okay
	// Collect the struct types in this slice.
	curMap := make(map[string]args.Map, 30)
	var rawErrSlice errcore.RawErrCollection

	// Use the Inspect function to walk AST looking for struct
	// type nodes.
	ast.Inspect(
		it.AstFile(), func(n ast.Node) bool {
			if n == nil {
				return true
			}

			toString, subsErr := it.SubstringByNode(n)

			if subsErr != nil {
				rawErrSlice.Add(subsErr)

				return true
			}

			recursive := recursiveRunner{
				maxTry:          30,
				SubstringByNode: it.SubstringByNode,
			}

			typeName := it.TypeName(n)
			m, isFound := curMap[typeName]

			if isFound {
				m[toString] = make(map[string]interface{})
				toMap := m[toString].(map[string]interface{})

				recursive.recursiveMapEntry(rawErrSlice, toMap, n)
			} else {
				curMap[typeName] = make(map[string]interface{}, 100)
				curMap[typeName][toString] = make(map[string]interface{}, 100)
				toMap := curMap[typeName][toString].(map[string]interface{})

				recursive.recursiveMapEntry(rawErrSlice, toMap, n)
			}

			return true
		},
	)

	return curMap, rawErrSlice.CompiledError()
}

func (it *AstReader) TypeName(n ast.Node) string {
	return astUtil.TypeName(n)
}

func (it *AstReader) StructTypes() ([]*ast.StructType, error) {
	if it.IsInvalid() {
		return []*ast.StructType{}, it.invalidErr()
	}

	// okay
	var rawErrSlice errcore.RawErrCollection
	var structTypes []*ast.StructType

	// Use the Inspect function to walk AST looking for struct
	// type nodes.
	ast.Inspect(
		it.AstFile(), func(n ast.Node) bool {
			if x, isOkay := n.(*ast.StructType); isOkay {
				structTypes = append(structTypes, x)
			}

			return true
		},
	)

	return structTypes, rawErrSlice.CompiledError()
}

func (it *AstReader) ChildNodes() *AstCollection {
	if it.IsEmpty() {
		return nil
	}

	if it.childNodes != nil {
		return it.childNodes
	}

	creatorFunc := New.AstElem.Create
	fullCode, _ := it.FullCode()
	var slice []AstElem
	var rawErr errcore.RawErrCollection

	ast.Inspect(
		it.AstFile(), func(n ast.Node) bool {
			if isany.Null(n) {
				return true
			}

			elem, err := creatorFunc(it, fullCode, n)
			rawErr.Add(err)

			if err == nil {
				slice = append(slice, *elem)
			}

			return true
		},
	)

	parent, _ := creatorFunc(it, fullCode, it.AstFile())

	collection := &AstCollection{
		Parent:     parent,
		childNodes: slice,
	}

	it.childNodes = collection

	return it.childNodes
}

func (it *AstReader) Filter(filter func(elem *AstElem) (isTake, isBreak bool)) *AstCollection {
	if it.IsEmpty() {
		return nil
	}

	if it.childNodes != nil {
		return it.childNodes
	}

	creatorFunc := New.AstElem.Create
	fullCode, _ := it.FullCode()
	var slice []AstElem
	var rawErr errcore.RawErrCollection

	ast.Inspect(
		it.AstFile(), func(n ast.Node) bool {
			if isany.Null(n) {
				return true
			}

			elem, err := creatorFunc(it, fullCode, n)
			rawErr.Add(err)
			isTake, isBreak := filter(elem)

			if err == nil && isTake {
				slice = append(slice, *elem)
			}

			if isBreak {
				return false
			}

			return true
		},
	)

	parent, _ := creatorFunc(it, fullCode, it.AstFile())

	collection := &AstCollection{
		Parent:     parent,
		childNodes: slice,
	}

	it.childNodes = collection

	return it.childNodes
}

func (it *AstReader) Functions() *AstFuncCollection {
	if it.IsEmpty() {
		return nil
	}

	astFuncCollection, err := New.AstFuncCollection.Create(
		it,
		it.AstFile(),
	)

	errcore.HandleErr(err)

	return astFuncCollection
}
