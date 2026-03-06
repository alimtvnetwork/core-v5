# Deprecated File Removal Plan — `conditional/` Package

> Planned for the next major version bump. All deprecated functions have replacements
> via generic base functions or typed convenience wrappers.

---

## Overview

The `conditional/` package contains **24 legacy per-type files** with deprecated functions
that delegate to generic equivalents. These were kept for backward compatibility but should
be removed in the next major version to reduce surface area and maintenance burden.

---

## Files Scheduled for Deletion

### Category 1: Per-Type Ternary Files (6 files)

| File | Deprecated Function | Replacement |
|------|---------------------|-------------|
| `Bool.go` | `Bool(cond, t, f)` | `IfBool(...)` or `If[bool](...)` |
| `Int.go` | `Int(cond, t, f)` | `IfInt(...)` or `If[int](...)` |
| `String.go` | `String(cond, t, f)` | `IfString(...)` or `If[string](...)` |
| `Byte.go` | `Byte(cond, t, f)` | `IfByte(...)` or `If[byte](...)` |
| `Interface.go` | `Interface(cond, t, f)` | `IfAny(...)` or `If[any](...)` |
| `StringPtr.go` | `StringPtr(cond, t, f)` | `IfPtrString(...)` or `IfPtr[string](...)` |

### Category 2: Per-Type Slice Files (6 files)

| File | Deprecated Function | Replacement |
|------|---------------------|-------------|
| `Booleans.go` | `Booleans(cond, t, f)` | `IfSliceBool(...)` |
| `Integers.go` | `Integers(cond, t, f)` | `IfSliceInt(...)` |
| `Strings.go` | `Strings(cond, t, f)` | `IfSliceString(...)` |
| `Bytes.go` | `Bytes(cond, t, f)` | `IfSliceByte(...)` |
| `Interfaces.go` | `Interfaces(cond, t, f)` | `IfSliceAny(...)` |
| `BooleansPtr.go` | `BooleansPtr(cond, t, f)` | `IfSliceBool(...)` |

### Category 3: Per-Type Slice Ptr Duplicates (4 files)

| File | Deprecated Function | Replacement |
|------|---------------------|-------------|
| `IntegersPtr.go` | `IntegersPtr(cond, t, f)` | `IfSliceInt(...)` |
| `StringsPtr.go` | `StringsPtr(cond, t, f)` | `IfSliceString(...)` |
| `BytesPtr.go` | `BytesPtr(cond, t, f)` | `IfSliceByte(...)` |
| `InterfacesPtr.go` | `InterfacesPtr(cond, t, f)` | `IfSliceAny(...)` |

### Category 4: Per-Type Func Files (4 files)

| File | Deprecated Function | Replacement |
|------|---------------------|-------------|
| `BoolFunc.go` | `BoolFunc(cond, tF, fF)` | `IfFuncBool(...)` |
| `StringFunc.go` | `StringFunc(cond, tF, fF)` | `IfFuncString(...)` |
| `InterfaceFunc.go` | `InterfaceFunc(cond, tF, fF)` | `IfFuncAny(...)` |
| `IntegersPtrFunc.go` | `IntegersPtrFunc(cond, tF, fF)` | `IfSlice[int]` with func wrappers |

### Category 5: Per-Type TrueFunc Files (3 files)

| File | Deprecated Function | Replacement |
|------|---------------------|-------------|
| `BooleanTrueFunc.go` | `BooleanTrueFunc(cond, tF)` | `IfTrueFuncBool(...)` |
| `StringTrueFunc.go` | `StringTrueFunc(cond, tF)` | `IfTrueFuncString(...)` |
| `BytesTrueFunc.go` | `BytesTrueFunc(cond, tF)` | `IfTrueFuncBytes(...)` |
| `StringsTrueFunc.go` | `StringsTrueFunc(cond, tF)` | `IfTrueFuncStrings(...)` |

### Category 6: Nil-Default Legacy Files (4 files)

| File | Deprecated Functions | Replacement |
|------|---------------------|-------------|
| `NilDefBool.go` | `NilDefBool`, `NilDefBoolPtr`, `NilBoolVal`, `NilBoolValPtr` | `ValueOrZero[bool]`, `PtrOrZero[bool]`, `NilDef[bool]`, `NilDefPtr[bool]` |
| `NilDefByte.go` | `NilDefByte`, `NilDefBytePtr`, `NilByteVal`, `NilByteValPtr` | `ValueOrZero[byte]`, `PtrOrZero[byte]`, `NilDef[byte]`, `NilDefPtr[byte]` |
| `NilDefInt.go` | `NilDefInt`, `NilDefIntPtr`, `NilDefValInt` | `ValueOrZero[int]`, `PtrOrZero[int]`, `NilDef[int]` |
| `NilDefStr.go` | `NilDefStr`, `NilDefStrPtr`, `NilStr` | `ValueOrZero[string]`, `PtrOrZero[string]`, `NilVal[string]` |

> **Note**: `NilOrEmptyStr` and `NilOrEmptyStrPtr` in `NilDefStr.go` are **not deprecated** —
> they have string-specific behavior with no generic equivalent. These must be extracted
> to a separate file before deletion.

### Category 7: Deprecated Generic Aliases (in `generic.go`)

| Function | Replacement |
|----------|-------------|
| `IfSlicePtr[T]` | `IfSlice[T]` |
| `IfSlicePtrFunc[T]` | `IfSlice[T]` with func wrappers |
| `NilDeref[T]` | `ValueOrZero[T]` |
| `NilDerefPtr[T]` | `PtrOrZero[T]` |

These should be removed from `generic.go` (not a separate file deletion).

---

## Files NOT Scheduled for Deletion

These files contain **non-deprecated**, unique functionality:

| File | Reason to Keep |
|------|---------------|
| `generic.go` | Core generic functions (remove deprecated aliases only) |
| `typed_*.go` (15 files) | Modern typed convenience wrappers |
| `typed_wrappers.go` | Additional typed wrappers |
| `Functions.go` | Generic batch function execution |
| `FunctionsExecuteResults.go` | Generic batch with isTake/isBreak |
| `AnyFunctions.go` | Any-typed batch execution |
| `AnyFunctionsExecuteResults.go` | Any-typed batch with control |
| `VoidFunctions.go` | Void batch execution |
| `ErrorFunc.go` | Error function selection |
| `ErrorFunctionResult.go` | Error function execution |
| `ErrorFunctionsExecuteResults.go` | Error batch with aggregation |
| `TypedErrorFunctionsExecuteResults.go` | Typed (T, error) batch |
| `Setter.go` | Conditional setter (uses `issetter.Value`) |
| `SetterDefault.go` | Conditional setter with default |
| `BoolByOrder.go` | First-true boolean selection |
| `BoolFunctionsByOrder.go` | First-true function selection |
| `StringsIndexVal.go` | Index-based string selection |
| `StringDefault.go` | String with empty default |
| `NilCheck.go` | Any-typed nil check (no generic equivalent) |
| `DefOnNil.go` | Any-typed default-on-nil |
| `Func.go` | Function selection (returns func) |
| `funcs.go` | Package-level helper functions |
| `executeErrorFunctions.go` | Internal error execution |
| `executeVoidFunctions.go` | Internal void execution |

---

## External Callers Requiring Migration

Before deletion, these callers must be updated:

| File | Current Usage | Migration |
|------|--------------|-----------|
| `chmodhelper/Attribute.go` | `conditional.Byte(...)` (5 calls) | → `conditional.IfByte(...)` |
| `coredata/corestr/ValidValues.go` | `conditional.String(...)` (1 call) | → `conditional.IfString(...)` |
| `tests/.../isanytests/ReflectionTypesVerify_test.go` | `conditional.String(...)` (1 call) | → `conditional.IfString(...)` |
| `tests/.../isanytests/Conclusive_test.go` | `conditional.String(...)` (1 call) | → `conditional.IfString(...)` |
| `tests/.../conditionaltests/If_test.go` | `conditional.NilCheck(...)` (1 call) | Keep — `NilCheck` is not deprecated |

---

## Execution Steps

### Step 1: Migrate External Callers
```bash
# Find all remaining deprecated usage
grep -rn "conditional\.\(Bool\|Int\|String\|Byte\|Interface\)(" --include="*.go" . \
  | grep -v "conditional/" | grep -v "If"
```

Update the 8 call sites listed above.

### Step 2: Extract Non-Deprecated Functions
- Move `NilOrEmptyStr` and `NilOrEmptyStrPtr` from `NilDefStr.go` → new `NilOrEmpty.go`

### Step 3: Delete Deprecated Files
```bash
# 27 files to delete
rm conditional/Bool.go conditional/Int.go conditional/String.go conditional/Byte.go
rm conditional/Interface.go conditional/StringPtr.go
rm conditional/Booleans.go conditional/Integers.go conditional/Strings.go
rm conditional/Bytes.go conditional/Interfaces.go conditional/BooleansPtr.go
rm conditional/IntegersPtr.go conditional/StringsPtr.go conditional/BytesPtr.go
rm conditional/InterfacesPtr.go
rm conditional/BoolFunc.go conditional/StringFunc.go conditional/InterfaceFunc.go
rm conditional/IntegersPtrFunc.go
rm conditional/BooleanTrueFunc.go conditional/StringTrueFunc.go
rm conditional/BytesTrueFunc.go conditional/StringsTrueFunc.go
rm conditional/NilDefBool.go conditional/NilDefByte.go
rm conditional/NilDefInt.go conditional/NilDefStr.go
```

### Step 4: Remove Deprecated Aliases from `generic.go`
Remove `IfSlicePtr`, `IfSlicePtrFunc`, `NilDeref`, `NilDerefPtr`.

### Step 5: Remove Deprecated Aliases from `typed_*.go`
Remove `IfSlicePtr*`, `IfSlicePtrFunc*`, `NilDeref*`, `NilDerefPtr*` from all 15 typed files.

### Step 6: Update Documentation
- Update `README.md` architecture tree
- Update `typed-wrappers.md` if affected
- Remove `migration-guide.md` (no longer needed post-removal)
- Update `spec/13-app-issues/golang/04-type-duplication-no-generics.md` → mark fully complete

### Step 7: Verify
```bash
go vet ./...
go test ./...
```

---

## Summary

| Category | Files | Functions |
|----------|:-----:|:---------:|
| Per-type ternary | 6 | 6 |
| Per-type slice | 6 | 6 |
| Per-type slice ptr | 4 | 4 |
| Per-type func | 4 | 4 |
| Per-type trueFunc | 4 | 4 |
| Nil-default legacy | 4 | 15 |
| Generic aliases | — | 4 |
| Typed aliases | — | ~60 |
| **Total** | **28 files** | **~99 functions** |

**Net result**: ~28 files deleted, ~99 deprecated functions removed, package reduced from ~69 files to ~41 files.

---

## Related Docs

- [conditional README](./README.md)
- [Migration Guide](./migration-guide.md) — How to update call sites
- [Typed Wrappers Reference](./typed-wrappers.md)
- [Type Duplication Issue](../spec/13-app-issues/golang/04-type-duplication-no-generics.md)
