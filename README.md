# `core` — Go Utility Framework

![Core logo](assets/core-250.png)

The foundational shared package for the **auk-go** ecosystem. It provides reusable primitives, data structures, interfaces, converters, validators, file-system helpers, and testing utilities that keep all downstream Go packages DRY and consistent.

## Quick Start

### Prerequisites

| Tool | Version |
|------|---------|
| Go   | **1.22+** |
| Git  | ≥ 2.29 |

### Install

```bash
go get gitlab.com/auk-go/core
```

### Clone

```bash
git clone https://gitlab.com/auk-go/core.git
```

### Build & Test

```bash
make                  # build and run default CLI
make build            # compile binary to build/cli
make run-tests        # run integration tests
make run-server       # start server entrypoint
make run-client       # start client entrypoint
make run-sample       # run sample/demo
```

## What This Framework Provides

| Category | Packages | What You Get |
|----------|----------|-------------|
| **Ternary helpers** | `conditional/` | Generic `If[T]`, `IfFunc[T]`, `IfSlice[T]` — replaces missing ternary operator |
| **Data structures** | `coredata/corestr/` | `Collection`, `Hashmap`, `Hashset`, `LinkedList`, `SimpleSlice` |
| **JSON** | `coredata/corejson/` | `Serialize.*`, `Deserialize.*` — full JSON pipeline |
| **Error building** | `errcore/` | Stack-traced errors, merge, expectations, Gherkins-style messages |
| **File permissions** | `chmodhelper/` | Parse, verify, and apply chmod on files and directories |
| **Interfaces** | `coreinterface/` | 100+ canonical interface contracts (`*Getter`, `*Checker`, `*Binder`) |
| **Converters** | `converters/` | Type conversions: strings ↔ bytes, maps, pointers |
| **Testing** | `coretests/` | Assertion helpers, `FuncWrap`, `CaseV1` for AAA-pattern tests |
| **Regex** | `regexnew/` | `LazyRegex` — compiles only on first use, with optional locking |
| **Validators** | `corevalidator/` | Line, slice, text, and range validators |
| **Sorting** | `coresort/` | Quick sort for strings and integers |
| **Math** | `coremath/` | Min/Max for all numeric types |
| **Versioning** | `coreversion/` | Semantic version data type (major.minor.patch) |
| **Constants** | `constants/` | OS line separators, empty values, capacity defaults |
| **Generics** | `core.go`, `generic.go` | `EmptySlicePtr[T]`, `SlicePtrByCapacity[T]`, `EmptyMapPtr[K,V]` |

## Design Philosophy

1. **One file per function** — each public function lives in its own `.go` file, named after the function.
2. **Struct-as-namespace** — related operations are grouped on unexported struct types exposed via package-level `var` (e.g., `corejson.Serialize.ToString()`).
3. **Interface-first** — all contracts are defined in `coreinterface/` using Go's `-er` suffix convention (e.g., `NameGetter`, `Csver`, `Serializer`).
4. **Zero-nil safety** — functions return empty slices/maps instead of nil wherever possible.
5. **Generics where clear** — generic versions (`If[T]`, `EmptySlicePtr[T]`) exist alongside backward-compatible type-specific wrappers.

## Examples

### Conditional (Ternary) Helpers

```go
import "gitlab.com/auk-go/core/conditional"

// Generic (Go 1.22+)
result := conditional.If[int](true, 2, 7)          // 2
name := conditional.If[string](len(s) > 0, s, "default")

// With lazy evaluation
val := conditional.IfFunc[string](expensive, func() string {
    return computeValue()
}, func() string {
    return "fallback"
})

// Legacy type-specific (still works, deprecated)
result := conditional.Int(true, 2, 7)   // 2
```

### Generic Slice/Map Factories

```go
import "gitlab.com/auk-go/core"

ints := core.EmptySlicePtr[int]()            // *[]int
strs := core.SlicePtrByLength[string](10)    // *[]string with len=10
m := core.EmptyMapPtr[string, int]()          // *map[string]int
```

### JSON Serialize / Deserialize

```go
import "gitlab.com/auk-go/core/coredata/corejson"

type Example struct {
    A       string
    B       int
    SomeMap map[string]string
}

from := &Example{A: "hello", B: 42, SomeMap: map[string]string{}}
to := &Example{}

err := corejson.Deserialize.FromTo(from, to)
// to is now a deep copy of from's public fields

jsonStr, err := corejson.Serialize.ToString(from)
// jsonStr = `{"A":"hello","B":42,"SomeMap":{}}`
```

### String Collections

```go
import (
    "gitlab.com/auk-go/core/coredata/corestr"
    "gitlab.com/auk-go/core/constants"
)

values := []string{"hello", "world", "something"}
collection := corestr.NewCollectionPtrUsingStrings(&values, constants.Zero)

fmt.Println(collection.Length())  // 3
fmt.Println(collection.IsEmpty()) // false

collection.AddsLock("else")
fmt.Println(collection.Length())  // 4
```

### Hashset

```go
import "gitlab.com/auk-go/core/coredata/corestr"

hs := corestr.NewHashset(2)
hs.Add("alpha")
hs.Add("beta")
fmt.Println(hs.Length()) // 2
fmt.Println(hs.Has("alpha")) // true
```

### Error Construction

```go
import "gitlab.com/auk-go/core/errcore"

// Rich error with stack trace
err := errcore.Expected.Error("config file", "/etc/app.conf")

// Merge multiple errors
combined := errcore.MergeErrors(err1, err2)
```

### Regex (Lazy Compiled)

```go
import "gitlab.com/auk-go/core/regexnew"

// Compiles only on first Match/Find call
lazy := regexnew.Create.New(`\d+`)
matched := lazy.IsMatch("abc123") // true, compiled once
```

### Sorting

```go
import "gitlab.com/auk-go/core/coresort/strsort"

fruits := []string{"banana", "mango", "apple"}
strsort.Quick(&fruits)    // [apple banana mango]
strsort.QuickDsc(&fruits) // [mango banana apple]
```

### 4-Valued Boolean (issetter)

```go
import "gitlab.com/auk-go/core/issetter"

val := issetter.False
fmt.Println(val.HasInitialized()) // true
fmt.Println(val.IsPositive())     // false

val2 := issetter.Uninitialized
fmt.Println(val2.HasInitialized()) // false
```

### File Permissions (chmodhelper)

```go
import "gitlab.com/auk-go/core/chmodhelper"

// Parse rwx string
rwx := chmodhelper.ExpandCharRwx("rwxr-xr--")
// Verify file permissions
isValid := chmodhelper.ChmodVerify.IsFileHasRwx(path, expectedRwx)
```

### CSV Formatting

```go
import "gitlab.com/auk-go/core/corecsv"

// Any type implementing Csver interface gets CSV output
line := corecsv.DefaultCsv(myStruct) // "field1,field2,field3"
```

## Unit Test Pattern

This project follows the **Arrange-Act-Assert (AAA)** pattern using `coretests.GetAssert`:

```go
import (
    "testing"
    "gitlab.com/auk-go/core/coretests"
)

func Test_MyFunction_ShouldReturnExpected(t *testing.T) {
    // Arrange
    input := "test"
    expected := "TEST"

    // Act
    result := strings.ToUpper(input)

    // Assert
    assert := coretests.GetAssert(t)
    assert.ShouldEqual(result, expected)
}
```

For table-driven tests, use `coretestcases.CaseV1`:

```go
import "gitlab.com/auk-go/core/coretests/coretestcases"

cases := []coretestcases.CaseV1{
    {Title: "empty input returns empty", ArrangeInput: []args.One{{First: ""}}},
    {Title: "non-empty returns processed", ArrangeInput: []args.One{{First: "hello"}}},
}

for i, tc := range cases {
    // arrange, act, assert using tc...
}
```

## Interface Conventions

All interfaces in `coreinterface/` follow Go's `-er` suffix convention:

| Pattern | Example | Purpose |
|---------|---------|---------|
| `*Getter` | `NameGetter`, `ValueGetter` | Read a single value |
| `*Checker` | `HasErrorChecker`, `IsEmptyChecker` | Boolean state check |
| `*Binder` | `ContractsBinder` | Compose multiple interfaces |
| `*er` | `Csver`, `Serializer`, `Stringer` | Action performer |

## Troubleshooting

| Problem | Solution |
|---------|----------|
| `go get` fails with auth error | Add SSH key to GitLab or use access token: `git config url."https://oauth2:TOKEN@gitlab.com".insteadOf "https://gitlab.com"` |
| `go mod tidy` reports version conflicts | Ensure `go.mod` specifies `go 1.22` and run `go mod tidy` |
| Tests fail after clone | Run `make run-tests` — some tests require the full module graph |
| Import path has typo | Known: `convertinteranl` → `convertinternal`, `refeflectcore` → `reflectcore` (being fixed) |

## Project Structure

For a complete folder-by-folder breakdown, see the [Folder Map](/spec/01-app/01-folder-map.md).

Key directories:

```
core.go / generic.go    ← root package with generic slice/map factories
conditional/            ← generic ternary helpers (If[T], IfFunc[T])
coredata/               ← data structures (corestr, corejson, coredynamic)
coreinterface/          ← 100+ canonical interface contracts
errcore/                ← rich error construction
chmodhelper/            ← file permission utilities
coretests/              ← testing helpers and assertion wrappers
```

## Specification Docs

Detailed architecture and conventions documentation for AI agents and contributors:

| Document | Path |
|----------|------|
| Repository Overview | [`/spec/01-app/00-repo-overview.md`](/spec/01-app/00-repo-overview.md) |
| Folder Map | [`/spec/01-app/01-folder-map.md`](/spec/01-app/01-folder-map.md) |
| Per-Folder Specs | [`/spec/01-app/folders/`](/spec/01-app/folders/) |
| Codegen Deprecation Plan | [`/spec/01-app/10-codegen-deprecation-plan.md`](/spec/01-app/10-codegen-deprecation-plan.md) |
| Go Modernization Plan | [`/spec/01-app/11-go-modernization.md`](/spec/01-app/11-go-modernization.md) |
| CMD Entrypoints | [`/spec/01-app/12-cmd-entrypoints.md`](/spec/01-app/12-cmd-entrypoints.md) |
| Testing Patterns | [`/spec/01-app/13-testing-patterns.md`](/spec/01-app/13-testing-patterns.md) |
| Interface Conventions | [`/spec/01-app/14-core-interface-conventions.md`](/spec/01-app/14-core-interface-conventions.md) |
| Code Review Report | [`/spec/01-app/15-code-review-report.md`](/spec/01-app/15-code-review-report.md) |
| Known Issues | [`/spec/13-app-issues/`](/spec/13-app-issues/) |

## Acknowledgement

External packages used:

- [`github.com/smartystreets/goconvey`](https://github.com/smartystreets/goconvey) — BDD-style testing
- [`github.com/smarty/assertions`](https://github.com/smarty/assertions) — assertion library
- [`golang.org/x/tools`](https://pkg.go.dev/golang.org/x/tools) — Go tooling support

## Reference Links

- [Go Slice Tricks Cheat Sheet](https://ueokande.github.io/go-slice-tricks/)
- [SliceTricks · golang/go Wiki](https://github.com/golang/go/wiki/SliceTricks)
- [Calling a method on a nil struct pointer](https://t.ly/aTp0)
- [Array of pointers to JSON](https://stackoverflow.com/questions/28101966/array-of-pointers-to-json)

## Issues

- [Create an issue](https://gitlab.com/auk-go/core/-/issues)

## Contributors

- [Md. Alim Ul Karim](https://www.google.com/search?q=Alim+Ul+Karim)

## License

See [LICENSE](LICENSE).
