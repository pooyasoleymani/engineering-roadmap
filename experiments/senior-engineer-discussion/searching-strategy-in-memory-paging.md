---
Related: "[[lesson-1.5-binary-search]]"
---
---
# Senior Engineer Challenge

Imagine you're implementing a storage engine for a database.

Each disk page contains **512 sorted keys**.

When searching for a key, would you:

1. Perform a linear scan of the page?
2. Use binary search?
3. Use another strategy?

Your answer should consider:

- CPU cache behavior
- Branch prediction
- Number of comparisons
- Typical B-tree page sizes
- Overall throughput


---

## A

## Suppose a page contains:

```
512 keys
```

Each key:

```
8 bytes
```

Total:

```
4096 bytes
```

Interesting...

That's exactly one memory page.

---

Now imagine searching inside that page.

### Binary Search

Comparisons:

```
log₂(512)

=

9
```

Very few comparisons.

Great.

But...

Every comparison jumps to a different location:

```
256

↓

128

↓

192

↓

160

↓

176
```

The access pattern is unpredictable.

---

The CPU struggles with:

- branch prediction
- hardware prefetching

---

### Linear Scan

Comparisons:

```
512
```

Terrible?

Not necessarily.

Memory access:

```
0

1

2

3

4

5
```

Perfectly sequential.

CPU prefetches everything.

Almost zero cache misses.

---

Which wins?

It depends.

---

## Real Database Systems

Many modern storage engines actually use:

### Small node

```
Linear Scan
```

Why?

Sequential memory is incredibly fast.

---

### Medium node

```
Binary Search
```

---

### Large node

Sometimes neither.

Instead they use

- SIMD search
- interpolation search
- branchless binary search
- cache-aware search

---

## Why?

Suppose each comparison causes:

```
Branch misprediction

↓

Pipeline flush

↓

15–20 CPU cycles
```

Nine branch mispredictions can cost more than scanning 30–40 integers sequentially.

Modern CPUs are weird.

Sometimes

```
O(n)
```

beats

```
O(log n)
```

for small `n`.

---

## PostgreSQL

PostgreSQL generally uses binary search inside B-tree pages because the page size and key count make it worthwhile.

---

## High-frequency Trading

Some trading systems use:

- branchless binary search
- SIMD comparisons

because every nanosecond matters.

---

# The Engineering Lesson

This is probably the most important sentence today:

> **Big O predicts how performance scales as input grows. It does not predict which implementation is fastest for a specific input size on a specific CPU.**

That single sentence separates algorithm theory from performance engineering.

---