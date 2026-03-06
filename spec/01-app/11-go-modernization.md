# Go Modernization Plan

## Upgrade Targets

### Current State ✅ COMPLETE

- **go.mod**: `go 1.24.0` (upgraded from 1.17.8)
- **makefile**: Updated
- **README**: Updated

### Migration Phases — Status

| Phase | Description | Status |
|-------|-------------|--------|
| Phase 1 | Update `go.mod` to 1.22+, fix compilation errors | ✅ Done (upgraded to 1.24.0) |
| Phase 2 | Replace `interface{}` with `any` project-wide | ✅ Done (zero matches remain) |
| Phase 3 | Add generic versions of high-duplication packages | ✅ Done (`conditional/`, `coredata/coregeneric/`, `corepayload/`) |
| Phase 4 | Deprecate old per-type functions | ✅ Done (deprecation comments on all legacy files) |
| Phase 5 | Remove deprecated functions in next major version | 🔄 Planned |

### Module and Tooling Updates ✅ COMPLETE

1. ~~Update `go.mod`: `go 1.22` (or latest).~~ → Done: `go 1.24.0`
2. ~~Update `makefile`: `GoVersion=v1.22`.~~ → Done
3. ~~Update `README.md` prerequisites.~~ → Done
4. ~~Update `go.sum` via `go mod tidy`.~~ → Done
5. ~~Verify all dependencies are compatible.~~ → Done
6. ~~Update CI/Docker images to use Go 1.22+.~~ → Done

## Generics Adoption Targets

### Rule: Prefer Clarity Over Clever Generics

Generics should be used when they:
- Eliminate clear code duplication across types.
- Maintain or improve readability.
- Don't add unnecessary abstraction.

### Packages Where Generics Reduce Duplication

| Package | Current Pattern | Generics Opportunity |
|---------|----------------|---------------------|
| `conditional/` | Separate functions per type: `Bool()`, `Int()`, `String()`, `Byte()`, etc. | Single `Ternary[T any](cond bool, t, f T) T` |
| `coremath/` | `MaxByte`, `MinByte`, `MaxInt`, etc. | `Max[T constraints.Ordered](a, b T) T` (stdlib has this in 1.21+) |
| `coredata/` | `Integers`, `IntegersDsc`, `PointerIntegers`, etc. | Generic slice types |
| `isany/` | Type-specific null/zero checks | Generic `IsZero[T comparable](v T) bool` |
| `issetter/` | Separate `Min`, `Max`, `MinByte`, `MaxByte` | Generic range checks |
| `converters/` | Per-type converter functions | Some generic converters possible |
| `core.go` | `EmptyIntsPtr()`, `EmptyStringsPtr()`, etc. | `EmptySlicePtr[T any]() *[]T` |
| `coreinterface/` | Many near-identical `Value*Getter` interfaces | `ValueGetter[T any]` |

### Acceptance Criteria ✅ ALL MET

- [x] Code compiles with `go 1.22+` (running 1.24.0).
- [x] All existing tests pass.
- [x] Readability is maintained or improved.
- [x] No gratuitous generics — each use eliminates real duplication.
- [x] Backward compatibility maintained for external consumers (major version bump if needed).

## Example Modernization Targets

### core.go — Root Package

**Before:**
```go
func EmptyIntsPtr() *[]int { return &([]int{}) }
func EmptyStringsPtr() *[]string { return &([]string{}) }
func EmptyFloat32Ptr() *[]float32 { return &([]float32{}) }
// ... 8 more functions
```

**After:**
```go
func EmptySlicePtr[T any]() *[]T {
    s := make([]T, 0)
    return &s
}
```

### conditional/ — Ternary Helpers

**Before:**
```go
func Int(cond bool, trueVal, falseVal int) int { ... }
func String(cond bool, trueVal, falseVal string) string { ... }
func Byte(cond bool, trueVal, falseVal byte) byte { ... }
```

**After:**
```go
func If[T any](cond bool, trueVal, falseVal T) T {
    if cond { return trueVal }
    return falseVal
}
```

### chmodhelper — Error Handling Improvements

- Replace `interface{}` with `any`.
- Use `errors.Join` (Go 1.20+) for multi-error combination instead of manual merge.
- Consider generic result types for operations that return value-or-error.

## Migration Strategy — Status

1. **Phase 1**: Update `go.mod` to 1.22, fix any compilation errors. ✅ Done (1.24.0)
2. **Phase 2**: Replace `interface{}` with `any` project-wide (mechanical). ✅ Done
3. **Phase 3**: Add generic versions of high-duplication packages (keep old functions as wrappers for compatibility). ✅ Done (`conditional/`, `coregeneric/`, `corepayload/`)
4. **Phase 4**: Deprecate old per-type functions, point to generic versions. ✅ Done
5. **Phase 5**: Remove deprecated functions in next major version. 🔄 Planned
