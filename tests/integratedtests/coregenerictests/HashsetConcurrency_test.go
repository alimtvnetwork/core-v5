package coregenerictests

import (
	"fmt"
	"sync"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================
// Test: Hashset AddLock — concurrent goroutines
// ==========================================

func Test_GenericHashset_AddLock_ConcurrentSafety(t *testing.T) {
	const goroutines = 500
	hs := coregeneric.NewHashset[int](goroutines)

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			hs.AddLock(idx)
			wg.Done()
		}(i)
	}

	wg.Wait()

	got := hs.Length()
	if got != goroutines {
		t.Errorf("AddLock concurrent: expected %d items, got %d", goroutines, got)
	}
}

// ==========================================
// Test: Hashset AddSliceLock — concurrent batch adds
// ==========================================

func Test_GenericHashset_AddSliceLock_ConcurrentSafety(t *testing.T) {
	const goroutines = 100
	const batchSize = 10
	hs := coregeneric.NewHashset[string](goroutines * batchSize)

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			batch := make([]string, batchSize)
			for b := 0; b < batchSize; b++ {
				batch[b] = fmt.Sprintf("g%d-b%d", idx, b)
			}
			hs.AddSliceLock(batch)
			wg.Done()
		}(i)
	}

	wg.Wait()

	expected := goroutines * batchSize
	got := hs.Length()
	if got != expected {
		t.Errorf("AddSliceLock concurrent: expected %d unique items, got %d", expected, got)
	}
}

// ==========================================
// Test: Hashset ContainsLock — concurrent reads with writes
// ==========================================

func Test_GenericHashset_ContainsLock_ConcurrentReadsWrites(t *testing.T) {
	const writers = 200
	const readers = 200
	hs := coregeneric.NewHashset[int](writers)

	wg := sync.WaitGroup{}
	wg.Add(writers + readers)

	// concurrent writers
	for i := 0; i < writers; i++ {
		go func(idx int) {
			hs.AddLock(idx)
			wg.Done()
		}(i)
	}

	// concurrent readers
	for i := 0; i < readers; i++ {
		go func(idx int) {
			_ = hs.ContainsLock(idx) // must not panic
			wg.Done()
		}(i)
	}

	wg.Wait()

	got := hs.Length()
	if got != writers {
		t.Errorf("After concurrent reads/writes: expected %d, got %d", writers, got)
	}
}

// ==========================================
// Test: Hashset RemoveLock — concurrent add and remove
// ==========================================

func Test_GenericHashset_RemoveLock_ConcurrentSafety(t *testing.T) {
	const items = 500
	hs := coregeneric.NewHashset[int](items)

	// Pre-populate
	for i := 0; i < items; i++ {
		hs.Add(i)
	}

	wg := sync.WaitGroup{}
	wg.Add(items)

	// Concurrently remove all items
	for i := 0; i < items; i++ {
		go func(idx int) {
			hs.RemoveLock(idx)
			wg.Done()
		}(i)
	}

	wg.Wait()

	got := hs.Length()
	if got != 0 {
		t.Errorf("RemoveLock concurrent: expected 0 items, got %d", got)
	}
}

// ==========================================
// Test: Hashset LengthLock — concurrent reads during mutations
// ==========================================

func Test_GenericHashset_LengthLock_ConcurrentSafety(t *testing.T) {
	const writers = 100
	const readers = 100
	hs := coregeneric.NewHashset[int](writers)

	wg := sync.WaitGroup{}
	wg.Add(writers + readers)

	for i := 0; i < writers; i++ {
		go func(idx int) {
			hs.AddLock(idx)
			wg.Done()
		}(i)
	}

	for i := 0; i < readers; i++ {
		go func() {
			length := hs.LengthLock()
			if length < 0 {
				t.Errorf("LengthLock returned negative: %d", length)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	if hs.Length() != writers {
		t.Errorf("Expected %d, got %d", writers, hs.Length())
	}
}

// ==========================================
// Test: Hashset IsEmptyLock — concurrent check
// ==========================================

func Test_GenericHashset_IsEmptyLock_ConcurrentSafety(t *testing.T) {
	const goroutines = 100
	hs := coregeneric.NewHashset[int](goroutines)

	wg := sync.WaitGroup{}
	wg.Add(goroutines * 2)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			hs.AddLock(idx)
			wg.Done()
		}(i)
		go func() {
			_ = hs.IsEmptyLock() // must not panic
			wg.Done()
		}()
	}

	wg.Wait()

	if hs.Length() != goroutines {
		t.Errorf("Expected %d items, got %d", goroutines, hs.Length())
	}
}
