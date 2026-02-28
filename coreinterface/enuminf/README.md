# enuminf — Enum Interface Contracts

Package `enuminf` defines all interface contracts for the enum subsystem. It provides composable interfaces for enum value access, type checking, formatting, comparison, marshalling/unmarshalling, and contracts binding across all supported numeric types (byte, int8, int16, int32, uint16) and string enums.

## Architecture

```
enuminf/
├── BaseEnumer.go               # BaseEnumer — foundation: Name, TypeName, Value, IsValid, IsInvalid
├── BasicEnumer.go              # BasicEnumer — BaseEnumer + Formatter + MinMax + Ranges + AllNameValues
├── SimpleEnumer.go             # SimpleEnumer — BaseEnumer + ToEnumString
├── BasicEnumWithComparer.go    # BasicEnumWithComparer — BasicEnumer + IsValueEqual
├── ByteEnumValuer.go           # ByteEnumValuer — byte-specific value getter
├── EnumFormatter.go            # EnumFormatter — format enum using template: {type-name}, {name}, {value}
├── EnumTypeChecker.go          # EnumTypeChecker — check enum type identity
├── EnumTyper.go                # EnumTyper — enum type metadata
├── OnlySupportedNamesErrorer.go # OnlySupportedNamesErrorer — generate "only supported" errors
├── PatternsSplitter.go         # PatternsSplitter — split enum patterns
├── all-checkers.go             # Checker interfaces: IsValid, IsInvalid, IsValueEqual, IsAnyValueEqual
├── all-cmd-typer.go            # Command typer interfaces
├── all-common-types.go         # Common type interfaces: CompletionStateTyper, StringCompareTyper, etc.
├── all-compare-enumer.go       # Comparison interfaces: CompareMethodsTyper
├── all-enumers.go              # Typed enumer interfaces: BasicByteEnumer, BasicInt8Enumer, ..., BasicInt64Enumer
├── all-equalers.go             # Equality interfaces: IsValueByteEqualer, IsValueInteger8Equaler, etc.
├── all-getters.go              # Getter interfaces: RangesDynamicMapGetter, IntegerEnumRangesGetter, etc.
├── all-namers.go               # Namer interfaces: ToNamer, TypeNameGetter, RangeNamesCsvGetter
├── all-stringers.go            # Stringer interfaces: ToEnumStringer, ToNumberStringer, Int*ToEnumStringer
├── all-unmarshal-values.go     # Unmarshalling interfaces: UnmarshallEnumToValue* per numeric type
└── enum-contracts.go           # Contracts binders: BasicByteEnumContractsBinder, BasicInt8EnumContractsBinder, etc.
```

## Interface Hierarchy

```
BaseEnumer                    ← Name, TypeName, Value, IsValid, IsInvalid, Stringer, Jsoner
    ↓
BasicEnumer                   ← BaseEnumer + EnumFormatter + MinMax + Ranges + AllNameValues
    ↓
StandardEnumer                ← BasicEnumer + StringRanges + RangeValidate + JsonContractsBinder
    ↓
BasicContractsEnumer          ← BasicEnumer + TypeNameWithRangeNamesCsv
    ↓
Basic*ContractsEnumer         ← per-type (Byte, Int8, Int16, Int32) + value equality
    ↓
Basic*EnumContractsBinder     ← per-type + AsBasic*EnumContractsBinder() self-reference
```

## Key Interfaces

| Interface | Description |
|-----------|-------------|
| `BaseEnumer` | Name, TypeName, Value, IsValid/IsInvalid, Stringer, Jsoner |
| `BasicEnumer` | Full enum: format, min/max, ranges, AllNameValues, OnlySupportedNamesError |
| `SimpleEnumer` | Lightweight: BaseEnumer + ToEnumString |
| `BasicByteEnumer` | Byte-specific: ValueByte, MinByte, MaxByte, RangesByte |
| `BasicInt8Enumer` | Int8-specific: ValueInt8, MinInt8, MaxInt8, RangesInt8 |
| `BasicInt16Enumer` | Int16-specific: ValueInt16, MinInt16, MaxInt16, RangesInt16 |
| `BasicInt32Enumer` | Int32-specific: ValueInt32, MinInt32, MaxInt32, RangesInt32 |
| `BasicIntEnumer` | Int-specific: ValueInt, MinInt, MaxInt, RangesInt |
| `BasicInt64Enumer` | Int64-specific: ValueInt64, MinInt64, MaxInt64, RangesInt64 |
| `EnumFormatter` | Format using template: `{type-name}`, `{name}`, `{value}` |
| `EnumTyper` | Enum type metadata |
| `StandardEnumer` | BasicEnumer + string ranges + range validation + JSON binding |

### Enum Formatter Template

```
Format: "Enum of {type-name} - {name} - {value}"
Output: "Enum of EnumFullName - Invalid - 0"
```

## Related Docs

- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [coreimpl/enumimpl](/coreimpl/README.md) — concrete implementations
