# Benchmark - Array with use built-in copy

Hardware

Intel(R) Core(TM) i5-10600 CPU @ 3.30GHz

OS

Ubuntu 22.04

Language

Go 1.26

---
## Goal 
Compare copy with built-in copy and manual copy in DynamicArray[T]

---
## Expected Result
Better performance for insert and reallocation

## Measurements

| Structure   | ns/op | allocs/op | bytes/op |
| ----------- | ----: | --------: | -------: |
| copy        |     34626392  |     0      |     20971     |
| manual copy |   43945273    |      0     |      20971    |

Conclusion