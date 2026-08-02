## Q

- Why is Merge Sort called a divide-and-conquer algorithm?


## A

Because merge sort divide data to small part for sorting

---

## Q

- Why does every recursion level perform O(n) work?


## A

every recursion level can have n element for sort

---
## Q

- Why is the overall complexity O(n log n)?


## A 

binary tree have log n level and every level have n work at the end of merging
```
O(n log n)
```

---

## Q

- Why is Merge Sort stable?


## A

because if value in left after compare is equal always choose the left element first 

---
## Q

- Why does Merge Sort require O(n) auxiliary memory?


## A

it need temporary buffer for dividing 


---

## Q

- Why is Merge Sort suitable for external sorting?


## A

1. Read a chunk.
2. Sort it in memory.
3. Write it back to disk.
4. Repeat.

---
## Q

- Why is Merge Sort easy to parallelize?


## A 
merge sort can work on left in one core and right on other core.
This divide-and-conquer structure maps well to multicore processors

---
## Q

- Why doesn't Go use Merge Sort as its default slice sorting algorithm?


## A

- Allocates extra memory
- Copies elements
- Has allocation overhead

Go prioritizes:

- Low allocations
- Cache efficiency
- Fast average-case performance