package codegen

import (
	"go/ast"
	"go/parser"
	"go/token"

	"gitlab.com/auk-go/core/errcore"
)

type astReader struct {
}

func (it astReader) Initialize(filePath string) (*ast.File, error) {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(
		fileSet,
		filePath,
		nil,
		parser.AllErrors,
	)

	if err != nil {
		return node, errcore.ParsingFailed.MsgCsvRefError(
			err.Error(),
			filePath,
		)
	}

}
