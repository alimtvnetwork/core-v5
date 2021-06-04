package coreinterface

import "gitlab.com/evatix-go/core/coredata/corejson"

type StandardSlicer interface {
	BasicSlicer
	ItemAtRemover
	ListStringsGetter
	JsonCombineStringer
	corejson.JsonContractsBinder
}
