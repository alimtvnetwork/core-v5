package coreinterface

type SimpleDynamicDataConverter interface {
	// GetConvertTo convert `type` from, to
	GetConvertTo(from any) (to any, err error)
	// Convert convert `type` from, to
	Convert(from, to any) error
	// ConvertPointers convert `*type` from, to
	ConvertPointers(from, to any) error
}
