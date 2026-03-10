package errcore

import (
	"reflect"
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

func typesNamesString(
	anyItems ...any,
) string {
	slice := make([]string, len(anyItems))

	for i, item := range anyItems {
		slice[i] = reflect.TypeOf(item).Name()
	}

	return strings.Join(slice, constants.CommaSpace)
}
