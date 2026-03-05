package corepayloadtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// testUser is a simple struct for testing TypedPayloadCollection[T].
type testUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var typedCollectionCreationTestCases = []coretestcases.CaseV1{
	{
		Title: "Empty collection has zero length",
		ArrangeInput: args.Map{
			"when":     "creating empty collection",
			"capacity": 0,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "0",    // length
			Second: "true", // isEmpty
		},
	},
	{
		Title: "Collection with capacity has zero length",
		ArrangeInput: args.Map{
			"when":     "creating collection with capacity 10",
			"capacity": 10,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "0",    // length
			Second: "true", // isEmpty
		},
	},
}

var typedCollectionAddTestCases = []coretestcases.CaseV1{
	{
		Title: "Add single item increases length to 1",
		ArrangeInput: args.Map{
			"when":  "adding one user",
			"name":  "Alice",
			"email": "alice@test.com",
			"age":   30,
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "1",     // length
			Second: "false", // isEmpty
			Third:  "Alice", // firstName
		},
	},
	{
		Title: "Add two items increases length to 2",
		ArrangeInput: args.Map{
			"when":   "adding two users",
			"name":   "Bob",
			"email":  "bob@test.com",
			"age":    25,
			"name2":  "Carol",
			"email2": "carol@test.com",
			"age2":   35,
		},
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "2",     // length
			Second: "false", // isEmpty
			Third:  "Bob",   // firstName
			Fourth: "Carol", // secondName
		},
	},
}

var typedCollectionFilterTestCases = []coretestcases.CaseV1{
	{
		Title: "FilterByData returns only matching items",
		ArrangeInput: args.Map{
			"when":      "filtering users by age >= 30",
			"minAge":    30,
			"userCount": 3,
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "2",     // filteredCount
			Second: "Alice", // match1
			Third:  "Carol", // match2
		},
	},
}

var typedCollectionMapTestCases = []coretestcases.CaseV1{
	{
		Title: "MapTypedPayloadData extracts names",
		ArrangeInput: args.Map{
			"when": "mapping users to names",
		},
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "3",     // count
			Second: "Alice", // name0
			Third:  "Bob",   // name1
			Fourth: "Carol", // name2
		},
	},
}

var typedCollectionReduceTestCases = []coretestcases.CaseV1{
	{
		Title: "ReduceTypedPayloadData sums ages",
		ArrangeInput: args.Map{
			"when": "reducing to sum of ages",
		},
		ExpectedInput: "90", // totalAge
	},
}

var typedCollectionGroupTestCases = []coretestcases.CaseV1{
	{
		Title: "GroupTypedPayloadData groups by category",
		ArrangeInput: args.Map{
			"when": "grouping by category name",
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "2", // groupCount
			Second: "1", // group1Size
			Third:  "2", // group2Size
		},
	},
}

var typedCollectionPartitionTestCases = []coretestcases.CaseV1{
	{
		Title: "PartitionTypedPayloads splits by age threshold",
		ArrangeInput: args.Map{
			"when": "partitioning by age >= 30",
		},
		ExpectedInput: args.Two[string, string]{
			First:  "2", // matchCount
			Second: "1", // nonMatchCount
		},
	},
}

var typedCollectionAllDataTestCases = []coretestcases.CaseV1{
	{
		Title: "AllData extracts all typed data",
		ArrangeInput: args.Map{
			"when": "extracting all data",
		},
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "3",     // count
			Second: "Alice", // data0
			Third:  "Bob",   // data1
			Fourth: "Carol", // data2
		},
	},
}

var typedCollectionElementAccessTestCases = []coretestcases.CaseV1{
	{
		Title: "First and Last return correct elements",
		ArrangeInput: args.Map{
			"when": "accessing first and last",
		},
		ExpectedInput: args.Two[string, string]{
			First:  "Alice", // firstName
			Second: "Carol", // lastName
		},
	},
}

var typedCollectionAnyAllTestCases = []coretestcases.CaseV1{
	{
		Title: "AnyTypedPayload returns true when match exists",
		ArrangeInput: args.Map{
			"when": "checking any user named Bob",
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "true",  // anyBob
			Second: "false", // anyDave
			Third:  "true",  // allHaveEmail
		},
	},
}
