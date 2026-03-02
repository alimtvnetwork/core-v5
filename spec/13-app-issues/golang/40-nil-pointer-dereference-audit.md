# Audit: Nil Pointer Dereference on Pointer Receivers

## Status: ✅ RESOLVED (Critical), 📋 DOCUMENTED (Medium/Low)

## Phase: 9.1 — Critical Bugs

## Date: 2026-03-02

## Methodology

Scanned 245+ Go files for pointer receiver methods (`func (it *T)`) that access receiver fields without first checking `it == nil`. Prioritized types where *some* methods already have nil guards (indicating nil pointers are expected at runtime).

---

## Critical — Fixed

### 1. `coredata/coredynamic/LeftRight.go` — 7 methods

| Method | Risk |
|---|---|
| `LeftReflectSet()` | Accesses `it.Left` |
| `RightReflectSet()` | Accesses `it.Right` |
| `DeserializeLeft()` | Accesses `it.Left` |
| `DeserializeRight()` | Accesses `it.Right` |
| `LeftToDynamic()` | Accesses `it.Left` |
| `RightToDynamic()` | Accesses `it.Right` |
| `TypeStatus()` | Accesses `it.Left` and `it.Right` |

**Context**: `IsEmpty()`, `HasLeft()`, `HasRight()`, `IsLeftEmpty()`, `IsRightEmpty()` all have nil guards. These 7 methods did not.

### 2. `coreinstruction/Specification.go` — 1 method

| Method | Risk |
|---|---|
| `FlatSpecification()` | Accesses `r.flatSpec` |

**Context**: `Clone()` has nil guard.

### 3. `coredata/coredynamic/MapAnyItems.go` — 1 method

| Method | Risk |
|---|---|
| `HasKey()` | Accesses `it.Items` map |

**Context**: `Length()`, `IsEmpty()` have nil guards.

### 4. Setter methods — 6 methods across 3 files

| File | Methods |
|---|---|
| `coreinstruction/FromTo.go` | `SetFromName()`, `SetToName()` |
| `coreinstruction/SourceDestination.go` | `SetFromName()`, `SetToName()` |
| `coreinstruction/Rename.go` | `SetFromName()`, `SetToName()` |

### 5. `coredata/corepayload/AttributesSetters.go` — 3 methods

| Method | Risk |
|---|---|
| `AddOrUpdateString()` | Accesses `it.KeyValuePairs` without nil check on `it` |
| `AddOrUpdateAnyItem()` | Accesses `it.AnyKeyValuePairs` without nil check on `it` |
| `Dispose()` | Calls `it.Clear()` without nil check on `it` |

### 6. `coredata/coredynamic/KeyValCollection.go` — 2 methods

| Method | Risk |
|---|---|
| `Items()` | Accesses `it.items` |
| `String()` | Accesses `it.items` |

**Context**: `Length()`, `IsEmpty()`, `ClonePtr()` have nil guards.

---

## Medium Risk — Documented (Not Fixed)

These types are typically constructed via factory functions and rarely nil, but callers could still pass nil pointers.

### `DynamicGetters.go` — ~20 methods

`Data()`, `Value()`, `Length()`, `StructStringPtr()`, `String()`, `IsNull()`, `IsValid()`, `IsInvalid()`, `IsPointer()`, `IsPrimitive()`, `IsNumber()`, `IsStringType()`, `IsStruct()`, `IsFunc()`, `IsSliceOrArray()`, `IsMap()`, `IntDefault()`, `Float64()`, `ValueInt()`, `ValueUInt()`, `ValueStrings()`, `ValueBool()`, `ValueInt64()`.

**Rationale**: `Dynamic` is almost always constructed via `NewDynamic`/`NewDynamicPtr` and embedded in structs by value. Nil `*Dynamic` is rare.

### `KeyVal.go` — ~10 methods

`KeyDynamic()`, `ValueDynamic()`, `KeyDynamicPtr()`, `ValueDynamicPtr()`, `IsKeyNull()`, `IsKeyNullOrEmptyString()`, `IsValueNull()`, `String()`, `ValueReflectValue()`, `ValueInt()`, `ValueUInt()`, etc.

**Rationale**: Many methods already have nil guards (`CastKeyVal`, `KeyString`, `ValueString`). The unsafe ones are value-extraction methods typically called in chains after nil checks.

### `SimpleResult.go` — 3 methods

`GetErrorOnTypeMismatch()`, `InvalidError()`, `Clone()`.

### `SimpleRequest.go` — 5 methods

`Message()`, `Request()`, `Value()`, `IsReflectKind()`, `IsPointer()`, `InvalidError()`.

### `TextValidator.go` — 5 methods

`ToString()`, `String()`, `SearchTextFinalized()`, `IsMatch()`, `MethodName()`.

**Context**: `IsMatchMany()`, `VerifyDetailError()`, `VerifyFirstError()`, `AllVerifyError()` all have nil guards.

### `PayloadsCollectionGetters.go` — 4 methods

`First()`, `Last()`, `FirstDynamic()`, `LastDynamic()` — access `it.Items[0]` without nil check. These will panic on nil receiver AND empty collection. Safe alternatives `FirstOrDefault()`/`LastOrDefault()` exist.

---

## Total

- **Critical fixed**: 20 methods across 8 files
- **Medium documented**: ~47 methods across 6 files
