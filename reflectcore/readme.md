# reflectcore — Reflection Utilities Facade

## Overview

Package `reflectcore` is a public facade that re-exports reflection utilities from `internal/reflectinternal` and `internal/convertinternal`. It provides type inspection, conversion, looping over reflected values, stack trace access, and structured reflection models via the `reflectmodel` sub-package.

## Architecture

```
reflectcore/
├── vars.go                          # Public variable exports (facade re-exports)
├── reflectmodel/                    # Reflection data models
│   ├── ReflectValue.go              # ReflectValue — type name, fields, methods, raw data
│   ├── ReflectValueKind.go          # ReflectValueKind — validated reflect.Value with kind and error
│   ├── FieldProcessor.go            # FieldProcessor — struct field metadata for iteration
│   ├── MethodProcessor.go           # MethodProcessor — method metadata for iteration
│   ├── isNull.go                    # Internal null check via reflection
│   └── utils.go                     # rvUtils — type verification, arg conversion, formatting
└── readme.md
```

## Exported Variables (Public API)

| Variable | Source | Description |
|----------|--------|-------------|
| `Converter` | `reflectinternal` | Type conversion utilities |
| `Utils` | `reflectinternal` | General reflection helpers |
| `Looper` | `reflectinternal` | Reflection-based iteration |
| `CodeStack` | `reflectinternal` | Code stack trace access |
| `GetFunc` | `reflectinternal` | Function reflection |
| `Is` | `reflectinternal` | Type/kind checking predicates |
| `TypeName` | `reflectinternal` | Single type name extraction |
| `TypeNames` | `reflectinternal` | Multiple type names |
| `TypeNamesString` | `reflectinternal` | Type names as string |
| `TypeNamesReferenceString` | `reflectinternal` | Type names with reference format |
| `ReflectType` | `reflectinternal` | `reflect.Type` utilities |
| `ReflectGetter` | `reflectinternal` | Reflection value getter |
| `ReflectGetterUsingReflectValue` | `reflectinternal` | Getter from `reflect.Value` |
| `SliceConverter` | `reflectinternal` | Slice reflection and conversion |
| `MapConverter` | `reflectinternal` | Map reflection and conversion |

> **Note**: Two unexported variables (`indexToPositionFunc`, `prependWithSpacesFunc`) delegate to `convertinternal.Util.String` for internal use only.

## reflectmodel Sub-Package

### ReflectValue

```go
type ReflectValue struct {
    TypeName     string
    FieldsNames  []string
    MethodsNames []string
    RawData      any
}
```

### ReflectValueKind

Validated reflection result with error handling.

| Method | Description |
|--------|-------------|
| `IsInvalid()` | Not valid or has error |
| `HasError()` / `IsEmptyError()` | Error state checks |
| `ActualInstance()` | Extract `any` from reflect.Value |
| `TypeName()` | String representation of type |
| `PkgPath()` | Package path of the type |
| `PointerRv()` | Reflect.Value as pointer |
| `PointerInterface()` | Pointer interface value |

### FieldProcessor

```go
type FieldProcessor struct {
    Name      string
    Index     int
    Field     reflect.StructField
    FieldType reflect.Type
}
```

| Method | Description |
|--------|-------------|
| `IsFieldType(reflect.Type)` | Type match check |
| `IsFieldKind(reflect.Kind)` | Kind match check |

### MethodProcessor

```go
type MethodProcessor struct {
    Name   string
    Index  int
    Method reflect.Method
}
```

### rvUtils

| Method | Description |
|--------|-------------|
| `ArgsToReflectValues([]any)` | Convert args to `[]reflect.Value` |
| `ReflectValuesToInterfaces([]reflect.Value)` | Convert back to `[]any` |
| `ReflectValueToAnyValue(reflect.Value)` | Single value conversion |
| `IsNull(any)` | Reflection-based nil check |
| `InterfacesToTypesNamesWithValues([]any)` | Debug-formatted type+value strings |
| `InterfacesToTypes([]any)` | Extract `[]reflect.Type` |
| `VerifyReflectTypes(name, expected, given)` | Type verification with error messages |
| `IsReflectTypeMatch(expected, given)` | Single type comparison |

## Dependencies

| Package | Usage |
|---------|-------|
| `internal/reflectinternal` | Core reflection implementations |
| `internal/convertinternal` | String/any conversion utilities |

## Usage

```go
import "gitlab.com/auk-go/core/reflectcore"

// Type name extraction
name := reflectcore.TypeName.OfAny(myStruct)

// Null checking via Is predicates
if reflectcore.Is.Null(someInterface) {
    // handle nil
}

// Slice conversion
converted := reflectcore.SliceConverter.ToStrings(anySlice)

// Iterate struct fields
reflectcore.Looper.StructFields(myStruct, func(fp reflectmodel.FieldProcessor) {
    fmt.Println(fp.Name, fp.FieldType)
})
```

## How to Extend Safely

- **New reflection utilities**: Add to `internal/reflectinternal` and re-export via `vars.go`.
- **New models**: Add to `reflectmodel/` — keep models as pure data structs with minimal methods.
- **Do not** add business logic to this package — it is a facade and model layer only.

## Related Docs

- [Coding Guidelines](../spec/01-app/17-coding-guidelines.md)
