---
Date: 2026-07-26
tags:
  - software_engineering
Next:
---
---
# Fundamentals of Sorting


## Objective

- What sorting actually is
- Why sorting is one of the most important algorithms
- Properties of sorting algorithms
- Stable vs unstable sorting
- In-place vs out-of-place sorting
- Comparison vs non-comparison sorting
- Why no comparison sort can beat **O(n log n)**
- Why different languages choose different algorithms

---

## Definition

Sorting is the process of arranging elements according to an ordering relation.

Example

Before

```
8 2 9 1 4
```

After

```
1 2 4 8 9
```

---
## Problem

Suppose you have:

```
100 million users
```

You want:

- search quickly
- remove duplicates
- merge datasets
- build indexes
- perform range queries

Sorting makes all of these easier.

Many algorithms first sort the data and then solve the actual problem.


---
## Real Systems That Depend on Sorting

Sorting is fundamental in:

- Database query execution
- Search engines
- Log aggregation
- Distributed systems
- MapReduce
- Git
- Linux kernel
- Memory allocators
- Compilers

Sorting is one of the most optimized algorithms in software engineering.


---

## Properties of a Sorting Algorithm

When evaluating a sorting algorithm, don't ask only:

> "How fast is it?"

Ask:

- Is it stable?
- Is it in-place?
- Worst-case complexity?
- Average-case complexity?
- Extra memory?
- Cache friendly?
- Parallelizable?
- Adaptive?
- Good for nearly sorted data?

A senior engineer evaluates algorithms across multiple dimensions.


---

## Stable Sort 

Suppose we sort employees by salary.

Before:

| Name    | Salary |
| ------- | ------ |
| Alice   | 5000   |
| Bob     | 4000   |
| Charlie | 5000   |

Sorted:

| Name    | Salary |
| ------- | ------ |
| Bob     | 4000   |
| Alice   | 5000   |
| Charlie | 5000   |

Notice:

Alice remains before Charlie.

The relative order of equal keys didn't change.

That is a **stable sort**.



#### Why Stability Matters

Imagine:

First sort by

```
Department
```

Then sort by

```
Salary
```

If the second sort is stable,

employees with equal salary remain grouped by department.

Database engines rely on this behavior.

---

## Unstable Sort

Suppose after sorting:

|Name|Salary|
|---|---|
|Bob|4000|
|Charlie|5000|
|Alice|5000|

Still sorted?

Yes.

But

Alice and Charlie swapped.

That's unstable.


---
## In-place Sorting

An algorithm is **in-place** if it uses only a small amount of additional memory.

Example:

```
Original Array

↓

Same Array Sorted
```

Examples:

- Heap Sort
- Quick Sort (mostly)
- Selection Sort

---
## Out-of-place Sorting

Merge Sort

```
Original

↓

Temporary Array

↓

Merged Result
```

Needs additional memory.

Trade-off:

More memory

↓

Better stability

↓

Often simpler merging


---

##  Comparison Sorting

Algorithms like:

- Merge Sort
- Quick Sort
- Heap Sort

work by asking questions like:

```
Is A < B?
```

Everything they know comes from comparisons.


---

## Why O(n log n) is the Limit

This is one of the most important theoretical results in computer science.

Imagine sorting:

```
3 elements
```

Possible orders:

```
ABC
ACB
BAC
BCA
CAB
CBA
```

There are

```
3!

=

6
```

possible outcomes.

Every comparison splits possibilities.

Binary decisions create a decision tree.

To distinguish all outcomes,

the tree must have at least

```
n!
```

leaves.

The minimum height is

```
log₂(n!)
```

Using Stirling's approximation:

```
log₂(n!)

≈

n log₂ n
```

Therefore:

> **No comparison-based sorting algorithm can have a better worst-case complexity than O(n log n).**

This is a mathematical proof, not an implementation limitation.

---

## Then How Can Counting Sort Be O(n)?

Because it doesn't compare elements.

Instead it uses knowledge about the data.

Examples:

- Counting Sort
- Radix Sort
- Bucket Sort

These algorithms trade memory or assumptions about the input for better time complexity.

We'll study them later.