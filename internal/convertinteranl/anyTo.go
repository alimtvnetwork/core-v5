package convertinteranl

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/tests/integratedtests/jsoninternal"
)

type anyTo struct{}

// String
//
// If nil then returns ""
// Or, returns %v of the value given.
func (it anyTo) String(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem,
	)
}

// SmartString
//
//   - If nil return ""
//   - If string return just returns
//   - Or, else return %v of value
func (it anyTo) SmartString(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	toStr, isSuccess := anyItem.(string)

	if isSuccess {
		return toStr
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem,
	)
}

func (it anyTo) SmartJson(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	switch v := anyItem.(type) {
	case []string:
		return strings.Join(
			v,
			constants.NewLineUnix,
		)
	case string:
		return v
	default:
		toPrettyJson := jsoninternal.Pretty.AnyTo.SafeString(anyItem)

		return toPrettyJson
	}
}

func (it anyTo) SmartPrettyJsonLines(anyItem interface{}) []string {
	if anyItem == nil {
		return []string{}
	}

	switch v := anyItem.(type) {
	case []string:
		return v
	case string:
		return strings.Split(
			v,
			constants.NewLineUnix,
		)

	default:
		toPrettyJson := jsoninternal.Pretty.AnyTo.SafeString(anyItem)

		return strings.Split(
			toPrettyJson,
			constants.NewLineUnix,
		)
	}
}

func (it anyTo) Strings(
	any interface{},
) []string {
	switch v := any.(type) {
	case string:
		return strings.Split(v, constants.NewLineUnix)
	case []string:
		return v
	case []interface{}:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))

		for i, line := range v {
			lines[i] = it.SmartString(line)
		}

		return lines

	case map[string]interface{}:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))
		index := 0

		for key, value := range v {
			lines[index] = fmt.Sprintf(
				"%s : %s",
				key,
				it.SmartJson(value),
			)

			index++
		}

		return lines
	case map[interface{}]interface{}:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))
		index := 0

		for key, value := range v {
			lines[index] = fmt.Sprintf(
				"%s : %s",
				it.SmartJson(key),
				it.SmartJson(value),
			)

			index++
		}

		return lines
	case map[string]string:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))
		index := 0

		for key, value := range v {
			lines[index] = fmt.Sprintf(
				"%s : %s",
				key,
				value,
			)

			index++
		}

		return lines

	case map[string]int:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))
		index := 0

		for key, value := range v {
			lines[index] = fmt.Sprintf(
				"%s : %s",
				key,
				it.String(value),
			)

			index++
		}

		return lines

	case map[int]string:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))
		index := 0

		for key, value := range v {
			lines[index] = fmt.Sprintf(
				"%s : %s",
				it.String(key),
				value,
			)

			index++
		}

		return lines
	default:
		toString := jsoninternal.
			Pretty.
			AnyTo.
			PrettyStringDefault(v)

		return strings.Split(toString, constants.NewLineUnix)
	}
}
