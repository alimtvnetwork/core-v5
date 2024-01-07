package codegen

import (
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/internal/convertinteranl"
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

	v := reflect.ValueOf(p)
	t := v.Type()

	if t.Kind() == reflect.Struct {
		return it.WriteStruct(p)
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
	sb.WriteString(t.Name() + "{\n")

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.CanInterface() { // check if field is exported
			sb.WriteString(
				fmt.Sprintf(
					"\t%s: %v,\n",
					t.Field(i).Name,
					f.Interface(),
				),
			)
		}
	}

	sb.WriteString("}")

	return sb.String()
}
