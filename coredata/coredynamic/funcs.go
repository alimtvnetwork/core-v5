package coredynamic

import "reflect"

type (
	SimpleInOutConverter                   func(in any, typeMust reflect.Type) *SimpleResult
	SimpleRequestConverter                 func(request SimpleRequest) *SimpleResult
	TypeToTypeConverterFunc                func(in any) (output any, err error)
	TypeToTypeConverterMustFunc            func(in any) (output any)
	TypeToTypeConverterWithOptionsFunc     func(in any, options any) (output any, err error)
	TypeToTypeConverterWithOptionsMustFunc func(in any, options any) (output any)
)
