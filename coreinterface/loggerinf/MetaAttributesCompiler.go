package loggerinf

type MetaAttributesCompiler interface {
	StringFinalizer
	IfStringCompiler
	Compiler
	FmtCompiler
	Comitter
	CompileStacks() []string
	CompileMap() map[string]interface{}
}
