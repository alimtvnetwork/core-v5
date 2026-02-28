# loggerinf — Logger Interface Contracts

Package `loggerinf` defines all interface contracts for the logging subsystem. It provides composable, level-based logging interfaces from atomic single-level loggers up to full multi-level loggers with format support, persistent attributes, conditional logging, and stack-skip control.

## Architecture

```
loggerinf/
├── atomic-interfaces.go                        # Atomic primitives: Committer, Compiler, Serializer, WriterConfigurer
├── Logger.go                                   # Logger — composite of Fatal/Error/Warning/Info/Debug loggers
├── FullLogger.go                               # FullLogger — Logger + FormatLogger
├── StandardLogger.go                           # StandardLogger — FullLogger + LogTypeWriter + AllLogWriter
├── all-logger.go                               # LogDefiner, LogDefinerWriter, LogTypeWriter, StandardLoggerGetter
├── FatalLogger.go                              # FatalLogger interface
├── ErrorLogger.go                              # ErrorLogger interface
├── WarningLogger.go                            # WarningLogger interface
├── InfoLogger.go                               # InfoLogger interface
├── DebugLogger.go                              # DebugLogger interface
├── FormatLogger.go                             # FormatLogger — composite of Format*Logger interfaces
├── FormatDebugLogger.go                        # FormatDebugLogger interface
├── FormatErrorLogger.go                        # FormatErrorLogger interface
├── FormatFatalLogger.go                        # FormatFatalLogger interface
├── FormatInfoLogger.go                         # FormatInfoLogger interface
├── FormatWarningLogger.go                      # FormatWarningLogger interface
├── VoidLogger.go                               # VoidLogger — void log functions (no return)
├── VoidFatalLogger.go                          # VoidFatalLogger — void fatal log
├── VoidStringLogger.go                         # VoidStringLogger — void string log
├── SingleLogger.go                             # SingleLogger — single log-type writer
├── AllLogWriter.go                             # AllLogWriter — write to all levels
├── BasePersistentLogger.go                     # BasePersistentLogger — logger with persistent attributes
├── AttrPersistentLogger.go                     # AttrPersistentLogger — attribute-based persistent logger
├── ConditionalBasePersistentLogger.go          # ConditionalBasePersistentLogger — conditional persistent logger
├── ConditionalStandardLogger.go                # ConditionalStandardLogger — conditional standard logger
├── LoggerWithFormatLogger.go                   # LoggerWithFormatLogger — Logger + FormatLogger composite
├── LoggerWithFormatLoggerContractsBinder.go     # Contracts binder for LoggerWithFormatLogger
├── LoggerContractsBinder.go                    # Contracts binder for Logger
├── LoggerStackSkip.go                          # LoggerStackSkip — stack-skip-aware logger
├── LoggerTyper.go                              # LoggerTyper — log type enum interface
├── LogTypeChecker.go                           # LogTypeChecker — check if log type is enabled
├── Caller.go                                   # Caller — caller info interface
├── FileWithLiner.go                            # FileWithLiner — file + line number
├── MetaAttributesCompiler.go                   # MetaAttributesCompiler — compile meta attributes
├── MetaAttributesStacker.go                    # MetaAttributesStacker — stack meta attributes
├── MetaAttributesWithoutTileStacker.go         # MetaAttributesWithoutTileStacker — stack without title
├── ModelCollectioner.go                        # ModelCollectioner — collect log models
├── SingleLogModeler.go                         # SingleLogModeler — single log model
├── ReflectSetter.go                            # ReflectSetter — reflection-based setter
├── SpecificValuer.go                           # SpecificValuer — specific value getter
├── all-checkers.go                             # Checker interfaces: IsEnabled, IsDisabled, etc.
├── all-getters.go                              # Getter interfaces: StackSkip, LogType, Writer
└── all-namers.go                               # Namer interfaces: TypeName, LogTypeName
```

## Interface Hierarchy

```
Logger                        ← Fatal + Error + Warning + Info + Debug
    ↓
FormatLogger                  ← FormatFatal + FormatError + FormatWarning + FormatInfo + FormatDebug
    ↓
FullLogger                    ← Logger + FormatLogger
    ↓
StandardLogger                ← FullLogger + LogTypeWriter + AllLogWriter
    ↓
BasePersistentLogger          ← StandardLogger + persistent attributes
    ↓
ConditionalStandardLogger     ← conditional logging based on log type
```

## Key Interfaces

| Interface | Description |
|-----------|-------------|
| `Logger` | Composite: Fatal + Error + Warning + Info + Debug |
| `FullLogger` | Logger + FormatLogger |
| `StandardLogger` | FullLogger + LogTypeWriter + AllLogWriter |
| `LogDefiner` | TypeName + Message + RawBytes + Deserializer |
| `LogDefinerWriter` | Write LogDefiner entities with optional stack skip |
| `LogTypeWriter` | Write by log type string with variadic args |
| `LoggerTyper` | Log type enum interface |

## Related Docs

- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
