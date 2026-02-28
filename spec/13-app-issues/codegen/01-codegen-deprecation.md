# Codegen Package Should Be Deprecated

## Issue Summary

The `codegen/` package adds significant complexity and maintenance burden. It generates unit tests from function signatures via reflection, but modern IDEs and AI tools can do this more effectively.

## Root Cause Analysis

Built as a productivity tool before AI-assisted code generation became prevalent.

## Fix Description

See [Codegen Deprecation Plan](/spec/01-app/10-codegen-deprecation-plan.md).

## TODO and Follow-Ups

- [ ] Mark as deprecated
- [ ] Audit consumers
- [ ] Remove after migration

## Done Checklist

- [ ] Codegen removed
- [ ] No broken imports
