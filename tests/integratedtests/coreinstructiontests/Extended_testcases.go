package coreinstructiontests

import (
	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

type stringSearchIsMatchCase struct {
	name     string
	search   *coreinstruction.StringSearch
	content  string
	expected bool
}

var stringSearchIsMatchCases = []stringSearchIsMatchCase{
	{
		name: "equal match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.Equal,
			Search:        "hello",
		},
		content:  "hello",
		expected: true,
	},
	{
		name: "equal no match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.Equal,
			Search:        "hello",
		},
		content:  "world",
		expected: false,
	},
	{
		name: "starts with match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.StartsWith,
			Search:        "hel",
		},
		content:  "hello",
		expected: true,
	},
	{
		name: "starts with no match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.StartsWith,
			Search:        "xyz",
		},
		content:  "hello",
		expected: false,
	},
	{
		name: "ends with match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.EndsWith,
			Search:        "llo",
		},
		content:  "hello",
		expected: true,
	},
	{
		name: "contains match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.Contains,
			Search:        "ell",
		},
		content:  "hello",
		expected: true,
	},
	{
		name: "regex match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.Regex,
			Search:        "^he.*o$",
		},
		content:  "hello",
		expected: true,
	},
	{
		name: "not equal match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.NotEqual,
			Search:        "hello",
		},
		content:  "world",
		expected: true,
	},
}
