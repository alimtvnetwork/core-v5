package corecreator

import (
	"sync"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

var (
	locker            = sync.Mutex{}
	creatorsMap       = map[string]Creator{}
	defaultCreatorMap = map[string]Item{
		"string": {
			Value: "some line input as a draft string to test out",
			Possibilities: []string{
				"",
				"hello world",
				"hello-world",
				"Hello-World",
				"John",
				"    ",
				"\n\t \v",
				"some name",
				"some value",
				"Jane Doe",
				"Jane;Doe;John, Doe",
				"1",
				"2",
				"3",
				"4",
				"5",
				"6",
				"7",
				"8",
				"   john doe",
			},
			StringOutput: "\"some line input as a draft string to test out\"",
		},
		"email": {
			Value: "alim@me.com",
			Possibilities: []string{
				"",
				"alim@me.com",
				" some john @email.com",
				"some.john @email.com       ",
				"some.john @email.com, some.john @email.com       ",
				"valid.email@me.com",
				"invalid#@(&*$.email@email.com",
			},
			StringOutput: "\"alim@me.com\"",
		},
		"valid-email": {
			Value: "alim@me.com",
			Possibilities: []string{
				"",
				"alim@me.com",
				"valid.email@me.com",
				"jane.doe@email.com",
				"john.doe@me.com",
			},
			StringOutput: "\"alim@valid-email.com\"",
		},
		"[]string": {
			Value: []string{
				"",
				"line 1",
				"line 2",
				"line 3",
				"line 4",
				"line 5",
			},
			Possibilities: [][]string{
				{
					"",
					"\t\v ",
					"\n\t\v ",
					"\n\t\v ",
					"\n\n ",
				},
				{
					"line 1",
					"line 2",
					"line 3",
				},
				{
					"john.doe@email.com",
					"alim.karim@email.com",
					"jane.doe@email.com",
					"alim@auk-go.com",
				},
			},
			StringOutput: `[]string{
				"",
				"line 1",
				"line 2",
				"line 3",
				"line 4",
				"line 5",
			}`,
		},
		"uint": {
			Value: uint(1),
			Possibilities: []uint{
				0,
				1,
				2,
				3,
				6,
			},
		},
		"uint8": {
			Value: uint8(255),
			Possibilities: []uint8{
				0,
				1,
				2,
				3,
				255,
			},
		},
		"uint16": {
			Value: uint16(255),
			Possibilities: []uint16{
				0,
				1,
				2,
				3,
				255,
				32000,
			},
		},
		"uint32": {
			Value: uint32(255),
			Possibilities: []uint32{
				0,
				1,
				2,
				3,
			},
		},
		"uint64": {
			Value: uint64(255),
			Possibilities: []uint64{
				-2,
				-1,
				0,
				1,
				2,
			},
		},
		"int": {
			Value: 1,
			Possibilities: []int{
				-1,
				0,
				1,
				2,
				3,
				4,
				5,
				6,
			},
		},
		"int16": {
			Value: int16(1),
			Possibilities: []int16{
				-3,
				-2,
				-1,
				0,
				1,
				2,
				15000,
			},
		},
		"int32": {
			Value: int32(1),
			Possibilities: []int32{
				-3,
				-2,
				-1,
				0,
				1,
				2,
				15000,
			},
		},
		"int64": {
			Value: int64(1),
			Possibilities: []int64{
				-3,
				-2,
				-1,
				0,
				1,
				2,
				3,
				5,
				7,
			},
		},
		"rune": {
			Value:         []rune("a")[0],
			Possibilities: []rune("some-runes"),
		},
		"float32": {
			Value: float32(1),
			Possibilities: []float32{
				-1.0,
				0,
				0.1,
				1,
				1.5,
				2.0,
			},
		},
		"float64": {
			Value: float64(1),
			Possibilities: []float64{
				-1.0,
				0,
				1.1,
				1,
				1.5,
				2.0,
			},
		},
		"map[string]string": {
			Value: map[string]string{
				"some-key1": "value-1",
				"some-key2": "value-2",
				"some-key3": "value-3",
			},
			Possibilities: []map[string]string{
				{
					"some-key-x-1": "value-1",
					"some-key-x-2": "value-2",
				},
				{
					"some-key-x-1": "value-1",
					"some-key-x-2": "value-2",
				},
			},
		},
	}

	getLenReflectFunc = reflectinternal.SliceConverter.Length
)
