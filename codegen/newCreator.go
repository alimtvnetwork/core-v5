package codegen

import "gitlab.com/auk-go/core/codegen/aukast"

type newCreator struct {
	GoCode    newGoCodeCreator
	FinalCode newFinalCodeCreator
	AstReader aukast.newAstReaderCreator
}
