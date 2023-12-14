package corecreator

import (
	"sync"

	"gitlab.com/auk-go/core/coretests/args"
)

var (
	locker                     = sync.Mutex{}
	creatorsMap                = map[string]Creator{}
	defaultCreatorMap args.Map = map[string]interface{}{
		"",
	}
)
