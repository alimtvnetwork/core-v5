# Issue: NewFlatSpecificationUsingSpec() Shallow Tags Reference

## Status: ✅ RESOLVED

## Phase: 7.2 — Code Quality

## File

`coreinstruction/FlatSpecification.go`

## Description

`NewFlatSpecificationUsingSpec()` assigned `Tags: spec.Tags` directly, sharing the slice backing array with the source `Specification`. Mutation of tags on either side would affect both.

## Fix

Deep-copied `Tags` using `make` + `copy`.
