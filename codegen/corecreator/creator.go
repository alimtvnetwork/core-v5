package corecreator

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/alimtvnetwork/core/coretests/args"
)

type Creator struct {
	scopeName string
	scopeMap  args.Map
}

func (it newCreator) Create(i any) any {
	return nil
}

func (it newCreator) CreateByType(rt reflect.Type) any {
	// Switch on the kind of the type
	switch rt.Kind() {
	case reflect.Bool:
		// Return a random bool value
		return rand.Intn(2) == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// Return a random int value
		return rand.Int63n(100)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		// Return a random uint value
		return rand.Uint64() % 100
	case reflect.Float32, reflect.Float64:
		// Return a random float value
		return rand.Float64() * 100
	case reflect.Complex64, reflect.Complex128:
		// Return a random complex value
		return complex(rand.Float64()*100, rand.Float64()*100)
	case reflect.String:
		// Return a random string value
		return fmt.Sprintf("string-%d", rand.Intn(100))
	case reflect.Array:
		// Create a new array value of the same type and length
		arr := reflect.New(rt).Elem()
		// Loop over the elements and fill them with random values
		for i := 0; i < rt.Len(); i++ {
			arr.Index(i).Set(reflect.ValueOf(it.CreateByType(rt.Elem())))
		}
		// Return the array value
		return arr.Interface()
	case reflect.Slice:
		// Create a new slice value of the same type and a random length
		slice := reflect.MakeSlice(rt, rand.Intn(10), rand.Intn(10))
		// Loop over the elements and fill them with random values
		for i := 0; i < slice.Len(); i++ {
			slice.Index(i).Set(reflect.ValueOf(it.CreateByType(rt.Elem())))
		}
		// Return the slice value
		return slice.Interface()
	case reflect.Map:
		// Create a new map value of the same type and a random length
		m := reflect.MakeMapWithSize(rt, rand.Intn(10))
		// Loop over the keys and values and fill them with random values
		for i := 0; i < m.Len(); i++ {
			m.SetMapIndex(reflect.ValueOf(it.CreateByType(rt.Key())), reflect.ValueOf(it.CreateByType(rt.Elem())))
		}
		// Return the map value
		return m.Interface()
	case reflect.Struct:
		// Create a new struct value of the same type
		s := reflect.New(rt).Elem()
		// Loop over the fields and fill them with random values
		for i := 0; i < rt.NumField(); i++ {
			s.Field(i).Set(reflect.ValueOf(it.CreateByType(rt.Field(i).Type)))
		}
		// Return the struct value
		return s.Interface()
	case reflect.Ptr:
		// Create a new pointer value of the same type
		p := reflect.New(rt.Elem())
		// Set the pointer to point to a random value of the underlying type
		p.Elem().Set(reflect.ValueOf(it.CreateByType(rt.Elem())))
		// Return the pointer value
		return p.Interface()
	case reflect.Chan:
		// Create a new channel value of the same type and a random buffer size
		c := reflect.MakeChan(rt, rand.Intn(10))
		// Return the channel value
		return c.Interface()
	case reflect.Func:
		// Create a new function value of the same type
		f := reflect.MakeFunc(
			rt, func(args []reflect.Value) (results []reflect.Value) {
				// Loop over the results and fill them with random values
				for i := 0; i < rt.NumOut(); i++ {
					results = append(results, reflect.ValueOf(it.CreateByType(rt.Out(i))))
				}
				// Return the results
				return
			},
		)
		// Return the function value
		return f.Interface()
	case reflect.Interface:
		// Return a nil interface value
		return nil
	default:
		// Return a zero value of the type
		return reflect.Zero(rt).Interface()
	}
}

func (it newCreator) CreateByTypeName(name string) any {
	return nil
}

func (it newCreator) CreateByFunc(typeName string, index int) any {
	return nil
}
