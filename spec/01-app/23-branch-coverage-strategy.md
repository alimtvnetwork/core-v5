# Branch Coverage Strategy — Systematic Test Path Identification

> **Purpose**: This document defines how to identify every branch in source code and ensure each path has at least one positive and one negative test case. Any AI agent or contributor can use this guide to audit coverage and write missing tests.

## Table of Contents

- [Philosophy](#philosophy)
- [Branch Identification Checklist](#branch-identification-checklist)
- [Positive vs Negative Test Cases](#positive-vs-negative-test-cases)
- [Step-by-Step Process](#step-by-step-process)
- [If/Else Branch Coverage](#ifelse-branch-coverage)
- [Nil Guard Coverage](#nil-guard-coverage)
- [Error Path Coverage](#error-path-coverage)
- [Switch/Case Coverage](#switchcase-coverage)
- [Early Return Coverage](#early-return-coverage)
- [Boundary and Edge Cases](#boundary-and-edge-cases)
- [Coverage Matrix Template](#coverage-matrix-template)
- [Examples: Source → Test Cases](#examples-source--test-cases)
- [Anti-Patterns in Coverage](#anti-patterns-in-coverage)
- [Related Docs](#related-docs)

---

## Philosophy

Every `if`, `switch`, `for`, early `return`, and nil guard in production code represents a **decision point**. Each decision has at least two outcomes. A well-tested codebase covers **both outcomes** for every decision:

1. **Positive path** — the expected/happy case (condition is true, input is valid, value exists).
2. **Negative path** — the alternative/failure case (condition is false, input is invalid, value is nil/missing).

**Goal**: No `if` statement in source code should lack a corresponding test for both its true and false branches.

---

## Branch Identification Checklist

When auditing a function or method, scan for these decision points:

| Branch Type | What to Look For | Test Cases Needed |
|-------------|-----------------|-------------------|
| `if condition` | Simple guard | True path + False path |
| `if/else` | Two-way decision | Both branches |
| `if/else if/else` | Multi-way decision | Each branch + default |
| `switch/case` | Dispatch | Each case + default |
| Nil guard (`if it == nil`) | Nil receiver safety | Nil receiver + non-nil receiver |
| Error check (`if err != nil`) | Error propagation | Error path + success path |
| Early return | Short-circuit | Triggering condition + fallthrough |
| `for` loop | Iteration | Empty collection + single item + multiple items |
| Boolean parameter | Function behavior toggle | `true` input + `false` input |

---

## Positive vs Negative Test Cases

### Definitions

| Term | Meaning | Example |
|------|---------|---------|
| **Positive test** | Input satisfies the method's intended use case. The happy path. | Valid email → returns true |
| **Negative test** | Input violates expectations. Tests robustness and error handling. | Empty string → returns false |
| **Boundary test** | Input sits at the edge of valid/invalid ranges. | Exactly 0 items, exactly max items |
| **Nil safety test** | Receiver or input is nil. | `(*MyStruct)(nil).IsValid()` → false |

### Naming Convention

```
Positive: Test_TypeName_Method_ValidScenario
Negative: Test_TypeName_Method_InvalidScenario
Boundary: Test_TypeName_Method_EmptyInput
Nil:      Test_TypeName_NilReceiver (via CaseNilSafe)
```

### Test Case Title Convention

```go
// Positive
Title: "Positive: valid email is accepted"
Title: "HasItems returns true when items exist"

// Negative
Title: "Negative: empty string returns error"
Title: "HasItems returns false when empty"

// Boundary
Title: "Boundary: single item collection"
Title: "Length returns 0 for empty list"
```

---

## Step-by-Step Process

### Step 1: Read the Source Code

Open the function/method under test. List every `if`, `switch`, early `return`, and nil guard.

### Step 2: Map Decision Points

For each decision, document:
- **Line number** and condition
- **True branch** behavior
- **False branch** behavior

### Step 3: Cross-Reference Existing Tests

Check `tests/integratedtests/{pkg}tests/` for existing test cases. Mark which branches are covered.

### Step 4: Identify Gaps

Any branch without a test case for **both** outcomes is a gap.

### Step 5: Write Missing Test Cases

Follow the [Testing Guidelines](./16-testing-guidelines.md) for format and structure.

---

## If/Else Branch Coverage

### Source Code

```go
func (it *Validator) IsValid(input string) bool {
    if it == nil {                    // Branch 1: nil guard
        return false
    }

    if input == "" {                  // Branch 2: empty check
        return false
    }

    if len(input) > it.MaxLength {   // Branch 3: max length
        return false
    }

    return true
}
```

### Required Test Cases

| Branch | Condition | Positive Test | Negative Test |
|--------|-----------|--------------|---------------|
| 1 | `it == nil` | Non-nil receiver → continues | Nil receiver → returns false |
| 2 | `input == ""` | Non-empty input → continues | Empty input → returns false |
| 3 | `len > MaxLength` | Short input → returns true | Long input → returns false |

### Test Case Data (`_testcases.go`)

```go
// Branch 1: Nil receiver (use CaseNilSafe)
var validatorNilSafeTestCases = []coretestcases.CaseNilSafe{
    {
        Title: "IsValid on nil returns false",
        Func:  (*Validator).IsValid,
        Args:  []any{"hello"},
        Expected: results.ResultAny{
            Value:    "false",
            Panicked: false,
        },
    },
}

// Branch 2+3: Positive and negative (use CaseV1)
var validatorIsValidEmptyTestCase = coretestcases.CaseV1{
    Title: "Negative: empty input returns false",
    ArrangeInput: args.Map{
        "input":     "",
        "maxLength": 100,
    },
    ExpectedInput: "false",
}

var validatorIsValidTooLongTestCase = coretestcases.CaseV1{
    Title: "Negative: input exceeds max length",
    ArrangeInput: args.Map{
        "input":     "this is way too long",
        "maxLength": 5,
    },
    ExpectedInput: "false",
}

var validatorIsValidHappyTestCase = coretestcases.CaseV1{
    Title: "Positive: valid input within bounds",
    ArrangeInput: args.Map{
        "input":     "hello",
        "maxLength": 100,
    },
    ExpectedInput: "true",
}
```

### Test Logic (`_test.go`)

```go
func Test_Validator_NilReceiver(t *testing.T) {
    for caseIndex, tc := range validatorNilSafeTestCases {
        // Arrange (implicit — nil receiver)

        // Act & Assert
        tc.ShouldBeSafe(t, caseIndex)
    }
}

func Test_Validator_IsValid_EmptyInput(t *testing.T) {
    tc := validatorIsValidEmptyTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    str, _ := input.GetAsString("input")
    maxLen, _ := input.GetAsInt("maxLength")
    v := NewValidator(maxLen)

    // Act
    result := v.IsValid(str)

    // Assert
    tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", result))
}
```

---

## Nil Guard Coverage

### Rule: Every nil guard → one CaseNilSafe entry

```go
// Source: every method with `if it == nil { return ... }`
func (it *MyStruct) IsEmpty() bool {
    if it == nil { return true }
    return len(it.items) == 0
}

func (it *MyStruct) Length() int {
    if it == nil { return 0 }
    return len(it.items)
}

func (it *MyStruct) String() string {
    if it == nil { return "" }
    return it.name
}
```

### Required CaseNilSafe entries (one per method with nil guard)

```go
var myStructNilSafeTestCases = []coretestcases.CaseNilSafe{
    {
        Title: "IsEmpty on nil returns true",
        Func:  (*MyStruct).IsEmpty,
        Expected: results.ResultAny{
            Value:    "true",
            Panicked: false,
        },
    },
    {
        Title: "Length on nil returns 0",
        Func:  (*MyStruct).Length,
        Expected: results.ResultAny{
            Value:    "0",
            Panicked: false,
        },
    },
    {
        Title: "String on nil returns empty",
        Func:  (*MyStruct).String,
        Expected: results.ResultAny{
            Value:    "",
            Panicked: false,
        },
    },
}
```

### Also test the non-nil path

For each method with a nil guard, there must also be a **positive test** that exercises the non-nil path:

```go
var myStructIsEmptyNonEmptyTestCase = coretestcases.CaseV1{
    Title:         "Positive: IsEmpty returns false when items exist",
    ExpectedInput: "false",
}

var myStructIsEmptyEmptyTestCase = coretestcases.CaseV1{
    Title:         "Positive: IsEmpty returns true when no items",
    ExpectedInput: "true",
}
```

---

## Error Path Coverage

### Source Code

```go
func (it *Parser) Parse(input string) (Result, error) {
    if input == "" {
        return Result{}, fmt.Errorf("input is required")
    }

    tokens, err := tokenize(input)
    if err != nil {
        return Result{}, fmt.Errorf("tokenize failed: %w", err)
    }

    return buildResult(tokens), nil
}
```

### Required Tests

| Path | Test Case | Type |
|------|-----------|------|
| Empty input | Returns error "input is required" | Negative |
| Tokenize fails | Returns wrapped tokenize error | Negative |
| Happy path | Returns Result, nil error | Positive |

### Test Case Data

```go
var parseEmptyInputTestCase = coretestcases.CaseV1{
    Title: "Negative: empty input returns error",
    ArrangeInput: args.Map{
        "input": "",
    },
    ExpectedInput: args.Map{
        "hasError":  true,
        "errorText": "input is required",
    },
}

var parseValidInputTestCase = coretestcases.CaseV1{
    Title: "Positive: valid input returns result",
    ArrangeInput: args.Map{
        "input": "valid data",
    },
    ExpectedInput: args.Map{
        "hasError": false,
        "hasValue": true,
    },
}
```

---

## Switch/Case Coverage

### Source Code

```go
func (it *Converter) Convert(kind string) string {
    switch kind {
    case "json":
        return it.toJSON()
    case "xml":
        return it.toXML()
    case "csv":
        return it.toCSV()
    default:
        return ""
    }
}
```

### Required Tests (one per case + default)

```go
var convertJsonTestCase = coretestcases.CaseV1{
    Title:         "Convert json produces JSON output",
    ArrangeInput:  args.Map{"kind": "json"},
    ExpectedInput: `{"name":"test"}`,
}

var convertXmlTestCase = coretestcases.CaseV1{
    Title:         "Convert xml produces XML output",
    ArrangeInput:  args.Map{"kind": "xml"},
    ExpectedInput: `<name>test</name>`,
}

var convertCsvTestCase = coretestcases.CaseV1{
    Title:         "Convert csv produces CSV output",
    ArrangeInput:  args.Map{"kind": "csv"},
    ExpectedInput: `name,test`,
}

var convertUnknownTestCase = coretestcases.CaseV1{
    Title:         "Negative: unknown kind returns empty",
    ArrangeInput:  args.Map{"kind": "yaml"},
    ExpectedInput: "",
}
```

---

## Early Return Coverage

### Source Code

```go
func (it *Cache) Get(key string) (string, bool) {
    if key == "" {             // Early return 1
        return "", false
    }

    if it.isExpired(key) {     // Early return 2
        it.delete(key)
        return "", false
    }

    val, exists := it.data[key]
    return val, exists          // Normal return
}
```

### Required Tests

| Path | Condition | Expected |
|------|-----------|----------|
| Early return 1 | `key == ""` | `("", false)` |
| Early return 2 | Key exists but expired | `("", false)` |
| Normal (found) | Key exists, not expired | `(value, true)` |
| Normal (not found) | Key doesn't exist | `("", false)` |

---

## Boundary and Edge Cases

### Common Boundaries

| Boundary | Test With |
|----------|-----------|
| Empty string | `""` |
| Empty slice | `[]string{}` |
| Nil pointer | `nil` |
| Zero value | `0`, `false`, `""` |
| Single item | `[]string{"one"}` |
| Max value | `math.MaxInt` |
| Unicode | `"こんにちは"`, `"🎉"` |
| Special chars | `"a\nb"`, `"a\tb"`, `"a\"b"` |

### Example: Collection Boundaries

```go
// Source
func (it *Collection) First() string {
    if it == nil || len(it.items) == 0 {
        return ""
    }
    return it.items[0]
}
```

### Required Tests

```go
// Nil receiver (CaseNilSafe)
{Title: "First on nil returns empty", Func: (*Collection).First, Expected: results.ResultAny{Value: "", Panicked: false}}

// Empty collection (CaseV1)
var collectionFirstEmptyTestCase = coretestcases.CaseV1{
    Title:         "Boundary: First on empty returns empty",
    ExpectedInput: "",
}

// Single item
var collectionFirstSingleTestCase = coretestcases.CaseV1{
    Title:         "Positive: First on single-item returns that item",
    ExpectedInput: "hello",
}

// Multiple items
var collectionFirstMultiTestCase = coretestcases.CaseV1{
    Title:         "Positive: First on multi-item returns first",
    ExpectedInput: "alpha",
}
```

---

## Coverage Matrix Template

Use this template when auditing a type's test coverage:

```markdown
## Coverage Matrix: TypeName

| Method | Nil Guard | Positive | Negative | Boundary | Error Path | Status |
|--------|-----------|----------|----------|----------|------------|--------|
| IsValid | ✅ | ✅ | ✅ | ✅ | N/A | ✅ |
| Parse | ✅ | ✅ | ❌ | ❌ | ✅ | ⚠️ |
| String | ✅ | ✅ | N/A | N/A | N/A | ✅ |
| Process | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |

Legend:
- ✅ = Covered
- ❌ = Missing test
- ⚠️ = Partially covered
- N/A = Not applicable
```

---

## Examples: Source → Test Cases

### Full Example: Method with Multiple Branches

```go
// Source: corestr/Hashset.go
func (it *Hashset) Add(item string) bool {
    if it == nil {           // Branch 1
        return false
    }

    if item == "" {          // Branch 2
        return false
    }

    if it.Contains(item) {   // Branch 3
        return false
    }

    it.items[item] = true
    return true
}
```

### Complete Test Coverage

```go
// _NilReceiver_testcases.go — Branch 1
var hashsetNilSafeTestCases = []coretestcases.CaseNilSafe{
    {
        Title: "Add on nil returns false",
        Func:  (*Hashset).Add,
        Args:  []any{"hello"},
        Expected: results.ResultAny{
            Value:    "false",
            Panicked: false,
        },
    },
}

// _testcases.go — Branches 2 and 3
var hashsetAddEmptyTestCase = coretestcases.CaseV1{
    Title: "Negative: Add empty string returns false",
    ArrangeInput: args.Map{
        "item": "",
    },
    ExpectedInput: "false",
}

var hashsetAddDuplicateTestCase = coretestcases.CaseV1{
    Title: "Negative: Add duplicate returns false",
    ArrangeInput: args.Map{
        "item":     "hello",
        "existing": []string{"hello"},
    },
    ExpectedInput: "false",
}

var hashsetAddNewTestCase = coretestcases.CaseV1{
    Title: "Positive: Add new item returns true",
    ArrangeInput: args.Map{
        "item":     "hello",
        "existing": []string{},
    },
    ExpectedInput: "true",
}
```

---

## Anti-Patterns in Coverage

### ❌ Testing only the happy path

```go
// ❌ Only tests positive — misses nil, empty, boundary
var tests = []coretestcases.CaseV1{
    {Title: "Works", ArrangeInput: args.Map{"input": "valid"}, ExpectedInput: "true"},
}
```

### ❌ Testing branches in wrong structure

```go
// ❌ Using CaseV1 for nil receiver — use CaseNilSafe instead
var nilTestCase = coretestcases.CaseV1{
    Title:         "Nil receiver returns false",
    ArrangeInput:  nil,
    ExpectedInput: "false",
}
```

### ❌ Missing default case in switch tests

```go
// ❌ Tests "json" and "xml" but not the default branch
```

### ❌ Not testing error propagation

```go
// ❌ Only tests that error is non-nil — doesn't verify error message
tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", err != nil))
// Missing: tc.ShouldContains(t, 0, err.Error()) for message validation
```

### ✅ Complete coverage pattern

For every method, ensure:
1. **CaseNilSafe** for nil receiver (if pointer receiver)
2. **Positive CaseV1** for happy path
3. **Negative CaseV1** for each failure branch
4. **Boundary CaseV1** for edge values
5. **Error CaseV1** for error return paths

---

## Related Docs

- [Testing Guidelines](./16-testing-guidelines.md) — assertion methods, AAA format, file organization
- [Test Quality Guide](./24-test-quality-guide.md) — good vs bad test examples
- [Test Structures Reference](./25-test-structures-reference.md) — CaseV1, CaseNilSafe, args.*, results.* detailed API
- [CaseNilSafe Design](./designs/CaseNilSafe-design.md) — architecture and edge cases
- [Coding Guidelines](./17-coding-guidelines.md) — formatting and naming conventions
