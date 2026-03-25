# Typed Wrappers Reference

Complete reference for all 15 typed convenience wrapper files in `conditional/`.
Each file provides the same function set, wrapping the generic base from `generic.go`.

## Function Set Template

Every typed file provides the following functions (shown with `{T}` placeholder and `{type}` for the Go type):

```
If{T}(cond bool, t, f {type}) {type}
IfFunc{T}(cond bool, tF, fF func() {type}) {type}
IfTrueFunc{T}(cond bool, tF func() {type}) {type}
IfSlice{T}(cond bool, t, f []{type}) []{type}
IfPtr{T}(cond bool, t, f *{type}) *{type}
NilDef{T}(ptr *{type}, def {type}) {type}              ¹
NilDefPtr{T}(ptr *{type}, def {type}) *{type}
NilVal{T}(ptr *{type}, onNil, onNonNil {type}) {type}
NilValPtr{T}(ptr *{type}, onNil, onNonNil {type}) *{type}
ValueOrZero{T}(ptr *{type}) {type}
PtrOrZero{T}(ptr *{type}) *{type}
```

¹ `NilDef{T}` is omitted for `bool`, `int`, and `byte` due to naming conflicts
with deprecated functions. Use `NilDef[T](ptr, def)` directly for those types.

### Deprecated Aliases (also present in typed files)

```
NilDeref{T}(ptr *{type}) {type}       → use ValueOrZero{T}
NilDerefPtr{T}(ptr *{type}) *{type}   → use PtrOrZero{T}
```

Some types also include deprecated `IfSlicePtr{T}` and `IfSlicePtrFunc{T}` for
backward compatibility with legacy code.

---

## bool (`typed_bool.go`)

```go
conditional.IfBool(cond, trueVal, falseVal)
conditional.IfFuncBool(cond, trueFunc, falseFunc)
conditional.IfTrueFuncBool(cond, trueFunc)
conditional.IfSliceBool(cond, trueSlice, falseSlice)
conditional.IfPtrBool(cond, truePtr, falsePtr)
conditional.NilDefPtrBool(ptr, defVal)          // NilDefBool omitted ¹
conditional.NilValBool(ptr, onNil, onNonNil)
conditional.NilValPtrBool(ptr, onNil, onNonNil)
conditional.ValueOrZeroBool(ptr)
conditional.PtrOrZeroBool(ptr)
```

> ¹ Use `NilDef[bool](ptr, defVal)` directly.

## byte (`typed_byte.go`)

```go
conditional.IfByte(cond, trueVal, falseVal)
conditional.IfFuncByte(cond, trueFunc, falseFunc)
conditional.IfTrueFuncByte(cond, trueFunc)
conditional.IfSliceByte(cond, trueSlice, falseSlice)
conditional.IfPtrByte(cond, truePtr, falsePtr)
conditional.NilDefPtrByte(ptr, defVal)          // NilDefByte omitted ¹
conditional.NilValByte(ptr, onNil, onNonNil)
conditional.NilValPtrByte(ptr, onNil, onNonNil)
conditional.ValueOrZeroByte(ptr)
conditional.PtrOrZeroByte(ptr)
```

> ¹ Use `NilDef[byte](ptr, defVal)` directly.

## string (`typed_string.go`)

```go
conditional.IfString(cond, trueVal, falseVal)
conditional.IfFuncString(cond, trueFunc, falseFunc)
conditional.IfTrueFuncString(cond, trueFunc)
conditional.IfSliceString(cond, trueSlice, falseSlice)
conditional.IfPtrString(cond, truePtr, falsePtr)
conditional.NilDefString(ptr, defVal)
conditional.NilDefPtrString(ptr, defVal)
conditional.NilValString(ptr, onNil, onNonNil)
conditional.NilValPtrString(ptr, onNil, onNonNil)
conditional.ValueOrZeroString(ptr)
conditional.PtrOrZeroString(ptr)
```

## int (`typed_int.go`)

```go
conditional.IfInt(cond, trueVal, falseVal)
conditional.IfFuncInt(cond, trueFunc, falseFunc)
conditional.IfTrueFuncInt(cond, trueFunc)
conditional.IfSliceInt(cond, trueSlice, falseSlice)
conditional.IfPtrInt(cond, truePtr, falsePtr)
conditional.NilDefPtrInt(ptr, defVal)           // NilDefInt omitted ¹
conditional.NilValInt(ptr, onNil, onNonNil)
conditional.NilValPtrInt(ptr, onNil, onNonNil)
conditional.ValueOrZeroInt(ptr)
conditional.PtrOrZeroInt(ptr)
```

> ¹ Use `NilDef[int](ptr, defVal)` directly.

## int8 (`typed_int8.go`)

```go
conditional.IfInt8(cond, trueVal, falseVal)
conditional.IfFuncInt8(cond, trueFunc, falseFunc)
conditional.IfTrueFuncInt8(cond, trueFunc)
conditional.IfSliceInt8(cond, trueSlice, falseSlice)
conditional.IfPtrInt8(cond, truePtr, falsePtr)
conditional.NilDefInt8(ptr, defVal)
conditional.NilDefPtrInt8(ptr, defVal)
conditional.NilValInt8(ptr, onNil, onNonNil)
conditional.NilValPtrInt8(ptr, onNil, onNonNil)
conditional.ValueOrZeroInt8(ptr)
conditional.PtrOrZeroInt8(ptr)
```

## int16 (`typed_int16.go`)

```go
conditional.IfInt16(cond, trueVal, falseVal)
conditional.IfFuncInt16(cond, trueFunc, falseFunc)
conditional.IfTrueFuncInt16(cond, trueFunc)
conditional.IfSliceInt16(cond, trueSlice, falseSlice)
conditional.IfPtrInt16(cond, truePtr, falsePtr)
conditional.NilDefInt16(ptr, defVal)
conditional.NilDefPtrInt16(ptr, defVal)
conditional.NilValInt16(ptr, onNil, onNonNil)
conditional.NilValPtrInt16(ptr, onNil, onNonNil)
conditional.ValueOrZeroInt16(ptr)
conditional.PtrOrZeroInt16(ptr)
```

## int32 (`typed_int32.go`)

```go
conditional.IfInt32(cond, trueVal, falseVal)
conditional.IfFuncInt32(cond, trueFunc, falseFunc)
conditional.IfTrueFuncInt32(cond, trueFunc)
conditional.IfSliceInt32(cond, trueSlice, falseSlice)
conditional.IfPtrInt32(cond, truePtr, falsePtr)
conditional.NilDefInt32(ptr, defVal)
conditional.NilDefPtrInt32(ptr, defVal)
conditional.NilValInt32(ptr, onNil, onNonNil)
conditional.NilValPtrInt32(ptr, onNil, onNonNil)
conditional.ValueOrZeroInt32(ptr)
conditional.PtrOrZeroInt32(ptr)
```

## int64 (`typed_int64.go`)

```go
conditional.IfInt64(cond, trueVal, falseVal)
conditional.IfFuncInt64(cond, trueFunc, falseFunc)
conditional.IfTrueFuncInt64(cond, trueFunc)
conditional.IfSliceInt64(cond, trueSlice, falseSlice)
conditional.IfPtrInt64(cond, truePtr, falsePtr)
conditional.NilDefInt64(ptr, defVal)
conditional.NilDefPtrInt64(ptr, defVal)
conditional.NilValInt64(ptr, onNil, onNonNil)
conditional.NilValPtrInt64(ptr, onNil, onNonNil)
conditional.ValueOrZeroInt64(ptr)
conditional.PtrOrZeroInt64(ptr)
```

## uint (`typed_uint.go`)

```go
conditional.IfUint(cond, trueVal, falseVal)
conditional.IfFuncUint(cond, trueFunc, falseFunc)
conditional.IfTrueFuncUint(cond, trueFunc)
conditional.IfSliceUint(cond, trueSlice, falseSlice)
conditional.IfPtrUint(cond, truePtr, falsePtr)
conditional.NilDefUint(ptr, defVal)
conditional.NilDefPtrUint(ptr, defVal)
conditional.NilValUint(ptr, onNil, onNonNil)
conditional.NilValPtrUint(ptr, onNil, onNonNil)
conditional.ValueOrZeroUint(ptr)
conditional.PtrOrZeroUint(ptr)
```

## uint8 (`typed_uint8.go`)

```go
conditional.IfUint8(cond, trueVal, falseVal)
conditional.IfFuncUint8(cond, trueFunc, falseFunc)
conditional.IfTrueFuncUint8(cond, trueFunc)
conditional.IfSliceUint8(cond, trueSlice, falseSlice)
conditional.IfPtrUint8(cond, truePtr, falsePtr)
conditional.NilDefUint8(ptr, defVal)
conditional.NilDefPtrUint8(ptr, defVal)
conditional.NilValUint8(ptr, onNil, onNonNil)
conditional.NilValPtrUint8(ptr, onNil, onNonNil)
conditional.ValueOrZeroUint8(ptr)
conditional.PtrOrZeroUint8(ptr)
```

## uint16 (`typed_uint16.go`)

```go
conditional.IfUint16(cond, trueVal, falseVal)
conditional.IfFuncUint16(cond, trueFunc, falseFunc)
conditional.IfTrueFuncUint16(cond, trueFunc)
conditional.IfSliceUint16(cond, trueSlice, falseSlice)
conditional.IfPtrUint16(cond, truePtr, falsePtr)
conditional.NilDefUint16(ptr, defVal)
conditional.NilDefPtrUint16(ptr, defVal)
conditional.NilValUint16(ptr, onNil, onNonNil)
conditional.NilValPtrUint16(ptr, onNil, onNonNil)
conditional.ValueOrZeroUint16(ptr)
conditional.PtrOrZeroUint16(ptr)
```

## uint32 (`typed_uint32.go`)

```go
conditional.IfUint32(cond, trueVal, falseVal)
conditional.IfFuncUint32(cond, trueFunc, falseFunc)
conditional.IfTrueFuncUint32(cond, trueFunc)
conditional.IfSliceUint32(cond, trueSlice, falseSlice)
conditional.IfPtrUint32(cond, truePtr, falsePtr)
conditional.NilDefUint32(ptr, defVal)
conditional.NilDefPtrUint32(ptr, defVal)
conditional.NilValUint32(ptr, onNil, onNonNil)
conditional.NilValPtrUint32(ptr, onNil, onNonNil)
conditional.ValueOrZeroUint32(ptr)
conditional.PtrOrZeroUint32(ptr)
```

## uint64 (`typed_uint64.go`)

```go
conditional.IfUint64(cond, trueVal, falseVal)
conditional.IfFuncUint64(cond, trueFunc, falseFunc)
conditional.IfTrueFuncUint64(cond, trueFunc)
conditional.IfSliceUint64(cond, trueSlice, falseSlice)
conditional.IfPtrUint64(cond, truePtr, falsePtr)
conditional.NilDefUint64(ptr, defVal)
conditional.NilDefPtrUint64(ptr, defVal)
conditional.NilValUint64(ptr, onNil, onNonNil)
conditional.NilValPtrUint64(ptr, onNil, onNonNil)
conditional.ValueOrZeroUint64(ptr)
conditional.PtrOrZeroUint64(ptr)
```

## float32 (`typed_float32.go`)

```go
conditional.IfFloat32(cond, trueVal, falseVal)
conditional.IfFuncFloat32(cond, trueFunc, falseFunc)
conditional.IfTrueFuncFloat32(cond, trueFunc)
conditional.IfSliceFloat32(cond, trueSlice, falseSlice)
conditional.IfPtrFloat32(cond, truePtr, falsePtr)
conditional.NilDefFloat32(ptr, defVal)
conditional.NilDefPtrFloat32(ptr, defVal)
conditional.NilValFloat32(ptr, onNil, onNonNil)
conditional.NilValPtrFloat32(ptr, onNil, onNonNil)
conditional.ValueOrZeroFloat32(ptr)
conditional.PtrOrZeroFloat32(ptr)
```

## float64 (`typed_float64.go`)

```go
conditional.IfFloat64(cond, trueVal, falseVal)
conditional.IfFuncFloat64(cond, trueFunc, falseFunc)
conditional.IfTrueFuncFloat64(cond, trueFunc)
conditional.IfSliceFloat64(cond, trueSlice, falseSlice)
conditional.IfPtrFloat64(cond, truePtr, falsePtr)
conditional.NilDefFloat64(ptr, defVal)
conditional.NilDefPtrFloat64(ptr, defVal)
conditional.NilValFloat64(ptr, onNil, onNonNil)
conditional.NilValPtrFloat64(ptr, onNil, onNonNil)
conditional.ValueOrZeroFloat64(ptr)
conditional.PtrOrZeroFloat64(ptr)
```

---

## Summary Matrix

| Type | `If` | `IfFunc` | `IfTrueFunc` | `IfSlice` | `IfPtr` | `NilDef` | `NilDefPtr` | `NilVal` | `NilValPtr` | `ValueOrZero` | `PtrOrZero` |
|------|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| `bool` | ✓ | ✓ | ✓ | ✓ | ✓ | — ¹ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `byte` | ✓ | ✓ | ✓ | ✓ | ✓ | — ¹ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `string` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `int` | ✓ | ✓ | ✓ | ✓ | ✓ | — ¹ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `int8` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `int16` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `int32` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `int64` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `uint` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `uint8` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `uint16` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `uint32` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `uint64` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `float32` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| `float64` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |

¹ Use `NilDef[T](ptr, defVal)` directly — typed wrapper omitted due to naming conflict with deprecated function.

## Using Generics for Non-Primitive Types

For types not in the 15 primitive set, use the generic base directly:

```go
// Slice types
items := conditional.If[[]string](ok, a, b)
nested := conditional.IfSlice[[]int](ok, a, b)        // [][]int

// Map types
m := conditional.If[map[string]int](ok, mapA, mapB)
p := conditional.NilDef[map[string]any](ptr, make(map[string]any))

// Func types
fn := conditional.If[func() error](ok, fnA, fnB)
handler := conditional.NilDef[func(int) string](ptr, defaultHandler)

// Custom structs
cfg := conditional.If[Config](ok, prodCfg, devCfg)
result := conditional.NilDef[*MyStruct](ptr, &MyStruct{})

// Interfaces
svc := conditional.If[io.Reader](ok, readerA, readerB)
```

## Related Docs

- [conditional README](./README.md) — full package overview
- [Repo Overview](../spec/01-app/00-repo-overview.md)
