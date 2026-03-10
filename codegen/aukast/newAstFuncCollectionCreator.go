package aukast

import (
	"go/ast"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/isany"
)

type newAstFuncCollectionCreator struct{}

func (it newAstFuncCollectionCreator) Create(
	astReader *AstReader,
	rootNode ast.Node,
) (*AstFuncCollection, error) {
	if isany.Null(rootNode) {
		return nil, nil
	}

	fullCode, codeErr := astReader.FullCode()

	if codeErr != nil {
		return nil, codeErr
	}

	funcMap := make(map[string]AstFunction, 10)
	astFuncCreator := New.AstFunction.Create
	var rawErr errcore.RawErrCollection

	ast.Inspect(
		rootNode, func(n ast.Node) bool {
			if isany.Null(n) {
				return true
			}

			toFunc, isOkay := n.(*ast.FuncDecl)

			if !isOkay || toFunc == nil {
				return true
			}

			astFunc, err := astFuncCreator(
				astReader,
				fullCode,
				n,
			)

			if err != nil {
				rawErr.Add(err)

				return true
			}

			if astFunc.IsValid() {
				funcMap[astFunc.Name] = *astFunc
			}

			return true
		},
	)

	parent, parentErr := New.AstElem.Create(astReader, fullCode, rootNode)
	namesSlice := corestr.New.SimpleSlice.Cap(len(funcMap))

	for key := range funcMap {
		namesSlice.Add(key)
	}

	collection := &AstFuncCollection{
		Names:  namesSlice,
		Map:    funcMap,
		Parent: parent,
	}

	rawErr.Add(parentErr)

	return collection, rawErr.CompiledErrorWithStackTraces()
}
