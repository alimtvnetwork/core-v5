package coredynamic

import (
	"reflect"

	"github.com/alimtvnetwork/core/errcore"
)

// MapAsKeyValSlice
//
//	expectation : map[key:any]any
func MapAsKeyValSlice(reflectVal reflect.Value) (*KeyValCollection, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return MapAsKeyValSlice(
			reflect.Indirect(reflect.ValueOf(reflectVal)),
		)
	}

	if reflectVal.Kind() != reflect.Map {
		return EmptyKeyValCollection(),
			errcore.TypeMismatchType.Error("Reflection is not Map", reflectVal)
	}

	mapKeys := reflectVal.MapKeys()
	keyValCollection := NewKeyValCollection(len(mapKeys))

	for _, key := range reflectVal.MapKeys() {
		value := reflectVal.MapIndex(key)

		dynamicKV := &KeyVal{
			Key:   key.Interface(),
			Value: value.Interface(),
		}

		keyValCollection.AddPtr(dynamicKV)
	}

	return keyValCollection, nil
}
