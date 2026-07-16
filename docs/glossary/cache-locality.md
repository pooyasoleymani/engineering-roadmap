## Definition

Cache locality (or locality of reference) is a principle where programs access the same data, or nearby data, repeatedly. It improves CPU performance by ensuring frequently used information is stored in fast, local cache memory instead of slower main memory (RAM).


---
## Why It Matters

Modern processors fetch data in blocks called "cache lines" (typically 64 bytes). If your code executes with poor locality, the CPU constantly waits on slow main memory—which can be 100× slower than accessing an `L1 cache`. Optimizing code for cache locality vastly reduces cache misses.

---
## Types

- **Temporal Locality:** The tendency to access the _same_ memory location repeatedly within a short period (e.g., in a loop, accessing the same counter variable multiple times).
- **Spatial Locality:** The tendency to access memory locations that are _close together_ (e.g., sequentially traversing elements in an array). 

## Real-World Optimization

- **Arrays vs. Linked Lists:** Arrays have excellent spatial locality because their elements are contiguous in memory, allowing the CPU to easily load subsequent elements. Linked lists have poor locality because nodes are scattered across random memory addresses. 
- **Matrix Multiplication:** In multi-dimensional array operations, looping through rows sequentially utilizes spatial locality, whereas column-major processing can ruin cache performance.
- **Data Structures:** Structuring your data into contiguous blocks (like using data-oriented design in game development) keeps related data tightly packed to maximize the 64-byte cache line.

---
## Related Topics

- Spatial Locality [[spatial-locality]]
- Temporal Locality [[temporal-locality]]
- Array [[array]]
- Cache miss [[cache-miss]]
- Cache line [[cache-line]]
- Pointer chasing [[docs/glossary/pointer-chasing|pointer-chasing]]

