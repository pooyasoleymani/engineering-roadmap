---
Date: 2026-08-04
tags:
  - software_engineering
Related:
Next:
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

That's why Quick Sort often wins despite the worse theoretical worst case.
