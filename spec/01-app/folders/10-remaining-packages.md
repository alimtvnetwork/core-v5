# Remaining Packages (Grouped)

## Comparison & Sorting

| Package | Purpose |
|---------|---------|
| `anycmp/` | Any-type comparison |
| `corecmp/` | Core comparison functions |
| `corecomparator/` | Comparator abstractions |
| `coresort/` | Quick sort and descending sort for strings/ints |

## Type Utilities

| Package | Purpose |
|---------|---------|
| `converters/` | Type converters: `StringTo`, `BytesTo`, `StringsTo`, `PointerStringsToStrings` |
| `typesconv/` | Additional type conversions |
| `bytetype/` | Byte-type utilities |
| `isany/` | Nil/zero/defined/type checks on `interface{}` |
| `iserror/` | `Defined()` check for errors |
| `issetter/` | 4-valued boolean: Uninitialized, True, False, Wildcard |

## String Utilities

| Package | Purpose |
|---------|---------|
| `coreutils/stringutil/` | Extensive string helpers: IsBlank, IsContains, SafeSubstring, ReplaceTemplate, etc. |
| `simplewrap/` | String wrapping with quotes, brackets, curly braces |
| `corecsv/` | CSV string building and formatting |

## Constants

| Package | Purpose |
|---------|---------|
| `constants/` | Global constants, OS-specific line separators, capacity defaults |
| `cmdconsts/` | Command-related constants |
| `defaultcapacity/` | Default slice/map capacity constants |
| `defaulterr/` | Default error factories |
| `dtformats/` | Date-time format strings |
| `extensionsconst/` | File extension constants (`.json`, `.csv`, etc.) |
| `filemode/` | File mode constants |
| `osconsts/` | OS-specific constants |
| `regconsts/` | Regex pattern constants |
| `testconsts/` | Test-only constants |

## Regex

| Package | Purpose |
|---------|---------|
| `regexnew/` | Lazy-compiled regex with lock/non-lock variants |

## Concurrency

| Package | Purpose |
|---------|---------|
| `mutexbykey/` | Per-key mutex for fine-grained locking |

## Misc

| Package | Purpose |
|---------|---------|
| `codestack/` | Runtime call-stack capture and trace formatting. Traces from Go standard library packages (e.g., `runtime`, `net`, `fmt`, `sync`, `reflect`) are automatically flagged as `IsSkippable` and filtered out at collection-creation time to reduce noise. The skippable packages are defined as a `map[string]bool` in `skippablePrefixes.go` for O(1) lookup. |
| `coreappend/` | Append utilities |
| `corefuncs/` | Function utilities |
| `coreindexes/` | Index-to-name mapping (First, Second, Third…) |
| `coreinstruction/` | Instruction abstractions |
| `coremath/` | Min/Max for all numeric types |
| `coretaskinfo/` | Task metadata |
| `coreunique/` | Unique value generators |
| `corevalidator/` | Line, slice, text, range validators |
| `coreversion/` | Semantic versioning type |
| `keymk/` | Key-maker utilities |
| `namevalue/` | Name-value pairs |
| `ostype/` | OS type detection |
| `pagingutil/` | Paging calculations |
| `refeflectcore/` | Reflection helpers (note: **typo** — should be `reflectcore`) |
| `reqtype/` | Request-type helpers |
| `enums/` | Enum helpers (`stringcompareas`, `versionindexes`) |
| `coreimpl/enumimpl/` | Concrete enum implementation, [`DifferChecker`](/coreimpl/enumimpl/readme.md#differchecker-interface) strategy interface |

## Related Docs

- [Folder Map](../01-folder-map.md)
