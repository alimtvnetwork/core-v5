package args

import (
	"fmt"
	"sort"
)

// CompileToStrings converts all map values to strings using %v format
// and returns sorted "key : value" lines.
//
// This enables test cases to store raw typed values (int, bool, etc.)
// while producing deterministic, human-readable string lines for comparison.
//
// Example:
//
//	m := args.Map{"isZero": false, "value": 5}
//	m.CompileToStrings()
//	// Returns: []string{"isZero : false", "value : 5"}
func (it Map) CompileToStrings() []string {
	if len(it) == 0 {
		return []string{}
	}

	keys := make([]string, 0, len(it))

	for k := range it {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	lines := make([]string, len(keys))

	for i, key := range keys {
		lines[i] = fmt.Sprintf(
			"%s : %v",
			key,
			it[key],
		)
	}

	return lines
}

// CompileToString converts all map values to a single sorted string
// with one "key : value" per line.
//
// Useful for quick debugging or single-string comparison.
func (it Map) CompileToString() string {
	lines := it.CompileToStrings()

	result := ""

	for i, line := range lines {
		if i > 0 {
			result += "\n"
		}

		result += line
	}

	return result
}
