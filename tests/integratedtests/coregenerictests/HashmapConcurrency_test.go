package coregenerictests

import (
	"fmt"
	"sync"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================
// Test: Hashmap SetLock — concurrent goroutines
// ==========================================

func Test_GenericHashmap_SetLock_ConcurrentSafety(t *testing.T) {
	const goroutines = 500
	hm := coregeneric.NewHashmap[int, string](goroutines)

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			hm.SetLock(idx, fmt.Sprintf("val-%d", idx))
			wg.Done()
		}(i)
	}

	wg.Wait()

	got := hm.Length()
	if got != goroutines {
		t.Errorf("SetLock concurrent: expected %d entries, got %d", goroutines, got)
	}
}

// ==========================================
// Test: Hashmap GetLock — concurrent reads with writes
// ==========================================

func Test_GenericHashmap_GetLock_ConcurrentReadsWrites(t *testing.T) {
	const writers = 200
	const readers = 200
	hm := coregeneric.NewHashmap[int, int](writers)

	wg := sync.WaitGroup{}
	wg.Add(writers + readers)

	// concurrent writers
	for i := 0; i < writers; i++ {
		go func(idx int) {
			hm.SetLock(idx, idx*10)
			wg.Done()
		}(i)
	}

	// concurrent readers
	for i := 0; i < readers; i++ {
		go func(idx int) {
			_, _ = hm.GetLock(idx) // must not panic
			wg.Done()
		}(i)
	}

	wg.Wait()

	got := hm.Length()
	if got != writers {
		t.Errorf("After concurrent reads/writes: expected %d, got %d", writers, got)
	}
}

// ==========================================
// Test: Hashmap ContainsLock — concurrent reads
// ==========================================

func Test_GenericHashmap_ContainsLock_ConcurrentSafety(t *testing.T) {
	const writers = 200
	const readers = 200
	hm := coregeneric.NewHashmap[string, int](writers)

	wg := sync.WaitGroup{}
	wg.Add(writers + readers)

	for i := 0; i < writers; i++ {
		go func(idx int) {
			hm.SetLock(fmt.Sprintf("key-%d", idx), idx)
			wg.Done()
		}(i)
	}

	for i := 0; i < readers; i++ {
		go func(idx int) {
			_ = hm.ContainsLock(fmt.Sprintf("key-%d", idx)) // must not panic
			wg.Done()
		}(i)
	}

	wg.Wait()

	if hm.Length() != writers {
		t.Errorf("Expected %d, got %d", writers, hm.Length())
	}
}

// ==========================================
// Test: Hashmap RemoveLock — concurrent removes
// ==========================================

func Test_GenericHashmap_RemoveLock_ConcurrentSafety(t *testing.T) {
	const items = 500
	hm := coregeneric.NewHashmap[int, string](items)

	// Pre-populate
	for i := 0; i < items; i++ {
		hm.Set(i, fmt.Sprintf("val-%d", i))
	}

	wg := sync.WaitGroup{}
	wg.Add(items)

	for i := 0; i < items; i++ {
		go func(idx int) {
			hm.RemoveLock(idx)
			wg.Done()
		}(i)
	}

	wg.Wait()

	got := hm.Length()
	if got != 0 {
		t.Errorf("RemoveLock concurrent: expected 0 entries, got %d", got)
	}
}

// ==========================================
// Test: Hashmap LengthLock — concurrent reads during mutations
// ==========================================

func Test_GenericHashmap_LengthLock_ConcurrentSafety(t *testing.T) {
	const writers = 100
	const readers = 100
	hm := coregeneric.NewHashmap[int, int](writers)

	wg := sync.WaitGroup{}
	wg.Add(writers + readers)

	for i := 0; i < writers; i++ {
		go func(idx int) {
			hm.SetLock(idx, idx)
			wg.Done()
		}(i)
	}

	for i := 0; i < readers; i++ {
		go func() {
			length := hm.LengthLock()
			if length < 0 {
				t.Errorf("LengthLock returned negative: %d", length)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	if hm.Length() != writers {
		t.Errorf("Expected %d, got %d", writers, hm.Length())
	}
}

// ==========================================
// Test: Hashmap IsEmptyLock — concurrent check
// ==========================================

func Test_GenericHashmap_IsEmptyLock_ConcurrentSafety(t *testing.T) {
	const goroutines = 100
	hm := coregeneric.NewHashmap[int, int](goroutines)

	wg := sync.WaitGroup{}
	wg.Add(goroutines * 2)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			hm.SetLock(idx, idx)
			wg.Done()
		}(i)
		go func() {
			_ = hm.IsEmptyLock() // must not panic
			wg.Done()
		}()
	}

	wg.Wait()

	if hm.Length() != goroutines {
		t.Errorf("Expected %d entries, got %d", goroutines, hm.Length())
	}
}

// ==========================================
// Test: Hashmap mixed SetLock+GetLock+RemoveLock
// ==========================================

func Test_GenericHashmap_MixedOperations_ConcurrentSafety(t *testing.T) {
	const items = 300
	hm := coregeneric.NewHashmap[int, string](items)

	// Pre-populate half
	for i := 0; i < items/2; i++ {
		hm.Set(i, fmt.Sprintf("initial-%d", i))
	}

	wg := sync.WaitGroup{}
	wg.Add(items * 3)

	// Writers: set keys items..items*2
	for i := items; i < items*2; i++ {
		go func(idx int) {
			hm.SetLock(idx, fmt.Sprintf("new-%d", idx))
			wg.Done()
		}(i)
	}

	// Readers: read keys 0..items
	for i := 0; i < items; i++ {
		go func(idx int) {
			_, _ = hm.GetLock(idx)
			wg.Done()
		}(i)
	}

	// Removers: remove keys 0..items/2 (the pre-populated ones)
	for i := 0; i < items; i++ {
		go func(idx int) {
			if idx < items/2 {
				hm.RemoveLock(idx)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	// After: pre-populated (0..149) removed, new (300..599) added
	got := hm.Length()
	if got != items {
		t.Errorf("Mixed ops concurrent: expected %d entries, got %d", items, got)
	}
}
