package converters

import (
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/jsoninternal"
)

var (
	StringsTo   = stringsTo{}
	AnyTo       = anyItemConverter{}
	Map         = convertinteranl.Map
	StringTo    = stringTo{}
	PrettyJson  = jsoninternal.Pretty
	BytesTo     = bytesTo{}
	Integers    = convertinteranl.Integers
	KeyValuesTO = convertinteranl.KeyValuesTo
)
