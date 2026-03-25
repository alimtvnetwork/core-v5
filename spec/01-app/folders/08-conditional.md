# conditional

## Folder Purpose

Provides generic ternary-style helper functions, nil-safe defaults, and batch function execution for Go, replacing the lack of a ternary operator.

## Responsibilities

1. Generic base functions: `If[T]`, `IfFunc[T]`, `IfTrueFunc[T]`, `IfSlice[T]`, `IfPtr[T]`, `NilDef[T]`, `ValueOrZero[T]`, etc.
2. Typed convenience wrappers for 15 primitive types (11 functions each): `IfInt`, `IfString`, `NilDefBool`, etc.
3. Batch function execution: `VoidFunctions`, `Functions`, `ErrorFunc`, `TypedErrorFunctionsExecuteResults`.
4. String-specific helpers: `NilOrEmptyStr`, `NilOrEmptyStrPtr`, `StringDefault`.
5. Conditional setters: `Setter`, `SetterDefault`.

## Usage Example

```go
result := conditional.IfInt(true, 2, 7)   // returns 2
result := conditional.IfInt(false, 2, 7)  // returns 7
name := conditional.IfString(len(s) > 0, s, "default")
val := conditional.NilDefInt(intPtr, 42)  // dereference or 42
```

## Related Docs

- [conditional README](../../conditional/README.md)
- [Typed Wrappers Reference](../../conditional/typed-wrappers.md)
- [Repo Overview](../00-repo-overview.md)
