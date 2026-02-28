# stringslice — String Slice Utilities

## Overview

The `stringslice` package provides a comprehensive set of pure functions for operating on `[]string` and `*[]string` slices. Unlike `corestr` which wraps slices in rich types, `stringslice` works directly with raw Go slices, making it ideal for low-level string manipulation without allocation overhead.

## Key Capabilities

### Creation & Cloning

| Function | Description |
|----------|-------------|
| `Make(cap)` / `MakePtr(cap)` | Create slice with capacity |
| `MakeLen(length)` / `MakeLenPtr(length)` | Create slice with length |
| `MakeDefault()` / `MakeDefaultPtr()` | Create with default capacity |
| `Empty()` / `EmptyPtr()` | Create empty slice |
| `Clone(slice)` | Deep copy a slice |
| `CloneIf(condition, slice)` | Conditional clone |
| `ClonePtr(slice)` | Clone from pointer |
| `CloneUsingCap(slice, cap)` | Clone with specific capacity |
| `CloneSimpleSliceToPointers(slice)` | Clone to pointer slice |

### Element Access

| Function | Description |
|----------|-------------|
| `First(slice)` / `FirstPtr(slice)` | First element (panics if empty) |
| `Last(slice)` / `LastPtr(slice)` | Last element (panics if empty) |
| `FirstOrDefault(slice)` / `FirstOrDefaultPtr(slice)` | First or empty string |
| `LastOrDefault(slice)` / `LastOrDefaultPtr(slice)` | Last or empty string |
| `FirstOrDefaultWith(slice, default)` | First or custom default |
| `IndexAt(slice, index)` | Element at index |
| `SafeIndexAt(slice, index)` | Safe element at index (returns empty if OOB) |
| `SafeIndexAtWith(slice, index, default)` | Safe element with custom default |
| `FirstLastDefault(slice)` | Returns (first, last) with defaults |
| `FirstLastDefaultStatus(slice)` | Returns (first, last, hasItems) |

### Filtering & Transformation

| Function | Description |
|----------|-------------|
| `NonEmpty(slice)` | Remove empty strings |
| `NonEmptyIf(condition, slice)` | Conditional non-empty filter |
| `NonWhitespace(slice)` | Remove whitespace-only strings |
| `NonNullStrings(slice)` | Remove null/empty strings |
| `NonEmptyJoin(slice, sep)` | Join non-empty with separator |
| `NonWhitespaceJoin(slice, sep)` | Join non-whitespace with separator |
| `TrimmedEachWords(slice)` | Trim each element |
| `TrimmedEachWordsIf(condition, slice)` | Conditional trim |
| `SortIf(condition, slice)` | Conditional sort |
| `InPlaceReverse(slice)` | Reverse in-place |

### Splitting & Expanding

| Function | Description |
|----------|-------------|
| `SplitTrimmedNonEmpty(input, sep)` | Split, trim, remove empty |
| `SplitTrimmedNonEmptyAll(input, seps)` | Split by multiple separators |
| `SplitContentsByWhitespace(input)` | Split by whitespace |
| `RegexTrimmedSplitNonEmptyAll(input, pattern)` | Regex-based split |
| `ExpandByFunc(slice, fn)` | Expand each element via function |
| `ExpandBySplit(slice, sep)` | Expand by splitting each element |
| `ExpandBySplits(slice, seps)` | Expand by multiple separators |

### Merging & Appending

| Function | Description |
|----------|-------------|
| `MergeNew(slices...)` | Merge multiple slices into new |
| `MergeNewSimple(a, b)` | Merge two slices |
| `MergeSlicesOfSlices(slices)` | Flatten slice of slices |
| `AppendLineNew(slice, line)` | Append and return new slice |
| `PrependNew(slice, items)` | Prepend items to new slice |
| `PrependLineNew(slice, line)` | Prepend single line |
| `AppendAnyItemsWithStrings(items, strings)` | Append any items with strings |
| `AppendStringsWithAnyItems(strings, items)` | Append strings with any items |
| `AppendStringsWithMainSlice(main, append)` | Append to main slice |

### Safe Indexing

| Function | Description |
|----------|-------------|
| `SafeIndexes(slice, indexes)` | Get elements at multiple indexes |
| `SafeIndexRanges(slice, start, end)` | Get elements in range |
| `SafeRangeItems(slice, start, end)` | Get items in range |
| `SafeIndexAtUsingLastIndex(slice, index)` | Index relative to last |

### Async Processing

| Function | Description |
|----------|-------------|
| `ProcessAsync(slice, fn)` | Process elements concurrently |
| `ProcessOptionAsync(slice, options, fn)` | Async with options |
| `LinesAsyncProcess(slice, fn)` | Process lines concurrently |
| `LinesProcess(slice, fn)` | Process lines sequentially |
| `LinesSimpleProcess(slice, fn)` | Simple line processing |
| `LinesSimpleProcessNoEmpty(slice, fn)` | Process non-empty lines |

### Queries

| Function | Description |
|----------|-------------|
| `IsEmpty(slice)` / `IsEmptyPtr(slice)` | Check if empty |
| `HasAnyItem(slice)` / `HasAnyItemPtr(slice)` | Check if non-empty |
| `LengthOfPointer(slice)` | Length of pointer slice |
| `AllElemLengthSlices(slices)` | Total length across slices |

## How to Extend Safely

- Add new functions as standalone exports following the `FuncName(slice, ...)` pattern.
- Pointer variants should be suffixed with `Ptr`.
- Async variants should use goroutines with proper synchronization.

## Contributors

## Issues for Future Reference
