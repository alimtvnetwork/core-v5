# Migration Guide — Deprecated → Modern `conditional` Functions

> Step-by-step guide for migrating from deprecated per-type functions to generic base functions or typed convenience wrappers.

---

## Quick Reference

| Migration Path | When to Use |
|---------------|-------------|
| Deprecated → **Typed Wrapper** | Primitive types (`int`, `string`, `bool`, etc.) — cleanest one-liner |
| Deprecated → **Generic Base** | Custom types, maps, funcs, or when you prefer explicit type params |

---

## Phase 1: Simple Ternary Functions

### `Bool`, `Int`, `String`, `Byte`, `Interface`

**Before:**
```go
result := conditional.Bool(isReady, true, false)
count  := conditional.Int(hasItems, len(items), 0)
name   := conditional.String(isSet, value, "default")
flag   := conditional.Byte(isOn, 1, 0)
val    := conditional.Interface(ok, data, nil)
```

**After (typed wrappers — recommended for primitives):**
```go
result := conditional.IfBool(isReady, true, false)
count  := conditional.IfInt(hasItems, len(items), 0)
name   := conditional.IfString(isSet, value, "default")
flag   := conditional.IfByte(isOn, 1, 0)
val    := conditional.IfAny(ok, data, nil)
```

**After (generics — for custom types):**
```go
result := conditional.If[bool](isReady, true, false)
config := conditional.If[MyConfig](useCustom, custom, defaultCfg)
```

### Find & Replace Patterns

```
conditional.Bool(   →  conditional.IfBool(
conditional.Int(    →  conditional.IfInt(
conditional.String( →  conditional.IfString(
conditional.Byte(   →  conditional.IfByte(
conditional.Interface( → conditional.IfAny(
```

---

## Phase 2: Function-Based Ternaries

### `BoolFunc`, `StringFunc`, `InterfaceFunc`

**Before:**
```go
result := conditional.BoolFunc(ok, trueFunc, falseFunc)
name   := conditional.StringFunc(ok, computeName, defaultName)
val    := conditional.InterfaceFunc(ok, loadData, fallbackData)
```

**After (typed wrappers):**
```go
result := conditional.IfFuncBool(ok, trueFunc, falseFunc)
name   := conditional.IfFuncString(ok, computeName, defaultName)
val    := conditional.IfFuncAny(ok, loadData, fallbackData)
```

**After (generics):**
```go
result := conditional.IfFunc[MyType](ok, computeA, computeB)
```

### Find & Replace Patterns

```
conditional.BoolFunc(      →  conditional.IfFuncBool(
conditional.StringFunc(    →  conditional.IfFuncString(
conditional.InterfaceFunc( →  conditional.IfFuncAny(
```

---

## Phase 3: True-Only Functions

### `BooleanTrueFunc`, `StringTrueFunc`, `BytesTrueFunc`, `StringsTrueFunc`

These return the zero value when the condition is false.

**Before:**
```go
active := conditional.BooleanTrueFunc(shouldCheck, checkFunc)
label  := conditional.StringTrueFunc(hasLabel, computeLabel)
data   := conditional.BytesTrueFunc(hasData, loadBytes)
items  := conditional.StringsTrueFunc(hasItems, loadStrings)
```

**After (typed wrappers):**
```go
active := conditional.IfTrueFuncBool(shouldCheck, checkFunc)
label  := conditional.IfTrueFuncString(hasLabel, computeLabel)
data   := conditional.IfTrueFuncBytes(hasData, loadBytes)
items  := conditional.IfTrueFuncStrings(hasItems, loadStrings)
```

**After (generics):**
```go
result := conditional.IfTrueFunc[MyType](ok, computeFunc)
```

### Find & Replace Patterns

```
conditional.BooleanTrueFunc( →  conditional.IfTrueFuncBool(
conditional.StringTrueFunc(  →  conditional.IfTrueFuncString(
conditional.BytesTrueFunc(   →  conditional.IfTrueFuncBytes(
conditional.StringsTrueFunc( →  conditional.IfTrueFuncStrings(
```

---

## Phase 4: Slice Ternaries

### `Booleans`, `Integers`, `Strings`, `Bytes`, `Interfaces`

**Before:**
```go
flags := conditional.Booleans(ok, trueFlags, falseFlags)
ids   := conditional.Integers(ok, activeIds, allIds)
names := conditional.Strings(ok, filteredNames, allNames)
raw   := conditional.Bytes(ok, trueBytes, falseBytes)
items := conditional.Interfaces(ok, trueItems, falseItems)
```

**After (typed wrappers):**
```go
flags := conditional.IfSliceBool(ok, trueFlags, falseFlags)
ids   := conditional.IfSliceInt(ok, activeIds, allIds)
names := conditional.IfSliceString(ok, filteredNames, allNames)
raw   := conditional.IfSliceByte(ok, trueBytes, falseBytes)
items := conditional.IfSliceAny(ok, trueItems, falseItems)
```

### Deprecated `*Ptr` Variants (Identical Behavior)

`BooleansPtr`, `IntegersPtr`, `StringsPtr`, `BytesPtr`, `InterfacesPtr` were identical
to their non-Ptr counterparts. Use the same replacements above.

```
conditional.BooleansPtr(   →  conditional.IfSliceBool(
conditional.IntegersPtr(   →  conditional.IfSliceInt(
conditional.StringsPtr(    →  conditional.IfSliceString(
conditional.BytesPtr(      →  conditional.IfSliceByte(
conditional.InterfacesPtr( →  conditional.IfSliceAny(
```

---

## Phase 5: Pointer Ternaries

### `StringPtr`

**Before:**
```go
ptr := conditional.StringPtr(ok, truePtr, falsePtr)
```

**After (typed wrapper):**
```go
ptr := conditional.IfPtrString(ok, truePtr, falsePtr)
```

**After (generics):**
```go
ptr := conditional.IfPtr[MyType](ok, truePtr, falsePtr)
```

---

## Phase 6: Nil-Default Helpers

### `NilDefStr`, `NilDefInt`, `NilDefBool`, `NilDefByte`

**Before:**
```go
name := conditional.NilDefStr(strPtr)          // "" if nil
id   := conditional.NilDefInt(intPtr)          // 0 if nil
ok   := conditional.NilDefBool(boolPtr)        // false if nil
b    := conditional.NilDefByte(bytePtr)        // 0 if nil
```

**After (generics — recommended):**
```go
name := conditional.ValueOrZero[string](strPtr)
id   := conditional.ValueOrZero[int](intPtr)
ok   := conditional.ValueOrZero[bool](boolPtr)
b    := conditional.ValueOrZero[byte](bytePtr)
```

### `NilDefStrPtr`, `NilDefIntPtr`, `NilDefBoolPtr`, `NilDefBytePtr`

**Before:**
```go
ptr := conditional.NilDefStrPtr(strPtr)    // guaranteed non-nil
ptr := conditional.NilDefIntPtr(intPtr)
```

**After (generics):**
```go
ptr := conditional.PtrOrZero[string](strPtr)
ptr := conditional.PtrOrZero[int](intPtr)
```

### `NilDefValInt`, `NilBoolVal`, `NilByteVal` (with custom default)

**Before:**
```go
val := conditional.NilDefValInt(intPtr, 42)
val := conditional.NilBoolVal(boolPtr, true)
val := conditional.NilByteVal(bytePtr, 0xFF)
```

**After (generics):**
```go
val := conditional.NilDef[int](intPtr, 42)
val := conditional.NilDef[bool](boolPtr, true)
val := conditional.NilDef[byte](bytePtr, 0xFF)
```

### `NilStr` (Nil branching)

**Before:**
```go
label := conditional.NilStr(strPtr, "unknown", "known")
```

**After (generics):**
```go
label := conditional.NilVal[string](strPtr, "unknown", "known")
```

---

## Phase 7: `NilCheck` and `DefOnNil`

### `NilCheck` (any-typed nil branch)

**Before:**
```go
result := conditional.NilCheck(canBeNil, onNil, onNonNil)
```

**After (generics — type-safe):**
```go
result := conditional.NilVal[string](strPtr, onNil, onNonNil)
```

> ⚠️ `NilCheck` accepts `any` and loses type safety. When migrating, prefer `NilVal[T]`
> with a typed pointer for compile-time safety.

### `DefOnNil` (return value or default)

`DefOnNil` has no direct generic replacement because it operates on `any` with interface
nil checks. It remains available for use with untyped values. For typed code, use `NilDef[T]`.

---

## Verification Checklist

After migrating, verify your changes:

```bash
# Compile check
go vet ./...

# Run all tests
go test ./...

# Search for remaining deprecated usage
grep -rn "conditional\.Bool(" --include="*.go" .
grep -rn "conditional\.Int(" --include="*.go" .
grep -rn "conditional\.String(" --include="*.go" .
grep -rn "conditional\.NilDefStr(" --include="*.go" .
grep -rn "conditional\.Interface(" --include="*.go" .
```

---

## Decision Guide

| Scenario | Recommendation |
|----------|---------------|
| Primitive type (`int`, `string`, etc.) | Use typed wrapper: `IfInt(...)`, `IfString(...)` |
| Custom struct or interface | Use generic: `If[MyConfig](...)` |
| Slice of primitives | Use typed wrapper: `IfSliceString(...)` |
| Slice of custom type | Use generic: `IfSlice[MyConfig](...)` |
| Map or func type | Use generic: `If[map[K]V](...)` |
| Nil-safe dereference | Use generic: `NilDef[T](ptr, def)` or `ValueOrZero[T](ptr)` |
| Existing code that works | No urgency — deprecated functions remain functional |

---

## Related Docs

- [conditional README](./README.md) — Full API reference
- [Typed Wrappers Reference](./typed-wrappers.md) — Complete per-type function matrix
- [Go Modernization Plan](../spec/01-app/11-go-modernization.md) — Project-wide upgrade roadmap
- [Type Duplication Issue](../spec/13-app-issues/golang/04-type-duplication-no-generics.md) — Background context
