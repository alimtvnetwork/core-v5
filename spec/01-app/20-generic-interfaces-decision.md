# Architecture Decision: Generic Interfaces in `coreinterface/`

## Status: ✅ DECIDED — Additive adoption

## Date: 2026-03-20

## Context

The `coreinterface/` package contains ~13 type-specific `Value*Getter` interfaces:
- `ValueIntegerGetter`, `ValueStringGetter`, `ValueFloat64Getter`, etc.
- Each defines `Value() T` for a specific type

With Go generics (1.18+), a single `TypedValueGetter[T]` can replace all of these.

## Decision

**Add generic interfaces alongside existing ones (non-breaking).**

### New generic interfaces (`TypedGetters.go`):

| Interface | Replaces | Signature |
|-----------|----------|-----------|
| `TypedValueGetter[T]` | All 13 `Value*Getter` interfaces | `Value() T` |
| `TypedValuesGetter[T]` | `ValueStringsGetter` + similar | `Values() []T` |
| `TypedKeyValueGetter[K,V]` | Future key-value patterns | `Key() K` + `Value() V` |

### Migration strategy:
1. **Phase 1 (done)**: Add generic interfaces — no existing code changes
2. **Phase 2 (future)**: New code uses `TypedValueGetter[T]` instead of type-specific getters
3. **Phase 3 (future, optional)**: Deprecate type-specific getters when usage drops to zero

## Rationale

- **Low usage**: Only 7 references to type-specific getters outside their definition files
- **Backward compatible**: Old interfaces remain — no breaking changes
- **Already precedented**: `Cloner[T]` already exists as a generic interface
- **Simplifies future code**: One interface instead of 13

## Risks

- **None for Phase 1**: Purely additive
- **Phase 3**: Removing old interfaces would require updating consumers (currently minimal)

## Files

- `coreinterface/TypedGetters.go` — new generic interfaces
- `coreinterface/all-getters.go` — existing type-specific getters (retained)
