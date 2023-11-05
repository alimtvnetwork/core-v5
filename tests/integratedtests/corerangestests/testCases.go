package corerangestests

import (
	"gitlab.com/auk-go/core/coredata/corerange"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/issetter"
)

var (
	validIntRangeTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "2-5, 7-10, 15-20",
				ArrangeInput: []corerange.MinMaxInt{
					{
						Min: 2,
						Max: 5,
					},
					{
						Min: 7,
						Max: 10,
					},
					{
						Min: 15,
						Max: 20,
					},
				},
				ExpectedInput: []int{
					2, 3, 4,
					5, 7, 8,
					9, 10, 15,
					16, 17, 18,
					19, 20,
				},
				VerifyTypeOf:    nil,
				IsEnable:        issetter.True,
				HasError:        false,
				IsValidateError: false,
			},
			IsExpectingError: false,
			HasPanic:         false,
		},
	}
)
