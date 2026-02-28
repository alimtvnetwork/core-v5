# Testing Patterns

## Dominant Testing Style

The repository uses a **table-driven test pattern** with an AAA (Arrange-Act-Assert) structure and the **goconvey** assertion library.

### Framework & Libraries

| Tool | Purpose |
|------|---------|
| `testing` (stdlib) | Test runner |
| `github.com/smartystreets/goconvey` | BDD-style assertions |
| `github.com/smarty/assertions` | Assertion functions |
| `coretests.GetAssert` | Custom assertion wrapper |
| `coretests/args.Map` | Typed test input map |
| `coretests/coretestcases.CaseV1` | Test case structure with ArrangeInput + expected |

### Test File Organization

- Integration tests live in `tests/integratedtests/`.
- Per-package test directories: `tests/integratedtests/{pkg}tests/`.
- Test case data files: `*_testcases.go` (separate from test logic `*_test.go`).

## Template Test Structure

```go
package sometests

import (
    "testing"
    "gitlab.com/auk-go/core/coretests"
    "gitlab.com/auk-go/core/coretests/args"
)

// Test cases defined in a separate _testcases.go file
var myTestCases = []coretestcases.CaseV1{
    {
        ArrangeInput: args.Map{
            "when":   "given valid input",
            "actual": "hello",
            "expect": "HELLO",
        },
        ExpectedLines: []string{"HELLO"},
    },
}

func Test_MyFunction(t *testing.T) {
    for caseIndex, testCase := range myTestCases {
        // Arrange
        input := testCase.ArrangeInput.(args.Map)

        // Act
        result := MyFunction(input.Actual())

        // Assert
        testCase.ShouldBeEqual(t, caseIndex, result)
    }
}
```

## Best Patterns Observed

1. **Separation of test data and test logic** — `_testcases.go` files keep data separate.
2. **Consistent AAA structure** — every test follows Arrange-Act-Assert.
3. **Index-based case tracking** — `caseIndex` enables precise failure identification.
4. **Formatted output** — `GetAssert.Quick` provides readable failure messages.

## Anti-Patterns to Avoid

1. **Inline test data in test functions** — always use separate testcases files.
2. **Skipping the Arrange comment** — always label AAA sections.
3. **Ignoring caseIndex** — always pass it for debugging.
4. **Direct `t.Fatal` without context** — use `ShouldBeEqual` or `GetAssert` for rich output.

## Coverage Expectations

No formal coverage requirements are documented. Recommended minimum: critical packages (`chmodhelper`, `errcore`, `coredata/corestr`, `converters`) should have ≥80% coverage.

## Related Docs

- [coretests folder spec](./folders/07-coretests.md)
- [Repo Overview](./00-repo-overview.md)
