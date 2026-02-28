# chmodhelper — File Permission Management

Package `chmodhelper` provides comprehensive file permission (chmod) management including parsing, applying, verifying, and comparing rwx permission strings. Supports Unix-style rwx notation, numeric chmod values, wildcard matching, and recursive directory operations.

## Architecture

```
chmodhelper/
├── vars.go                              # Package singletons: New, ChmodApply, ChmodVerify, SimpleFileWriter, TempDirDefault
├── consts.go                            # Constants: rwx values, default file/dir modes, format strings
├── newCreator.go                        # New.* — factory for RwxWrapper, SimpleFileReaderWriter, Attribute
├── newRwxWrapperCreator.go              # New.RwxWrapper.* — RwxWrapper construction
├── newSimpleFileReaderWriterCreator.go  # New.SimpleFileReaderWriter.* — file I/O construction
├── newAttributeCreator.go              # New.Attribute.* — attribute construction
├── chmodApplier.go                      # ChmodApply.* — apply chmod to paths
├── chmodVerifier.go                     # ChmodVerify.* — verify chmod expectations
├── RwxWrapper.go                        # RwxWrapper — owner/group/other rwx representation
├── SingleRwx.go                         # SingleRwx — single rwx triplet (e.g., "rwx", "r-x")
├── VarAttribute.go                      # VarAttribute — variable rwx attribute with wildcards
├── Variant.go                           # Variant — permission variant type
├── Attribute.go                         # Attribute — permission attribute
├── AttributeValue.go                    # AttributeValue — numeric permission value
├── AttrVariant.go                       # AttrVariant — attribute variant enum
├── RwxVariableWrapper.go                # RwxVariableWrapper — wildcard-aware rwx wrapper
├── RwxMatchingStatus.go                 # RwxMatchingStatus — match result type
├── RwxMismatchInfo.go                   # RwxMismatchInfo — mismatch detail
├── RwxInstructionExecutor.go            # RwxInstructionExecutor — single instruction executor
├── RwxInstructionExecutors.go           # RwxInstructionExecutors — batch executor
├── PathExistStat.go                     # PathExistStat — path existence + os.FileInfo
├── LocationFileInfoRwxWrapper.go        # LocationFileInfoRwxWrapper — path + FileInfo + RwxWrapper
├── FilteredPathFileInfoMap.go           # FilteredPathFileInfoMap — filtered path results
├── DirWithFiles.go                      # DirWithFiles — directory with file list
├── DirFilesWithContent.go               # DirFilesWithContent — directory with file contents
├── DirFilesRwxPermission.go             # DirFilesRwxPermission — directory with permissions
├── FileWithContent.go                   # FileWithContent — file path + content
├── SimpleFileReaderWriter.go            # SimpleFileReaderWriter — simple read/write operations
├── all-models.go                        # Model type definitions
├── Parse*.go                            # Parsing functions: rwx strings → executors, file modes, var wrappers
├── Get*.go                              # Getter functions: existing chmod, paths, stats, recursive paths
├── Is*.go                               # Check functions: IsChmod, IsDirectory, IsPathExists, etc.
├── Create*.go                           # Creation functions: directories, files with permissions
├── Expand*.go                           # Expansion: rwx chars, wildcard merging
├── Merge*.go                            # Merging: wildcard + fixed rwx
├── fwChmodApplier.go                    # Internal: file-writer chmod applier
├── fwChmodVerifier.go                   # Internal: file-writer chmod verifier
├── fileReader.go / fileWriter.go        # Internal: file I/O
├── fileBytesWriter.go / fileStringWriter.go # Internal: typed file writers
├── anyItemWriter.go                     # Internal: any-item writer
├── simpleFileWriter.go                  # Internal: simple file writer singleton
├── dirCreator.go                        # Internal: directory creator
├── tempDirGetter.go                     # Internal: temp directory getter
├── errorCreator.go                      # Internal: error construction
├── pathErrorMessage.go                  # Internal: path error messages
├── removeDirIf.go                       # Internal: conditional directory removal
├── vars-errors.go                       # Internal: error variables
├── chmodclasstype/                      # Chmod class type enum (Owner, Group, Other)
│   ├── Variant.go                       # Variant enum type
│   └── vars.go                          # Enum instances
└── chmodins/                            # Chmod instruction parsing
    ├── RwxInstruction.go                # Single rwx instruction
    ├── RwxOwnerGroupOther.go            # Owner/group/other instruction set
    ├── BaseRwxInstructions.go           # Base instruction collection
    ├── Condition.go                     # Instruction condition
    ├── Parse*.go                        # JSON-based instruction parsing
    ├── Expand*.go                       # Rwx string expansion
    ├── Fix*.go                          # Rwx string wildcard fixing
    ├── Get*.go                          # Rwx length validation
    ├── expandCharsRwx.go                # Char-level rwx expansion
    └── consts.go                        # Instruction constants
```

## Entry Points

| Namespace | Description |
|-----------|-------------|
| `chmodhelper.New.*` | Create `RwxWrapper`, `SimpleFileReaderWriter`, `Attribute` |
| `chmodhelper.ChmodApply.*` | Apply chmod permissions to file paths |
| `chmodhelper.ChmodVerify.*` | Verify chmod expectations against actual permissions |
| `chmodhelper.SimpleFileWriter.*` | Simple file write operations |
| `chmodhelper.TempDirDefault` | OS temp directory path |
| `chmodhelper.TempDirGetter.*` | Temp directory resolution |

## Usage

### Checking Permissions

```go
import "gitlab.com/auk-go/core/chmodhelper"

// Check if path exists
exists := chmodhelper.IsPathExists("/tmp/myfile")

// Get existing chmod
rwxWrapper, err := chmodhelper.GetExistingChmodRwxWrapperPtr("/tmp/myfile")

// Get human-readable chmod string
friendly := chmodhelper.FileModeFriendlyString(fileMode)
// e.g., `{chmod : "0755 (-rwxr-xr-x)"}`
```

### Applying Permissions

```go
// Create directory with specific permissions
chmodhelper.CreateDefaultPathsMust([]string{"/tmp/mydir"})
```

### Recursive Operations

```go
// Get all paths recursively
paths := chmodhelper.GetRecursivePaths("/tmp/mydir")
```

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
