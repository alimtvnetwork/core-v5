package chmodhelper

import "sync"

var (
	SimpleFileWriter = simpleFileWriter{}
	New              = newCreator{}
	ChmodApply       = chmodApplier{}
	ChmodVerify      = chmodVerifier{}
	globalMutex      = sync.Mutex{}
)
