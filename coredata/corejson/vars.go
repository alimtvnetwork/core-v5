package corejson

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
)

var (
	Empty                   = emptyCreator{}
	Serialize               = serializerLogic{}
	Deserialize             = deserializerLogic{}
	NewResult               = newResultCreator{}
	NewResultsCollection    = newResultsCollectionCreator{}
	NewBytesCollection      = newBytesCollectionCreator{}
	NewResultsPtrCollection = newResultsPtrCollectionCreator{}
	NewMapResults           = newMapResultsCreator{}
	StaticJsonError         = errcore.
				EmptyResultCannotMakeJsonType.
				ErrorNoRefs(constants.EmptyString)
)
