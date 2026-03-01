# coreversion — Semantic Version Parsing & Comparison

## Overview

Package `coreversion` provides a structured `Version` type for parsing, comparing, and displaying semantic version strings (e.g., `"v1.2.3"`, `"0.1.0.4"`). It supports up to 4 components (Major, Minor, Patch, Build), hierarchical comparison via `corecomparator.Compare`, and a `VersionsCollection` for managing sets of versions.

## Architecture

```
coreversion/
├── Version.go                      # Version struct — display, validation, comparison, JSON
├── VersionsCollection.go           # VersionsCollection — slice management, search, equality
├── all-compare.go                  # Compare, CompareVersionString, IsAtLeast, IsLower
├── consts.go                       # VSymbol ("v"), InvalidVersionValue
├── vars.go                         # New (newCreator singleton), skipValuesMap
├── newCreator.go                   # Factory methods — Default, Create, MajorMinor, SpreadIntegers, etc.
├── Empty.go                        # Empty() Version — zero-value factory
├── EmptyUsingCompactVersion.go     # EmptyUsingCompactVersion(string) Version
├── InvalidCompactVersion.go        # InvalidCompactVersion(string) Version
├── hasDeductUsingNilNess.go        # Nil-safety deduction for comparison
└── readme.md
```

## Version Struct

```go
type Version struct {
    VersionCompact string `json:"Compact,omitempty"`   // e.g. "1.0.1"
    Compiled       string `json:"Compiled,omitempty"`  // e.g. "v1.0.1"
    IsInvalid      bool   `json:"IsInvalid,omitempty"`
    VersionMajor   int    `json:"Major,omitempty"`
    VersionMinor   int    `json:"Minor,omitempty"`
    VersionPatch   int    `json:"Patch,omitempty"`
    VersionBuild   int    `json:"Build,omitempty"`
}
```

## Factory Methods (`New.*`)

### Primary Constructors

| Method | Signature | Description |
|--------|-----------|-------------|
| `Default` / `Create` / `Version` | `(string) Version` | Parse from `"v1.2.3"` or `"1.2.3"` |
| `DefaultPtr` | `(string) *Version` | Pointer variant |
| `Empty` / `Invalid` | `() Version` | Zero-value / invalid version |

### Component Constructors (String)

| Method | Signature | Description |
|--------|-----------|-------------|
| `Major` | `(string) Version` | Major-only |
| `MajorMinor` | `(major, minor string) Version` | Two-component |
| `MajorMinorPatch` | `(major, minor, patch string) Version` | Three-component |
| `MajorMinorPatchBuild` / `All` | `(major, minor, patch, build string) Version` | Full four-component |
| `MajorPatch` | `(major, patch string) Version` | Major + Patch |
| `MajorBuild` | `(major, build string) Version` | Major + Build |
| `MajorMinorBuild` | `(major, minor, build string) Version` | Major + Minor + Build |
| `MinorBuild` | `(minor, build string) Version` | Minor + Build |
| `PatchBuild` | `(patch, build string) Version` | Patch + Build |

### Component Constructors (Integer)

| Method | Signature | Description |
|--------|-----------|-------------|
| `MajorMinorInt` | `(major, minor int) Version` | Two-component integers |
| `MajorMinorPatchInt` | `(major, minor, patch int) Version` | Three-component integers |
| `AllInt` | `(major, minor, patch, build int) Version` | Full four-component integers |
| `AllByte` | `(major, minor, patch, build byte) Version` | Full four-component bytes |
| `MajorPatchInt` | `(major, patch int) Version` | Major + Patch integers |
| `MajorBuildInt` | `(major, build int) Version` | Major + Build integers |
| `MinorBuildInt` | `(minor, build int) Version` | Minor + Build integers |
| `PatchBuildInt` | `(patch, build int) Version` | Patch + Build integers |

### Variadic Constructors

| Method | Signature | Description |
|--------|-----------|-------------|
| `SpreadStrings` | `(...string) Version` | Variadic string components |
| `SpreadIntegers` | `(...int) Version` | Variadic integer components |
| `SpreadBytes` | `(...byte) Version` | Variadic byte components |
| `SpreadUnsignedIntegers` | `(...uint) Version` | Variadic uint components |

### Collection Constructors

| Method | Signature | Description |
|--------|-----------|-------------|
| `Many` / `Collection` | `(...string) *VersionsCollection` | Parse multiple version strings |
| `CollectionUsingCap` | `(int) *VersionsCollection` | Empty collection with capacity |
| `EmptyCollection` | `() *VersionsCollection` | Empty collection (capacity 0) |

## Version Methods

### Display

| Method | Description |
|--------|-------------|
| `String()` | Delegates to `CompiledVersion()` |
| `VersionDisplay()` | `"v1.2.3"` — with `v` prefix |
| `CompiledVersion()` | Compiled version string from parsed components |
| `VersionDisplayMajor()` | `"v1"` |
| `VersionDisplayMajorMinor()` | `"v1.2"` |
| `VersionDisplayMajorMinorPatch()` | `"v1.2.3"` |
| `MajorString()` / `MinorString()` / `PatchString()` / `BuildString()` | Individual component strings |

### Validation

| Method | Description |
|--------|-------------|
| `HasMajor()` / `HasMinor()` / `HasPatch()` / `HasBuild()` | Component exists and valid |
| `IsMajorInvalid()` / `IsMinorInvalid()` / `IsPatchInvalid()` / `IsBuildInvalid()` | Component invalid |
| `IsMajorInvalidOrZero()` / `IsMinorInvalidOrZero()` / ... | Invalid or zero |
| `IsEmptyOrInvalid()` / `IsInvalidOrEmpty()` / `IsDefined()` / `HasAnyItem()` | Overall validity |

### Component Comparison (returns `corecomparator.Compare`)

| Method | Description |
|--------|-------------|
| `Major(int)` | Compare major component |
| `MajorMinor(int, int)` | Compare major + minor |
| `MajorMinorPatch(int, int, int)` | Compare major + minor + patch |
| `MajorMinorPatchBuild(int, int, int, int)` | Compare all four |
| `MajorBuild(int, int)` | Compare major + build |
| `MajorPatch(int, int)` | Compare major + patch |
| `Patch(int)` | Compare patch component |
| `Build(int)` | Compare build component |
| `MajorMinorPatchBuildString(string, string, string, string)` | Compare all four (string inputs) |
| `MajorBuildString(string, string)` | Compare major + build (string inputs) |
| `ComparisonValueIndexes(*Version, ...versionindexes.Index)` | Compare by selective indexes |

### Version-Level Comparison

| Method | Description |
|--------|-------------|
| `Compare(*Version)` | Full version comparison |
| `IsEqual(*Version)` | Equality check |
| `IsLeftLessThan(*Version)` | `it < right` |
| `IsLeftGreaterThan(*Version)` | `it > right` |
| `IsLeftLessThanOrEqual(*Version)` | `it <= right` |
| `IsLeftGreaterThanOrEqual(*Version)` | `it >= right` |
| `IsExpectedComparison(Compare, *Version)` | Matches expected comparison result |
| `IsExpectedComparisonRawVersion(Compare, string)` | Same with raw version string |

### At-Least / Threshold Checks

| Method | Description |
|--------|-------------|
| `IsMajorAtLeast(int)` | Major >= given |
| `IsMajorStringAtLeast(string)` | Major >= given (string input) |
| `IsMajorMinorAtLeast(int, int)` | Major.Minor >= given |
| `IsMajorBuildAtLeast(int, int)` | Major.Build >= given |
| `IsMajorMinorPatchAtLeast(int, int, int)` | Major.Minor.Patch >= given |
| `IsAtLeast(string)` | Version >= given (raw version string) |
| `IsVersionCompareEqual(string)` | Compact string equality |
| `IsVersionCompareNotEqual(string)` | Compact string inequality |
| `IsEqualVersionString(string)` | Deduced equality with raw version string |
| `IsLowerVersionString(string)` | `it < given` (raw version string) |
| `IsLowerEqualVersionString(string)` | `it <= given` (raw version string) |

### Value Access

| Method | Description |
|--------|-------------|
| `ValueByIndex(versionindexes.Index)` | Get component by index enum |
| `ValueByIndexes(...versionindexes.Index)` | Get multiple components |
| `AllVersionValues()` | All four component values |
| `AllValidVersionValues()` | Non-invalid component values |

### Lifecycle & Serialization

| Method | Description |
|--------|-------------|
| `Clone()` / `ClonePtr()` | Deep copy |
| `NonPtr()` / `Ptr()` | Value/pointer conversion |
| `Json()` / `JsonPtr()` | JSON serialization via `corejson` |
| `JsonParseSelfInject(*corejson.Result)` | Deserialize into self |
| `AsJsonContractsBinder()` | Return as `JsonContractsBinder` interface |

### Package-Level Functions

| Function | Description |
|----------|-------------|
| `Compare(left, right *Version)` | Hierarchical Major→Minor→Patch→Build comparison |
| `CompareVersionString(left, right string)` | Parse and compare two version strings |
| `IsAtLeast(left, right string)` | Left >= Right |
| `IsLower(left, right string)` | Left < Right |
| `IsLowerOrEqual(left, right string)` | Left <= Right |
| `IsExpectedVersion(expected, left, right)` | Compare matches expected result |
| `Empty()` | Zero-value Version |
| `EmptyUsingCompactVersion(string)` | Zero-value with compact string set |
| `InvalidCompactVersion(string)` | Invalid Version with compact string |

## VersionsCollection

| Method | Description |
|--------|-------------|
| `Add(string)` | Parse and add version |
| `AddSkipInvalid(string)` | Add only if valid |
| `AddVersionsRaw(...string)` | Bulk add from strings |
| `AddVersions(...Version)` | Bulk add Version values |
| `Length()` / `IsEmpty()` / `HasAnyItem()` | Collection state |
| `IndexOf(string)` | Find version index |
| `IsContainsVersion(string)` | Contains check |
| `IsEqual(*VersionsCollection)` | Collection equality |
| `VersionCompactStrings()` | Compact string slice |
| `VersionsStrings()` | Display string slice |

## Dependencies

| Package | Usage |
|---------|-------|
| `corecmp` | Integer comparison for version components |
| `corecomparator` | `Compare` result type |
| `corejson` | JSON serialization |
| `versionindexes` | Component index enum (Major, Minor, Patch, Build) |
| `converters` | String-to-integer conversion |

## Usage

```go
import "gitlab.com/auk-go/core/coreversion"

// Parse version
v := coreversion.New.Create("v1.2.3")
fmt.Println(v.VersionDisplay()) // "v1.2.3"

// Integer factory
v2 := coreversion.New.AllInt(2, 0, 1, 0)
fmt.Println(v2.VersionDisplay()) // "v2.0.1"

// Compare
cmp := coreversion.CompareVersionString("1.2.3", "1.3.0")
if cmp.IsLeftLess() {
    // 1.2.3 < 1.3.0
}

// At-least check (package-level)
if coreversion.IsAtLeast("2.0.0", "1.5.0") {
    // 2.0.0 >= 1.5.0
}

// At-least check (instance method)
v3 := coreversion.New.Create("3.1.0")
if v3.IsMajorMinorAtLeast(3, 0) {
    // 3.1 >= 3.0
}

// Collection
col := coreversion.New.Many("1.0.0", "2.0.0", "1.5.0")
fmt.Println(col.IsContainsVersion("2.0.0")) // true
```

## How to Extend Safely

- **New version formats**: Add parsing methods to `newCreator` — do not modify `Default`.
- **New comparison modes**: Add as separate methods on `Version` that compose existing `Major`/`Minor`/`Patch`/`Build` comparisons.
- **New display formats**: Add as new methods (e.g., `VersionDisplayShort`) — do not modify existing display methods.

## Related Docs

- [corecomparator readme](../corecomparator/readme.md)
- [corecmp readme](../corecmp/readme.md)
- [Coding Guidelines](../spec/01-app/17-coding-guidelines.md)
