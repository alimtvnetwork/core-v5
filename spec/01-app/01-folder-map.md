# Folder Map

## Top-Level Folders

| # | Folder | Purpose | Primary Entrypoints |
|---|--------|---------|-------------------|
| 1 | `chmodhelper/` | File permission parsing, verification, chmod application | `RwxWrapper`, `SimpleFileReaderWriter`, `ChmodApply`, `ChmodVerify` |
| 2 | `cmd/` | CLI entry-points for running, testing, and demos | `cmd/main/main.go` |
| 3 | ~~`codegen/`~~ | **Removed** (v1.6.0) — was unit-test code generation | — |
| 4 | `codestack/` | Runtime call-stack capture and trace formatting | `Trace`, `TraceCollection` |
| 5 | `conditional/` | Ternary-style helpers for all primitive types | `Bool()`, `Int()`, `String()`, `Interface()` |
| 6 | `constants/` | Global constants, OS line separators, capacity defaults | `constants.go`, `line_*.go` |
| 7 | `converters/` | Type conversion utilities (strings, bytes, pointers, maps) | `StringTo`, `BytesTo`, `StringsTo` |
| 8 | `coreappend/` | Append helper utilities | — |
| 9 | `anycmp/` | Any-type comparison helpers | — |
| 10 | `corecmp/` | Core comparison functions | — |
| 11 | `corecomparator/` | Comparator abstractions | — |
| 12 | `corecsv/` | CSV formatting and string building | `DefaultCsv`, `StringsToCsvString` |
| 13 | `coredata/` | Rich data structures umbrella | See sub-packages below |
| 14 | `coredata/corestr/` | String collections: `Collection`, `Hashmap`, `Hashset`, `LinkedList`, `SimpleSlice` | `New.SimpleSlice`, `New.Collection` |
| 15 | `coredata/corejson/` | JSON serialize/deserialize logic | `Serialize.*`, `Deserialize.*` |
| 16 | `coredata/coredynamic/` | Dynamic/reflection-based data manipulation | `AnyCollection`, `Dynamic`, `KeyVal` |
| 17 | `coredata/coreonce/` | Compute-once value holders | — |
| 18 | `coredata/corepayload/` | Enhanced payload structures | — |
| 19 | `coredata/corerange/` | Range data types | — |
| 20 | `coredata/coreapi/` | API-related data models | — |
| 21 | `corefuncs/` | Function-level utilities | — |
| 22 | `coreimpl/` | Concrete implementations (e.g. `enumimpl/`) | — |
| 23 | `coreindexes/` | Index-to-name mapping (First, Second, Third…) | `NameByIndex()` |
| 24 | `coreinstruction/` | Instruction abstractions | — |
| 25 | `coreinterface/` | Canonical interface contracts for the ecosystem | All `*Getter`, `*Checker`, `*Binder` interfaces |
| 26 | `coremath/` | Min/Max for all numeric types | `MaxByte`, `MinByte` |
| 27 | `coresort/` | String/int sorting | `Quick`, `QuickDsc` |
| 28 | `coretaskinfo/` | Task metadata structures | — |
| 29 | `coretests/` | Test helpers, assertion wrappers, Gherkins-style test cases | `GetAssert`, `ShouldAsserter`, `SimpleTestCase` |
| 30 | `coreunique/` | Unique value generators | — |
| 31 | `coreutils/` | String utilities and template replacers | `stringutil.ReplaceTemplate` |
| 32 | `corevalidator/` | Line, slice, text, and range validators | `LineValidator`, `SliceValidator`, `TextValidator` |
| 33 | `coreversion/` | Semantic versioning data type | `Version` |
| 34 | `defaultcapacity/` | Default capacity constants | — |
| 35 | `defaulterr/` | Default error factory functions | — |
| 36 | `dtformats/` | Date-time format strings | — |
| 37 | `enums/` | Enum helpers (string compare areas, version indexes) | — |
| 38 | `errcore/` | Rich error building: stack traces, merge, formatting, expectations | `Expected`, `ShouldBe`, `StackEnhance` |
| 39 | `extensionsconst/` | File-extension constants | — |
| 40 | `filemode/` | File-mode constants | — |
| 41 | `internal/` | Private implementation helpers | `reflectinternal`, `convertinteranl`, `pathinternal` |
| 42 | `isany/` | Nil/zero/type checks on `interface{}` | `Null()`, `Defined()`, `Zero()` |
| 43 | `iserror/` | Error-defined checks | `Defined()` |
| 44 | `issetter/` | 4-valued boolean (Uninitialized/True/False/Wildcard) | `Value`, `New`, `NewBool` |
| 45 | `keymk/` | Key-maker utilities | — |
| 46 | `mutexbykey/` | Per-key mutex locking | — |
| 47 | `namevalue/` | Name-value pair structures | — |
| 48 | `osconsts/` | OS-specific constants | — |
| 49 | `ostype/` | OS type detection | — |
| 50 | `pagingutil/` | Paging calculation utilities | — |
| 51 | `refeflectcore/` | Reflection core helpers (note: typo in folder name) | — |
| 52 | `regconsts/` | Regex constant strings | — |
| 53 | `regexnew/` | Lazy-compiled regex with lock support | `LazyRegex`, `Create`, `CreateLock` |
| 54 | `reqtype/` | Request-type helpers | — |
| 55 | `scripts/` | Deployment and Docker scripts | — |
| 56 | `simplewrap/` | String wrapping (quotes, brackets, curly) | `WithDoubleQuote`, `WithCurly`, `WithBrackets` |
| 57 | `testconsts/` | Test-only constants | — |
| 58 | `tests/` | Integration tests and test wrappers | `tests/integratedtests/` |
| 59 | `typesconv/` | Additional type conversion utilities | — |
| 60 | `bytetype/` | Byte-type utilities | — |
| 61 | `cmdconsts/` | Command constants | — |
| 62 | `configs/` | Configuration files | — |

## Sub-Package Directories of Note

| Parent | Sub-Package | Purpose |
|--------|------------|---------|
| `coreinterface/` | `baseactioninf/` | Executor, Initializer, Planner interfaces |
| `coreinterface/` | `entityinf/` | Entity definer interfaces |
| `coreinterface/` | `enuminf/` | Enum interfaces |
| `coreinterface/` | `errcoreinf/` | Error-core interfaces |
| `coreinterface/` | `loggerinf/` | Logger interfaces |
| `coreinterface/` | `serializerinf/` | Serializer interfaces |
| `chmodhelper/` | `chmodclasstype/` | Chmod class type enums |
| `chmodhelper/` | `chmodins/` | Chmod instruction types |
| `internal/` | `reflectinternal/` | Reflection helpers |
| `internal/` | `convertinteranl/` | Conversion internals (note: typo in name) |
| `internal/` | `pathinternal/` | Path manipulation internals |
| `internal/` | `csvinternal/` | CSV internals |
| `internal/` | `jsoninternal/` | JSON internals |

## Related Docs

- [Repo Overview](./00-repo-overview.md)
- [Per-folder specs](./folders/)
