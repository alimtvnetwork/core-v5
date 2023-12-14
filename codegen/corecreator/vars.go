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
			Value: "",
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
				"john doe",
				"John Doe",
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
			CreatorFunc: nil,
		},
		"email": {
			Value: "alim@me.com",
			Possibilities: []string{
				"",
				"alim@me.com",
				" some john @email.com",
				"some.john @email.com       ",
				"some.john @email.com, some.john @email.com       ",
				"some.john @email.com, some.john @email.com       ",
				"valid.email@me.com",
				"invalid#@(&*$.email@email.com",
			},
			CreatorFunc: nil,
		},
	}

	getLenReflectFunc = reflectinternal.SliceConverter.Length
)
