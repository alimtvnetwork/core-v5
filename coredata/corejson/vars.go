package corejson

import (
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/jsoninternal"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

var (
	resultTypeName          = reflectinternal.TypeName(Result{})
	Empty                   = emptyCreator{}
	Serialize               = serializerLogic{}   // deals with json.Marshal
	Deserialize             = deserializerLogic{} // deals with json.Unmarshal
	NewResult               = newResultCreator{}
	NewResultsCollection    = newResultsCollectionCreator{}
	NewBytesCollection      = newBytesCollectionCreator{}
	NewResultsPtrCollection = newResultsPtrCollectionCreator{}
	NewMapResults           = newMapResultsCreator{}
	CastAny                 = castingAny{}
	AnyTo                   = anyTo{}
	Pretty                  = jsoninternal.Pretty
	StaticJsonError         = errcore.
				EmptyResultCannotMakeJsonType.
				ErrorNoRefs(constants.EmptyString)
)
