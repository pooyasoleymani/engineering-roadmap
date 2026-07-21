# Benchmark - Array vs Linked List Traversal

Hardware

Intel Core `i5-4200U`

OS

Ubuntu 22.04

Language

Go 1.26

---

## Goal

Compare sequential traversal of a slice against traversal of a linked list.

---

## Expected Result

The slice should outperform the linked list due to:

- Better cache locality
- Hardware prefetching
- Fewer pointer dereferences
- Reduced cache misses

---

## Measurements

| Structure   | ns/op | allocs/op | bytes/op |
| ----------- | ----: | --------: | -------: |
| Slice       |       |           |          |
| Linked List |       |           |          |

---

## Conclusion

Although both traversals have O(n) complexity, contiguous memory gives slices significantly better real-world performance.