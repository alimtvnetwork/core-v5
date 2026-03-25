# Issue: BaseTags.TagsHashset() Value Receiver Caching No-Op

## Status: ✅ RESOLVED

## Phase: 7.2 — Code Quality

## File

`coreinstruction/BaseTags.go`

## Description

`TagsHashset()` used a value receiver, so `it.tagsHashset = ...` was assigned on a copy and discarded after return. The hashset cache was recomputed on every call.

## Fix

Changed to pointer receiver `*BaseTags`.
