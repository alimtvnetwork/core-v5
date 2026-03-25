# coreinterface

## Folder Purpose

Defines the canonical interface contracts for the entire auk-go ecosystem. Every interface that downstream packages should implement or accept lives here. This is the **contract layer** — it contains zero implementation.

## Responsibilities

1. Define getter interfaces (`*Getter` suffix) for every common data accessor pattern.
2. Define checker interfaces (`Is*Checker`, `Has*Checker`) for boolean queries.
3. Define action interfaces (`Executor`, `Initializer`, `Planner`) in `baseactioninf/`.
4. Define binder interfaces (`*ContractsBinder`) for type-safe contract binding.
5. Define entity interfaces in `entityinf/`.
6. Define serializer contracts in `serializerinf/`.

## Key Files and Entrypoints

| File | Purpose |
|------|---------|
| `all-getters.go` | ~80 getter interfaces (Name, Value, Error, Length, …) |
| `all-is-checkers.go` | ~60 checker interfaces (IsValid, IsEmpty, IsNull, …) |
| `all-binders.go` | Contract binder interfaces |
| `all-stringers.go` | String-output interfaces |
| `all-serializer.go` | Serializer / MustSerializer |
| `all-namers.go` | Naming interfaces |
| `baseactioninf/` | Executor, SimpleExecutor, StandardExecutor |
| `entityinf/` | Entity definer interfaces |
| `CsvLiner.go` | CSV line generation interface |
| `Csver.go` | CSV string generation interface |

## Naming Conventions (CRITICAL)

Go convention: interface names use the **-er** suffix.

| Pattern | Example | Rule |
|---------|---------|------|
| Single-method getter | `NameGetter` | `{Field}Getter` |
| Single-method checker | `IsValidChecker` | `Is{Condition}Checker` |
| Action interface | `Executor` | Verb + `-er` |
| Binder | `BasicSlicerContractsBinder` | `{Type}ContractsBinder` |
| Combined | `SimpleValidInvalidChecker` | Compose smaller interfaces |

## How to Extend Safely

- One interface per logical concept; compose via embedding.
- Never add implementations to this package.
- Follow the `-er` suffix convention consistently.
- Group related interfaces in `all-*.go` files.

## Common Mistakes and Prevention

| Mistake | Prevention |
|---------|-----------|
| Adding implementation code | This package is contracts-only |
| Breaking `-er` convention | Review all new interface names |
| Duplicating method signatures across interfaces | Embed existing interfaces instead |

## Related Docs

- [Core Interface Conventions](../14-core-interface-conventions.md)
- [Repo Overview](../00-repo-overview.md)
