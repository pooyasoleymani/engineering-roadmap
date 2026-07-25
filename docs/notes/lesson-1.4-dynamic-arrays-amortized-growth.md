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
