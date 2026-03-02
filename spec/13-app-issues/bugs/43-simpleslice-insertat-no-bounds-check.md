# SimpleSlice.InsertAt Missing Bounds Check

## Issue Summary

`coregeneric.SimpleSlice.InsertAt()` does not validate the `index` parameter. Negative indices or indices beyond the slice length cause a runtime panic with no meaningful error message.

## Fix Description

Add bounds validation similar to `Collection.RemoveAt()`:

```go
func (it *SimpleSlice[T]) InsertAt(index int, item T) *SimpleSlice[T] {
    if index < 0 || index > it.Length() {
        return it // or panic with descriptive message
    }
    // ... existing logic
}
```

## Done Checklist

- [ ] Add bounds check to `InsertAt`
- [ ] Add test cases for negative index, out-of-bounds index
