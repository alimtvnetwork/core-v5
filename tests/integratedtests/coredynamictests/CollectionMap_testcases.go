package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// Map — int to string
// ==========================================

var mapIntToStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Map transforms int collection to string collection",
		ArrangeInput: args.Map{
			"when":  "given ints [1, 2, 3]",
			"items": []int{1, 2, 3},
		},
		ExpectedInput: []string{
			"3",
			"#1",
			"#2",
			"#3",
		},
	},
}

var mapEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Map on empty collection returns empty",
		ArrangeInput: args.Map{
			"when": "given empty int collection",
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}

var mapNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Map on nil collection returns empty",
		ArrangeInput: args.Map{
			"when": "given nil collection",
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}

// ==========================================
// Map — string to int (length)
// ==========================================

var mapStringToIntTestCases = []coretestcases.CaseV1{
	{
		Title: "Map transforms strings to their lengths",
		ArrangeInput: args.Map{
			"when":  "given strings [hi, hello, x]",
			"items": []string{"hi", "hello", "x"},
		},
		ExpectedInput: []string{
			"3",
			"2",
			"5",
			"1",
		},
	},
}

// ==========================================
// FlatMap — string to chars
// ==========================================

var flatMapTestCases = []coretestcases.CaseV1{
	{
		Title: "FlatMap flattens nested slices into single collection",
		ArrangeInput: args.Map{
			"when":  "given strings split to chars",
			"items": []string{"ab", "cd"},
		},
		ExpectedInput: []string{
			"4",
			"a",
			"b",
			"c",
			"d",
		},
	},
}

var flatMapEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "FlatMap on empty collection returns empty",
		ArrangeInput: args.Map{
			"when": "given empty collection",
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}

// ==========================================
// Reduce — sum
// ==========================================

var reduceSumTestCases = []coretestcases.CaseV1{
	{
		Title: "Reduce sums int collection",
		ArrangeInput: args.Map{
			"when":  "given ints [10, 20, 30]",
			"items": []int{10, 20, 30},
		},
		ExpectedInput: []string{
			"60",
		},
	},
}

var reduceEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Reduce on empty returns initial value",
		ArrangeInput: args.Map{
			"when":    "given empty int collection",
			"initial": 99,
		},
		ExpectedInput: []string{
			"99",
		},
	},
}

// ==========================================
// Reduce — string concat
// ==========================================

var reduceConcatTestCases = []coretestcases.CaseV1{
	{
		Title: "Reduce concatenates strings",
		ArrangeInput: args.Map{
			"when":  "given strings [a, b, c]",
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: []string{
			"a-b-c",
		},
	},
}

// ==========================================
// Map chained — Map then Filter
// ==========================================

var mapThenFilterTestCases = []coretestcases.CaseV1{
	{
		Title: "Map then Filter produces correct subset",
		ArrangeInput: args.Map{
			"when":  "given ints mapped to doubled then filtered > 5",
			"items": []int{1, 2, 3, 4, 5},
		},
		ExpectedInput: []string{
			"3",
			"6",
			"8",
			"10",
		},
	},
}
