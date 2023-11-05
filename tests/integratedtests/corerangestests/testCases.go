package corerangestests

import (
	"reflect"

	"gitlab.com/auk-go/core/coredata/corerange"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/issetter"
)

var (
	validIntRangeTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "2-5, 7-10, 15-20 --- ranges generate for int",
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
				VerifyTypeOf: &coretests.VerifyTypeOf{
					ArrangeInput:  reflect.TypeOf([]corerange.MinMaxInt{}),
					ActualInput:   reflect.TypeOf([]int{}),
					ExpectedInput: reflect.TypeOf([]int{}),
				},
				IsEnable: issetter.True,
			},
		},
	}

	validInt8RangeTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "2-5, 7-10, 15-20 --- ranges generate for int8",
				ArrangeInput: []corerange.MinMaxInt8{
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
				ExpectedInput: []int8{
					2, 3, 4,
					5, 7, 8,
					9, 10, 15,
					16, 17, 18,
					19, 20,
				},
				VerifyTypeOf: &coretests.VerifyTypeOf{
					ArrangeInput:  reflect.TypeOf([]corerange.MinMaxInt8{}),
					ActualInput:   reflect.TypeOf([]int8{}),
					ExpectedInput: reflect.TypeOf([]int8{}),
				},
				IsEnable: issetter.True,
			},
		},
	}
)
