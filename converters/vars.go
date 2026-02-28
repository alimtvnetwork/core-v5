package converters

import (
	"gitlab.com/auk-go/core/internal/convertinternal"
	"gitlab.com/auk-go/core/internal/jsoninternal"
)

var (
	StringsTo     = stringsTo{}
	AnyTo         = anyItemConverter{}
	Map           = convertinternal.Map
	StringTo      = stringTo{}
	PrettyJson    = jsoninternal.Pretty
	JsonString    = jsoninternal.String
	BytesTo       = bytesTo{}
	Integers      = convertinternal.Integers
	KeyValuesTo   = convertinternal.KeyValuesTo
	CodeFormatter = convertinternal.CodeFormatter
)
