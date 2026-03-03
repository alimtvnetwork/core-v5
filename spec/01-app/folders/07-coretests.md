# coretests

## Folder Purpose

Testing utilities, assertion wrappers, and test-case structures for the auk-go ecosystem. Provides a consistent AAA (Arrange-Act-Assert) pattern with formatted output.

## Responsibilities

1. Provide `GetAssert` with comparison and formatting helpers.
2. Define `SimpleTestCase` and `BaseTestCase` wrappers.
3. Support Gherkins-style test output.
4. Provide `args/` sub-package for function wrapping and argument handling.
5. Provide `coretestcases/` for test case definitions.

## Key Files

| File | Purpose |
|------|---------|
| `vars.go` | `GetAssert` singleton |
| `ShouldAsserter.go` | Assertion helpers using goconvey |
| `SimpleTestCase.go` | Test case with input/expected |
| `BaseTestCaseWrapper.go` | Base wrapper for test cases |
| `args/` | `FuncWrap`, `Map` — function reflection and argument handling |
| `coretestcases/` | `CaseV1` test case structure, `GenericGherkins` (proposed) |

## Testing Pattern (Dominant Style)

```go
func Test_Something(t *testing.T) {
    for caseIndex, testCase := range testCases {
        // Arrange
        input := testCase.ArrangeInput.(args.Map)
        
        // Act
        result := someFunc(input.When(), input.Actual())
        
        // Assert
        testCase.ShouldBeEqual(t, caseIndex, result...)
    }
}
```

## Related Docs

- [Testing Patterns](../13-testing-patterns.md)
