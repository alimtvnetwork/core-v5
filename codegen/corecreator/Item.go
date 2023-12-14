package corecreator

import (
	"math/rand"
	"reflect"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Item struct {
	Value           interface{}
	Possibilities   interface{}
	CreatorFunc     func(i Item, index int) interface{}
	rvPossibilities *reflect.Value
	length          int
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
		return it.At(i)
	}

	return nil
}

func (it Item) CreateBySafeIndexDefault(i int) interface{} {
	if it.HasIndex(i) {
		return it.At(i)
	}

	return it.Value
}

func (it *Item) Length() int {
	if it.Possibilities == nil {
		return 0
	}

	if it.length > 0 {
		return it.length
	}

	it.length = getLenReflectFunc(it.Possibilities)

	return it.length
}

func (it *Item) PossibilitiesRv() reflect.Value {
	if it.rvPossibilities != nil {
		return *it.rvPossibilities
	}

	*it.rvPossibilities = reflect.ValueOf(it.Possibilities)

	return *it.rvPossibilities
}

func (it Item) At(index int) interface{} {
	return it.PossibilitiesRv().Index(index)
}

func (it Item) Count() int {
	return it.Length()
}

func (it Item) HasIndex(i int) bool {
	return it.Length()-1 >= i
}

func (it Item) IsBoolean() bool {
	return reflectinternal.Is.Boolean(it.Value)
}

func (it Item) IsPrimitiveType() bool {
	return reflectinternal.Is.Primitive(it.Value)
}

func (it Item) IsNumber() bool {
	return reflectinternal.Is.Number(it.Value)
}

func (it Item) IsString() bool {
	return reflectinternal.Is.String(it.Value)
}

func (it Item) IsPointer() bool {
	return reflectinternal.Is.Pointer(it.Value)
}

func (it Item) IsSliceOrArray() bool {
	return reflectinternal.Is.SliceOrArray(it.Value)
}

func (it Item) IsStruct() bool {
	return reflectinternal.Is.Struct(it.Value)
}

func (it Item) IsFunc() bool {
	return reflectinternal.Is.Function(it.Value)
}

func (it Item) IsZero() bool {
	return reflectinternal.Is.Zero(it.Value)
}

func (it Item) CreateByFunc(i int) interface{} {
	if it.CreatorFunc == nil {
		return nil
	}

	return it.CreatorFunc(it, i)
}
