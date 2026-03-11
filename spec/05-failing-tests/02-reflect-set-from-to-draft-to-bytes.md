# 02 — ReflectSetFromTo DraftToBytes Marshal-then-Unmarshal Bug

## Test
`Test_ReflectSetFromTo_DraftToBytes` — Case 0: "(otherType, *[]byte) -- try marshal, reflect"

## Root Cause
**Production code bug** in `ReflectSetFromTo` (`coredata/coredynamic/ReflectSetFromTo.go`).

The documented supported case `(otherType, *[]byte)` is: marshal `From` (a struct) into JSON
bytes, then store those bytes into the `*[]byte` destination. The code correctly performed
step 1 (`json.Marshal(from)` → `rawBytes`), but then fell through to a generic
`json.Unmarshal(rawBytes, toPointer)` call.

`json.Unmarshal` into a `*[]byte` expects a **base64-encoded JSON string**, not a JSON object.
Since `rawBytes` contained `{"SampleString1":"Expected",...}` (a JSON object), the unmarshal
failed, returning an error.

### Flow Before Fix
```
From: *DraftType{Expected}  →  json.Marshal  →  rawBytes (JSON object)
                                                      ↓
                              json.Unmarshal(rawBytes, *[]byte)  ← FAILS (expects base64 string)
```

## Solution
After successful `json.Marshal` and when the destination is `*[]byte`, **set the bytes
directly** via `*toPointer.(*[]byte) = rawBytes` instead of falling through to `json.Unmarshal`.

### Flow After Fix
```
From: *DraftType{Expected}  →  json.Marshal  →  rawBytes (JSON object)
                                                      ↓
                              *toPointer.(*[]byte) = rawBytes  ← direct assignment ✓
```

The `json.Unmarshal` fallback remains for other non-byte destination types that reach that
code path (currently none, but kept for defensive completeness).

## Iteration Details
1. Initial attempt set `From` from `JsonBytesPtr()` to `&struct` — fixed the test case setup
   but revealed the production bug (both lines now `false` instead of just line 2).
2. Root-caused to the `Unmarshal` step — struct JSON cannot be unmarshalled into `[]byte`.
3. Fixed production code to short-circuit with direct byte assignment for `*[]byte` targets.

## Learnings
- `json.Unmarshal` into `*[]byte` only works for **base64-encoded JSON strings**, not
  arbitrary JSON objects/arrays.
- When a function documents supported type combinations, each combination needs its own
  explicit handling path — generic fallthrough logic may not cover all cases correctly.

## What Not to Repeat
- Do not assume `json.Unmarshal` is a universal setter — its behavior is type-dependent.
- When adding new supported type combinations to `ReflectSetFromTo`, verify the full
  marshal→assign path with a dedicated test case before declaring it supported.
