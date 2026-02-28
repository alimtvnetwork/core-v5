# errcoreinf — Error Core Interface Contracts

Package `errcoreinf` defines all interface contracts for the structured error handling subsystem. It provides composable interfaces for error wrapping, error type classification, reference tracking, stack traces, assertion chains (`ShouldBe`), completion semantics, and error collection aggregation.

## Architecture

```
errcoreinf/
├── all-errors-related.go     # Core error interfaces: BasicErrWrapper, BaseErrorOrCollectionWrapper,
│                              #   BaseErrorTyper, BasicErrorTyper, Referencer, ReferenceCollectionDefiner,
│                              #   StackTracer, GenericErrorCompiler, ErrorCompleter, etc.
└── all-should-related.go     # Test assertion interfaces: ShouldBeMessager, ShouldBeChainer,
                               #   ShouldBeChainCollectionDefiner, stringShouldBer, stringsShouldBer,
                               #   errorShouldBer, checkerShouldBer
```

## Key Interfaces

### Error Wrapping

| Interface | Description |
|-----------|-------------|
| `BaseErrorOrCollectionWrapper` | Root error/collection wrapper — JSON, stack traces, collect, reflect |
| `BasicErrWrapper` | Full error wrapper — type, references, merge, clone |
| `BaseErrorWrapperCollectionDefiner` | Error collection — add errors, compile, list lines |
| `BaseRawErrCollectionDefiner` | Raw error collection base |

### Error Types

| Interface | Description |
|-----------|-------------|
| `BaseErrorTyper` | Error type with code, name, category, JSON, serialization |
| `BasicErrorTyper` | BaseErrorTyper + BasicEnumer + ErrTypeDetailDefiner |
| `BaseErrorTypeGetter` | Getter for BaseErrorTyper |

### References

| Interface | Description |
|-----------|-------------|
| `Referencer` | Variable name + value + string/JSON/stringer representation |
| `ReferenceCollectionDefiner` | Collection of Referencer with Add/Clone/Compile |

### Stack Traces

| Interface | Description |
|-----------|-------------|
| `StackTracer` | StackTraces, NewStackTraces, JSON stack trace results |

### Assertion Chain (ShouldBe)

| Interface | Description |
|-----------|-------------|
| `ShouldBeChainer` | Fluent assertion builder — 50+ typed assertion methods |
| `ShouldBeMessager` | Single assertion: Title + Actual + Expected |
| `ShouldBeChainCollectionDefiner` | Collection of ShouldBeMessager |
| `stringShouldBer` | String-specific assertions: ShouldBe, Contains, StartsWith, EndsWith |
| `stringsShouldBer` | String slice assertions: ShouldBe, Length, Distinct, RegardlessOrder |
| `errorShouldBer` | Error assertions: ShouldBeDefined, ShouldBeEmpty |
| `checkerShouldBer` | Checker assertions: IsEmpty, IsEnabled, IsDisabled |

### Completion

| Interface | Description |
|-----------|-------------|
| `ErrorCompleter` | CompleteSuccess/Failure, CompleteUsingErr, CompleteJson |
| `GenericErrorCompiler` | CompiledMessage, JsonString, CompileError, Length, IsEmpty |
| `MustBeEmptier` | MustBeSuccess, MustBeEmpty, HandleError |

### Logging

| Interface | Description |
|-----------|-------------|
| `VoidLogger` | Log without return value |
| `CompiledVoidLogger` | Compiled void logger |
| `ErrWrapperLogger` | Error wrapper logger |

## Related Docs

- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
