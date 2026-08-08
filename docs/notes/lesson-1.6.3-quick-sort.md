---
Date: 2026-08-04
tags:
  - software_engineering
Next: "[[lesson-1.6.4-heap-sort]]"
---
---
## Objective 

- Divide and Conquer revisited
- Partitioning
- Lomuto Partition
- Hoare Partition
- Pivot selection
- Worst-case behavior
- Randomized Quick Sort
- Median-of-three
- Tail recursion elimination
- Cache behavior
- Why Quick Sort beats Merge Sort in practice
- How Introsort works
- How PDQSort improves Quick Sort


---

## Core Idea

Quick Sort:

```
Choose Pivot

↓

Partition

↓

Sort Left

↓

Sort Right
```

Unlike Merge Sort,

Quick Sort never merges.

Instead,

it moves elements to the correct side of a pivot.


---

## Example

Array

```
9 4 8 3 1 2 5
```

Choose pivot

```
5
```

Partition

```
4 3 1 2

5

9 8
```

Notice

Everything left of the pivot

```
<
```

Everything right

```
>
```

The pivot is now in its **final sorted position**.

That is the key property of Quick Sort.


---

## Partitioning

Partitioning is the heart of Quick Sort.

Everything else is recursion.

There are two famous partition algorithms:

- Lomuto
- Hoare

Understanding their trade-offs is far more important than memorizing code.


---

## Lomuto Partition

Algorithm

```
Pivot = last element

Scan left to right

Move smaller values forward

Finally place pivot
```

Example

```
7 3 8 4 2 5
```

Pivot

```
5
```

Result

```
3 4 2 5 8 7
```

Simple.

Easy to implement.

Popular in textbooks.


### Advantages

- Easy to understand
- Easy to verify
- Good for teaching

### Disadvantages

- Lots of swaps.
- Even when unnecessary.
- That increases memory writes.


---

## Hoare Partition

Invented by Tony Hoare.

Algorithm

Two pointers

```
←

→
```

Move inward.

Swap only when necessary.

---

Example

```
9 3 8 2 5 7
```

Pivot

```
5
```

Pointers move until they find misplaced elements.

Swap.

Repeat.


### Advantages

- Fewer swaps
- Faster
- Better cache behavior
- Used in many production implementations


### Disadvantages

- Harder to reason about.
- Many beginners introduce subtle bugs.


---

## Choosing the Pivot

Suppose the pivot is always the last element.

Already sorted input

```
1 2 3 4 5 6
```

Pivot

```
6
```

Partition

```
1 2 3 4 5

6
```

Left side

```
5 elements
```

Right side

```
0
```

Repeat.

Recursion tree

```
n

↓

n-1

↓

n-2

↓

...
```

Complexity

```
O(n²)
```


---

## Better Pivot Strategies

### 1. Random Pivot

Choose a random element.

Probability of worst-case becomes extremely small.

### 2. Median of Three

Choose

```
First

Middle

Last
```

Take their median.

Example

```
2

50

100
```

Pivot

```
50
```

Produces more balanced partitions.

Most real implementations use something similar.


---

## Complexity

Average

```
O(n log n)
```

Worst

```
O(n²)
```

Space

```
O(log n)
```

(recursion stack)


---

## Why Quick Sort Is Usually Faster Than Merge Sort

Interesting question.

Merge Sort

```
Allocate buffer

↓

Copy

↓

Merge
```

Quick Sort

```
Partition in-place
```

Less memory traffic.

Better cache locality.

Fewer allocations.

Modern CPUs reward these characteristics.

That's why Quick Sort often wins despite the worse theoretical worst case.


---

## Tail Recursion Elimination

Naive Quick Sort

```
quick(left)
quick(right)
```

Maximum recursion

```
O(n)
```

Stack overflow possible.

Optimization

Always recurse into the smaller partition.

Loop over the larger one.

Maximum recursion

```
O(log n)
```

This is used in production libraries.

---

## Cache Behavior

Partition scans memory sequentially.

```
0

1

2

3

4

5
```

Sequential memory

↓

Excellent cache locality

↓

Hardware prefetching

↓

High throughput

Quick Sort's cache behavior is one reason it performs so well.

---

## Why C++ Doesn't Use Plain Quick Sort

Imagine malicious input.

Always worst-case.

```
O(n²)
```

Unacceptable.

C++ uses

## Introsort

Algorithm

```
Quick Sort

↓

Too deep?

↓

Heap Sort
```

The recursion depth is monitored.

If it exceeds a threshold (typically proportional to `2*log₂(n)`), the algorithm switches to Heap Sort, guaranteeing **O(n log n)** worst-case performance while preserving Quick Sort's excellent average-case speed.

---

## Why Go Uses PDQSort

PDQ

```
Pattern Defeating Quick Sort
```

Improves Quick Sort by detecting bad partition patterns.

Features

- Better pivot selection
- Handles duplicates efficiently
- Detects nearly sorted input
- Reduces branch mispredictions
- Avoids common Quick Sort worst cases

Result

Very fast on real workloads.

---

## Stability

Quick Sort

```
NOT Stable
```

During partitioning,

equal elements may move relative to each other.

Example

```
Alice 5000

Bob 5000
```

After Quick Sort

```
Bob 5000

Alice 5000
```

Still sorted.

Not stable.

---

## Complexity Summary

| Property       | Quick Sort |
| -------------- | ---------- |
| Best           | O(n log n) |
| Average        | O(n log n) |
| Worst          | O(n²)      |
| Stable         | ❌          |
| In-place       | ✅          |
| Cache Friendly | Excellent  |
| Extra Memory   | O(log n)   |
| Parallelizable | Good       |

---

# Merge Sort vs Quick Sort

| Feature          | Merge Sort | Quick Sort            |
| ---------------- | ---------- | --------------------- |
| Stable           | ✅          | ❌                     |
| Extra Memory     | O(n)       | O(log n) stack        |
| Worst Case       | O(n log n) | O(n²)                 |
| Average Speed    | Very Good  | Excellent             |
| Cache Locality   | Good       | Excellent             |
| External Sorting | Excellent  | Poor                  |
| Default in Go    | ❌          | Foundation of PDQSort |

---

## Real-World Applications

Quick Sort (or its descendants) is used in:

- Go runtime (`slices.Sort`)
- Rust standard library
- C++ `std::sort` (via Introsort)
- Java primitive sorting (dual-pivot Quick Sort)
- Many embedded systems

---


## Common Production Bugs

Avoid these pitfalls:

1. Poor pivot selection leading to O(n²).
2. Infinite recursion when partition boundaries are wrong.
3. Stack overflow from deep recursion.
4. Mishandling duplicate values.
5. Assuming Quick Sort is stable.
6. Forgetting to optimize for small partitions (many runtimes switch to Insertion Sort).

---

## Exercises

### Theory

Answer these:

1. Why is partitioning the core operation of Quick Sort?
2. Compare Lomuto and Hoare partition schemes. What are the trade-offs?
3. Why does choosing the last element as pivot degrade performance on sorted input?
4. Why does randomized pivot selection reduce the chance of worst-case behavior?
5. Why is Quick Sort generally faster than Merge Sort on modern CPUs?
6. Why is Quick Sort not stable?
7. What problem does Introsort solve?
8. What improvements does PDQSort add over classic Quick Sort?

---

## Go

Implement:

```
func QuickSort[T cmp.Ordered](arr []T)
```

Requirements:

- Start with Lomuto partition.
- Then implement Hoare partition.
- Benchmark both.
- Compare against:
    - Merge Sort
    - `slices.Sort`
- Run:

```
go test -bench=. -benchmem
```

Record:

- ns/op
- B/op
- allocs/op

---

## Python

Implement Quick Sort using:

- Lomuto partition
- Hoare partition

Compare both with:

```
sorted(...)
```

Observe:

- runtime
- recursion depth
- behavior on already sorted input

---

## C++

Implement:

```
template<typename T>
void QuickSort(std::vector<T>& arr);
```

Implement both partition strategies and compare them with:

```
std::sort(...)
```

Analyze differences in performance and behavior.

---

## Reading Assignment

Study:

1. Go's `slices.Sort` documentation.
2. The PDQSort paper (high-level overview is sufficient for now).
3. C++ `std::sort` complexity guarantees.
4. Tony Hoare's original partition algorithm.

Focus on **why** runtime libraries evolved beyond textbook Quick Sort.

---

## The Most Important Lesson

Your Quick Sort lesson isn't really about Quick Sort.

It's this:

> **A theoretically good algorithm can become a bad production algorithm when its assumptions don't match the workload.**

Examples:

```
Quick Sort
    +
better pivot selection
    +
duplicate handling
    +
small-array optimization
    +
pattern detection
    +
fallback strategy
    ↓
PDQSort / Introsort
```


