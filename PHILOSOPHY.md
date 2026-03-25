# Design Philosophy

> **Why does this codebase look the way it does?**

This document explains the architectural decisions behind `core` — not just _what_ the conventions are, but _why_ they exist and the problems they solve. If the [README](README.md) is the "what" and the [Coding Guidelines](/spec/01-app/17-coding-guidelines.md) are the "how," this document is the "why."

---

## Table of Contents

- [The Core Problem](#the-core-problem)
- [One File Per Function](#one-file-per-function)
- [Struct-as-Namespace](#struct-as-namespace)
- [The `newCreator` Convention](#the-newcreator-convention)
- [Interface-First Design](#interface-first-design)
- [Zero-Nil Safety](#zero-nil-safety)
- [Value Receivers by Default](#value-receivers-by-default)
- [Large Type Splitting](#large-type-splitting)
- [Testing as Architecture](#testing-as-architecture)
- [Why Not Just Use the Standard Library?](#why-not-just-use-the-standard-library)

---

## The Core Problem

Go's simplicity is its greatest strength — and its most subtle weakness. The language deliberately omits features like ternary operators, generics (until 1.18), enums, and inheritance. This forces clarity at the statement level, but creates a different kind of complexity at the _system_ level:

1. **Verbosity compounds.** A single nil check is fine. A thousand nil checks across fifty packages create a maintenance tax that slowly erodes velocity.
2. **Flat packages fracture consistency.** Without hierarchical namespacing, large codebases accumulate dozens of `NewFoo()`, `NewBar()`, `NewBaz()` functions with no discoverability.
3. **Convention gaps invite drift.** Go prescribes `gofmt` but says nothing about file organization, factory patterns, or test data separation. Each team invents its own, and onboarding becomes archaeology.

`core` exists to solve these problems _once_, at the foundation layer, so that every package built on top inherits the same structure, the same naming, and the same developer experience.

---

## One File Per Function

### The Rule

Each public function or method group lives in its own `.go` file, named after the function.

### Why

**Navigability scales linearly.** In a 500-line file with ten functions, finding the one you need requires scanning or searching. In ten 50-line files, file names _are_ the index. This matters most when you're debugging at 2 AM or onboarding a new contributor who doesn't have your mental model of the codebase.

**Merge conflicts vanish.** When two developers modify different functions, they touch different files. No more resolving conflicts in a shared `utils.go` that everyone edits.

**Blame is precise.** `git blame` on a single-function file tells you exactly who wrote that logic and when — no noise from unrelated changes in the same file.

**AI agents navigate instantly.** Modern AI-assisted development tools work best when they can read a single file and understand its complete purpose. A 50-line file with one function is the ideal unit of context.

### The Trade-off

Yes, this means more files. A package might have 30+ `.go` files instead of 3. But Go tooling handles this gracefully — `gopls` indexes by symbol, not by file. The IDE's file tree becomes a table of contents rather than a wall of text.

### Example

```
coretaskinfo/
├── Info.go                    # struct definition
├── InfoGetters.go             # Name(), Description(), Url()
├── InfoSetters.go             # SetSecure(), SetPlainText()
├── InfoJson.go                # Json(), PrettyJsonString()
├── InfoClone.go               # Clone()
├── InfoNilSafe.go             # SafeName(), SafeDescription()
├── newInfoCreator.go          # New.Info.Default(), New.Info.Secure.*
└── vars.go                    # var New = newCreator{}
```

Every file is short, focused, and self-documenting through its name alone.

---

## Struct-as-Namespace

### The Rule

Related operations are grouped on unexported struct types, exposed via package-level `var` declarations. This creates a hierarchical namespace that IDE autocompletion can traverse.

### Why

Go has no classes, no modules, and no nested namespaces. The package is the only organizational boundary. For small packages, this is fine. For larger ones — a JSON library, a payload system, a testing framework — flat function names become an unnavigable list.

The struct-as-namespace pattern solves this by turning Go's method dispatch into a namespace tree:

```go
// Instead of:
corejson.SerializeToString(v)
corejson.SerializeToBytes(v)
corejson.DeserializeFromBytes(b, v)
corejson.DeserializeFromTo(a, b)

// You get:
corejson.Serialize.ToString(v)
corejson.Serialize.Raw(v)
corejson.Deserialize.UsingBytes(b, v)
corejson.Deserialize.FromTo(a, b)
```

The difference is subtle in two functions but transformative in twenty. When you type `corejson.Serialize.`, the IDE shows you _only_ serialization methods — not the full package API. This is **progressive disclosure**: you see what's relevant at each level of the hierarchy.

### How It Works

```go
// 1. Define unexported struct types
type serializeCreator struct{}
type deserializeCreator struct{}

// 2. Expose as package-level vars
var Serialize serializeCreator
var Deserialize deserializeCreator

// 3. Attach methods
func (it serializeCreator) ToString(v any) (string, error) { ... }
func (it serializeCreator) Raw(v any) ([]byte, error) { ... }
```

The structs are empty — they carry no state. They exist purely as method receivers that create a navigable API surface.

### Why Not Just Use Sub-packages?

Sub-packages create import dependencies and force you to decide upfront which operations belong together. Struct-as-namespace is lighter: it's a grouping mechanism within a single package that can be reorganized without breaking import paths.

---

## The `newCreator` Convention

### The Rule

Object construction uses a hierarchical tree of factory structs, exposed via a package-level `var New`. This replaces flat `NewFoo()` functions with a discoverable creation API.

### Why

Consider a package with ten types, each with three or four constructors. That's 30–40 top-level `New*` functions — an undifferentiated list in autocompletion. The `newCreator` pattern organizes them into a tree:

```go
corepayload.New.PayloadWrapper.Empty()
corepayload.New.PayloadWrapper.UsingInstruction(instr)
corepayload.New.Attributes.FromMap(m)
corepayload.New.PagingInfo.Create(size, index, total)
```

When you type `corepayload.New.`, you see the type categories. When you select `PayloadWrapper`, you see the construction strategies. The API teaches you how to use it.

### The Three-Layer Architecture

1. **Root aggregator** (`newCreator`) — exposed as `var New`. Contains sub-creator fields.
2. **Sub-creators** (`newPayloadWrapperCreator`, `newAttributesCreator`) — one per type or type family.
3. **Factory methods** (`.Empty()`, `.Create(...)`, `.UsingInstruction(...)`) — the actual constructors.

Each layer lives in its own file following the one-file-per-function rule, making the entire creation API navigable through the file tree.

### Generic Typed Creators

For packages like `coregeneric` that provide the same data structure across many types, a `typedXCreator[T]` pattern avoids duplicating factory code:

```go
// Single generic creator supports 16+ primitive types
New.Collection.String.Cap(10)
New.Collection.Int.From(items)
New.Collection.Any.Empty()
```

---

## Interface-First Design

### The Rule

All behavioral contracts are defined as interfaces in `coreinterface/`, following Go's `-er` suffix convention. Packages depend on interfaces, not concrete types.

### Why

**Decoupling.** When package A depends on `NameGetter` instead of `coretaskinfo.Info`, you can swap implementations without touching A's code. This isn't theoretical — it's how the testing framework plugs into the payload system without importing it.

**Discoverability.** The `coreinterface/` package is a catalog of every capability in the ecosystem. Need to know which types can serialize? Search for `Serializer`. Need to know which types have names? Search for `NameGetter`. The interface names _are_ the documentation.

**Composability.** Go interfaces compose implicitly. A type that implements `NameGetter`, `Serializer`, and `IsEmptyChecker` automatically satisfies any interface that embeds all three — no explicit declaration needed. This means small, focused interfaces naturally combine into larger contracts.

### The `-er` Convention

| Pattern | Meaning | Examples |
|---------|---------|----------|
| `*Getter` | Read access to a property | `NameGetter`, `ValueGetter`, `IdGetter` |
| `*Setter` | Write access to a property | `NameSetter`, `ValueSetter` |
| `*Checker` | Boolean state query | `HasErrorChecker`, `IsEmptyChecker`, `IsValidChecker` |
| `*Binder` | Capability attachment | `ContractsBinder`, `AttributesBinder` |
| `*er` | General behavior | `Serializer`, `Cloner`, `Stringer` |

Over 100 interfaces follow this pattern, creating a consistent vocabulary across the entire ecosystem.

---

## Zero-Nil Safety

### The Rule

Functions return empty slices/maps instead of nil. Pointer-receiver methods include nil guards. Types provide explicit state-checking methods: `IsNull()`, `IsEmpty()`, `IsDefined()`.

### Why

Nil panics are Go's most common runtime error. They're also entirely preventable. The cost of returning `[]string{}` instead of `nil` is negligible. The cost of a nil panic in production is not.

**The hierarchy of emptiness matters.** There's a semantic difference between "no value exists" (`nil`), "a value exists but is empty" (empty struct), and "a value exists with content" (populated struct). The `IsNull()` / `IsEmpty()` / `IsDefined()` trio makes this distinction explicit instead of overloading `== nil`.

### In Practice

```go
// ✅ Return empty, not nil
func (it Collection) Items() []string {
    if it.items == nil {
        return []string{}
    }
    return it.items
}

// ✅ Nil-safe method with guard
func (it *Info) SafeName() string {
    if it == nil {
        return ""
    }
    return it.RootName
}

// ✅ Explicit state checking
if payload.IsNull() { ... }    // no value at all
if payload.IsEmpty() { ... }   // value exists, no content
if payload.IsDefined() { ... } // value exists with content
```

---

## Value Receivers by Default

### The Rule

New code uses value receivers for all read-only methods. Pointer receivers are reserved for mutation, large structs, or interface satisfaction.

### Why

Value receivers communicate intent: _this method does not modify the receiver_. This is the most common case — getters, formatters, serializers, cloners. Using a value receiver makes immutability the default and mutation the exception.

It also eliminates an entire class of nil-safety guards. A value receiver can never be nil — the method operates on a copy. Only pointer receivers need `if it == nil` checks.

```go
// Value receiver: no nil guard needed, communicates immutability
func (it Info) Name() string { return it.RootName }
func (it Info) Clone() Info  { return Info{RootName: it.RootName} }

// Pointer receiver: mutation, needs nil guard
func (it *Info) SetName(name string) {
    if it == nil {
        return
    }
    it.RootName = name
}
```

---

## Large Type Splitting

### The Rule

Types exceeding ~400 lines are split into multiple files by responsibility: `{Type}.go` (struct + identity), `{Type}Getters.go`, `{Type}Setters.go`, `{Type}Json.go`, `{Type}Validation.go`.

### Why

This is the one-file-per-function principle applied at the type level. A 1,000-line type file is just as hard to navigate as a 1,000-line function file. Splitting by responsibility means each file has a clear purpose:

| File | Contains |
|------|----------|
| `Attributes.go` | Struct definition, constructor |
| `AttributesGetters.go` | All read-only accessors |
| `AttributesSetters.go` | All mutators |
| `AttributesJson.go` | `MarshalJSON`, `UnmarshalJSON`, `Json()` |
| `AttributesValidation.go` | `IsValid()`, `Validate()` |

The public API stays unified — callers don't know or care that the type is split across files. But contributors can find and modify the serialization logic without scrolling past fifty getter methods.

---

## Testing as Architecture

### The Rule

Test data lives in `_testcases.go` files, separate from test logic in `_test.go` files. Expectations use `args.Map` with semantic keys. All tests follow the AAA (Arrange-Act-Assert) pattern.

### Why

**Separation of data and logic.** When a test fails, you need to answer two questions: "Is the expected value wrong?" and "Is the test logic wrong?" When data and logic are interleaved, answering either question requires reading both. Separating them means you can review all test cases as a table without parsing assertion code.

**Semantic keys over positional indices.** `args.Map{"hasError": true, "keyCount": 3}` is self-documenting. `[]string{"true", "3"}` requires you to check which index means what — and that mapping lives only in the test logic file, not in the data file.

**Diagnostic clarity.** When `ShouldBeEqualMap` fails, it reports _which key_ mismatched: `"hasError": expected true, got false`. When a `[]string` assertion fails, it reports `index 0: expected "true", got "false"` — meaningless without cross-referencing the test logic.

```go
// _testcases.go — pure data, no logic
var createUserTestCase = coretestcases.CaseV1{
    Title: "valid user creation",
    Input: args.Map{"name": "Alice", "email": "alice@example.com"},
    ExpectedInput: args.Map{
        "hasError": false,
        "userName": "Alice",
        "isActive": true,
    },
}

// _test.go — pure logic, no data
func Test_CreateUser(t *testing.T) {
    tc := createUserTestCase
    user, err := CreateUser(tc.Input)
    actual := args.Map{
        "hasError": err != nil,
        "userName": user.Name,
        "isActive": user.Active,
    }
    tc.ShouldBeEqualMapFirst(t, actual)
}
```

---

## Why Not Just Use the Standard Library?

The standard library is excellent for what it provides. `core` doesn't replace it — it builds on top of it to fill the gaps that appear in real-world, large-scale projects:

| Gap | Standard Library | `core` |
|-----|-----------------|--------|
| Ternary expressions | `if/else` blocks | `conditional.If[T](cond, a, b)` |
| Nil-safe access | Manual guards everywhere | Built-in `Safe*()` methods |
| JSON pipeline | `json.Marshal` + error handling | `corejson.Serialize.ToString(v)` |
| Factory patterns | Flat `New*()` functions | Hierarchical `New.Type.Method()` |
| Test assertions | `t.Errorf` with manual messages | `ShouldBeEqualMap` with semantic diffs |
| Interface catalog | Defined per-package, inconsistently | 100+ canonical contracts in `coreinterface/` |
| Compute-once caching | `sync.Once` + manual wiring | `coreonce.StringOnce`, `IntegerOnce`, etc. |

The goal isn't to abstract away Go — it's to eliminate the repetitive scaffolding so you can focus on the logic that matters.

---

## Summary

Every convention in this codebase exists to solve a specific, recurring problem:

| Convention | Problem It Solves |
|------------|-------------------|
| One file per function | Navigation, merge conflicts, blame precision |
| Struct-as-namespace | Flat namespace clutter, poor discoverability |
| `newCreator` | Constructor sprawl, no progressive disclosure |
| Interface-first | Tight coupling, poor testability |
| Zero-nil safety | Runtime panics, ambiguous empty states |
| Value receivers | Accidental mutation, unnecessary nil guards |
| Large type splitting | Monolithic files, tangled responsibilities |
| Test data separation | Interleaved data/logic, poor diagnostics |

These aren't aesthetic preferences. They're engineering decisions made to keep a large, growing codebase maintainable by both humans and AI agents over years of evolution.

---

## Related Docs

- [README](README.md) — what the framework provides
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md) — formatting rules and code conventions
- [newCreator Convention](/spec/01-app/18-new-creator-convention.md) — full factory pattern specification
- [Interface Conventions](/spec/01-app/14-core-interface-conventions.md) — `-er` suffix rules
- [Testing Guidelines](/spec/01-app/16-testing-guidelines.md) — AAA pattern and `args.Map` mandate
- [Repo Overview](/spec/01-app/00-repo-overview.md) — folder structure and responsibilities

## Contributors

- [Md. Alim Ul Karim](https://www.linkedin.com/in/alimkarim) — Architecture & design decisions

## Issues for Future Reference
