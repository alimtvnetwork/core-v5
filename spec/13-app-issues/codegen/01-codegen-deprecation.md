# Codegen Package Should Be Deprecated

## Issue Summary

The `codegen/` package adds significant complexity and maintenance burden. It generates unit tests from function signatures via reflection, but modern IDEs and AI tools can do this more effectively.

## Root Cause Analysis

Built as a productivity tool before AI-assisted code generation became prevalent.

## Fix Description

See [Codegen Deprecation Plan](/spec/01-app/10-codegen-deprecation-plan.md).

## TODO and Follow-Ups

- [x] Mark as deprecated → Done (doc.go deprecation notices in all sub-packages)
- [x] Audit consumers → Done (only `cmd/main/unitTestGenerator.go` and codegen's own tests)
- [ ] Remove after migration → Pending (S-006, awaiting external audit confirmation)

## Done Checklist

- [ ] Codegen removed → Pending removal
- [ ] No broken imports → Will verify post-removal
