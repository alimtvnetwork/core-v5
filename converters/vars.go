package converters

import (
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/tests/integratedtests/jsoninternal"
)

var (
	StringsTo  = stringsTo{}
	Any        = anyItemConverter{}
	Map        = convertinteranl.Map
	PrettyJson = jsoninternal.Pretty
)
