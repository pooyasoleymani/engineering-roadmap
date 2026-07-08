
## Definition

>"Every high-performance system is built on arrays."

An array is a contiguous block of memory that stores elements of the same type.

Each element occupies the same amount of memory, allowing the CPU to calculate the address of any element directly using its index.

---

## Example

**Example** :  array of `int64` each element occupies exactly *8 bytes* .

```text
Index :     0      1      2      3      4

Value :     9      8      2      4      5

Address: 1000   1008   1016   1024   1032
```

---
## Advantages

- Contiguous memory improves cache locality.
- Constant-time indexing (O(1)).
- CPU can calculate element addresses directly.
- Low memory overhead because no pointer is stored per element.
- Hardware prefetchers work efficiently with sequential access.

---
## Disadvantages

- Inserting in the middle requires shifting subsequent elements (O(n)).
- Deleting from the middle requires shifting remaining elements (O(n)).
- Static arrays have fixed size.
- Dynamic arrays occasionally resize, requiring allocation and copying.

---
## Trade-offs


### Linked List

Advantages

- Constant-time insertion when the node is already known.

Disadvantages

- Poor cache locality.
- Frequent heap allocations.
- Pointer chasing.
- Higher memory overhead.
- Difficult for CPU prefetchers.

---

### Dynamic Array

Advantages

- Contiguous memory.
- Excellent cache locality.
- Fewer allocations.
- Better throughput.
- Lower memory overhead.

Disadvantages

- Middle insertion is O(n).
- Occasional resize requires copying.

---
## Complexity

| Operation       | Time           |
| --------------- | -------------- |
| Index           | O(1)           |
| Search          | O(n)           |
| Insert (End)    | Amortized O(1) |
| Insert (Middle) | O(n)           |
| Delete          | O(n)           |


---
## Real-world Usage

Arrays are the foundation of many high-performance systems, including:

- Go slices
- C++ std::vector
- Python list
- Ring buffers
- Hash table buckets
- CPU caches
- Database pages
- Image processing
- Machine learning tensors


---
## References

- Efficient Go 
- Computer Systems: A Programmer's Perspective
- Designing Data-Intensive Applications

---
## Related Topics

- Dynamic Array [[dynamic-array]]
- Slice 
- Cache Line [[cache-line]]
- Spatial Locality [[spatial-locality]]
- Amortized Analysis [[amortized-analysis]]
- Pointer Chasing [[docs/glossary/pointer-chasing|pointer-chasing]]