package reflectinternal

import "gitlab.com/auk-go/core/internal/convertinteranl"

var (
	Converter = reflectConverter{}
	Utils     = reflectUtils{}
	Looper    = looper{}

	CodeStack = codeStack{}
	GetFunc   = getFunc{}

	indexToPositionFunc   = convertinteranl.Util.String.IndexToPosition
	prependWithSpacesFunc = convertinteranl.Util.String.PrependWithSpacesDefault
)
