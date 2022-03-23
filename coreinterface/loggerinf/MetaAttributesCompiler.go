package loggerinf

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
)

type MetaAttributesCompiler interface {
	coreinterface.Disposer

	StringFinalizer
	IfStringCompiler
	Compiler
	FmtCompiler
	// Committer
	//
	// logs and clears
	Committer
	CompileAnyTo(toPointer interface{}) error
	CompileAny() interface{}
	CompileStacks() []string
	ReflectSetter
	CompileMap() map[string]interface{}
	CompileToJsonResult() *corejson.Result

	BytesCompiler
	BytesCompilerIf
	MustBytesCompiler
}
