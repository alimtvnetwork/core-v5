package codegen

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

type printer struct{}

func (it printer) WriteProperty(p interface{}) string {
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
		return fmt.Sprintf("%d", casted)
	case args.String:
		return fmt.Sprintf("%s", casted)
	}

	rv := reflect.ValueOf(p)
	t := rv.Type()
	switch t.Kind() {
	case reflect.Struct:
		return it.WriteStruct(p)
	case reflect.Slice, reflect.Array:
		var slice corestr.SimpleSlice
		_ = reflectinternal.Looper.SliceForRv(
			rv,
			func(total int, index int, item interface{}) (err error) {
				expand := it.WriteProperty(item)

				slice.Add(expand)

				return nil
			},
		)

		toJoined := slice.Join(ArgsJoinerEachLine)

		return fmt.Sprintf("%T {\n\t%s\n}\n", p, toJoined)
	case reflect.Ptr:
		if isany.Null(rv.Interface()) {
			return "nil"
		}

		elem := rv.Elem().Interface()
		expandProperties := it.WriteProperty(elem)

		return fmt.Sprintf("&%s", expandProperties)
	}

	// TODO fix this for https://prnt.sc/SNvDVD9KBDs7
	return convertinteranl.AnyTo.FullPropertyString(p)
}

func (it printer) WriteStruct(p interface{}) string {
	if isany.Null(p) {
		return "nil"
	}

	v := reflect.ValueOf(p)
	t := v.Type()
	var sb strings.Builder
	sb.WriteString(t.String() + "{\n")

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.CanInterface() { // check if field is exported
			switch f.Type().Kind() {
			case reflect.Struct:
				sb.WriteString(
					fmt.Sprintf(
						"\t%s: %s,\n",
						t.Field(i).Name,
						it.WriteStruct(f.Interface()),
					),
				)
			case reflect.Ptr:
				if isany.Null(f.Interface()) {
					sb.WriteString(
						fmt.Sprintf(
							"\t%s: nil,\n",
							t.Field(i).Name,
						),
					)

					continue
				}
			}

			sb.WriteString(
				fmt.Sprintf(
					"\t%s: %s,\n",
					t.Field(i).Name,
					it.WriteProperty(f.Interface()),
				),
			)
		}
	}

	sb.WriteString("}")

	return sb.String()
}
