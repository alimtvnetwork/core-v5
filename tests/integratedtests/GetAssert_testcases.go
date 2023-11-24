package integratedtests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	arrangeTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]interface{}{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	twoArgsTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.Two{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	oneFuncTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.OneFunc{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	commonType = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(args.Map{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	quickTestCases = []coretestcases.CaseV1{
		{
			Title: "Quick output as gherkins format",
			ArrangeInput: args.Map{
				"when":    "some title, or case when",
				"actual":  "actual rec",
				"expect":  "expected item",
				"counter": 3,
			},
			ExpectedInput: []string{
				"----------------------",
				"3 )  When: some title, or case when,",
				"   Actual: actual rec,",
				" Expected: expected item",
			},
			VerifyTypeOf: commonType,
		},
	}

	sortedArrayTestCases = []coretestcases.CaseV1{
		{
			Title: "SortedArray output verification",
			ArrangeInput: args.Map{
				"isPrint": false,
				"isSort":  true,
				"message": "some message alim, knows, who, do you know --- #alim",
			},
			ExpectedInput: []string{
				"#alim",
				"---",
				"alim,",
				"do",
				"know",
				"knows,",
				"message",
				"some",
				"who,",
				"you",
			},
			VerifyTypeOf: commonType,
		},
	}

	sortedMessageTestCases = []coretestcases.CaseV1{
		{
			Title: "SortedMessage output verification",
			ArrangeInput: args.Map{
				"isPrint": false,
				"message": "some message alim, knows, who, do you know --- #alim",
				"joiner":  " | ",
			},
			ExpectedInput: []string{
				"#alim | --- | alim, | do | know | knows, | message | some | who, | you",
			},
			VerifyTypeOf: commonType,
		},
	}

	stringsToSpaceStringTestCases = []coretestcases.CaseV1{
		{
			Title: "StringsToSpaceString output verification",
			ArrangeInput: args.Map{
				"spaceCount": 4,
				"lines": []string{
					"#alim",
					"---",
					"alim,",
					"do",
					"",
					"know",
					"when",
					"you",
					"type,",
					"lines",
				},
			},
			ExpectedInput: []string{
				"    #alim",
				"    ---",
				"    alim,",
				"    do",
				"    ",
				"    know",
				"    when",
				"    you",
				"    type,",
				"    lines",
			},
			VerifyTypeOf: commonType,
		},
	}

	toStringsTestCases = []coretestcases.CaseV1{
		{
			Title: "giving string - output split to lines by newlines",
			ArrangeInput: args.Map{
				"any": "some string contains\nnewline\nin between",
			},
			ExpectedInput: []string{
				"some string contains",
				"newline",
				"in between",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving []string or slice string - outputs as is.",
			ArrangeInput: args.Map{
				"any": []string{
					"having exact lines will output",
					"as the lines",
					"were.",
					"no change.",
				},
			},
			ExpectedInput: []string{
				"having exact lines will output",
				"as the lines",
				"were.",
				"no change.",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving []string{} outputs as it is - empty string has no issues.",
			ArrangeInput: args.Map{
				"any": []string{},
			},
			ExpectedInput: []string{},
			VerifyTypeOf:  commonType,
		},
		{
			Title: "giving []interface - json convert and returns as it is.",
			ArrangeInput: args.Map{
				"any": []interface{}{
					"passed []interface, which is",
					"any but lines of any",
					"gets no converted and",
					"returns as it is",
				},
			},
			ExpectedInput: []string{
				"passed []interface, which is",
				"any but lines of any",
				"gets no converted and",
				"returns as it is",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving map[string]interface{} - converts to lines and returns sorted lines.",
			ArrangeInput: args.Map{
				"any": map[string]interface{}{
					"line 1": "passed map[string]interface, which is",
					"line 2": "any but keys as is but converts",
					"line 3": "value to SmartJSON and",
					"line 4": map[string]interface{}{
						"sub line 1": "returns",
						"sub line 2": -5,
					},
					"line 5": []string{
						"some line 1",
						"some line 2",
					},
					"line 6": []interface{}{
						args.One{
							First:  "line 6.1 first",
							Expect: "line 6.1 expect",
						},
						"some line 2",
					},
				},
			},
			ExpectedInput: []string{
				"line 1 : passed map[string]interface, which is",
				"line 2 : any but keys as is but converts",
				"line 3 : value to SmartJSON and",
				"line 4 : {\n  \"sub line 1\": \"returns\",\n  \"sub line 2\": -5\n}",
				"line 5 : some line 1\nsome line 2",
				"line 6 : [\n  {\n    \"First\": \"line 6.1 first\",\n    \"Expect\": \"line 6.1 expect\"\n  },\n  \"some line 2\"\n]",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving map[string]string - converts to lines and returns sorted lines.",
			ArrangeInput: args.Map{
				"any": map[string]string{
					"line 1": "passed map[string]string, which is",
					"line 2": "any but keys as is but converts",
					"line 3": "value to as is and",
					"line 4": "returns simple line",
				},
			},
			ExpectedInput: []string{
				"line 1 : passed map[string]string, which is",
				"line 2 : any but keys as is but converts",
				"line 3 : value to as is and",
				"line 4 : returns simple line",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving map[interface{}]interface{} - converts to lines and returns sorted lines.",
			ArrangeInput: args.Map{
				"any": map[interface{}]interface{}{
					0:        "it is 0",
					1:        "it is 1",
					"line 1": "passed map[interface{}]interface{}, which is",
					"line 2": "converts both keys and values to",
					"line 3": "SmartJSON and returns it.",
					"line 4": map[string]interface{}{
						"sub line 1": "returns",
						"sub line 2": -5,
					},
					"line 5": []string{
						"some line 1",
						"some line 2",
					},
					"line 6": []interface{}{
						args.One{
							First:  "line 6.1 first",
							Expect: "line 6.1 expect",
						},
						"some line 2",
					},
					args.One{
						First: "line 7 - key",
					}: args.One{
						First:  "line 7 - value",
						Expect: "line 7 - value.expect",
					},
				},
			},
			ExpectedInput: []string{
				"0 : it is 0",
				"1 : it is 1",
				"line 1 : passed map[interface{}]interface{}, which is",
				"line 2 : converts both keys and values to",
				"line 3 : SmartJSON and returns it.",
				"line 4 : {\n  \"sub line 1\": \"returns\",\n  \"sub line 2\": -5\n}",
				"line 5 : some line 1\nsome line 2",
				"line 6 : [\n  {\n    \"First\": \"line 6.1 first\",\n    \"Expect\": \"line 6.1 expect\"\n  },\n  \"some line 2\"\n]",
				"{\n  \"First\": \"line 7 - key\"\n} : {\n  \"First\": \"line 7 - value\",\n  \"Expect\": \"line 7 - value.expect\"\n}",
			},
			VerifyTypeOf: commonType,
		},
	}
)
