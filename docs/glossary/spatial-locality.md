# Spatial Locality

## Definition

Spatial locality means that memory locations near recently accessed memory are likely to be accessed soon.

---

## Example

Array traversal

```go
for i := range arr
```

---

## Benefits

- Better cache utilization
- Fewer cache misses
- Improved performance

---

## Opposite

Random memory access