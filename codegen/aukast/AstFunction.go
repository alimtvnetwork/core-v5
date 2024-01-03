package aukast

import (
	"fmt"
	"go/ast"
	"reflect"
)

type AstFunction struct {
	AstReader      *AstReader
	Name           string
	StructVarName  string
	StructName     string
	IsAttached     bool
	IsPublic       bool
	IsPrivate      bool
	ReceiverCount  int
	Parent         *AstElem
	ReceiverStruct *AstElem
	Comments       *AstElem
	Type           *ast.FuncType
	FuncArg        *AstFuncArgsRoot
	Node           *ast.FuncDecl
	DefCode, Code  string
}

func (it *AstFunction) IsNotAttached() bool {
	return it.IsValid() && !it.IsAttached
}

func (it *AstFunction) IsValid() bool {
	return !it.IsInvalid()
}

func (it *AstFunction) IsInvalid() bool {
	return it == nil || it.Name == "" || it.Type == nil
}

func (it *AstFunction) HasInArgs() bool {
	return it.IsValid() && len(it.FuncArg.InArgs) > 0
}

func (it *AstFunction) HasOutArgs() bool {
	return it.IsValid() && len(it.FuncArg.OutArgs) > 0
}

func (it *AstFunction) IsEmptyInArgs() bool {
	return !it.HasInArgs()
}

func (it *AstFunction) IsEmptyOutArgs() bool {
	return !it.HasOutArgs()
}

func (it *AstFunction) IsAttachToStructOf(structName string) bool {
	if it.IsInvalid() {
		return false
	}

	return it.StructName == structName
}

func (it AstFunction) String() string {
	if it.IsInvalid() {
		return ""
	}

	fields := fmt.Sprintf(
		"%s:%s\n"+
			"%s:%s\n"+
			"%s:%x\n"+
			"%s:%s\n"+
			"%s:%s\n",
		"Name", it.DisplayName(),
		"Code", it.DefCode,
		"IsPublic", it.IsPublic,
		"Args", it.FuncArg.String(),
		"Childs", it.ChildNodes().String(),
	)

	return "AstFunction: {\n" + fields + "\n}"
}

func (it *AstFunction) DisplayName() string {
	if it == nil {
		return ""
	}

	if it.IsAttached {
		return fmt.Sprintf("%s.%s", it.StructName, it.Name)
	}

	return it.Name
}

func (it *AstFunction) ChildNodes() *AstCollection {
	return AstFilter{
		AstReader:  it.AstReader,
		ParentNode: it.Node,
		Node:       it.Node,
		fullCode:   it.AstReader.SafeFullCode(),
	}.ChildNodes()
}

func (it *AstFunction) Filter(
	filterFunc AstWithBreakFilterFunc,
) *AstCollection {
	astFilter := AstFilter{
		AstReader:  it.AstReader,
		ParentNode: it.Node,
		Node:       it.Node,
		fullCode:   it.AstReader.SafeFullCode(),
	}

	return astFilter.Filter(it.Node, filterFunc)
}

func (it *AstFunction) ChildOf(
	typeMatches ...reflect.Type,
) *AstCollection {
	return it.Filter(
		func(elem *AstElem) (isTake, isBreak bool) {
			if elem.IsAnyNodeTypeMatches(typeMatches...) {
				return true, false
			}

			return false, false
		},
	)
}

func (it *AstFunction) CodeTakeMax(charsCount int) string {
	if it.IsEmpty() {
		return ""
	}

	return astUtil.MaxSubstringTrimSpaces(it.Code, charsCount)
}

func (it *AstFunction) IsEmpty() bool {
	return it == nil || it.Node == nil
}
