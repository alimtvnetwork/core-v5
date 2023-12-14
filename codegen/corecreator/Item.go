package corecreator

import (
	"math/rand"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Item struct {
	Value         interface{}
	Possibilities []interface{}
	CreatorFunc   func(i Item, index int) interface{}
}

func (it Item) Create() interface{} {
	return it.Value
}

func (it Item) CreateRandom() interface{} {
	rndIndex := rand.Intn(it.Length())

	return it.CreateByIndex(rndIndex)
}

func (it Item) CreateByIndex(i int) interface{} {
	if it.HasIndex(i) {
		return it.Possibilities[i]
	}

	return nil
}

func (it Item) CreateBySafeIndexDefault(i int) interface{} {
	if it.HasIndex(i) {
		return it.Possibilities[i]
	}

	return it.Value
}

func (it Item) Length() int {
	return len(it.Possibilities)
}

func (it Item) Count() int {
	return len(it.Possibilities)
}

func (it Item) HasIndex(i int) bool {
	return len(it.Possibilities)-1 >= i
}

func (it Item) IsPrimitiveType() bool {
	return reflectinternal.Is.Primitive(it.Value)
}

func (it Item) IsNumber() bool {
	return reflectinternal.Is.NumberKind(it.Value)
}

func (it Item) CreateByFunc(i int) interface{} {
	if it.CreatorFunc == nil {
		return nil
	}

	return it.CreatorFunc(it, i)
}
