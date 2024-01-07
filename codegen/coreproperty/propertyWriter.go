package coreproperty

import (
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/isany"
	"gitlab.com/auk-go/core/simplewrap"
)

type propertyWriter struct{}

func (it propertyWriter) Write(p interface{}) string {
	if isany.Null(p) {
		return "nil"
	}

	return it.WritePropertyOptions(false, p)
}

func (it propertyWriter) WritePropertyOptions(
	isSubRequest bool,
	p interface{},
) string {
	if isany.Null(p) {
		return "nil"
	}

	switch casted := p.(type) {
	case string:
		return simplewrap.WithDoubleQuote(casted)
	case bool, int, int32, int64,
		float64, float32, byte,
		int8, uint16, uint32,
		uint64:
		return fmt.Sprintf("%v", casted)
	case args.String:
		return fmt.Sprintf("%s", casted)
	}

	rv := reflect.ValueOf(p)
	t := rv.Type()
	kind := t.Kind()

	switch kind {
	case reflect.Struct:
		return it.WriteStruct(p)
	case reflect.Slice, reflect.Array:
		return it.WriteArrayOrSlice(isSubRequest, p)
	case reflect.Ptr:
		return it.WritePointerRv(isSubRequest, rv)
	case reflect.Map:
		return it.WriteMap(isSubRequest, p)
	}

	// TODO fix this for https://prnt.sc/SNvDVD9KBDs7
	return convertinteranl.AnyTo.FullPropertyString(p)
}

func (it propertyWriter) WriteMap(
	isSubRequest bool,
	p interface{},
) string {
	var slice corestr.SimpleSlice
	_ = reflectinternal.Looper.Map(
		p,
		func(total int, index int, key interface{}, value interface{}) (err error) {
			expandKey := it.Write(key)
			expandValue := it.Write(value)

			slice.AppendFmt("%s: %s", expandKey, expandValue)

			return nil
		},
	)

	toJoined := slice.Join(ArgsJoinerEachLineTab)

	return fmt.Sprintf("%T {\n\t%s,\n}", p, toJoined)
}

func (it propertyWriter) WritePointerRv(
	isSubRequest bool,
	rv reflect.Value,
) string {
	if isany.Null(rv.Interface()) {
		return "nil"
	}

	elem := rv.Elem().Interface()
	expandProperties := it.Write(elem)

	return fmt.Sprintf("&%s", expandProperties)
}

func (it propertyWriter) WriteArrayOrSlice(
	isSubRequest bool,
	p interface{},
) string {
	var slice corestr.SimpleSlice
	_ = reflectinternal.Looper.Slice(
		p,
		func(total int, index int, item interface{}) (err error) {
			expand := it.Write(item)

			slice.Add(expand)

			return nil
		},
	)

	toJoined := slice.Join(ArgsJoinerEachLineTab)

	return fmt.Sprintf("%T {\n\t%s,\n}", p, toJoined)
}

func (it propertyWriter) WriteStruct(p interface{}) string {
	if isany.Null(p) {
		return "nil"
	}

	v := reflect.ValueOf(p)
	t := v.Type()
	var sb strings.Builder
	sb.WriteString(t.String() + "{\n")

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		if !f.CanInterface() {
			// cannot export
			continue
		}

		fieldName := t.Field(i).Name
		fValue := f.Interface()

		sb.WriteString(
			fmt.Sprintf(
				"\t%s: %s,\n",
				fieldName,
				it.Write(fValue),
			),
		)
	}

	sb.WriteString("}")

	return sb.String()
}
