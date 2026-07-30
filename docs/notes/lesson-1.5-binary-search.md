---
Date: 2026-07-26
tags:
  - software_engineering
Next: "[[lesson-1.6.1-sorting]]"
---
---

# Objective

- Binary search from first principles
- Loop invariants
- Correctness proofs
- Integer overflow
- Lower Bound
- Upper Bound
- Equal Range
- Monotonic predicates
- Binary search on answers
- Floating-point binary search
- Branch prediction
- Cache behavior
- How B-Trees search nodes
- Why databases use binary search

---
## Problem

Suppose we have

```
Index

0 1 2 3 4 5 6 7

Value

2 5 8 11 17 21 40 55
```

Find

```
21
```

How many comparisons?

Linear search

Worst case

```
8
```

Now imagine

```
1 Billion
```

records.

Linear search

Worst case

```
1 Billion comparisons
```

Impossible for many workloads.


---

## Idea

Instead of finding the answer immediately, we **eliminate impossible regions**.
Each comparison halves the search space.

---


## Complexity


```
O(log₂ n)
```


---

## The Real Secret

Binary search is about **maintaining an invariant**.


---

## What is an Invariant?

An invariant is a condition that is always true throughout the algorithm.

Example

```
If the target exists,

it is inside

[low, high]
```

Every iteration must preserve this property.

If you violate the invariant, your algorithm is incorrect—even if it seems to work for many cases.

This mindset is fundamental in algorithm design.


---
## Example

Initial state

```
low = 0
high = 7
```

Invariant

```
Target ∈ [low, high]
```

Middle

```
mid = 3
```

Value

```
11
```

Target

```
21
```

Update

```
low = mid + 1
```

Why not

```
low = mid
```

?

Because we already know

```
arr[mid] != target
```

Keeping `mid` would violate progress and may cause an infinite loop.


---
## Overflow

Classic implementation

```
mid := (low + high) / 2
```

Looks fine.

Imagine

```
low = 2,000,000,000
high = 2,100,000,000
```

On a 32-bit integer

```
low + high
```

overflows.

The safe version

```
mid := low + (high-low)/2
```

This is now the standard implementation across many languages.


---

## Lower Bound

Suppose

```
1 2 2 2 2 5 9
```

Find

```
2
```

Which one?

There are four.

Lower Bound means

> Find the **first** element that is **greater than or equal to** the target.

Result

```
Index 1
```


---

## Upper Bound

Upper Bound means

> Find the **first** element that is **greater than** the target.

Same array

```
1 2 2 2 2 5 9
```

Upper Bound

```
Index 5
```


Notice something interesting.

Neither function necessarily returns the exact element.

They return a **position**.

That position is often more useful.

---

## Why Databases Love Lower Bound

Suppose PostgreSQL stores

```
100
101
105
109
120
125
130
```

Query

```sql
SELECT *
WHERE id >= 109;
```

The storage engine doesn't search for 109 directly.

Instead it performs

```
LowerBound(109)
```

and starts scanning from there.

This is why *lower_bound* is fundamental to [[b-tree]] indexes.


---
## Binary Search on Answers

This surprises many engineers.

Suppose you need to answer:

> What is the minimum bandwidth required to finish all uploads in one hour?

Search space

```
Bandwidth

1 Mbps

↓

10 Gbps
```

We test:

```
500 Mbps

Enough?

YES
```

Then

```
250 Mbps

Enough?

NO
```

The answer space is **monotonic**.

---
## Monotonic Predicate

Suppose

```
Bandwidth

100 Mbps

❌

200 Mbps

❌

300 Mbps

✅

400 Mbps

✅

500 Mbps

✅
```

Notice

```
False False False True True True
```

Once it becomes true,

it stays true.

Binary search works whenever the predicate is *monotonic*—not just on sorted arrays.


---
# Real Engineering Examples

Find minimum:

- Number of servers
- Buffer size
- Cache size
- Timeout
- Maximum packet size
- Compression level

If the answer space is monotonic, binary search applies.

---
# Branch Prediction

Modern CPUs try to guess:

```
if target < arr[mid]
```

before the comparison completes.

Random searches are harder to predict.
*Mispredictions* flush the CPU pipeline and reduce performance.
This is one reason *branchless* search techniques exist in *high-performance* code.

---
# Binary Search Inside a B-Tree

Suppose a [[b-tree]] node stores

```
5
11
19
22
35
48
60
```

To decide which child to visit,

the database performs a binary search within the node.

This reduces comparisons while keeping the tree shallow.

So binary search is used _inside_ larger data structures, not just on standalone arrays.


---

# Common Mistakes

1. Incorrect loop condition (`low < high` vs. `low <= high`).
2. Overflow when computing `mid`.
3. Infinite loops due to not advancing `low` or `high`.
4. Failing on empty slices.
5. Returning an arbitrary duplicate instead of the first or last occurrence.
6. Not considering integer overflow on 32-bit platforms.
7. Ignoring branch prediction effects in hot code paths.

---

# Trade-offs

| Algorithm     | Lookup   | Ordered Data | Range Queries |
| ------------- | -------- | ------------ | ------------- |
| Linear Search | O(n)     | No           | No            |
| Binary Search | O(log n) | Yes          | Yes           |
| Hash Table    | O(1) avg | No           | No            |
| B-tree        | O(log n) | Yes          | Excellent     |

This table should immediately suggest why databases favor B-trees over hash tables for general-purpose indexing.

---

# Summery

1. Binary search is about eliminating impossible regions
2. Loop invariants guarantee correctness
3.  Progress is mandatory:  `low = mid` is wrong  and causing an infinite loop  use `low = mid + 1`
4. Use the overflow-safe midpoint `mid := low + (high-low)/2`
5. `lower_bound` and `upper_bound` are more useful than ordinary binary search
6. Binary search works on monotonic predicates
7. Big O isn't the whole story
8. Binary search is everywhere