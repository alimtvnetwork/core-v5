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
    VersionCompact string // "1.0.1"
    Compiled       string // "v1.0.1"
    IsInvalid      bool
    VersionMajor   int
    VersionMinor   int
    VersionPatch   int
    VersionBuild   int
}
```

## Factory Methods (`New.*`)

| Method | Signature | Description |
|--------|-----------|-------------|
| `Default` / `Create` / `Version` | `(string) Version` | Parse from `"v1.2.3"` or `"1.2.3"` |
| `DefaultPtr` | `(string) *Version` | Same as pointer |
| `Major` | `(string) Version` | Parse major-only string |
| `MajorMinor` | `(major, minor string) Version` | Two-component version |
| `MajorMinorPatch` | `(major, minor, patch string) Version` | Three-component version |
| `MajorMinorPatchBuild` / `All` | `(major, minor, patch, build string) Version` | Full four-component |
| `MajorMinorInt` | `(major, minor int) Version` | Integer inputs |
| `MajorMinorPatchInt` | `(major, minor, patch int) Version` | Integer inputs |
| `AllInt` | `(major, minor, patch, build int) Version` | All integer inputs |
| `AllByte` | `(major, minor, patch, build byte) Version` | Byte inputs |
| `SpreadStrings` | `(...string) Version` | Variadic string components |
| `SpreadIntegers` | `(...int) Version` | Variadic integer components |
| `SpreadBytes` | `(...byte) Version` | Variadic byte components |
| `SpreadUnsignedIntegers` | `(...uint) Version` | Variadic uint components |

## Version Methods

### Display

| Method | Description |
|--------|-------------|
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
| `IsEmptyOrInvalid()` / `IsDefined()` / `HasAnyItem()` | Overall validity |

### Comparison

| Method | Description |
|--------|-------------|
| `Major(int)` | Compare major component |
| `MajorMinor(int, int)` | Compare major + minor |
| `MajorMinorPatch(int, int, int)` | Compare major + minor + patch |
| `MajorMinorPatchBuild(int, int, int, int)` | Compare all four |
| `MajorBuild(int, int)` | Compare major + build |
| `Compare(*Version)` | Full version comparison |
| `IsEqual(*Version)` | Equality check |
| `IsMajorAtLeast(int)` | Major >= given |
| `IsMajorMinorAtLeast(int, int)` | Major.Minor >= given |
| `IsMajorMinorPatchAtLeast(int, int, int)` | Major.Minor.Patch >= given |
| `IsVersionCompareEqual(string)` | Compare compact string |

### Package-Level Functions

| Function | Description |
|----------|-------------|
| `Compare(left, right *Version)` | Hierarchical Major→Minor→Patch→Build comparison |
| `CompareVersionString(left, right string)` | Parse and compare two version strings |
| `IsAtLeast(left, right string)` | Left >= Right |
| `IsLower(left, right string)` | Left < Right |
| `IsLowerOrEqual(left, right string)` | Left <= Right |
| `IsExpectedVersion(expected, left, right)` | Compare matches expected result |

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

// Compare
cmp := coreversion.CompareVersionString("1.2.3", "1.3.0")
if cmp.IsLeftLess() {
    // 1.2.3 < 1.3.0
}

// At-least check
if coreversion.IsAtLeast("2.0.0", "1.5.0") {
    // 2.0.0 >= 1.5.0
}
```

## How to Extend Safely

- **New version formats**: Add parsing methods to `newCreator` — do not modify `Default`.
- **New comparison modes**: Add as separate methods on `Version` that compose existing `Major`/`Minor`/`Patch`/`Build` comparisons.
- **New display formats**: Add as new methods (e.g., `VersionDisplayShort`) — do not modify existing display methods.

## Related Docs

- [corecomparator readme](../corecomparator/readme.md)
- [corecmp readme](../corecmp/readme.md)
- [Coding Guidelines](../spec/01-app/17-coding-guidelines.md)
