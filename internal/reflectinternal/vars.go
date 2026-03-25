package reflectinternal

import (
	"github.com/alimtvnetwork/core/internal/convertinternal"
)

var (
	Converter = reflectConverter{}
	Utils     = reflectUtils{}
	Looper    = looper{}
	CodeStack = codeStack{}
	GetFunc   = getFunc{}
	Is        = isChecker{}

	ReflectGetter                  = reflectGetter{}
	ReflectType                    = reflectTypeConverter{}
	ReflectGetterUsingReflectValue = reflectGetUsingReflectValue{}
	SliceConverter                 = sliceConverter{}
	MapConverter                   = mapConverter{}
	Path                           = reflectPath{}

	indexToPositionFunc   = convertinternal.Util.String.IndexToPosition
	prependWithSpacesFunc = convertinternal.Util.String.PrependWithSpacesDefault
	repoDir               *string // will be updated, dangerous but this is the way for this now
	typeReplacerMap       = map[string]string{
		"*":  "Ptr",
		"[]": "Slice",
	}
)
