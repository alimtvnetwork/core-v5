package convertinternal

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/internal/jsoninternal"
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
		if isNilPointer(anyItem) {
			return ""
		}

		return v.Name()
	case fmt.Stringer:
		if isNilPointer(anyItem) {
			return ""
		}

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

// mapToSortedLines converts a map to sorted "key : value" formatted lines.
func mapToSortedLines[K comparable, V any](
	m map[K]V,
	keyFmt func(K) string,
	valFmt func(V) string,
) []string {
	if len(m) == 0 {
		return []string{}
	}

	lines := make([]string, 0, len(m))

	for key, value := range m {
		lines = append(lines, fmt.Sprintf(
			"%s : %s",
			keyFmt(key),
			valFmt(value),
		))
	}

	sort.Strings(lines)

	return lines
}

// stringsFromSlice converts typed slices to string slices.
func stringsFromSlice[T any](s []T, fmtFunc func(T) string) []string {
	lines := make([]string, len(s))
	for i, elem := range s {
		lines[i] = fmtFunc(elem)
	}
	return lines
}

func (it anyTo) Strings(
	item any,
) []string {
	switch v := item.(type) {
	case string:
		if v == "" {
			return []string{}
		}
		return strings.Split(v, constants.NewLineUnix)
	case error:
		if v == nil {
			return []string{}
		}
		return strings.Split(v.Error(), constants.NewLineUnix)
	case []string:
		return v
	case []any:
		return it.stringsFromAnySlice(v)
	case map[string]any:
		return mapToSortedLines(v, func(k string) string { return k }, it.SmartJson)
	case map[any]any:
		return mapToSortedLines(v, it.SmartJson, it.SmartJson)
	case map[string]string:
		return mapToSortedLines(v, func(k string) string { return k }, func(v string) string { return v })
	case map[string]int:
		return mapToSortedLines(v, func(k string) string { return k }, func(v int) string { return fmt.Sprintf("%d", v) })
	case map[int]string:
		return mapToSortedLines(v, func(k int) string { return it.ValueString(k) }, func(v string) string { return v })
	default:
		return it.stringsFromPrimitiveOrFallback(item)
	}
}

func (it anyTo) stringsFromAnySlice(v []any) []string {
	if len(v) == 0 {
		return []string{}
	}

	lines := make([]string, len(v))
	for i, line := range v {
		lines[i] = it.SmartJson(line)
	}
	return lines
}

func (it anyTo) stringsFromPrimitiveOrFallback(item any) []string {
	switch v := item.(type) {
	case int, int32, int64,
		uint8,
		uint16, uint32, uint64,
		float32, float64:
		return []string{fmt.Sprintf("%d", v)}
	case fmt.Stringer:
		return strings.Split(v.String(), constants.NewLineUnix)
	case bool:
		return []string{strconv.FormatBool(v)}
	case []int:
		return stringsFromSlice(v, strconv.Itoa)
	case []bool:
		return stringsFromSlice(v, strconv.FormatBool)
	case []int64:
		return stringsFromSlice(v, func(n int64) string { return strconv.FormatInt(n, 10) })
	case []float64:
		return stringsFromSlice(v, func(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) })
	case []byte:
		return stringsFromSlice(v, func(b byte) string { return strconv.Itoa(int(b)) })
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

func isNilPointer(item any) bool {
	rv := reflect.ValueOf(item)

	return rv.Kind() == reflect.Ptr && rv.IsNil()
}
