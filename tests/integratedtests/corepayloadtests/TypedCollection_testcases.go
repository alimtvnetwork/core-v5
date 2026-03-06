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
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
	{
		Title: "Collection with capacity has zero length",
		ArrangeInput: args.Map{
			"when":     "creating collection with capacity 10",
			"capacity": 10,
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
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
		ExpectedInput: args.Map{
			"length":    1,
			"isEmpty":   false,
			"firstName": "Alice",
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
		ExpectedInput: args.Map{
			"length":     2,
			"isEmpty":    false,
			"firstName":  "Bob",
			"secondName": "Carol",
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
		ExpectedInput: args.Map{
			"filteredCount": 2,
			"match1":        "Alice",
			"match2":        "Carol",
		},
	},
}

var typedCollectionMapTestCases = []coretestcases.CaseV1{
	{
		Title: "MapTypedPayloadData extracts names",
		ArrangeInput: args.Map{
			"when": "mapping users to names",
		},
		ExpectedInput: args.Map{
			"count": 3,
			"name0": "Alice",
			"name1": "Bob",
			"name2": "Carol",
		},
	},
}

var typedCollectionReduceTestCases = []coretestcases.CaseV1{
	{
		Title: "ReduceTypedPayloadData sums ages",
		ArrangeInput: args.Map{
			"when": "reducing to sum of ages",
		},
		ExpectedInput: args.Map{
			"totalAge": 90,
		},
	},
}

var typedCollectionGroupTestCases = []coretestcases.CaseV1{
	{
		Title: "GroupTypedPayloadData groups by category",
		ArrangeInput: args.Map{
			"when": "grouping by category name",
		},
		ExpectedInput: args.Map{
			"groupCount":      2,
			"juniorGroupSize": 1,
			"seniorGroupSize": 2,
		},
	},
}

var typedCollectionPartitionTestCases = []coretestcases.CaseV1{
	{
		Title: "PartitionTypedPayloads splits by age threshold",
		ArrangeInput: args.Map{
			"when": "partitioning by age >= 30",
		},
		ExpectedInput: args.Map{
			"seniorCount": 2,
			"juniorCount": 1,
		},
	},
}

var typedCollectionAllDataTestCases = []coretestcases.CaseV1{
	{
		Title: "AllData extracts all typed data",
		ArrangeInput: args.Map{
			"when": "extracting all data",
		},
		ExpectedInput: args.Map{
			"count": 3,
			"data0": "Alice",
			"data1": "Bob",
			"data2": "Carol",
		},
	},
}

var typedCollectionElementAccessTestCases = []coretestcases.CaseV1{
	{
		Title: "First and Last return correct elements",
		ArrangeInput: args.Map{
			"when": "accessing first and last",
		},
		ExpectedInput: args.Map{
			"firstName": "Alice",
			"lastName":  "Carol",
		},
	},
}

var typedCollectionAnyAllTestCases = []coretestcases.CaseV1{
	{
		Title: "AnyTypedPayload returns true when match exists",
		ArrangeInput: args.Map{
			"when": "checking any user named Bob",
		},
		ExpectedInput: args.Map{
			"anyBob":        true,
			"anyNonexistent": false,
			"allAreParsed":  true,
		},
	},
}
