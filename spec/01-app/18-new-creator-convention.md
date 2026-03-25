# The `newCreator` Convention — Hierarchical Factory Pattern

> This is the core architectural pattern for object construction in the `core` framework.
> It decomposes complex packages into a tree of small, focused factory structs exposed via a single `New` (or similar) entry point.

## Table of Contents

1. [Problem It Solves](#problem-it-solves)
2. [Core Concept](#core-concept)
3. [Anatomy of the Pattern](#anatomy-of-the-pattern)
4. [Real Examples from the Codebase](#real-examples-from-the-codebase)
5. [Step-by-Step: How to Implement](#step-by-step-how-to-implement)
6. [Naming Conventions](#naming-conventions)
7. [Design Rules](#design-rules)
8. [Advanced: Multi-Level Nesting](#advanced-multi-level-nesting)
9. [Advanced: Namespace Structs for Non-Factory Operations](#advanced-namespace-structs-for-non-factory-operations)
10. [Anti-Patterns](#anti-patterns)
11. [Summary Cheat Sheet](#summary-cheat-sheet)

---

## Problem It Solves

Go has no constructors, no classes, and no method overloading. In a package with many types (e.g., `corepayload` has `PayloadWrapper`, `Attributes`, `PayloadsCollection`, `User`), you end up with:

```go
// ❌ Flat function soup — hard to discover, no grouping
func NewPayloadWrapper() *PayloadWrapper { ... }
func NewPayloadWrapperFromBytes(b []byte) (*PayloadWrapper, error) { ... }
func NewPayloadWrapperFromInstruction(i *Instruction) (*PayloadWrapper, error) { ... }
func NewAttributes() *Attributes { ... }
func NewAttributesFromAuth(a *AuthInfo) *Attributes { ... }
func NewUser(name string) *User { ... }
// ... 30 more functions with no discoverability
```

The `newCreator` convention solves this by creating a **tree of factory structs** that turns construction into a **navigable autocomplete path**:

```go
// ✅ Tree-structured — IDE autocomplete guides you
payload := corepayload.New.PayloadWrapper.Empty()
payload := corepayload.New.PayloadWrapper.Deserialize(bytes)
attrs   := corepayload.New.Attributes.UsingAuthInfo(auth)
user    := corepayload.New.User.Default("alice")
```

---

## Core Concept

```
Package var    → Root creator struct  → Sub-creator structs → Factory methods
  New          →   newCreator{}      →   newXCreator{}     → .Empty(), .Create(), .Deserialize()
```

The pattern has **three layers**:

1. **Package-level `var`** — single entry point (e.g., `var New = newCreator{}`)
2. **Root creator struct** — aggregates sub-creators as fields (e.g., `type newCreator struct { PayloadWrapper newPayloadWrapperCreator; Attributes newAttributesCreator }`)
3. **Sub-creator structs** — each owns factory methods for one type (e.g., `type newPayloadWrapperCreator struct{}` with methods `Empty()`, `Deserialize()`, `Create()`)

---

## Anatomy of the Pattern

### File: `vars.go` — Package-level entry point

```go
package mypkg

var New = newCreator{}
```

### File: `newCreator.go` — Root aggregator (one per package)

```go
package mypkg

type newCreator struct {
    Widget    newWidgetCreator
    Config    newConfigCreator
    Session   newSessionCreator
}
```

**Rules**:
- The struct is **unexported** (`newCreator`, not `NewCreator`)
- Fields are **exported** (`Widget`, not `widget`) — this enables `New.Widget.Create()`
- Field types are **unexported** (`newWidgetCreator`) — users never reference the type directly
- The root struct has **no methods** (it's purely an aggregator) OR has convenience shortcut methods

### File: `newWidgetCreator.go` — Sub-creator (one per type)

```go
package mypkg

type newWidgetCreator struct{}

func (it newWidgetCreator) Empty() *Widget {
    return &Widget{
        Items: []string{},  // zero-nil safety: never nil
    }
}

func (it newWidgetCreator) Create(name string, size int) *Widget {
    return &Widget{
        Name:  name,
        Size:  size,
        Items: []string{},
    }
}

func (it newWidgetCreator) Deserialize(rawBytes []byte) (*Widget, error) {
    w := &Widget{}
    err := json.Unmarshal(rawBytes, w)
    return w, err
}
```

---

## Real Examples from the Codebase

### Example 1: `regexnew` — Two-level tree with shortcut methods

```
regexnew.New.Lazy("pattern")              // shortcut on root
regexnew.New.LazyRegex.New("pattern")     // full path via sub-creator
regexnew.New.LazyRegex.NewLock("pattern") // thread-safe variant
regexnew.New.LazyRegex.TwoLock(p1, p2)   // create two at once
regexnew.New.Default("pattern")           // compiled *regexp.Regexp
```

**File structure:**
```
regexnew/
├── vars.go                  → var New = newCreator{}
├── newCreator.go            → type newCreator struct { LazyRegex newLazyRegexCreator }
│                              + shortcut methods: Lazy(), LazyLock(), Default()
├── newLazyRegexCreator.go   → type newLazyRegexCreator struct{}
│                              + methods: New(), NewLock(), TwoLock(), ManyUsingLock()
├── LazyRegex.go             → the actual LazyRegex struct
└── ...
```

**Key insight**: The root `newCreator` has **both** sub-creators as fields AND shortcut methods that delegate to them. This gives users a choice:

```go
// Short path (common usage):
regex := regexnew.New.Lazy(`\d+`)

// Full path (when you need the sub-creator's specific methods):
allPatterns := regexnew.New.LazyRegex.AllPatternsMap()
```

### Example 2: `corepayload` — Four sub-creators

```
corepayload.New.PayloadWrapper.Empty()
corepayload.New.PayloadWrapper.Deserialize(bytes)
corepayload.New.PayloadWrapper.UsingCreateInstruction(instr)
corepayload.New.PayloadWrapper.Create(name, id, task, cat, record)
corepayload.New.Attributes.Empty()
corepayload.New.Attributes.UsingAuthInfo(auth)
corepayload.New.Attributes.UsingBasicError(err)
corepayload.New.PayloadsCollection.Deserialize(bytes)
corepayload.New.User.Default(name)
```

**File structure:**
```
coredata/corepayload/
├── vars.go                         → var New = newCreator{}
├── newCreator.go                   → type newCreator struct {
│                                       PayloadWrapper     newPayloadWrapperCreator
│                                       PayloadsCollection newPayloadsCollectionCreator
│                                       Attributes         newAttributesCreator
│                                       User               newUserCreator
│                                     }
├── newPayloadWrapperCreator.go     → 500+ lines of factory methods
├── newAttributesCreator.go         → factory methods for Attributes
├── newPayloadsCollectionCreator.go → factory methods for PayloadsCollection
├── newUserCreator.go               → factory methods for User
├── PayloadWrapper.go               → the actual struct + behavior methods
├── Attributes.go                   → the actual struct + behavior methods
└── ...
```

### Example 3: `corejson` — Multiple namespace vars (not just `New`)

```go
// vars.go
var (
    Empty       = emptyCreator{}
    Serialize   = serializerLogic{}
    Deserialize = deserializerLogic{}
    NewResult   = newResultCreator{}
    CastAny     = castingAny{}
    AnyTo       = anyTo{}
)
```

**Usage:**
```go
corejson.Serialize.ToString(obj)
corejson.Deserialize.UsingBytes(bytes, &target)
corejson.NewResult.UsingTypeBytesPtr("name", bytes)
corejson.CastAny.FromToDefault(src, dst)
corejson.Empty.ResultPtr()
```

**Key insight**: When a package has logically distinct operation groups (serialize vs deserialize vs create vs cast), use **multiple namespace vars** instead of nesting everything under a single `New`.

### Example 4: `coretaskinfo` — Nested sub-creators with variants

```
coretaskinfo.New.Info.Default(name, desc, url)
coretaskinfo.New.Info.Examples(name, desc, url, examples...)
coretaskinfo.New.Info.Plain.Default(name, desc, url)
coretaskinfo.New.Info.Secure.Default(name, desc, url)
```

**File structure:**
```
coretaskinfo/
├── vars.go                       → var New = newCreator{}
├── newCreator.go                 → type newCreator struct { Info newInfoCreator }
├── newInfoCreator.go             → type newInfoCreator struct {
│                                     Plain  newInfoPlainTextCreator
│                                     Secure newInfoSecureTextCreator
│                                   }
│                                   + methods: Default(), Examples(), Create()
├── newInfoPlainTextCreator.go    → type newInfoPlainTextCreator struct{}
├── newInfoSecureTextCreator.go   → type newInfoSecureTextCreator struct{}
└── Info.go                       → the actual Info struct
```

**Key insight**: This is a **three-level tree**. `newInfoCreator` is both a sub-creator (has its own `Default()`, `Create()` methods) AND an aggregator (has `Plain` and `Secure` sub-creators). This allows:

```go
// Generic creation:
info := coretaskinfo.New.Info.Default("name", "desc", "url")

// Variant-specific creation:
secure := coretaskinfo.New.Info.Secure.Default("name", "desc", "url")
plain  := coretaskinfo.New.Info.Plain.Default("name", "desc", "url")
```

### Example 5: `keymk` — Multiple entry-point vars

```go
var (
    NewKey           = &newKeyCreator{}
    NewKeyWithLegend = &newKeyWithLegendCreator{}
    FixedLegend      = fixedLegend{}
)
```

**Key insight**: When a package has **multiple independent creatable types** that don't share a common root concept, use separate vars instead of forcing them under one `New`.

---

## Step-by-Step: How to Implement

### Scenario: Package `notification` with types `Email`, `SMS`, `Template`

**Step 1**: Create `vars.go`

```go
package notification

var New = newCreator{}
```

**Step 2**: Create `newCreator.go`

```go
package notification

type newCreator struct {
    Email    newEmailCreator
    SMS      newSMSCreator
    Template newTemplateCreator
}
```

**Step 3**: Create `newEmailCreator.go`

```go
package notification

type newEmailCreator struct{}

func (it newEmailCreator) Empty() *Email {
    return &Email{
        Recipients: []string{},
        Headers:    map[string]string{},
    }
}

func (it newEmailCreator) Simple(to, subject, body string) *Email {
    return &Email{
        Recipients: []string{to},
        Subject:    subject,
        Body:       body,
        Headers:    map[string]string{},
    }
}

func (it newEmailCreator) UsingTemplate(
    to string,
    tmpl *Template,
    data map[string]string,
) (*Email, error) {
    body, err := tmpl.Render(data)
    if err != nil {
        return nil, err
    }
    return &Email{
        Recipients: []string{to},
        Subject:    tmpl.Subject,
        Body:       body,
        Headers:    map[string]string{},
    }, nil
}
```

**Step 4**: Usage

```go
email := notification.New.Email.Simple("alice@example.com", "Hello", "World")
sms := notification.New.SMS.Default("+1234567890", "Hello")
tmpl := notification.New.Template.FromFile("welcome.html")
```

---

## Naming Conventions

| Item | Convention | Example |
|------|-----------|---------|
| Package var | `New` (or descriptive: `Serialize`, `Empty`) | `var New = newCreator{}` |
| Root struct | `newCreator` (unexported) | `type newCreator struct{}` |
| Sub-creator struct | `new{Type}Creator` (unexported) | `type newWidgetCreator struct{}` |
| Root struct file | `newCreator.go` | — |
| Sub-creator file | `new{Type}Creator.go` | `newWidgetCreator.go` |
| Factory methods | Descriptive verb/noun | `.Empty()`, `.Create()`, `.Deserialize()`, `.UsingX()` |
| Shortcut methods | Same name as sub-creator method | Root `.Lazy()` delegates to `.LazyRegex.New()` |

### Common Factory Method Names

| Method | When to Use |
|--------|-------------|
| `Empty()` | Zero-value instance (non-nil, safe to use) |
| `Default(...)` | Most common creation path |
| `Create(...)` | Full creation with all required params |
| `All(...)` | Creation with all possible params |
| `Deserialize(bytes)` | From JSON bytes |
| `DeserializeUsingJsonResult(r)` | From `corejson.Result` |
| `UsingX(x)` | Creation from a specific input type |
| `CastOrDeserializeFrom(any)` | Try cast first, fallback to JSON roundtrip |

---

## Design Rules

1. **One sub-creator per type** — `newPayloadWrapperCreator` only creates `PayloadWrapper` instances.
2. **Unexported structs, exported fields** — users access via `New.Widget`, never `newWidgetCreator{}`.
3. **Zero-nil safety in factories** — always initialize slices/maps to empty (not nil).
4. **One file per creator** — `newWidgetCreator.go` contains only `newWidgetCreator` and its methods.
5. **Value receivers on creators** — creator structs are stateless (empty structs), use value receivers.
6. **No business logic in creators** — they only construct and return. Business logic lives on the created type.
7. **Shortcut methods are optional** — only add them for the most common creation paths.

---

## Advanced: Multi-Level Nesting

When a type has **variants** (like `Secure` vs `Plain`), nest sub-creators:

```go
// Level 1: Root
type newCreator struct {
    Info newInfoCreator    // ← has its own sub-creators
}

// Level 2: Type creator (also an aggregator)
type newInfoCreator struct {
    Plain  newInfoPlainTextCreator   // ← Level 3
    Secure newInfoSecureTextCreator  // ← Level 3
}

// Level 2 also has its own methods:
func (it newInfoCreator) Default(name, desc, url string) *Info { ... }

// Level 3: Variant creators
type newInfoPlainTextCreator struct{}
func (it newInfoPlainTextCreator) Default(name, desc, url string) *Info { ... }

type newInfoSecureTextCreator struct{}
func (it newInfoSecureTextCreator) Default(name, desc, url string) *Info { ... }
```

**Result:**
```go
info   := coretaskinfo.New.Info.Default(...)        // generic
plain  := coretaskinfo.New.Info.Plain.Default(...)   // plain variant
secure := coretaskinfo.New.Info.Secure.Default(...)  // secure variant
```

**Rule**: Don't nest deeper than 3 levels. If you need more, your package needs splitting.

---

## Advanced: Namespace Structs for Non-Factory Operations

The same pattern works for **operation grouping**, not just factories:

```go
// corejson/vars.go
var (
    Serialize   = serializerLogic{}   // groups marshal operations
    Deserialize = deserializerLogic{} // groups unmarshal operations
    CastAny     = castingAny{}        // groups casting operations
    Empty       = emptyCreator{}      // groups empty-value factories
)
```

This gives you:
```go
corejson.Serialize.ToString(obj)
corejson.Serialize.Raw(obj)
corejson.Deserialize.UsingBytes(bytes, &target)
corejson.Deserialize.FromTo(src, dst)
```

**When to use multiple vars vs single `New`**:
- **Single `New`**: When everything is a factory for types in this package
- **Multiple vars**: When the package has logically distinct operation groups (serialize/deserialize/cast)
- **Both**: Use `New` for factories, plus additional vars for operation groups

---

## Anti-Patterns

| ❌ Don't | ✅ Do |
|----------|-------|
| `func NewWidget() *Widget` (flat function) | `New.Widget.Empty()` (tree-structured) |
| Export the creator type: `type NewWidgetCreator struct{}` | Keep unexported: `type newWidgetCreator struct{}` |
| Put creation + business logic in same methods | Creators only construct; behavior lives on the type |
| Return nil slices/maps from factories | Always return empty: `Items: []string{}` |
| Nest deeper than 3 levels | Split the package if it's that complex |
| Mix unrelated types in one sub-creator | One sub-creator per type |
| Put all creators in one file | One file per creator: `new{Type}Creator.go` |

---

## Summary Cheat Sheet

```
📁 mypkg/
├── vars.go                    → var New = newCreator{}
├── newCreator.go              → type newCreator struct { A newACreator; B newBCreator }
├── newACreator.go             → type newACreator struct{} + factory methods
├── newBCreator.go             → type newBCreator struct{} + factory methods
├── A.go                       → type A struct{} + behavior methods
└── B.go                       → type B struct{} + behavior methods
```

**The autocomplete path:**
```
mypkg.New.          → shows: A, B
mypkg.New.A.        → shows: Empty(), Create(), Deserialize(), ...
mypkg.New.B.        → shows: Default(), FromConfig(), ...
```

**This is why it works**: IDE autocompletion turns construction into a guided, discoverable experience. Users don't need to read docs — they just type `New.` and follow the tree.
