---
Date: 2026-08-02
tags:
  - software_engineering
Related: "[[sorting-merge-sort]]"
Next: "[[lesson-1.6.3-quick-sort]]"
---
---
## Objective

- Divide and Conquer
- Recursion Trees
- Merge operation
- Why Merge Sort is stable
- Time complexity proof
- Space complexity
- Cache behavior
- External Merge Sort
- Parallel Merge Sort
- Why databases love Merge Sort
- Why Go doesn't use Merge Sort for slices

---

##  Divide and Conquer

Many difficult problems become easy if we:

1. Divide the problem
2. Solve smaller problems
3. Combine the results

Instead of sorting:

```text
8 2 5 1 9 6 3 4
```

we divide:

```text
8 2 5 1

9 6 3 4
```

then again:

```text
8 2

5 1

9 6

3 4
```

until:

```text
8

2

5

1

9

6

3

4
```

A single element is already sorted.

This is the **base case**.

---

## Merge

Now the interesting part begins.

We don't sort while dividing.

We sort while **merging**.

Example:

```text
2

8
```

↓

```text
2 8
```

---

Another

```text
1

5
```

↓

```text
1 5
```

Now merge:

```text
2 8

1 5
```

Compare

```text
2

1
```

Take smaller

```text
1
```

Now compare

```text
2

5
```

Take

```text
2
```

Continue

Result

```text
1 2 5 8
```

Notice:

Neither half is re-sorted.

They are already sorted.

The merge operation only combines them.

---

##  The Recursion Tree

For eight elements:

```text
                8
           /         \
         4             4
       /   \         /   \
      2     2       2     2
     / \   / \     / \   / \
    1   1 1   1    1  1  1  1
```

Height

```text
log₂(n)
```

At every level,

the total work is

```text
n
```

Therefore

```text
n

×

log n
```

Final complexity

```text
O(n log n)
```

This proof is more important than memorizing the result.

---

##  Why Merge Sort Is Stable

Suppose

```text
Alice 5000

Bob 5000
```

Alice is left.

Bob is right.

During merge

Values equal.

Which one should we take?

Always choose the left element first.

Result

```text
Alice

Bob
```

Relative order preserved.

That's exactly what makes Merge Sort stable.

---

## Space Complexity

Merge Sort needs temporary memory.

Example

```text
Original

↓

Temporary Buffer

↓

Merged Output
```

Auxiliary memory

```text
O(n)
```

This is Merge Sort's biggest disadvantage.

---

##  Why Merge Sort Is Predictable

Quick Sort

Worst case

```text
O(n²)
```

Merge Sort

Always

```text
O(n log n)
```

No pathological input.

This predictability is extremely valuable in production systems.

---

## Cache Behavior

During merge we scan:

```text
Left

↓

Right

↓

Output
```

Sequential memory access.

Modern CPUs love sequential access.

Hardware prefetchers work efficiently.

Merge Sort is surprisingly cache-friendly despite using extra memory.

---

## External Merge Sort

Suppose you need to sort

```text
5 TB
```

RAM

```text
64 GB
```

Impossible?

No.

Algorithm:

1. Read a chunk.
2. Sort it in memory.
3. Write it back to disk.
4. Repeat.

Now we have many sorted files.

Merge them:

```text
File A

↓

File B

↓

File C

↓

Merged Output
```

This is **External Merge Sort**.

Databases and distributed systems use this technique routinely.

---

## Parallel Merge Sort

Merge Sort naturally supports parallelism.

```text
Left Half

↓

CPU 1
```

```text
Right Half

↓

CPU 2
```

After both finish,

merge the results.

This divide-and-conquer structure maps well to multicore processors.

---

## Why Go Doesn't Use Merge Sort for Slices

Go's `sort.Sort` and `slices.Sort` do **not** use Merge Sort.

Why?

Because Merge Sort:

- Allocates extra memory
- Copies elements
- Has allocation overhead

Go prioritizes:

- Low allocations
- Cache efficiency
- Fast average-case performance

So it uses **PDQSort**, a highly optimized Quick Sort variant.

However,

Go **does** provide stable sorting:

```go
slices.SortStableFunc(...)
```

That implementation trades memory and speed for stability.

This is an example of exposing different algorithms because they solve different problems.

---

## Complexity Summary

| Property         | Merge Sort |
| ---------------- | ---------- |
| Best             | O(n log n) |
| Average          | O(n log n) |
| Worst            | O(n log n) |
| Extra Memory     | O(n)       |
| Stable           | ✅          |
| In-place         | ❌          |
| Parallelizable   | Excellent  |
| External Sorting | Excellent  |
| Cache Behavior   | Good       |

---

## Real-World Uses

Merge Sort is widely used in:

- Database external sorting
- Distributed systems (MapReduce, Spark)
- Log processing pipelines
- Stable sorting libraries
- File sorting utilities
- Data warehousing

---

## Trade-offs

| Merge Sort Strength | Cost              |
| ------------------- | ----------------- |
| Stable              | Extra memory      |
| Predictable         | Allocations       |
| Parallel            | Merge overhead    |
| Sequential access   | Additional buffer |

No algorithm is universally best. The right choice depends on workload.



---

## Summery

- Merge Sort recursively divides the problem into smaller `subproblems` until each contains one element (already sorted), then combines the sorted `subproblems` by merging them into larger sorted sequences.
- Merge sort  Recursion Tree have `log(n)` level and every level have `n` work then complexity in merge sort is `n log(n)` 