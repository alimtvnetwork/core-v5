package loggerinf

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coreinterface"
	"github.com/alimtvnetwork/core/coreinterface/errcoreinf"
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
	CompileAnyTo(toPointer any) error
	CompileAny() any
	CompileStacks() []string
	ReflectSetter
	CompileMap() map[string]any
	CompileToJsonResult() *corejson.Result

	CompiledAsBasicErr(
		basicErrTyper errcoreinf.BasicErrorTyper,
	) errcoreinf.BasicErrWrapper

	BytesCompiler
	BytesCompilerIf
	MustBytesCompiler
}
