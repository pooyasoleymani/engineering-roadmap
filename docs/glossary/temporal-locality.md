# Temporal Locality

## Definition

Temporal locality means recently accessed data is likely to be accessed again.

---

## Example

```go
counter++

counter++

counter++
```
---

## Benefit

CPU keeps frequently used data in cache.