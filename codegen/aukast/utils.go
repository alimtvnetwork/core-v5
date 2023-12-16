package aukast

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

	"gitlab.com/auk-go/core/errcore"
)

type utils struct{}

func (it utils) TypeName(n ast.Node) string {
	return fmt.Sprintf("%T", n)
}

func (it *utils) NodeToString(fullCode string, n ast.Node) (string, error) {
	if n == nil {
		return "", errcore.FailedToParseType.ErrorNoRefs("node is nil")
	}

	if len(fullCode) == 0 {
		return "", errcore.FailedToParseType.ErrorNoRefs("full code cannot be empty")
	}

	start := int(n.Pos() - 1)
	end := int(n.End() - 1)

	if start < 0 {
		return fullCode[:end], errcore.FailedToParseType.ErrorNoRefs("start cannot be less than 0")
	}

	if end > len(fullCode) {
		return fullCode, errcore.FailedToParseType.ErrorNoRefs("end cannot be larger than full len")
	}

	return fullCode[start:end], nil
}

func (it *utils) NodeToStringSafe(fullCode string, n ast.Node) string {
	code, _ := it.NodeToString(fullCode, n)

	return code
}

func (it *utils) AstFileToCode(fSet *token.FileSet, file *ast.File) (string, error) {
	myWriter := &BytesWriter{}
	err := printer.Fprint(myWriter, fSet, file)

	return myWriter.String(), errcore.StackEnhance.Error(err)
}

func (it *utils) IdentifiersToString(identifiers []*ast.Ident) string {
	if len(identifiers) == 0 {
		return ""
	}

	var slice []string

	for _, identifier := range identifiers {
		if identifier == nil {
			continue
		}

		slice = append(slice, identifier.Name)
	}

	return strings.Join(slice, ", ")
}

func (it *utils) FieldsListToString(fieldsList *ast.FieldList) string {
	if fieldsList == nil || len(fieldsList.List) == 0 {
		return ""
	}

	var slice []string

	for _, field := range fieldsList.List {
		if field == nil {
			continue
		}

		slice = append(slice, it.IdentifiersToString(field.Names))
	}

	toStr := strings.Join(slice, "; ")

	return fmt.Sprintf("[]%s", toStr)
}

func (it utils) Name(fullCode string, n ast.Node) string {
	switch v := n.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.BasicLit:
		return v.Kind.String()
	case *ast.FuncDecl:
		return it.NodeToStringSafe(fullCode, v.Name)
	case *ast.FuncType:
		return it.FieldsListToString(v.Results)
	case *ast.SelectorExpr:
		return it.NodeToStringSafe(fullCode, v.X)
	case *ast.KeyValueExpr:
		return it.NodeToStringSafe(fullCode, v.Key)
	case *ast.Field:
		return it.IdentifiersToString(v.Names)
	case *ast.FieldList:
		return it.FieldsListToString(v)

	case *ast.ExprStmt, *ast.RangeStmt, *ast.CompositeLit:
		// https://prnt.sc/48i_Cuko_J5r

		return it.NodeToStringSafe(fullCode, v)
	}

	return ""
}

func (it utils) Kind(fullCode string, n ast.Node) string {
	switch v := n.(type) {
	case *ast.RangeStmt:
		return ""
	case *ast.BasicLit:
		return v.Kind.String()
	}

	return ""
}

func (it utils) ToIdent(n ast.Node) *ast.Ident {
	switch v := n.(type) {
	case *ast.Ident:
		return v
	case *ast.Field:
		if len(v.Names) == 0 {
			return nil
		}

		return v.Names[0]
	case *ast.FieldList:
		if len(v.List) == 0 {
			return nil
		}

		if len(v.List[0].Names) == 0 {
			return nil
		}

		return v.List[0].Names[1]
	}

	return nil
}
