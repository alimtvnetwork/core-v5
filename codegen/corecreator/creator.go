package corecreator

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests/args"
)

type Creator struct {
	scopeName string
	scopeMap  args.Map
}

func (it newCreator) Create(i interface{}) interface{} {
	return nil
}

func (it newCreator) CreateByType(rt reflect.Type) interface{} {
	return nil
}

func (it newCreator) CreateByTypeName(name string) interface{} {
	return nil
}

func (it newCreator) CreateByFunc(typeName string, index int) interface{} {
	return nil
}
