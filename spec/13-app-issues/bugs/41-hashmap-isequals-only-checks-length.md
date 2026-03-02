# Hashmap.IsEquals Only Checks Length

## Issue Summary

`coregeneric.Hashmap.IsEquals()` only compares `Length()` between two hashmaps. Two hashmaps with the same size but completely different key-value pairs are incorrectly considered equal.

## Root Cause Analysis

The comment says "values not compared due to any constraint" but keys are also not compared.

## Fix Description

At minimum, iterate over keys and check `other.Has(k)` for key equality. Value comparison is impossible without a `comparable` constraint on V, but key membership can and should be checked.

```go
func (it *Hashmap[K, V]) IsEquals(other *Hashmap[K, V]) bool {
    if it == nil && other == nil {
		return true
	}

	if it == nil || other == nil {
		return false
	}
    if it.Length() != other.Length() { return false }
    for k := range it.items {
        if !other.Has(k) { return false }
    }
    return true
}
```

## Done Checklist

- [ ] Fix `IsEquals` to check key membership
- [ ] Add tests for IsEquals with same-length different-key maps
