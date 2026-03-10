package corepayload

import (
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

var (
	Empty              = emptyCreator{}
	New                = newCreator{}
	attributesTypeName = reflectinternal.TypeName(Attributes{})
)
