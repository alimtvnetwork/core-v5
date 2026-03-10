# coreinstruction

Composable instruction structs for describing operations, specifications, identifiers, and metadata. These building blocks are embedded into higher-level command and request types throughout the codebase.

## Architecture

```
coreinstruction/
├── Identification
│   ├── BaseIdentifier.go              # Id string + match/search helpers
│   ├── IdentifierWithIsGlobal.go      # Identifier + IsGlobal flag
│   ├── Identifiers.go                 # Collection of BaseIdentifier
│   ├── IdentifiersWithGlobals.go      # Collection of IdentifierWithIsGlobal
│   ├── ParentIdentifier.go            # Parent ID reference
│   └── LineIdentifier.go / BaseLineIdentifier.go
│
├── Specification
│   ├── Specification.go               # Full spec: Id + Display + Type + Tags + IsGlobal
│   ├── FlatSpecification.go           # Flat JSON-friendly version of Specification
│   ├── BaseSpecification.go           # Wrapper with HasSpec/IsEmptySpec helpers
│   └── RequestSpecification.go        # Spec + ContinueOnError + RunAll flags
│
├── Composite Types
│   ├── BaseIdDisplayType.go           # Identifier + Display + Type
│   ├── BaseFromTo.go / FromTo.go      # From/To string pair
│   ├── BaseSourceDestination.go / SourceDestination.go  # Source/Destination pair
│   ├── Rename.go / BaseIsRename.go    # Rename instruction
│   ├── DependsOn.go                   # Version + dependency name + IsLatest
│   ├── StringCompare.go               # String comparison with method selection
│   ├── StringSearch.go                # String search with compare method
│   ├── NameList.go                    # Named string list
│   ├── NameListCollection.go          # Collection of NameList
│   ├── NameRequests.go                # Named request type list
│   ├── NameRequestsCollection.go      # Collection of NameRequests
│   └── ById.go / BaseByIds.go        # ID-based lookup structures
│
├── Request IDs
│   ├── BaseRequestIds.go             # Collection of IdentifierWithIsGlobal for requests
│   └── BaseSpecPlusRequestIds.go     # Spec + RequestIds combined
│
├── Boolean Flags
│   ├── BaseIsGlobal.go               # IsGlobal flag
│   ├── BaseIsLatest.go               # IsLatest flag
│   ├── BaseIsRecursive.go            # IsRecursive flag
│   ├── BaseIsRunAll.go               # IsRunAll flag
│   ├── BaseIsSkipOnError.go          # IsSkipOnError flag
│   ├── BaseIsContinueOnError.go      # IsContinueOnError flag (JSON)
│   ├── BaseContinueOnError.go        # IsContinueOnError + IsExitOnError()
│   ├── BaseEnabler.go                # IsEnabled + Set/Toggle helpers
│   └── BaseIsSecure.go               # IsSecure flag
│
├── Metadata
│   ├── BaseDisplay.go                # Display name
│   ├── BaseType.go                   # Type name
│   ├── BaseTags.go                   # Tags slice + hashset-based matching
│   ├── BaseTypeDotFilter.go          # Dot-separated type filter
│   ├── BaseUsername.go               # Username field
│   ├── BaseModifyAs.go               # Modification identity
│   ├── DependencyName.go             # Dependency name field
│   └── SpecificVersion.go            # Version string field
```

## Key Types

### BaseIdentifier

| Method | Description |
|---|---|
| `IdString()` | Returns the ID |
| `IsId(id)` | Exact match |
| `IsIdCaseInsensitive(id)` | Case-insensitive match |
| `IsIdContains(sub)` | Substring check |
| `IsIdRegexMatches(regex)` | Regex match |
| `Clone()` | Deep copy |

### Specification

| Method | Description |
|---|---|
| `NewSpecification(id, display, type, tags, isGlobal)` | Full constructor |
| `NewSpecificationSimple(id, display, type)` | No tags, not global |
| `Clone()` | Deep copy |
| `FlatSpecification()` | Lazy-cached flat representation |

### BaseTags

| Method | Description |
|---|---|
| `HasAllTags(tags...)` | True if all tags present |
| `HasAnyTags(tags...)` | True if any tag present |
| `IsAnyTagMatchesRegex(regex)` | Regex match against tags |
| `TagsHashset()` | Lazy-cached hashset for fast lookups |

### StringCompare / StringSearch

| Method | Description |
|---|---|
| `IsMatch()` / `IsMatch(content)` | Compare using selected method (Equal, Contains, StartsWith, EndsWith, Regex) |
| `VerifyError()` | Returns error if match fails |

## Usage Examples

```go
import "github.com/alimtvnetwork/core/coreinstruction"

// Create a specification
spec := coreinstruction.NewSpecification(
    "user-service",
    "User Service",
    "microservice",
    []string{"auth", "core"},
    true,
)

// Identifier matching
id := coreinstruction.NewIdentifier("my-resource")
id.IsIdContains("resource")   // true
id.IsIdCaseInsensitive("MY-RESOURCE") // true

// String comparison
cmp := coreinstruction.NewStringCompareContains(true, "hello", "Hello World")
cmp.IsMatch() // true

// From/To pair
ft := coreinstruction.NewBaseFromTo("/old/path", "/new/path")
```

## Related Docs

- [corecomparator](../corecomparator/readme.md) — `BaseIsIgnoreCase` used by StringSearch
- [stringcompareas](../enums/stringcompareas/readme.md) — compare method variants
- [corestr](../coredata/corestr/readme.md) — hashset used by BaseTags
