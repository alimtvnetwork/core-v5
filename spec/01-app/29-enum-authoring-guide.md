# Enum Authoring Guide — Reusing `enumimpl` for Byte Enums

## Goal

This spec explains how to create a new enum package that matches the existing core style, reuses the shared enum building blocks, and is easy for another AI or engineer to extend safely.

Use this guide when you want a package like `reqtype`, `ostype`, or `enums/versionindexes`.

---

## Default Recommendation

For a normal byte-backed enum, use this pattern:

1. `type MyEnum byte`
2. `const (...)` with contiguous values
3. `Ranges = [...]string{...}`
4. `RangesMap = map[string]MyEnum{...}`
5. `BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(...)`
6. Delegate shared enum behavior to `BasicEnumImpl`
7. Keep only domain-specific logic on the enum type itself

This is the standard path.

---

## When to Use `enumimpl.New.BasicByte`

Use `BasicByte` when all of these are true:

- The enum is a **single byte value**
- The values are **contiguous** or can be treated as a simple ordered range
- You want shared behaviors like:
  - `Name()` / `String()`
  - `MarshalJSON()` / `UnmarshalJSON()`
  - `RangeNamesCsv()`
  - `MinMaxAny()`
  - string-to-value mapping
  - enum formatting helpers

Examples:

- request types
- OS variants
- version index positions
- state/status enums

---

## When **Not** to Use a Standard Byte Enum

Do **not** model bit-flags or combinable masks as a normal `BasicByte` enum.

Examples:

- permission bits: `4 = read`, `2 = write`, `1 = execute`
- flags built from `1 << n`
- values that can be OR-combined like `1 | 2 | 4`

Why:

- `BasicByte` is designed around a **single enum value**
- its range validation is min/max based, not full set-membership validation
- bitmask/flags usually need composition logic, not only enum lookup

For flags, use a dedicated helper/value object instead.

**Example in this repo:** `chmodhelper/newAttributeCreator.go` converts the byte formula `4/2/1` into a richer `Attribute` object instead of pretending that combinations are ordinary enum members.

---

## Package Shape

For a reusable enum package, prefer this structure:

```text
mypackage/
├── MyEnum.go         # enum type + methods
├── vars.go           # Ranges, RangesMap, grouping maps, BasicEnumImpl
├── consts.go         # package constants if needed
├── readme.md         # package overview and examples
└── extra files       # range helpers, logical groups, result structs, etc.
```

If the main enum file becomes large, split it by responsibility.

Good split examples:

- `MyEnum.naming.go`
- `MyEnum.json.go`
- `MyEnum.checkers.go`
- `MyEnum.logical-groups.go`

This matters because large single enum files become hard for AI handoff.

---

## Required Reuse Points

Most byte enums in this codebase should reuse these components:

- `coreimpl/enumimpl.BasicByte`
- `enumimpl.New.BasicByte.*`
- `coreinterface/enuminf`
- `reflectinternal.TypeName(...)` for the type name

Core idea:

```go
var BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
    reflectinternal.TypeName(Invalid),
    Ranges[:],
)
```

Then the enum methods mostly delegate to `BasicEnumImpl`.

---

## Canonical Sequential Byte Enum Pattern

### 1) Define the enum type

```go
package status

type Status byte
```

### 2) Define contiguous constants

```go
const (
    Invalid Status = iota
    Pending
    Ready
    Failed
)
```

### 3) Define the shared lookup data in `vars.go`

```go
package status

import (
    "github.com/alimtvnetwork/core/coreimpl/enumimpl"
    "github.com/alimtvnetwork/core/internal/reflectinternal"
)

var (
    Ranges = [...]string{
        Invalid: "Invalid",
        Pending: "Pending",
        Ready:   "Ready",
        Failed:  "Failed",
    }

    RangesMap = map[string]Status{
        "Invalid": Invalid,
        "Pending": Pending,
        "Ready":   Ready,
        "Failed":  Failed,
    }

    BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
        reflectinternal.TypeName(Invalid),
        Ranges[:],
    )
)
```

### 4) Implement the minimum reusable methods

```go
package status

import "github.com/alimtvnetwork/core/coreinterface/enuminf"

func (it Status) Value() byte {
    return byte(it)
}

func (it Status) ValueByte() byte {
    return it.Value()
}

func (it Status) ValueInt() int {
    return int(it)
}

func (it Status) Name() string {
    return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Status) String() string {
    return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Status) TypeName() string {
    return BasicEnumImpl.TypeName()
}

func (it Status) RangeNamesCsv() string {
    return BasicEnumImpl.RangeNamesCsv()
}

func (it Status) MinMaxAny() (min, max any) {
    return BasicEnumImpl.MinMaxAny()
}

func (it Status) RangesDynamicMap() map[string]any {
    return BasicEnumImpl.RangesDynamicMap()
}

func (it Status) IsValid() bool {
    return it != Invalid
}

func (it Status) IsInvalid() bool {
    return it == Invalid
}

func (it Status) MarshalJSON() ([]byte, error) {
    return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}

func (it *Status) UnmarshalJSON(data []byte) error {
    dataConv, err := it.UnmarshallEnumToValue(data)

    if err == nil {
        *it = Status(dataConv)
    }

    return err
}

func (it Status) EnumType() enuminf.EnumTyper {
    return BasicEnumImpl.EnumType()
}
```

### 5) Add domain checkers as plain methods

```go
func (it Status) IsPending() bool {
    return it == Pending
}

func (it Status) IsReady() bool {
    return it == Ready
}

func (it Status) IsFailed() bool {
    return it == Failed
}
```

### 6) Add logical grouping maps only when they add meaning

```go
var finalStates = map[Status]bool{
    Ready:  true,
    Failed: true,
}

func (it Status) IsFinal() bool {
    return finalStates[it]
}
```

---

## Alias-Aware Enum Pattern

If JSON or user input should accept aliases, use an alias map when building `BasicEnumImpl`.

```go
var BasicEnumImpl = enumimpl.New.BasicByte.CreateUsingSlicePlusAliasMapOptions(
    true,
    Invalid,
    Ranges[:],
    map[string]byte{
        "ok":    Ready.Value(),
        "error": Failed.Value(),
    },
)
```

Use this when you want inputs such as:

- `"Ready"`
- `"ready"`
- `"READY"`
- `"ok"`

all to resolve to the same enum value.

---

## Explicit Byte Values Pattern

Use this only when values must be assigned explicitly and still behave like a single enum value.

Example:

```go
type Priority byte

const (
    Low    Priority = 1
    Medium Priority = 2
    High   Priority = 3
)

var BasicEnumImpl = enumimpl.New.BasicByte.CreateUsingMapPlusAliasMapOptions(
    false,
    Low,
    map[byte]string{
        byte(Low):    "Low",
        byte(Medium): "Medium",
        byte(High):   "High",
    },
    nil,
)
```

Use this pattern when values are explicit but still represent **one chosen state**, not combined flags.

---

## Formula Rule for Byte Enums

Use this rule before generating code:

### Safe for `BasicByte`

- `0, 1, 2, 3, ...`
- `1, 2, 3, ...`
- any byte values that represent **one selected member**

### Not safe as a normal enum

- `1 << 0`, `1 << 1`, `1 << 2` when values are meant to be combined
- permission-style formulas like `4`, `2`, `1`, `7`
- bitwise flag sets

If the design is combinable, build a **flags helper** instead of a normal enum package.

---

## Methods That Should Usually Delegate to `BasicEnumImpl`

Prefer delegation for all generic enum behavior.

Typical delegation set:

- `Name()`
- `String()`
- `NameValue()`
- `TypeName()`
- `RangeNamesCsv()`
- `RangesDynamicMap()`
- `MinMaxAny()`
- `IntegerEnumRanges()`
- `MarshalJSON()`
- `UnmarshalJSON()`
- `Format(...)`
- `EnumType()`

Only keep custom business semantics local, such as:

- `IsCreateLogically()`
- `IsFinal()`
- `IsHttpMethod()`
- `IsRetryAction()`

---

## AI Authoring Checklist

When an AI creates a new enum package, it should follow this checklist:

1. Choose `byte` only if the enum is byte-sized and single-valued
2. Prefer contiguous constants with `iota`
3. Put shared lookup data in `vars.go`
4. Build `BasicEnumImpl` with `enumimpl.New.BasicByte.*`
5. Delegate generic enum behavior to `BasicEnumImpl`
6. Add custom `IsX()` methods only for domain meaning
7. Add logical grouping maps only when reused by 2+ methods
8. Keep read-only methods on value receivers
9. Split large enum files by responsibility
10. Do not model bitmask flags as a plain enum package

---

## Quick Copy Template

```go
package myenum

import (
    "github.com/alimtvnetwork/core/coreimpl/enumimpl"
    "github.com/alimtvnetwork/core/coreinterface/enuminf"
    "github.com/alimtvnetwork/core/internal/reflectinternal"
)

type MyEnum byte

const (
    Invalid MyEnum = iota
    First
    Second
)

var (
    Ranges = [...]string{
        Invalid: "Invalid",
        First:   "First",
        Second:  "Second",
    }

    RangesMap = map[string]MyEnum{
        "Invalid": Invalid,
        "First":   First,
        "Second":  Second,
    }

    BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
        reflectinternal.TypeName(Invalid),
        Ranges[:],
    )
)

func (it MyEnum) Value() byte { return byte(it) }

func (it MyEnum) ValueByte() byte { return it.Value() }

func (it MyEnum) Name() string { return BasicEnumImpl.ToEnumString(it.Value()) }

func (it MyEnum) String() string { return BasicEnumImpl.ToEnumString(it.Value()) }

func (it MyEnum) TypeName() string { return BasicEnumImpl.TypeName() }

func (it MyEnum) MarshalJSON() ([]byte, error) {
    return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}

func (it *MyEnum) UnmarshalJSON(data []byte) error {
    dataConv, err := it.UnmarshallEnumToValue(data)

    if err == nil {
        *it = MyEnum(dataConv)
    }

    return err
}

func (it MyEnum) EnumType() enuminf.EnumTyper { return BasicEnumImpl.EnumType() }
```

---

## Related Docs

- [coreimpl/enumimpl/readme.md](/coreimpl/enumimpl/readme.md)
- [coreinterface/enuminf/README.md](/coreinterface/enuminf/README.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [newCreator Convention](/spec/01-app/18-new-creator-convention.md)