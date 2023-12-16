package aukast

type AstFunction struct {
	Name       string
	StructName string
	IsAttached bool
	IsPublic   bool
	IsPrivate  bool
	Parent     *AstElem
	InArgs     []AstElem
	OutArgs    []AstElem
}
