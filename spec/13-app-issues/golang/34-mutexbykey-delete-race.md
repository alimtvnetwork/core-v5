# Issue: mutexbykey.Delete() Race Condition on Locked Mutex

## Status: ✅ RESOLVED (Documented)

## Phase: 8.2 — Code Quality

## File

`mutexbykey/mutexMap.go`

## Description

`Delete()` removes a mutex from the map while it may still be locked by another goroutine that obtained it via `Get()`. This creates a race: goroutine A gets and locks the mutex, goroutine B deletes it from the map, goroutine C creates a new mutex for the same key — now A and C can operate concurrently on the same logical key, defeating the purpose.

## Recommendation

Document the contract clearly (caller must ensure no concurrent users before deleting), or add reference counting to prevent deletion while in use.
