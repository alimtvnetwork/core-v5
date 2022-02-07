package internalinterface

type SimpleEnumer interface {
	IsNameEqual(name string) bool
	ValueByteWithValueEqualer
	ToNamer
}
