package aukast

type AstFunction struct {
	Name       string
	StructName string
	Parent     *AstElem
	InArgs     []AstElem
	OutArgs    []AstElem
}
