package aukast

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
	"unicode"

	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/isany"
)

type utils struct{}

func (it utils) TypeName(n ast.Node) string {
	return fmt.Sprintf("%T", n)
}

func (it *utils) NodeToString(fullCode string, n ast.Node) (string, error) {
	if isany.Null(n) {
		return "", errcore.FailedToParseType.ErrorNoRefs("node is nil")
	}

	if len(fullCode) == 0 {
		return "", errcore.FailedToParseType.ErrorNoRefs("full code cannot be empty")
	}

	pos := n.Pos()
	end := n.End()

	start := pos - 1
	end = end - 1

	if start < 0 {
		return fullCode[:end], errcore.FailedToParseType.ErrorNoRefs("start cannot be less than 0")
	}

	if int(end) > len(fullCode) {
		return fullCode, errcore.FailedToParseType.ErrorNoRefs("end cannot be larger than full len")
	}

	return fullCode[start:end], nil
}

func (it *utils) NodeToStringSafe(fullCode string, n ast.Node) string {
	if isany.Null(n) {
		return ""
	}

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

	toStr := strings.Join(slice, ",")

	return fmt.Sprintf("%s", toStr)
}

func (it *utils) TypesNamesOfFieldList(code string, fieldsList *ast.FieldList) string {
	if fieldsList == nil || len(fieldsList.List) == 0 {
		return ""
	}

	var slice []string

	for _, field := range fieldsList.List {
		if field == nil {
			continue
		}

		slice = append(slice, it.ExprToString(code, field.Type))
	}

	toStr := strings.Join(slice, ",")

	return fmt.Sprintf("%s", toStr)
}

func (it utils) Name(fullCode string, n ast.Node) string {
	if isany.Null(n) {
		return ""
	}

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
	case ast.Node:
		// https://prnt.sc/48i_Cuko_J5r

		return it.NodeToStringSafe(fullCode, v)
	}

	return ""
}

func (it utils) NodeTypeName(fullCode string, n ast.Node) string {
	if isany.Null(n) {
		return ""
	}

	switch v := n.(type) {
	case *ast.Ident:
		switch casted := v.Obj.Type.(type) {
		case *ast.Ident:
			return casted.Name
		case *ast.Expr:
			return it.ExprToString(fullCode, *casted)
		}

		return it.Name(fullCode, v.Obj.Type.(ast.Node))
	case *ast.KeyValueExpr:
		return it.NodeToStringSafe(fullCode, v.Value)
	case *ast.Field:
		return it.ExprToString(fullCode, v.Type)
	case *ast.FieldList:
		return it.TypesNamesOfFieldList(fullCode, v)
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
	if isany.Null(n) {
		return nil
	}

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

		return v.List[0].Names[0]
	}

	return nil
}

func (it utils) IdentNameTypeName(code string, v *ast.Ident) (name, typeName string) {
	if v == nil {
		return "", ""
	}

	switch casted := v.Obj.Type.(type) {
	case *ast.Ident:
		return v.Name, casted.Name
	}

	return v.Name, it.Name(code, v.Obj.Type.(ast.Node))
}

func (it utils) ExprToString(code string, expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.StarExpr:
		return it.ExprToStringDirect(v.X)
	}

	return it.NodeToStringSafe(code, expr)
}

func (it utils) ExprToStringDirect(expr ast.Expr) string {
	if isany.Null(expr) {
		return ""
	}

	switch v := expr.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.StarExpr:
		return it.ExprToStringDirect(v.X)
	case *ast.ArrayType:
		// https://prnt.sc/kUiAu-WAzU42
		return it.ExprToStringDirect(v.Elt)
	}

	return ""
}

func (it utils) ExprToIdent(expr ast.Expr) *ast.Ident {
	if isany.Null(expr) {
		return nil
	}

	switch v := expr.(type) {
	case *ast.Ident:
		return v
	case *ast.StarExpr:
		return it.ExprToIdent(v.X)
	case *ast.ArrayType:
		// https://prnt.sc/kUiAu-WAzU42
		return it.ExprToIdent(v.Elt)
	}

	return nil
}

func (it utils) HasAnyPrefix(s string, prefixes ...string) bool {
	if len(s) == 0 || len(prefixes) == 0 {
		return false
	}

	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}

	return false
}

func (it utils) IsPublicFuncByName(s string) bool {
	if len(s) == 0 {
		return false
	}

	return unicode.IsUpper(rune(s[0]))
}

func (it utils) IsPrivateFuncByName(s string) bool {
	if len(s) == 0 {
		return false
	}

	return unicode.IsLower(rune(s[0]))
}
