package coregenerictests

import (
	"fmt"
	"sync"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================
// Test: AddLock — concurrent goroutines
// ==========================================

func Test_GenericCollection_AddLock_ConcurrentSafety(t *testing.T) {
	const goroutines = 500
	col := coregeneric.EmptyCollection[int]()

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			col.AddLock(idx)
			wg.Done()
		}(i)
	}

	wg.Wait()

	got := col.Length()
	if got != goroutines {
		t.Errorf("AddLock concurrent: expected %d items, got %d", goroutines, got)
	}
}

// ==========================================
// Test: AddsLock — concurrent batch appends
// ==========================================

func Test_GenericCollection_AddsLock_ConcurrentSafety(t *testing.T) {
	const goroutines = 200
	const batchSize = 5
	col := coregeneric.EmptyCollection[string]()

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			batch := make([]string, batchSize)
			for b := 0; b < batchSize; b++ {
				batch[b] = fmt.Sprintf("g%d-b%d", idx, b)
			}
			col.AddsLock(batch...)
			wg.Done()
		}(i)
	}

	wg.Wait()

	expected := goroutines * batchSize
	got := col.Length()
	if got != expected {
		t.Errorf("AddsLock concurrent: expected %d items, got %d", expected, got)
	}
}

// ==========================================
// Test: LengthLock — concurrent reads during writes
// ==========================================

func Test_GenericCollection_LengthLock_ConcurrentReadsWrites(t *testing.T) {
	const writers = 100
	const readers = 100
	col := coregeneric.EmptyCollection[int]()

	wg := sync.WaitGroup{}
	wg.Add(writers + readers)

	// concurrent writers
	for i := 0; i < writers; i++ {
		go func(idx int) {
			col.AddLock(idx)
			wg.Done()
		}(i)
	}

	// concurrent readers
	for i := 0; i < readers; i++ {
		go func() {
			length := col.LengthLock()
			if length < 0 {
				t.Errorf("LengthLock returned negative: %d", length)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	got := col.Length()
	if got != writers {
		t.Errorf("After concurrent reads/writes: expected %d, got %d", writers, got)
	}
}

// ==========================================
// Test: IsEmptyLock — concurrent check with writes
// ==========================================

func Test_GenericCollection_IsEmptyLock_ConcurrentSafety(t *testing.T) {
	const goroutines = 100
	col := coregeneric.EmptyCollection[int]()

	wg := sync.WaitGroup{}
	wg.Add(goroutines * 2)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			col.AddLock(idx)
			wg.Done()
		}(i)
		go func() {
			_ = col.IsEmptyLock() // must not panic
			wg.Done()
		}()
	}

	wg.Wait()

	if col.Length() != goroutines {
		t.Errorf("Expected %d items, got %d", goroutines, col.Length())
	}
}
