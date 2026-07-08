## Definition

Cache locality (or locality of reference) is a principle where programs access the same data, or nearby data, repeatedly. It improves CPU performance by ensuring frequently used information is stored in fast, local cache memory instead of slower main memory (RAM).


---
## Why It Matters

Modern processors fetch data in blocks called "cache lines" (typically 64 bytes). If your code executes with poor locality, the CPU constantly waits on slow main memory—which can be 100× slower than accessing an `L1 cache`. Optimizing code for cache locality vastly reduces cache misses.

---
