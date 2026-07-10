---
Date: 2026-07-08
---
---
# Objectives

- Physical memory layout
- Static vs Dynamic arrays
- How arrays are stored in RAM
- Why indexing is O(1)
- Why insertion is O(n)
- Why deletion is O(n)
- Dynamic array growth
- Why Go slices are powerful
- Why `std::vector` exists
- Python list internals (high level)

---
# Summary

Every high performance system build with array


---
## Array

An array is a **contiguous block of memory** containing elements of the same type.

Example:
```text
Index:     0    1    2    3   4
Value:     8    2    5    9   1
Address: 1000 1008 1016 1024 1032 
```

---

## Address Calculation

CPU can Compute address of array directly and no search is required.
That's why indexing is O(1).

```text
Base Address = 1000
Element Size = 8 bytes
--------------------------------
Address = base + index * size
--------------------------------
arr[4] -> 1000 + 4 * 8
```

---

## Insertion

For insert into the array every element must move one position.
Time complexity for this work is O(n).

Example: Insert 99 at index 2
```text
arr : 1 2 3 4 5
element move: 3 4 5
arr : 1 2 99 3 4 5 
```

---
## Deletion

Delete:

```
2
```

Array becomes:

```
1 3 4 5
```

Again:

Every remaining element shifts left.

Complexity:

```
O(n)
```


---
## Dynamic Arrays

Dynamic array is a array that can insert and delete element dynamically.
Suppose capacity:

```
4
```

Current:

```
1 2 3 4
```

Append:

```
5
```

No space.

The runtime:

1. Allocates a larger array.
2. Copies existing elements.
3. Frees the old array (or lets the GC reclaim it in Go).
4. Appends the new element.
---

## Go Slice Growth

A slice contains:

```
PointerLengthCapacity
```

Example:

```
numbers := []int{1,2,3}
```

Internally:

```
Pointer -----> ArrayLength = 3Capacity = 3
```

Append:

```
numbers = append(numbers,4)
```

If capacity is exhausted:

A new backing array is allocated and elements are copied.

---
## Why Capacity Exists

Without extra capacity:

Every append would require:

- allocation
- copy
- free/GC

Complexity:

```
O(n)
```

Instead:

Go over-allocates.

Result:

Most appends are:

```
O(1)Amortized
```

---
## Real-World Example

Imagine building a TCP packet queue.
Packets:

```
Packet1Packet2Packet3
```

Store them in:

### Linked List

Pros:

Easy insertion.

Cons:

- many allocations
- cache misses
- pointer chasing

---

### Dynamic Array

Pros:

- contiguous
- cache-friendly
- fewer allocations
- higher throughput

This is why many high-performance networking systems use arrays or ring buffers.

---

## Trade-offs

|Operation|Array|Linked List|
|---|---|---|
|Index|O(1)|O(n)|
|Search|O(n)|O(n)|
|Insert End (with spare capacity)|Amortized O(1)|O(1)|
|Insert Middle|O(n)|O(1) (if node known)|
|Delete Middle|O(n)|O(1) (if node known)|
|Cache Locality|Excellent|Poor|
|Memory Overhead|Low|Higher (pointer per node)|

Notice that complexity alone doesn't determine the better choice.