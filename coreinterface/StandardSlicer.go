package coreinterface

import "github.com/alimtvnetwork/core/coredata/corejson"

type StandardSlicer interface {
	BasicSlicer
	ListStringsGetter
	JsonCombineStringer
	corejson.JsonContractsBinder
}
