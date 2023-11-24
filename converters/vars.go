package converters

import (
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/jsoninternal"
)

var (
	StringsTo  = stringsTo{}
	AnyTo      = anyItemConverter{}
	Map        = convertinteranl.Map
	PrettyJson = jsoninternal.Pretty
)
