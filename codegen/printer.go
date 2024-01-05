package codegen

import (
	"fmt"

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

	return convertinteranl.AnyTo.FullPropertyString(p)
}
