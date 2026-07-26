---
Date: 2026-07-18
tags:
  - software_engineering
Next: "[[lesson-1.5-binary-search]]"
---
---
# Objectives

- Why static arrays are insufficient
- How dynamic arrays grow
- Amortized O(1)
- Reallocation
- Growth strategies
- Memory fragmentation
- Why Go, Python, Java and C++ all implement growth differently
- Why `append()` is not always O(1)

---
# Summary

Dynamic arrays solve the primary limitation of static arrays: **fixed capacity**. Instead of rejecting new elements when the array is full, a dynamic array allocates a larger block of contiguous memory, copies the existing elements, and continues operating transparently.

Unlike a static array, a dynamic array maintains two separate values:

- **Length** – the number of valid elements currently stored.    
- **Capacity** – the total number of elements that can be stored before another allocation is required.


Most append operations simply place the new element into unused capacity, making them **O(1)**. Only when the array is full does the runtime perform an expensive reallocation and copy, which costs **O(n)**.

To avoid reallocating on every append, dynamic arrays grow **geometrically** (for example, doubling their capacity or using an adaptive growth factor). This dramatically reduces the number of reallocations over time, making the **average (amortized)** cost of appending an element **O(1)**.

Modern programming languages all implement some form of dynamic array:

| Language | Dynamic Array Implementation |
| -------- | ---------------------------- |
| Go       | Slice                        |
| Python   | `list`                       |
| C++      | `std::vector`                |
| Java     | `ArrayList`                  |
| Rust     | `Vec<T>`                     |

Although these languages expose different APIs, they all rely on the same fundamental design:

```text
Pointer
Length
Capacity
```

Dynamic arrays represent an engineering trade-off between **performance** and **memory usage**. Allocating extra capacity wastes some memory, but it significantly reduces expensive reallocations and copying. This trade-off enables high-performance append operations while preserving contiguous memory, excellent cache locality, and efficient CPU utilization.

Understanding dynamic arrays is essential because they form the foundation of many modern data structures and runtime implementations, including Go slices, Python lists, C++ vectors, ring buffers, stacks, queues, and numerous high-performance systems.


---
# Why Growth Strategy is important ?

## Problem

Imagine:

```
Capacity = 4

+----+----+----+----+
| 10 | 20 | 30 | 40 |
+----+----+----+----+
```

Now execute

```
Push(50)
```

Where should 50 go?

There is no free memory.

## Solution

Move every element into a larger array.

```
Old

+----+----+----+----+
|10|20|30|40|
+----+----+----+----+

↓

New

+----+----+----+----+----+----+----+----+
|10|20|30|40|50|
+----+----+----+----+----+----+----+----+
```

This is exactly what most dynamic arrays do.

---

## Reallocation

Reallocation consists of three steps.

## Step 1

Allocate new memory.

```
Old

Capacity = 4
```

↓

```
New

Capacity = 8
```


## Step 2

Copy data.

```
10

↓

10
```

```
20

↓

20
```

...


## Step 3

Free the old memory.

Now the new array becomes the backing storage.


Complexity:

```text
O(n)
```


---
# Definition 
A **dynamic array** is a contiguous block of memory that automatically grows when its capacity is exhausted. Unlike a static array, a dynamic array separates: - Logical Size - Allocated Capacity Example: 

```text
Length = 3 
Capacity = 8 

+----+----+----+----+----+----+----+----+ 
| 10 | 20 | 30 | | | | | | 
+----+----+----+----+----+----+----+----+
```

Only the first three elements are valid.

---

# Why is append() O(1) ?

Suppose we append one million values.

Most appends cost:

```
1
```

Only a few cost:

```
1000

5000

20000
```

Average:

```
O(1)
```

This is **amortized analysis**.


## Visualizing Growth

Imagine capacities:

```
1

↓

2

↓

4

↓

8

↓

16

↓

32

↓

64
```

Number of reallocations:

```
log₂(n)
```

Only about 20 reallocations are needed to reach roughly one million elements.

---
## Formal Intuition

Suppose we insert:

```
1024 elements
```

How many copies happen?

```
1

2

4

8

16

32

64

128

256

512
```

Total:

```
1023
```

Interesting.

To insert **1024** elements,

we copy only about **1023** elements in total.

Average cost remains constant.


---
