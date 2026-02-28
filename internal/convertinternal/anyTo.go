package convertinternal

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/jsoninternal"
)

type anyTo struct{}

func (it anyTo) ValueString(anyItem any) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem,
	)
}

func (it anyTo) FullPropertyString(anyItem any) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintFullPropertyNameValueFormat,
		anyItem,
	)
}

func (it anyTo) TypeName(anyItem any) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintTypeFormat,
		anyItem,
	)
}

func (it anyTo) SmartString(anyItem any) string {
	if anyItem == nil {
		return ""
	}

	switch v := anyItem.(type) {
	case string:
		return v
	case Namer:
		return v.Name()
	case fmt.Stringer:
		return v.String()
	case error:
		if v == nil {
			return ""
		}

		return v.Error()
	case []string:
		return strings.Join(
			v,
			constants.NewLineUnix,
		)
	case []any:
		if len(v) == 0 {
			return ""
		}

		var slice []string

		for _, elem := range v {
			slice = append(slice, it.SmartString(elem))
		}

		return strings.Join(
			slice,
			",",
		)
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem,
	)
}

func (it anyTo) SmartJson(anyItem any) string {
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
	case int, int32, byte, int64, float64, float32, bool, uint, uint32, uint64:
		return fmt.Sprintf(
			constants.SprintValueFormat,
			anyItem,
		)
	case error:
		if v == nil {
			return ""
		}

		return v.Error()
	default:
		toPrettyJson := jsoninternal.Pretty.
			AnyTo.
			SafeString(anyItem)

		return toPrettyJson
	}
}

func (it anyTo) SmartPrettyJsonLines(anyItem any) []string {
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
		return it.PrettyJsonLines(anyItem)
	}
}

func (it anyTo) PrettyJsonLines(anyItem any) []string {
	if anyItem == nil {
		return []string{}
	}

	toPrettyJson := jsoninternal.Pretty.
		AnyTo.
		PrettyStringDefault(anyItem)

	return strings.Split(
		toPrettyJson,
		constants.NewLineUnix,
	)
}

func (it anyTo) Strings(
	item any,
) []string {
	switch v := item.(type) {
	case string:
		if v == "" {
			return []string{}
		}

		return strings.Split(
			v,
			constants.NewLineUnix,
		)
	case error:
		if v == nil {
			return []string{}
		}

		return strings.Split(
			v.Error(),
			constants.NewLineUnix,
		)
	case []string:
		return v
	case []any:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))

		for i, line := range v {
			lines[i] = it.SmartJson(line)
		}

		return lines

	case map[string]any:
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

		sort.Strings(lines)

		return lines
	case map[any]any:
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

		sort.Strings(lines)

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

		sort.Strings(lines)

		return lines
	case map[string]int:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))
		index := 0

		for key, value := range v {
			lines[index] = fmt.Sprintf(
				"%s : %d",
				key,
				value,
			)

			index++
		}

		sort.Strings(lines)

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
				it.ValueString(key),
				value,
			)

			index++
		}

		sort.Strings(lines)

		return lines
	case int, int32, int64,
		uint8,
		uint16, uint32, uint64,
		float32, float64:
		return []string{
			fmt.Sprintf("%d", v),
		}
	case fmt.Stringer:
		return strings.Split(v.String(), constants.NewLineUnix)
	case bool:
		return []string{
			strconv.FormatBool(v),
		}
	default:
		return it.PrettyJsonLines(item)
	}
}

func (it anyTo) String(
	item any,
) string {
	switch v := item.(type) {
	case string:
		return v
	case *string:
		if v == nil {
			return ""
		}

		return *v
	case error:
		if v == nil {
			return ""
		}

		return v.Error()
	case int, int32, int64,
		uint8,
		uint16, uint32, uint64,
		float32, float64:
		return fmt.Sprintf("%d", v)
	case bool:
		return strconv.FormatBool(v)
	}

	toLines := it.Strings(item)

	return strings.Join(toLines, constants.NewLineUnix)
}
