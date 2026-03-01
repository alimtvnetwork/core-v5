# keymk — Composable Key Builder

## Overview

Package `keymk` provides a `Key` type for building structured, joinable key strings from chain segments. It supports configurable joiners, bracket wrapping, template replacement, and a `KeyWithLegend` variant for hierarchical legend-based keys (root-package-group-state-user-item). Pre-configured `Option` presets cover common patterns like pipe-delimited, curly-brace, square-bracket, and hyphen-joined keys.

## Architecture

```
keymk/
├── Key.go                         # Key struct — chain building, compile, clone, JSON
├── KeyWithLegend.go               # KeyWithLegend — legend-based hierarchical key (1058 lines)
├── KeyLegendCompileRequest.go     # KeyLegendCompileRequest — compile request params
├── LegendName.go                  # LegendName — Root, Package, Group, State, User, Item
├── Option.go                      # Option struct — joiner, brackets, skip-empty config
├── Range.go                       # Range utilities for key generation
├── TempReplace.go                 # Temporary replacement operations
├── newKeyCreator.go               # NewKey factory — All, Default, Curly, Parenthesis, etc.
├── newKeyWithLegendCreator.go     # NewKeyWithLegend factory
├── keyModel.go                    # keyModel — JSON-serializable key representation
├── templateReplacer.go            # Template replacement engine
├── fixedLegend.go                 # FixedLegend utilities
├── consts.go                      # DefaultJoiner, DefaultCap, LegendChainSample
├── vars.go                        # Pre-configured Options, LegendNames, factory singletons
├── appendAnyItemsWithBaseStrings.go  # Internal append helpers
├── appendStringsWithBaseAnyItems.go  # Internal append helpers
├── curlyWrapIf.go                 # Internal conditional curly wrapping
└── readme.md
```

## Option Struct

```go
type Option struct {
    Joiner                          string
    IsSkipEmptyEntry, IsUseBrackets bool
    StartBracket, EndBracket        string
}
```

## Pre-Configured Options

| Variable | Joiner | Brackets | Example Output |
|----------|--------|----------|----------------|
| `JoinerOption` | `-` | None | `root-pkg-item` |
| `BracketJoinerOption` | `-` | `[` `]` | `[root]-[pkg]-[item]` |
| `CurlyBraceJoinerOption` | `-` | `{` `}` | `{root}-{pkg}-{item}` |
| `CurlyBracePathJoinerOption` | `/` | `{` `}` | `{root}/{pkg}/{item}` |
| `ParenthesisJoinerOption` | `-` | `(` `)` | `(root)-(pkg)-(item)` |
| `PipeJoinerOption` | `-` | `\|` `\|` | `\|root\|-\|pkg\|` |
| `PipeCurlyJoinerOption` | `-` | `\|{` `}\|` | `\|{root}\|-\|{pkg}\|` |
| `PipeSquareJoinerOption` | `-` | `\|[` `]\|` | `\|[root]\|-\|[pkg]\|` |

## Factory Methods (`NewKey.*`)

| Method | Description |
|--------|-------------|
| `All(option, main, ...chains)` | Full constructor with option and variadic chains |
| `AllStrings(option, main, ...strings)` | String-only variant |
| `Create(option, main)` | Minimal constructor |
| `Default(main, ...chains)` | Hyphen-joined, no brackets |
| `DefaultStrings(main, ...strings)` | String variant of Default |
| `DefaultMain(main)` | Main-only with default option |
| `OptionMain(option, main)` | Option + main only |
| `Curly(main, ...chains)` / `CurlyStrings(...)` | Curly-brace wrapped |
| `SquareBrackets(main, ...chains)` / `SquareBracketsStrings(...)` | Square-bracket wrapped |
| `Parenthesis(main, ...chains)` / `ParenthesisStrings(...)` | Parenthesis wrapped |
| `PathTemplate(root, ...chains)` | Path-style with curly braces |
| `PathTemplateDefault(...chains)` | Path with default root template |

## Key Methods

### Chain Building

| Method | Description |
|--------|-------------|
| `AppendChain(...any)` | Append items to chain (panics if finalized) |
| `AppendChainStrings(...string)` | Append string items |
| `AppendChainKeys(...*Key)` | Append from other keys |
| `Finalized(...any)` | Append + compile + lock |

### Compilation

| Method | Description |
|--------|-------------|
| `Compile(...any)` | Compile with optional additional items |
| `CompileStrings(...string)` | String-only compile |
| `CompileDefault()` | Compile with no extras |
| `CompileKeys(...*Key)` | Compile merging other keys |
| `CompileReplaceCurlyKeyMap(map)` | Compile with `{key}` template replacement |
| `CompileReplaceCurlyKeyMapUsingItems(map, ...any)` | Same with additional items |
| `JoinUsingJoiner(joiner, ...any)` | Compile with custom joiner |
| `JoinUsingOption(option, ...any)` | Compile with temporary option |

### Inspection

| Method | Description |
|--------|-------------|
| `MainName()` | Root name |
| `KeyChains()` | Chain segments |
| `AllRawItems()` | Main + chains |
| `Length()` / `IsEmpty()` | Chain state |
| `IsComplete()` | Whether finalized |
| `HasInChains(string)` | Search chains |
| `String()` / `Name()` / `KeyCompiled()` | Compiled output |
| `Strings()` | Raw items |

### Lifecycle

| Method | Description |
|--------|-------------|
| `ClonePtr(...any)` | Deep clone with optional appends |
| `ConcatNewUsingKeys(...*Key)` | Clone + merge keys |
| `IntRange(start, end)` | Generate range of compiled keys |
| `IntRangeEnding(end)` | Range from 0 |

## LegendName

```go
type LegendName struct {
    Root, Group, Package, State, User, Item string
}
```

Pre-configured: `FullLegends`, `FullCategoryLegends`, `FullEventLegends`, `ShortLegends`, `ShortEventLegends`.

## Usage

```go
import "gitlab.com/auk-go/core/keymk"

// Simple hyphen key
key := keymk.NewKey.Default("app", "users", "create")
fmt.Println(key.Compile()) // "app-users-create"

// Bracket key
key2 := keymk.NewKey.All(keymk.BracketJoinerOption, "cache")
key2.AppendChainStrings("region", "us-east")
fmt.Println(key2.Compile()) // "[cache]-[region]-[us-east]"

// Template replacement
tmpl := keymk.NewKey.Curly("api", "version", "resource")
result := tmpl.CompileReplaceCurlyKeyMap(map[string]string{
    "version":  "v2",
    "resource": "users",
})

// Legend-based key
legendKey := keymk.NewKeyWithLegend.All(
    keymk.JoinerOption,
    "myapp",
)
```

## Related Docs

- [Coding Guidelines](../spec/01-app/17-coding-guidelines.md)
