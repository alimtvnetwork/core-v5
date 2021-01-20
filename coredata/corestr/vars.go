package corestr

import "gitlab.com/evatix-go/core/msgtype"

var (
	StaticEmptyCollection           = *EmptyCollection()
	StaticEmptyCollectionPtr        = &StaticEmptyCollection
	StaticEmptyCharCollectionMap    = *EmptyCharCollectionMap()
	StaticEmptyCharCollectionMapPtr = &StaticEmptyCharCollectionMap
	StaticJsonError                 = msgtype.
					EmptyResultCannotMakeJson.
					Error("", "")
)
