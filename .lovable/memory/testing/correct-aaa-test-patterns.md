# Memory: testing/correct-aaa-test-patterns
Updated: 2026-03-22

## Correct AAA Test Pattern — Learned from `anycmptests` and `spec/testing-guidelines`

### Core Architecture: Separate Data from Logic

Every test package has TWO file types:
- `*_testcases.go` — Named `CaseV1` variables with `ArrangeInput`/`ExpectedInput`. Never contains assertions.
- `*_test.go` — `Test_*` functions with AAA structure. Never contains expected values.

### Example: Correct Test Case File (`_testcases.go`)

```go
package mypkgtests

import (
    "github.com/alimtvnetwork/core/coretests/args"
    "github.com/alimtvnetwork/core/coretests/coretestcases"
)

// Branch: valid input → success
var serializeApplyValidTestCases = []coretestcases.CaseV1{
    {
        Title: "Serialize Apply returns no error -- valid string input",
        ArrangeInput: args.Map{
            "input": "hello",
        },
        ExpectedInput: args.Map{
            "hasError": false,
            "hasBytes": true,
        },
    },
    {
        Title: "Serialize Apply returns error -- unmarshalable channel input",
        ArrangeInput: args.Map{
            "useChannel": true,
        },
        ExpectedInput: args.Map{
            "hasError": true,
            "hasBytes": false,
        },
    },
}
```

### Example: Correct Test Runner File (`_test.go`)

```go
package mypkgtests

import (
    "testing"

    "github.com/alimtvnetwork/core/coredata/corejson"
    "github.com/alimtvnetwork/core/coretests/args"
)

func Test_Serialize_Apply_Verification(t *testing.T) {
    for caseIndex, tc := range serializeApplyValidTestCases {
        // Arrange
        input := tc.ArrangeInput.(args.Map)
        useChannel := input.GetAsBoolDefault("useChannel", false)

        // Act
        var r corejson.Result
        if useChannel {
            r = *corejson.Serialize.Apply(make(chan int))
        } else {
            inputStr, _ := input.GetAsString("input")
            r = *corejson.Serialize.Apply(inputStr)
        }
        actual := args.Map{
            "hasError": r.HasError(),
            "hasBytes": r.Length() > 0,
        }

        // Assert
        tc.ShouldBeEqualMap(t, caseIndex, actual)
    }
}
```

### Key Rules (from spec/testing-guidelines)

1. **AAA comments mandatory**: `// Arrange`, `// Act`, `// Assert` — even if section is empty
2. **Title format**: `"{Function} returns {Result} -- {Input Context}"`
3. **Native types in args.Map**: Use `true`/`false`/`42` — never `"true"`/`"42"`
4. **No branching in test body** (exception: choosing input based on ArrangeInput flags is OK)
5. **One test function per scenario** OR loop over homogeneous cases
6. **No raw `t.Error`/`t.Fatal`** — use `ShouldBeEqualMap`, `ShouldBeEqual`, etc.
7. **Named variables for single test cases**: `var myFeatureNilCase = coretestcases.CaseV1{...}`
8. **Named slices for grouped cases**: `var myFeatureBranchTestCases = []coretestcases.CaseV1{...}`
9. **params.go** for key constants (optional but recommended for larger packages)
10. **Comments above case slices** documenting which branch they cover

### Anti-Patterns (BANNED)

- ❌ `expected.ShouldBeEqual(t, 0, "title", actual)` on raw `args.Map` — use `CaseV1.ShouldBeEqualMap`
- ❌ Expected values inside `_test.go` — move to `_testcases.go`
- ❌ Inline `args.Map` construction in test functions for expected values
- ❌ `defer recover()` without using `CaseNilSafe` or the `callPanics` helper pattern
- ❌ `fmt.Sprintf("%v", boolVal)` — use native bool in args.Map instead
- ❌ Stringified booleans/numbers in args.Map
- ❌ Multiple unrelated scenarios in one test function

### File Structure Per Package

```
tests/integratedtests/{pkg}tests/
├── params.go                          # Key constants (optional)
├── Feature_testcases.go               # CaseV1 variables (data)
├── Feature_test.go                    # Test_ functions (logic)
├── Coverage{N}_{Label}_testcases.go   # Coverage-specific data
├── Coverage{N}_{Label}_test.go        # Coverage-specific logic
└── helpers.go                         # Shared helpers (callPanics, etc.)
```
