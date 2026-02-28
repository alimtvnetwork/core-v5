# Go Modernization Plan

## Upgrade Targets

### Current State

- **go.mod**: `go 1.17.8`
- **makefile**: `GoVersion=v1.17.8`
- **README**: References Go 1.17.8 as prerequisite

### Desired Target

- **Go 1.22+** (latest stable as of 2026)
- This enables: generics (1.18+), `any` alias, improved error handling, range-over-int (1.22), enhanced stdlib.

### Module and Tooling Updates Needed

1. Update `go.mod`: `go 1.22` (or latest).
2. Update `makefile`: `GoVersion=v1.22`.
3. Update `README.md` prerequisites.
4. Update `go.sum` via `go mod tidy`.
5. Verify all dependencies are compatible.
6. Update CI/Docker images to use Go 1.22+.

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

### Acceptance Criteria

- [ ] Code compiles with `go 1.22+`.
- [ ] All existing tests pass.
- [ ] Readability is maintained or improved.
- [ ] No gratuitous generics — each use must eliminate real duplication.
- [ ] Backward compatibility maintained for external consumers (major version bump if needed).

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

## Migration Strategy

1. **Phase 1**: Update `go.mod` to 1.22, fix any compilation errors.
2. **Phase 2**: Replace `interface{}` with `any` project-wide (mechanical).
3. **Phase 3**: Add generic versions of high-duplication packages (keep old functions as wrappers for compatibility).
4. **Phase 4**: Deprecate old per-type functions, point to generic versions.
5. **Phase 5**: Remove deprecated functions in next major version.
