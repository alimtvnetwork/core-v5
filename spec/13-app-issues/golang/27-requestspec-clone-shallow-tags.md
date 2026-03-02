# Issue: RequestSpecification.Clone() Shallow Tags Copy

## Status: ✅ RESOLVED

## Phase: 7.2 — Code Quality

## File

`coreinstruction/RequestSpecification.go`

## Description

`Clone()` assigned `Tags: r.Tags` directly, sharing the backing array between original and clone. Mutating tags on either would corrupt the other. Same class of bug previously fixed in `Specification.Clone()`.

## Fix

Deep-copied `Tags` using `make` + `copy`.
