---
Date: 2026-07-04
---


---

# Objectives

- Understand CPU memory hierarchy
- Learn cache lines
- Learn locality of reference
- Understand pointer chasing
- Understand stack vs heap
- Explain why contiguous memory is faster

---

# Summary

Modern processors are much faster than RAM.

To reduce latency, CPUs contain several cache levels (L1, L2, L3).

Programs that access memory sequentially achieve much higher performance because they maximize cache utilization and allow hardware prefetching.

Understanding memory layout is essential for designing high-performance systems.

---

# Important Concepts

## Memory Hierarchy

Registers

↓

L1 Cache

↓

L2 Cache

↓

L3 Cache

↓

RAM

↓

SSD

↓

HDD

The lower the level, the larger the storage but the higher the latency.

---

## Cache Line

A cache line is the smallest block of memory transferred between RAM and CPU cache.

Typical size

64 Bytes

Reading one integer often loads several nearby integers into cache.

---

## Spatial Locality

Programs perform better when accessing nearby memory locations sequentially.

Example

Array traversal.

---

## Temporal Locality

Recently accessed data is likely to be accessed again.

The CPU keeps frequently used data in cache.

---

## Pointer Chasing

Linked lists require following pointers scattered across memory.

Each pointer dereference may trigger another cache miss.

---

## Stack Memory

Characteristics

- Automatic allocation
- Automatic deallocation
- Very fast
- Function-local variables

---

## Heap Memory

Characteristics

- Dynamic allocation
- Managed by allocator or GC
- Slower than stack
- Larger capacity

---

# Trade-offs Learned

#### Array

Advantages

- Cache friendly
- Sequential access
- Hardware prefetching
- Better throughput

Disadvantages

- Expensive insertions in middle
- Fixed contiguous allocation

---

#### Linked List

Advantages

- Fast insertion
- Dynamic size

Disadvantages

- Pointer chasing
- Cache misses
- Poor locality
- Extra memory per node

---

# Real-World Examples

**Networking**

Ring buffers outperform linked lists for packet processing.

**Databases**

B-Trees store keys contiguously inside nodes to reduce disk and cache misses.

**Go**

Slices are fast because their backing arrays are contiguous.

---

# Common Mistakes

- Thinking RAM is as fast as CPU.
- Believing O(n) algorithms always perform similarly.
- Ignoring cache locality.
- Allocating on the heap unnecessarily.
- Assuming linked lists are faster because insertion is O(1).

---

# What Surprised Me

CPU performance is often limited by memory access rather than computation.

Sequential memory access can significantly outperform random access due to cache behavior.

---

# Questions

- How does Go decide stack vs heap allocation?
- What is escape analysis?
- How does a CPU prefetch data? [[cpu-prefetch]]