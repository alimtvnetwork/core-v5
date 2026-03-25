# Issue: stringsTo.BytesWithDefaults() Missing Negative Value Check

## Status: ✅ RESOLVED

## Phase: 9.1 — Critical Bugs

## File

`converters/stringsTo.go`

## Description

`BytesWithDefaults()` checked for `vInt > MaxUnit8AsInt` but not for `vInt < 0`. Negative integers would be silently cast to `byte(vInt)`, wrapping around (e.g., `-1` → `255`). The sister function `stringTo.Byte()` correctly checks `vInt < 0`.

## Fix

Changed condition to `vInt < 0 || vInt > constants.MaxUnit8AsInt`.
