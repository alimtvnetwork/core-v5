package aukast

type newCreator struct {
	AstReader         newAstReaderCreator
	AstElem           newAstElemCreator
	ArgsParams        newArgsParamsCreator
	AstFunction       newAstFunctionCreator
	AstFuncCollection newAstFuncCollectionCreator
}
