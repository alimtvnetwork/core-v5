# Test Quality Guide — Good Tests vs Bad Tests

> **Purpose**: A comprehensive reference defining what makes a high-quality test in this codebase, with concrete good and bad examples for every pattern. Portable enough for any AI agent or contributor to produce correct tests without prior codebase knowledge.

## Table of Contents

- [What Makes a Good Test](#what-makes-a-good-test)
- [What Makes a Bad Test](#what-makes-a-bad-test)
- [Side-by-Side Comparisons](#side-by-side-comparisons)
- [Test File Organization Rules](#test-file-organization-rules)
- [AAA Comments Are Mandatory](#aaa-comments-are-mandatory)
- [Assertion Framework Usage](#assertion-framework-usage)
- [Test Case Naming and Titling](#test-case-naming-and-titling)
- [Input and Expected Value Formatting](#input-and-expected-value-formatting)
- [One Scenario Per Test Function](#one-scenario-per-test-function)
- [Error and Panic Testing](#error-and-panic-testing)
- [Nil Receiver Testing](#nil-receiver-testing)
- [Loop Tests vs Named Tests](#loop-tests-vs-named-tests)
- [Complete Quality Checklist](#complete-quality-checklist)
- [Related Docs](#related-docs)

---

## What Makes a Good Test

A good test in this codebase:

1. **Has a clear, descriptive Title** — anyone reading the title knows the scenario.
2. **Follows AAA pattern** with explicit `// Arrange`, `// Act`, `// Assert` comments.
3. **Separates data from logic** — test cases in `_testcases.go`, logic in `_test.go`.
4. **Uses framework assertions** — `ShouldBeEqual`, `ShouldBeEqualMap`, or `errcore.AssertDiffOnMismatch`.
5. **Has no branching** — no `if/else` or `switch` inside test functions.
6. **Covers both positive and negative paths** for every branch.
7. **Uses the right structure** — `CaseV1` for functional tests, `CaseNilSafe` for nil receivers.
8. **Produces self-documenting failure output** — failures show what went wrong with context.
9. **Is deterministic** — runs identically every time.
10. **Tests one thing** — each test function validates exactly one scenario.

---

## What Makes a Bad Test

A bad test:

1. **Uses raw `t.Error` / `t.Errorf` / `t.Fatal`** — bypasses the assertion framework.
2. **Has branching logic** — `if`, `switch`, or conditional assertions inside the test.
3. **Inlines test data** — expected values hardcoded in the test function.
4. **Has no AAA comments** — unclear where arrange ends and act begins.
5. **Tests multiple scenarios** in one function — mixes positive and negative in one test.
6. **Uses indexed slice access** — `testCases[0]` instead of named variables.
7. **Ignores getter errors** — `str, _ := input.GetAsString("key")` without validation.
8. **Has non-descriptive titles** — `"Test 1"`, `"works"`, `"basic"`.
9. **Only tests happy path** — no negative, boundary, or nil cases.
10. **Calls `convey.Convey` directly** — framework handles GoConvey wrapping.

---

## Side-by-Side Comparisons

### 1. Test File Structure

```go
// ❌ BAD — everything in one file, data mixed with logic
func Test_Hashset_Add(t *testing.T) {
    hs := NewHashset()
    result := hs.Add("hello")
    if !result {
        t.Error("expected true")
    }
}

// ✅ GOOD — data separated, framework assertions, AAA comments
// File: Hashset_testcases.go
var hashsetAddNewTestCase = coretestcases.CaseV1{
    Title: "Positive: Add new item returns true",
    ArrangeInput: args.Map{
        "item": "hello",
    },
    ExpectedInput: "true",
}

// File: Hashset_test.go
func Test_Hashset_Add_NewItem(t *testing.T) {
    tc := hashsetAddNewTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    item, _ := input.GetAsString("item")
    hs := NewHashset()

    // Act
    result := hs.Add(item)

    // Assert
    tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", result))
}
```

### 2. Assertions

```go
// ❌ BAD — raw t.Error, no diagnostic context
if result != "hello" {
    t.Errorf("got %s, want hello", result)
}

// ❌ BAD — direct convey call
convey.Convey("test", t, func() {
    convey.So(result, convey.ShouldEqual, "hello")
})

// ✅ GOOD — framework assertion with full context
tc.ShouldBeEqual(t, caseIndex, result)

// ✅ GOOD — diff-based assertion for custom wrappers
errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
```

### 3. Multi-Property Assertions

```go
// ❌ BAD — positional strings, magic indices, no labels
tc.ExpectedInput = []string{"5", "false", "false", "true"}
tc.ShouldBeEqual(t, 0, 
    fmt.Sprintf("%v", v.Value),
    fmt.Sprintf("%v", v.IsZero),
    fmt.Sprintf("%v", v.IsInvalid),
    fmt.Sprintf("%v", v.IsValid),
)

// ✅ GOOD — args.Map with semantic keys, self-documenting failure output
// In _testcases.go:
ExpectedInput: args.Map{
    "value":     5,
    "isZero":    false,
    "isInvalid": false,
    "isValid":   true,
},

// In _test.go:
actual := args.Map{
    "value":     v.ValueInt(),
    "isZero":    v.IsZero(),
    "isInvalid": v.IsInvalid(),
    "isValid":   v.IsValid(),
}
tc.ShouldBeEqualMap(t, caseIndex, actual)
```

### 4. Branching in Tests

```go
// ❌ BAD — branching inside test function
func Test_Parse(t *testing.T) {
    for _, tc := range testCases {
        result, err := Parse(tc.Input)
        if tc.HasError {
            if err == nil {
                t.Error("expected error")
            }
        } else {
            if result != tc.Expected {
                t.Errorf("got %s, want %s", result, tc.Expected)
            }
        }
    }
}

// ✅ GOOD — separate functions, no branching
func Test_Parse_ValidInput(t *testing.T) {
    tc := parseValidTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    str, _ := input.GetAsString("input")

    // Act
    result, _ := Parse(str)

    // Assert
    tc.ShouldBeEqual(t, 0, result)
}

func Test_Parse_InvalidInput(t *testing.T) {
    tc := parseInvalidTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    str, _ := input.GetAsString("input")

    // Act
    _, err := Parse(str)

    // Assert
    tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", err != nil))
}
```

### 5. Test Case References

```go
// ❌ BAD — indexed access, fragile
tc := testCases[0]
tc2 := testCases[3]

// ✅ GOOD — named variables, self-documenting
tc := hashsetAddDuplicateTestCase
tc := hashsetAddNewItemTestCase
```

### 6. Nil Receiver Testing

```go
// ❌ BAD — raw t.Error for nil receiver test
func Test_MyStruct_IsValid_Nil(t *testing.T) {
    var s *MyStruct
    if s.IsValid() != false {
        t.Error("expected false")
    }
}

// ✅ GOOD — CaseNilSafe with compile-time safety
// _NilReceiver_testcases.go
var myStructNilSafeTestCases = []coretestcases.CaseNilSafe{
    {
        Title: "IsValid on nil returns false",
        Func:  (*MyStruct).IsValid,
        Expected: results.ResultAny{
            Value:    "false",
            Panicked: false,
        },
    },
}

// NilReceiver_test.go
func Test_MyStruct_NilReceiver(t *testing.T) {
    for caseIndex, tc := range myStructNilSafeTestCases {
        // Arrange (implicit — nil receiver)

        // Act & Assert
        tc.ShouldBeSafe(t, caseIndex)
    }
}
```

### 7. AAA Comments

```go
// ❌ BAD — no section markers
func Test_X(t *testing.T) {
    tc := xTestCase
    result := doX()
    tc.ShouldBeEqual(t, 0, result)
}

// ✅ GOOD — every section labeled
func Test_X(t *testing.T) {
    tc := xTestCase

    // Arrange
    // (no setup needed)

    // Act
    result := doX()

    // Assert
    tc.ShouldBeEqual(t, 0, result)
}

// ✅ GOOD — CaseNilSafe combined Act & Assert
func Test_X_NilReceiver(t *testing.T) {
    for caseIndex, tc := range xNilSafeTestCases {
        // Arrange (implicit — nil receiver)

        // Act & Assert
        tc.ShouldBeSafe(t, caseIndex)
    }
}
```

---

## Test File Organization Rules

| Rule | Correct | Incorrect |
|------|---------|-----------|
| Test data location | `Feature_testcases.go` | Inline in `_test.go` |
| Test logic location | `Feature_test.go` | Mixed with data |
| Nil receiver data | `Feature_NilReceiver_testcases.go` | In general `_testcases.go` |
| Nil receiver logic | `NilReceiver_test.go` (shared per package) | Separate per type |
| Package name | `{pkg}tests` | Same as production package |
| Directory | `tests/integratedtests/{pkg}tests/` | Anywhere else |

---

## Assertion Framework Usage

### Decision Tree

```
Is this a nil-receiver test?
├── Yes → CaseNilSafe + ShouldBeSafe
└── No
    ├── Is this using CaseV1?
    │   ├── Single value? → ShouldBeEqual
    │   ├── Multiple values with semantic keys? → ShouldBeEqualMap
    │   ├── Order-independent? → ShouldBeSortedEqual
    │   └── Substring match? → ShouldContains
    └── Is this using a custom wrapper?
        └── errcore.AssertDiffOnMismatch
```

### Non-Loop Tests: Use *First Variants

```go
// ❌ Magic 0 in non-loop test
tc.ShouldBeEqual(t, 0, result)
tc.ShouldBeEqualMap(t, 0, actual)

// ✅ Named First variant
tc.ShouldBeEqualFirst(t, result)
tc.ShouldBeEqualMapFirst(t, actual)
```

---

## Test Case Naming and Titling

### Variable Naming

```
{lowerCamelType}{Method}{Scenario}TestCase

Examples:
  hashsetAddNewItemTestCase       ✅
  hashsetAddDuplicateTestCase     ✅
  linkedListLengthEmptyTestCase   ✅
  test1                           ❌
  tc                              ❌
  myTest                          ❌
```

### Title Format

```go
// ✅ Good titles — descriptive, states condition and outcome
Title: "Positive: valid email is accepted"
Title: "Negative: empty input returns error"
Title: "Boundary: single item collection returns correct length"
Title: "IsEmpty on nil returns true"
Title: "Add duplicate returns false"

// ❌ Bad titles
Title: "test 1"
Title: "works"
Title: "basic test"
Title: "error case"
Title: ""  // empty title
```

---

## Input and Expected Value Formatting

### args.Map (ArrangeInput) — multi-line for 2+ entries

```go
// ✅ Good — one key-value per line
ArrangeInput: args.Map{
    "input":  "hello world",
    "sep":    " ",
    "maxLen": 100,
},

// ❌ Bad — inline for multiple entries
ArrangeInput: args.Map{"input": "hello world", "sep": " ", "maxLen": 100},

// ✅ OK — single entry can be inline
ArrangeInput: args.Map{"input": "hello"},
```

### ExpectedInput format selection

```go
// Single value → string
ExpectedInput: "true",
ExpectedInput: "42",
ExpectedInput: "",

// Multiple positional values → []string
ExpectedInput: []string{"alice", "30", "true"},

// Multiple semantic values → args.Map (preferred)
ExpectedInput: args.Map{
    "name":   "alice",
    "age":    30,
    "active": true,
},
```

### Native types in args.Map (not stringified)

```go
// ✅ Good — native types, framework handles conversion
ExpectedInput: args.Map{
    "count":   3,
    "isValid": true,
    "name":    "alice",
},

// ❌ Bad — pre-stringified values
ExpectedInput: args.Map{
    "count":   "3",
    "isValid": "true",
    "name":    "alice",
},
```

---

## One Scenario Per Test Function

### What counts as "one scenario"

- One specific input condition (e.g., empty input, nil receiver, valid input)
- One assertion path (no if/else in the test body)
- One method under test (don't test Add and Remove in the same function)

### Exception: Homogeneous loop tests

When ALL cases share **identical** Arrange/Act/Assert logic (no branching), a loop is acceptable:

```go
// ✅ OK — all cases use identical logic path
func Test_Converter_Output(t *testing.T) {
    for caseIndex, tc := range converterOutputTestCases {
        // Arrange
        input := tc.ArrangeInput.(args.Map)
        str, _ := input.GetAsString("input")

        // Act
        result := Convert(str)

        // Assert
        tc.ShouldBeEqual(t, caseIndex, result)
    }
}
```

---

## Error and Panic Testing

### Error existence (good pattern)

```go
// _testcases.go
var parseErrorTestCase = coretestcases.CaseV1{
    Title: "Negative: empty input returns error",
    ArrangeInput: args.Map{
        "input": "",
    },
    ExpectedInput: args.Map{
        "hasError":     true,
        "errorContains": "required",
    },
}

// _test.go
func Test_Parse_EmptyInput_Error(t *testing.T) {
    tc := parseErrorTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    str, _ := input.GetAsString("input")

    // Act
    _, err := Parse(str)
    actual := args.Map{
        "hasError":     err != nil,
        "errorContains": err != nil && strings.Contains(err.Error(), "required"),
    }

    // Assert
    tc.ShouldBeEqualMap(t, 0, actual)
}
```

### Panic testing (good pattern)

```go
// Use helper
func callPanics(fn func()) bool {
    defer func() { recover() }()
    fn()

    return false
}

// _testcases.go
var mustPanicTestCase = coretestcases.CaseV1{
    Title:         "MustHaveItems panics when empty",
    ExpectedInput: "true",
}

// _test.go
func Test_MustHaveItems_PanicsWhenEmpty(t *testing.T) {
    tc := mustPanicTestCase

    // Arrange
    col := NewEmptyCollection()

    // Act
    panicked := callPanics(func() { col.MustHaveItems() })

    // Assert
    tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", panicked))
}
```

### CaseNilSafe for panic testing (best for nil receivers)

```go
{
    Title: "UnsafeMethod on nil panics",
    Func:  (*MyStruct).UnsafeMethod,
    Expected: results.ResultAny{
        Panicked: true,
    },
},
```

---

## Nil Receiver Testing

### When to use CaseNilSafe

- **Always** for testing nil-receiver safety on pointer receiver methods
- **Never** for functional tests with non-nil receivers

### Structure

```
Package_NilReceiver_testcases.go  — CaseNilSafe slice variable
NilReceiver_test.go               — shared test runner per package
```

### CaseNilSafe for non-generic types

```go
Func: (*MyStruct).MethodName,  // direct method expression
```

### CaseNilSafe for generic types (function literal wrapper)

```go
Func: func(c *Collection[string]) bool {
    return c.IsEmpty()
},
```

### CaseNilSafe for package functions (not methods)

```go
Func: func(_ *struct{}) bool {
    return mypackage.SomeFunc("input", nil) == nil
},
```

### CaseNilSafe Expected patterns

```go
// Method returns bool
Expected: results.ResultAny{Value: "false", Panicked: false}

// Method returns string
Expected: results.ResultAny{Value: "", Panicked: false}

// Method returns int
Expected: results.ResultAny{Value: "0", Panicked: false}

// Method returns error
Expected: results.ResultAny{Panicked: false, Error: results.ExpectAnyError}

// Method returns nil (pointer/interface)
Func: func(w *MyStruct) bool { return w.SomeMethod() == nil },
Expected: results.ResultAny{Value: "true", Panicked: false}

// Method panics on nil
Expected: results.ResultAny{Panicked: true}

// Void method (no return)
Expected: results.ResultAny{Panicked: false},
CompareFields: []string{"panicked", "returnCount"},
```

---

## Loop Tests vs Named Tests

### Named test (single case, non-loop)

```go
func Test_Hashset_Add_Duplicate(t *testing.T) {
    tc := hashsetAddDuplicateTestCase

    // Arrange
    ...

    // Act
    result := hs.Add(item)

    // Assert
    tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", result))  // ← *First variant, no magic 0
}
```

### Loop test (multiple homogeneous cases)

```go
func Test_Hashset_NilReceiver(t *testing.T) {
    for caseIndex, tc := range hashsetNilSafeTestCases {
        // Arrange (implicit — nil receiver)

        // Act & Assert
        tc.ShouldBeSafe(t, caseIndex)           // ← caseIndex, not 0
    }
}
```

---

## Complete Quality Checklist

When writing or reviewing a test, verify ALL of the following:

### Structure
- [ ] Test data in `_testcases.go`, logic in `_test.go`
- [ ] Package directory under `tests/integratedtests/{pkg}tests/`
- [ ] Package named `{pkg}tests`

### AAA Format
- [ ] `// Arrange` comment present
- [ ] `// Act` comment present  
- [ ] `// Assert` comment present
- [ ] Or `// Arrange (implicit — nil receiver)` + `// Act & Assert` for CaseNilSafe

### Assertions
- [ ] No raw `t.Error` / `t.Errorf` / `t.Fatal`
- [ ] No direct `convey.Convey` calls
- [ ] Uses `ShouldBeEqual`, `ShouldBeEqualMap`, `ShouldBeSafe`, or `errcore.AssertDiffOnMismatch`
- [ ] `*First` variants for non-loop tests

### Test Cases
- [ ] Named variables (not indexed slice access)
- [ ] Descriptive Title field
- [ ] `args.Map` multi-line for 2+ entries
- [ ] Native types in `ExpectedInput` (not pre-stringified)

### Coverage
- [ ] Positive path tested
- [ ] Negative path tested for each branch
- [ ] Nil receiver tested via CaseNilSafe (if pointer receiver)
- [ ] Boundary cases tested (empty, single, max)
- [ ] Error paths tested

### Logic
- [ ] No branching in test functions
- [ ] One scenario per test function
- [ ] `caseIndex` passed to all loop assertions

### Coding Guidelines
- [ ] Blank line before `return` statements
- [ ] Blank line before/after control flow blocks
- [ ] Multi-argument function calls formatted one-per-line
- [ ] No numbered variable names (`val1`, `val2`)
- [ ] Positive boolean variables (`isInvalid := !isValid`)

---

## Related Docs

- [Testing Guidelines](./16-testing-guidelines.md) — full API reference for assertion methods
- [Branch Coverage Strategy](./23-branch-coverage-strategy.md) — how to identify and cover all branches
- [Test Structures Reference](./25-test-structures-reference.md) — CaseV1, CaseNilSafe, args.*, results.* API
- [Coding Guidelines](./17-coding-guidelines.md) — formatting and naming conventions
- [CaseNilSafe Design](./designs/CaseNilSafe-design.md) — architecture and edge cases
