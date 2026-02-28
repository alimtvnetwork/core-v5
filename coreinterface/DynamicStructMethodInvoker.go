package coreinterface

import "gitlab.com/auk-go/core/coredata/coredynamic"

type DynamicStructMethodInvoker interface {
	DynamicMethodInvoke(structInput any, args ...any) *coredynamic.SimpleResult
}
