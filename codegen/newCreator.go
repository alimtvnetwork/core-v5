package codegen

type newCreator struct {
	GoCode    newGoCodeCreator
	FinalCode newFinalCodeCreator
	AstReader newAstReaderCreator
}
