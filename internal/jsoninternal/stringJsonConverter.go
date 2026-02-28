package jsoninternal

import (
	"encoding/json"
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

type stringJsonConverter struct{}

func (it stringJsonConverter) SafeDefault(anyItem any) string {
	s, _ := it.Default(anyItem)

	return s
}

// Default
//
// It is not pretty JSON
func (it stringJsonConverter) Default(
	anyItem any,
) (string, error) {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return "", err
	}

	return string(b), err
}

// Pretty
//
// Default pretty json
func (it stringJsonConverter) Pretty(
	anyItem any,
) (string, error) {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return "", err
	}

	return Pretty.Bytes.Prefix("", b)
}

func (it stringJsonConverter) StringValue(name string) []byte {
	doubleQuoted := fmt.Sprintf(
		constants.SprintDoubleQuoteFormat,
		name,
	)

	return []byte(doubleQuoted)
}
