package corecreator

import "gitlab.com/auk-go/core/internal/reflectinternal"

type Item struct {
	Value         interface{}
	Possibilities []interface{}
	CreatorFunc   func(i Item, index int) interface{}
}

func (it Item) Create() interface{} {

}

func (it Item) CreateRandom() interface{} {

}

func (it Item) CreateByIndex(i int) interface{} {

}

func (it Item) IsPrimitiveType() bool {
	return reflectinternal.Is.PrimitiveKind()
}

func (it Item) CreateByFunc(i int) interface{} {
	if it.CreatorFunc == nil {
		return nil
	}

	return it.CreatorFunc(it, i)
}

func (it Item) CreateRandom() interface{} {

}
