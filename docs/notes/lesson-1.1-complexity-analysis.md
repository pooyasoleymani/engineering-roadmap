# Lesson 1.1 - Complexity Analysis

Date:
2026-07-04

---

# Objectives

- Understand Big O notation
- Understand Big Ω and Big Θ
- Learn amortized analysis
- Learn why Big O alone is insufficient
- Introduce cache locality

---

# Summary

Algorithm complexity measures how resource usage grows as input size increases.

Big O is useful for comparing scalability, but it ignores hardware effects such as cache locality, branch prediction, memory allocation, and constant factors.

Real software performance depends on both algorithmic complexity and hardware behavior.

---

# Important Concepts

## Big O

Worst-case upper bound.

Example:

Binary Search

O(log n)

---

## Big Ω

Best-case lower bound.

Example:

Linear Search

Ω(1)

---

## Big Θ

Tight bound.

Example:

Traversing an array

Θ(n)

---

## Amortized Analysis

Some operations are occasionally expensive but inexpensive on average.

Example:

Appending to a dynamic array.

Single append

O(n)

Average append

O(1)

---

# Trade-offs Learned

Hash Table

Pros

- O(1) lookup

Cons

- Extra memory
- No ordering
- Hash computation

Binary Search

Pros

- Sorted data
- Range queries
- No hash function

Cons

- Requires sorted collection

---

# Common Mistakes

- Thinking O(1) is always faster
- Ignoring cache locality
- Ignoring memory allocations
- Ignoring constant factors

---

# What Surprised Me

Even though arrays and linked lists both have O(n) traversal, arrays are usually much faster because of contiguous memory and CPU caches.

---

# Questions

- Why are B-Trees preferred over hash tables?
- How does CPU cache actually work?
