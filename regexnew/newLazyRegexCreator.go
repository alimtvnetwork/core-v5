package regexnew

import "regexp"

type newLazyRegexCreator struct{}

// New
//
//	used to create as vars
func (it newLazyRegexCreator) New(
	pattern string,
) *LazyRegex {
	lazyRegex, _ := lazyRegexOnceMap.CreateOrExisting(
		pattern)

	return lazyRegex
}

// NewLock
//
//	used to generate inside method
func (it newLazyRegexCreator) NewLock(
	pattern string,
) *LazyRegex {
	lazyRegex, _ := lazyRegexOnceMap.CreateOrExistingLock(
		pattern)

	return lazyRegex
}

func (it newLazyRegexCreator) TwoLock(
	pattern, secondPattern string,
) (first, second *LazyRegex) {
	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	first = it.New(pattern)
	second = it.New(secondPattern)

	return first, second
}

func (it newLazyRegexCreator) ManyUsingLock(
	patterns ...string,
) (patternsKeyAsMap map[string]*LazyRegex) {
	if len(patterns) == 0 {
		return map[string]*LazyRegex{}
	}

	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	patternsKeyAsMap = make(
		map[string]*LazyRegex,
		len(patterns))

	for _, pattern := range patterns {
		patternsKeyAsMap[pattern] = it.New(pattern)
	}

	return patternsKeyAsMap
}

func (it newLazyRegexCreator) AllPatternsMap() map[string]*LazyRegex {
	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	return lazyRegexOnceMap.items
}

// NewLockIf
//
//	used to generate inside method
//	lock must be performed when called from method.
func (it newLazyRegexCreator) NewLockIf(
	isLock bool,
	pattern string,
) *LazyRegex {
	lazyRegex, _ := lazyRegexOnceMap.CreateOrExistingLockIf(
		isLock,
		pattern)

	return lazyRegex
}

// NewWithCompiler creates a LazyRegex using a custom compiler function.
func (it newLazyRegexCreator) NewWithCompiler(
	pattern string,
	compiler func(pattern string) (*regexp.Regexp, error),
) *LazyRegex {
	return lazyRegexOnceMap.createLazyRegex(pattern, compiler)
}

// Count returns the number of cached lazy regex entries.
func (it newLazyRegexCreator) Count() int {
	return lazyRegexOnceMap.Length()
}

// CountLock returns the count with mutex lock.
func (it newLazyRegexCreator) CountLock() int {
	return lazyRegexOnceMap.LengthLock()
}

// IsEmpty returns true if no lazy regex entries are cached.
func (it newLazyRegexCreator) IsEmpty() bool {
	return lazyRegexOnceMap.IsEmpty()
}

// IsEmptyLock returns IsEmpty with mutex lock.
func (it newLazyRegexCreator) IsEmptyLock() bool {
	return lazyRegexOnceMap.IsEmptyLock()
}

// HasAnyItem returns true if at least one lazy regex is cached.
func (it newLazyRegexCreator) HasAnyItem() bool {
	return lazyRegexOnceMap.HasAnyItem()
}

// HasAnyItemLock returns HasAnyItem with mutex lock.
func (it newLazyRegexCreator) HasAnyItemLock() bool {
	return lazyRegexOnceMap.HasAnyItemLock()
}

// Has returns true if a lazy regex with the given pattern exists.
func (it newLazyRegexCreator) Has(pattern string) bool {
	return lazyRegexOnceMap.Has(pattern)
}

// HasLock returns Has with mutex lock.
func (it newLazyRegexCreator) HasLock(pattern string) bool {
	return lazyRegexOnceMap.HasLock(pattern)
}
