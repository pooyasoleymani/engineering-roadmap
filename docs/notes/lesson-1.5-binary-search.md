---
Date: 2026-07-26
tags:
  - software_engineering
Next:
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
