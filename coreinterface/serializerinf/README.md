# serializerinf — Serializer Interface Contracts

Package `serializerinf` defines all interface contracts for serialization, deserialization, and JSON-related operations. Provides composable atomic interfaces for bytes/string serialization, JSON formatting, error inspection, and self-serializing types.

## Architecture

```
serializerinf/
├── Serializer.go                           # Serializer — Serialize() ([]byte, error)
├── MustSerializer.go                       # MustSerializer — SerializeMust() []byte
├── BytesInToSelfDeserializer.go            # BytesInToSelfDeserializer — deserialize bytes into self
├── MustBytesInToSelfDeserializer.go        # MustBytesInToSelfDeserializer — must-deserialize bytes into self
├── ByteToJsonMustStringer.go               # ByteToJsonMustStringer — bytes → JSON string (must)
├── FieldBytesToPointerDeserializer.go      # FieldBytesToPointerDeserializer — field-level deserialization
├── SerializerDeserializer.go               # SerializerDeserializer — Serializer + Must + Deserializer
├── SerializerDeserializerBinder.go         # SerializerDeserializerBinder — contracts binder
├── SelfBytesSerializerDeserializer.go      # SelfBytesSerializerDeserializer — self-referencing serialize/deserialize
├── all-json-releated-methods.go            # 35+ atomic JSON/bytes interfaces
└── all-json-resulter.go                    # JSON result interfaces
```

## Key Interfaces

### Core Serialization

| Interface | Description |
|-----------|-------------|
| `Serializer` | `Serialize() ([]byte, error)` |
| `MustSerializer` | `SerializeMust() []byte` |
| `SelfSerializer` | `Serialize() ([]byte, error)` — self-referencing |
| `MustSelfSerializer` | `SerializeMust() []byte` — self-referencing |
| `SerializerDeserializer` | Serializer + MustSerializer + BytesInToSelfDeserializer |

### Deserialization

| Interface | Description |
|-----------|-------------|
| `BytesInToSelfDeserializer` | Deserialize bytes into the receiver |
| `MustBytesInToSelfDeserializer` | Must-variant of bytes deserialization |
| `Deserializer` | `Deserialize(anyPointer any) error` |
| `MustDeserializer` | `DeserializeMust(anyPointer any)` |
| `Unmarshaler` | `Unmarshal(anyPointer any) error` |
| `MustUnmarshaler` | `UnmarshalMust(anyPointer any)` |

### JSON String/Bytes Access

| Interface | Description |
|-----------|-------------|
| `JsonStringGetter` | `JsonString() string` |
| `JsonStringPtrGetter` | `JsonStringPtr() *string` |
| `PrettyJsonStringGetter` | `PrettyJsonString() string` |
| `PrettyJsonBufferGetter` | `PrettyJsonBuffer(prefix, indent) (*bytes.Buffer, error)` |
| `SafeStringGetter` | `SafeString() string` — nil-safe |
| `SafeBytesGetter` | `SafeBytes() []byte` — nil-safe |

### Raw Serialization

| Interface | Description |
|-----------|-------------|
| `RawSerializeGetter` | `Raw() ([]byte, error)` |
| `MustRawSerializeGetter` | `RawMust() []byte` |
| `RawStringSerializeGetter` | `RawString() (string, error)` |
| `MustRawStringSerializeGetter` | `RawStringMust() string` |
| `RawPrettyStringGetter` | `RawPrettyString() (string, error)` |

### Inspection & Checks

| Interface | Description |
|-----------|-------------|
| `HasErrorChecker` | `HasError() bool` |
| `IsEmptyChecker` | `IsEmpty() bool` |
| `HasAnyItemChecker` | `HasAnyItem() bool` |
| `IsEmptyErrorChecker` | `IsEmptyError() bool` |
| `HasIssuesOrEmptyChecker` | `HasIssuesOrEmpty() bool` |
| `IsAnyNullChecker` | `IsAnyNull() bool` |
| `HasBytesChecker` / `HasJsonBytesChecker` | Check for non-empty bytes |
| `ErrorHandler` | `HandleError()` — panic on error |
| `MustBeSafer` | `MustBeSafe()` — panic if unsafe |

## Related Docs

- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [corejson](/coredata/corejson/README.md) — concrete JSON implementation
