package mutexbykey

import (
	"sync"

	"github.com/alimtvnetwork/core/constants"
)

type mutexMap struct {
	items map[string]*sync.Mutex
}

var globalMutex = sync.Mutex{}

var items = make(
	map[string]*sync.Mutex,
	constants.ArbitraryCapacity10)

var internalMap = mutexMap{
	items: items,
}

func Get(key string) *sync.Mutex {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	mutex, has := internalMap.items[key]

	if has {
		return mutex
	}

	// not there
	newMutex := &sync.Mutex{}
	internalMap.items[key] = newMutex

	return newMutex
}

// Delete removes the mutex for the given key from the registry.
//
// WARNING: The caller MUST ensure no other goroutine currently holds or is
// waiting on the mutex for this key. Deleting a mutex while it is in use
// will cause a race condition — a new mutex may be created for the same key
// via Get(), allowing two goroutines to operate concurrently on the same
// logical key.
func Delete(key string) bool {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	_, has := internalMap.items[key]

	if has {
		delete(internalMap.items, key)
	}

	return has
}
