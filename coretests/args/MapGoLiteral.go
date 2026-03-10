package args

import (
	"fmt"
	"sort"
	"strings"
)

// GoLiteralLines returns the map entries formatted as Go literal lines,
// suitable for copy-pasting into test case definitions.
//
// Output format:
//
//	"key": value,
//
// Values are formatted using %#v for strings (quoted) and %v for others.
//
// Example:
//
//	m := args.Map{"isZero": false, "value": 5, "name": "hello"}
//	m.GoLiteralLines()
//	// Returns:
//	//   "isZero": false,
//	//   "name":   "hello",
//	//   "value":  5,
func (it Map) GoLiteralLines() []string {
	if len(it) == 0 {
		return []string{}
	}

	keys := make([]string, 0, len(it))

	for k := range it {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// Find max key length for alignment
	maxKeyLen := 0
	for _, k := range keys {
		if len(k) > maxKeyLen {
			maxKeyLen = len(k)
		}
	}

	lines := make([]string, len(keys))

	for i, key := range keys {
		padding := strings.Repeat(" ", maxKeyLen-len(key))
		val := it[key]

		var valStr string

		switch v := val.(type) {
		case string:
			valStr = fmt.Sprintf("%q", v)
		default:
			valStr = fmt.Sprintf("%v", v)
		}

		lines[i] = fmt.Sprintf(
			"%q: %s%s,",
			key,
			padding,
			valStr,
		)
	}

	return lines
}

// GoLiteralString returns the map formatted as a multi-line Go literal block.
//
// Example output:
//
//	"isZero": false,
//	"name":   "hello",
//	"value":  5,
func (it Map) GoLiteralString() string {
	lines := it.GoLiteralLines()

	return strings.Join(lines, "\n")
}
