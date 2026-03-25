package coreinterface

import "github.com/alimtvnetwork/core/coredata/coredynamic"

type DynamicStructMethodInvoker interface {
	DynamicMethodInvoke(structInput any, args ...any) *coredynamic.SimpleResult
}
