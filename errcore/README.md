# errcore — Error Construction & Formatting

## Overview

The `errcore` package provides a comprehensive toolkit for creating, combining, formatting, and handling errors throughout the codebase. It offers structured error messages with variable context, expectation comparisons, stack trace enhancement, reference annotations, and batch error merging.

## Entry Points

| Variable | Type | Description |
|----------|------|-------------|
| `ShouldBe` | `shouldBe` | Assertion-style error messages (`"actual" should be "expected"`) |
| `Expected` | `expected` | Expectation comparison messages with type info |
| `StackEnhance` | `stackTraceEnhance` | Wraps errors with code stack traces |

## Key Capabilities

### Error Creation

- **`ToError(msg)`** / **`ToExitError(msg, code)`** — Create errors from strings
- **`MeaningFulError(msg, ref)`** — Error with contextual reference
- **`MeaningFulErrorWithData(msg, data)`** — Error with attached data
- **`PathMeaningfulError(path, msg)`** — Error scoped to a file path
- **`EnumRangeNotMeet(...)`** — Enum validation range error

### Error Combining

- **`MergeErrors(errs...)`** — Merge multiple errors into one
- **`MergeErrorsToString(errs...)`** — Merge and stringify
- **`ManyErrorToSingle(errs)`** — Reduce a slice to a single error
- **`SliceToError(errs)`** / **`SliceToErrorPtr(errs)`** — Convert error slices
- **`CombineWithMsgType(msg, err)`** — Prepend a message to an error

### Variable Formatting

- **`VarTwo(n1, v1, n2, v2)`** — Format two named variables
- **`VarThree(n1, v1, n2, v2, n3, v3)`** — Format three named variables
- **`VarMap(msg, map)`** — Format a message with a variable map
- **`MessageNameValues(msg, names, values)`** — Structured name-value pairs
- **`MessageWithRef(msg, ref)`** / **`MessageWithRefToError(...)`** — Annotate with references

### Expectation Messages

- **`Expecting(title, expect, actual)`** — Full type-aware comparison
- **`ExpectingSimple(...)`** / **`ExpectingSimpleNoType(...)`** — Simplified comparisons
- **`ExpectingError(...)`** — Returns comparison as an error
- **`ExpectingRecord(...)`** — Record-style expectation

### Stack Traces

- **`StackEnhance`** — Enhance errors with code stacks
- **`StackTracesCompiled(...)`** — Compile stack traces into formatted output
- **`ErrorWithCompiledTraceRef(...)`** — Error with compiled traces and references

### Testing Support

- **`GetActualAndExpectProcessedMessage(...)`** — Diff-style messages for tests
- **`GetSearchTermExpectationMessage(...)`** — Search term matching messages
- **`GherkinsString(...)`** — Gherkin-format test messages
- **`PrintError(err)`** / **`PrintErrorWithTestIndex(i, err)`** — Debug printing

### Error Handling

- **`HandleErr(err)`** — Panic on error
- **`HandleErrMessage(msg, err)`** — Panic with message context
- **`HandleErrorGetter(getter)`** — Handle from interface
- **`SimpleHandleErr(err)`** / **`SimpleHandleErrMany(errs...)`** — Simplified panic handlers
- **`MustBeEmpty(errs)`** — Panic if error slice is non-empty

## File Organization

| File Pattern | Responsibility |
|-------------|---------------|
| `Expecting*.go` | Expectation message builders |
| `Var*.go`, `Message*.go` | Variable and message formatting |
| `Handle*.go` | Error handlers (panic on error) |
| `Merge*.go`, `Slice*.go` | Error combining and conversion |
| `ErrorWith*.go` | Error wrapping with traces/refs |
| `ToString.go`, `ToError.go` | Type conversions |
| `shouldBe.go`, `expected.go` | Assertion-style constructors |

## Contributors

## Issues for Future Reference
