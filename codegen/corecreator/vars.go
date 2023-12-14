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
			Value:         "",
			Possibilities: []string{},
			CreatorFunc:   nil,
		},
	}

	getLenReflectFunc = reflectinternal.SliceConverter.Length
)
