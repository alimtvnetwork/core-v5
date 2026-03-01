# errcore — Error Construction & Formatting

## Overview

Package `errcore` provides a comprehensive toolkit for creating, combining, formatting, and handling errors throughout the codebase. It offers structured error messages with variable context, expectation comparisons, stack trace enhancement, reference annotations, Gherkin-style test output, and batch error merging.

## Entry Points

| Variable | Type | Description |
|----------|------|-------------|
| `ShouldBe` | `shouldBe` | Assertion-style error messages (`"actual" should be "expected"`) |
| `Expected` | `expected` | Expectation comparison messages with type info and reflection support |
| `StackEnhance` | `stackTraceEnhance` | Wraps errors/messages with code stack traces |

### `ShouldBe` Methods

| Method | Signature | Description |
|--------|-----------|-------------|
| `StrEqMsg` | `(actual, expecting string) string` | Format `"actual should be expecting"` |
| `StrEqErr` | `(actual, expecting string) error` | Same as error |
| `AnyEqMsg` | `(actual, expecting any) string` | Any-typed should-be message |
| `AnyEqErr` | `(actual, expecting any) error` | Any-typed should-be error |
| `JsonEqMsg` | `(actual, expecting any) string` | JSON-serialized should-be message |
| `JsonEqErr` | `(actual, expecting any) error` | JSON-serialized should-be error |

### `Expected` Methods

| Method | Signature | Description |
|--------|-----------|-------------|
| `But` | `(title, expecting, actual any) error` | Simple "expected X but got Y" error |
| `ButFoundAsMsg` | `(title, expecting, actual any) string` | Same as message string |
| `ButFoundWithTypeAsMsg` | `(title, expecting, actual any) string` | With type information |
| `ButUsingType` | `(title, expecting, actual any) error` | With type information as error |
| `ReflectButFound` | `(expected, found reflect.Kind) error` | Reflection kind mismatch |
| `PrimitiveButFound` | `(found reflect.Kind) error` | Expected primitive but found other kind |
| `ValueHasNoElements` | `(typ reflect.Kind) error` | Generic value is nil or empty |

### `StackEnhance` Methods

| Method | Signature | Description |
|--------|-----------|-------------|
| `Error` | `(err error) error` | Wrap error with stack trace |
| `ErrorSkip` | `(skip int, err error) error` | Wrap with custom stack skip |
| `Msg` | `(msg string) string` | Enhance message with stack trace |
| `MsgSkip` | `(skip int, msg string) string` | Enhance with custom stack skip |
| `MsgToErrSkip` | `(skip int, msg string) error` | Message to error with stack |
| `FmtSkip` | `(skip int, format string, v ...any) error` | Formatted error with stack |
| `MsgErrorSkip` | `(skip int, msg string, err error) string` | Combine message + error with stack |
| `MsgErrorToErrSkip` | `(skip int, msg string, err error) error` | Same as error |

## Key Capabilities

### Error Creation

| Function | Description |
|----------|-------------|
| `ToError(msg)` | Create error from string |
| `ToExitError(msg, code)` | Create error with exit code |
| `MeaningFulError(msg, ref)` | Error with contextual reference |
| `MeaningFulErrorWithData(msg, data)` | Error with attached data |
| `PathMeaningfulError(path, msg)` | Error scoped to a file path |
| `EnumRangeNotMeet(...)` | Enum validation range error |
| `SourceDestination(src, dst)` | Source-destination context message |
| `SourceDestinationErr(src, dst)` | Source-destination as error |
| `SourceDestinationNoType(src, dst)` | Without type annotations |

### Error Combining

| Function | Description |
|----------|-------------|
| `MergeErrors(errs...)` | Merge multiple errors into one |
| `MergeErrorsToString(errs...)` | Merge and stringify |
| `MergeErrorsToStringDefault(errs...)` | Merge with default separator |
| `ManyErrorToSingle(errs)` | Reduce a slice to a single error |
| `ManyErrorToSingleDirect(errs)` | Direct reduction without nil filtering |
| `CombineWithMsgType(msg, err)` | Prepend a message to an error |
| `ConcatMessageWithErr(msg, err)` | Concatenate message with error |
| `SliceToError(errs)` | Convert error slice to single error |
| `SliceToErrorPtr(errs)` | Convert to error pointer |
| `SliceError(errs)` | Error slice operations |
| `SliceErrorDefault(errs)` | With default formatting |
| `SliceErrorsToStrings(errs)` | Convert error slice to string slice |

### Variable Formatting

| Function | Description |
|----------|-------------|
| `VarTwo(n1, v1, n2, v2)` | Format two named variables |
| `VarTwoNoType(n1, v1, n2, v2)` | Without type annotations |
| `VarThree(n1, v1, n2, v2, n3, v3)` | Format three named variables |
| `VarThreeNoType(n1, v1, n2, v2, n3, v3)` | Without type annotations |
| `VarMap(msg, map)` | Format a message with a variable map |
| `VarMapStrings(msg, map)` | String-typed variable map |
| `VarNameValues(msg, names, values)` | Structured name-value pairs |
| `VarNameValuesStrings(msg, names, values)` | String-typed name-value pairs |
| `VarNameValuesJoiner(msg, names, values, joiner)` | With custom joiner |
| `MessageNameValues(msg, names, values)` | Message with name-value context |
| `MessageVarMap(msg, map)` | Message with variable map |
| `MessageVarTwo(msg, n1, v1, n2, v2)` | Message with two variables |
| `MessageVarThree(msg, n1, v1, n2, v2, n3, v3)` | Message with three variables |
| `MessageWithRef(msg, ref)` | Annotate with reference |
| `MessageWithRefToError(msg, ref)` | Reference annotation as error |

### Expectation Messages

| Function | Description |
|----------|-------------|
| `Expecting(title, expect, actual)` | Full type-aware comparison message |
| `ExpectingSimple(title, expect, actual)` | Simplified comparison |
| `ExpectingSimpleNoType(title, expect, actual)` | Without type information |
| `ExpectingNotEqualSimpleNoType(title, left, right)` | Not-equal expectation |
| `ExpectingError(title, expect, actual)` | Returns comparison as error |
| `ExpectingRecord(...)` | Record-style expectation |
| `ExpectingFuture(...)` | Future-based expectation |
| `ExpectingErrorSimpleNoType(title, expect, actual)` | Simple no-type as error |

### Stack Traces & References

| Function | Description |
|----------|-------------|
| `StackTracesCompiled(...)` | Compile stack traces into formatted output |
| `ErrorWithCompiledTraceRef(...)` | Error with compiled traces and references |
| `ErrorWithCompiledTraceRefToError(...)` | Same as error type |
| `ErrorWithRef(msg, ref)` | Error with reference annotation |
| `ErrorWithRefToError(msg, ref)` | Reference annotation as error |
| `ErrorWithTracesRefToError(...)` | Traces + references as error |
| `Ref(name, value)` | Create a reference string |
| `RefToError(name, value)` | Reference as error |
| `getReferenceMessage(...)` | Internal reference message builder |

### Testing Support

| Function | Description |
|----------|-------------|
| `GetActualAndExpectProcessedMessage(...)` | Diff-style messages for tests |
| `GetActualAndExpectSortedMessage(...)` | Sorted diff messages |
| `GetSearchTermExpectationMessage(...)` | Search term matching messages |
| `GetSearchTermExpectationSimpleMessage(...)` | Simplified search messages |
| `GetSearchLineNumberExpectationMessage(...)` | Line-number search messages |
| `GherkinsString(...)` | Gherkin-format test messages |
| `GherkinsStringWithExpectation(...)` | Gherkin with expectation data |
| `PrintError(err)` | Debug print error |
| `PrintErrorWithTestIndex(i, err)` | Debug print with test case index |
| `StringLinesToQuoteLines(...)` | Quote each line of output |
| `StringLinesToQuoteLinesWithTabs(...)` | Quote with tab indentation |
| `StringLinesToQuoteLinesToSingle(...)` | Quote and join to single string |

### Error Handling (Panic on Error)

| Function | Description |
|----------|-------------|
| `HandleErr(err)` | Panic if error is non-nil |
| `HandleErrMessage(msg, err)` | Panic with message context |
| `HandleErrorGetter(getter)` | Handle from `errorGetter` interface |
| `HandleCompiledErrorGetter(getter)` | Handle from `compiledErrorGetter` interface |
| `HandleCompiledErrorWithTracesGetter(getter)` | Handle compiled error with traces |
| `HandleFullStringsWithTracesGetter(getter)` | Handle full strings with traces |
| `SimpleHandleErr(err)` | Simplified panic handler |
| `SimpleHandleErrMany(errs...)` | Panic on any non-nil error |
| `MustBeEmpty(errs)` | Panic if error slice is non-empty |
| `MeaningFulErrorHandle(msg, err)` | Handle with meaningful context |

### Miscellaneous

| Function | Description |
|----------|-------------|
| `FmtDebug(...)` | Debug formatted output |
| `FmtDebugIf(condition, ...)` | Conditional debug output |
| `MsgHeader(header, msg)` | Prepend header to message |
| `MsgHeaderIf(condition, header, msg)` | Conditional header prepend |
| `MsgHeaderPlusEnding(header, msg, ending)` | Header + ending wrapper |
| `PanicOnIndexOutOfRange(...)` | Panic for index bounds violations |
| `PanicRangeNotMeet(...)` | Panic for range validation failures |
| `RangeNotMeet(...)` | Range validation error (non-panic) |
| `ToString(err)` | Error to string conversion |
| `ToStringPtr(err)` | Error to string pointer |
| `ToValueString(v)` | Any value to string |
| `CompiledError(...)` | Compile error with details |
| `CompiledErrorString(...)` | Compile error to string |
| `CountStateChangeTracker(...)` | Track state changes with count |
| `ErrorToSplitLines(err)` | Split error message into lines |
| `ErrorToSplitNonEmptyLines(err)` | Split into non-empty lines |
| `RawErrCollection(...)` | Raw error collection operations |
| `RawErrorType(...)` | Raw error type utilities |
| `typesNamesString(...)` | Type names formatting |

## Internal Interfaces

| Interface | Description |
|-----------|-------------|
| `errorGetter` | `Error() error` |
| `compiledErrorGetter` | `CompiledError() error` |
| `compiledErrorWithTracesGetter` | `CompiledErrorWithStackTraces() error` |
| `fullStringWithTracesGetter` | `FullStringWithTraces() error` |
| `lengthGetter` | `Length() int` |

## File Organization

| File Pattern | Responsibility |
|-------------|---------------|
| `Expecting*.go` | Expectation message builders (7 files) |
| `Var*.go` | Variable formatting (8 files) |
| `Message*.go` | Message formatting with context (5 files) |
| `Handle*.go` | Error handlers — panic on error (6 files) |
| `Merge*.go` | Error combining and merging (3 files) |
| `Slice*.go` | Error slice operations (5 files) |
| `ErrorWith*.go` | Error wrapping with traces/refs (5 files) |
| `GherkinsString*.go` | Gherkin-style test output (2 files) |
| `GetActualAndExpect*.go` | Test diff messages (2 files) |
| `GetSearchTerm*.go` / `GetSearchLine*.go` | Search expectation messages (3 files) |
| `StringLinesToQuote*.go` | Line quoting utilities (3 files) |
| `SourceDestination*.go` | Source/destination context (3 files) |
| `shouldBe.go` | `shouldBe` type — assertion-style constructors |
| `expected.go` | `expected` type — expectation comparisons with reflection |
| `stackTraceEnhance.go` | `stackTraceEnhance` type — stack trace wrapping |
| `vars.go` | Package singletons: `ShouldBe`, `Expected`, `StackEnhance` |
| `consts.go` | Format string constants |
| `funcs.go` | Internal helper functions |
| `combine.go` | Internal combine logic |
| `all-interfaces.go` | Internal interface definitions |
| `ToString.go` / `ToError.go` / `ToExitError.go` | Type conversions |

## Related Docs

- [errcore Folder Spec](../spec/01-app/folders/06-errcore.md)
- [errcoreinf readme](../coreinterface/errcoreinf/readme.md)
- [Coding Guidelines](../spec/01-app/17-coding-guidelines.md)
