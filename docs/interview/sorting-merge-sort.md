## Q

- Why is Merge Sort called a divide-and-conquer algorithm?


## A

Merge Sort recursively divides the problem into smaller `subproblems` until each contains one element (already sorted), then combines the sorted `subproblems` by merging them into larger sorted sequences.

---

## Q

- Why does every recursion level perform O(n) work?


## A

At one recursion level:

```
8

↓

4 + 4

↓

2 + 2 + 2 + 2

↓

1 + 1 + 1 + 1 + ...
```

Every element is copied exactly once during merging.

Total work per level:

```
O(n)
```

---
## Q

- Why is the overall complexity O(n log n)?


## A 

Binary tree have `log n` level and every level have `n work` at the end of merging

```
O(n log n)
```

---

## Q

- Why is Merge Sort stable?


## A

Because if value in left after compare is equal always choose the **left element** first 

The rule

```
Equal

↓

Take left first
```

creates stability.

---
## Q

- Why does Merge Sort require O(n) auxiliary memory?


## A

Merge Sort allocates temporary arrays to merge two sorted halves into one sorted sequence.


---

## Q

- Why is Merge Sort suitable for external sorting?


## A

1. Read a chunk.
2. Sort it in memory.
3. Write it back to disk.
4. Repeat.

Because the complete dataset does **not** fit in RAM.

---
## Q

- Why is Merge Sort easy to parallelize?


## A 
Merge sort can work on left in one core and right on other core.
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


---
## Senior Engineer Challenge

You're implementing a backend service that must sort **50 million transaction records** every night.

The records are stored on disk, and the server has:

- 32 GB RAM
- 2 TB of transaction data

Questions:

1. Can an in-memory Merge Sort solve this problem?
2. If not, what algorithmic strategy would you use?
3. Why is sequential disk I/O preferred over random disk access in this scenario?
4. How would you take advantage of an 8-core CPU during the sort?


## A

1. No because Even before Merge Sort allocates its `temporary buffer`, the dataset cannot fit into a machine with 32 GB RAM. So an `in-memory algorithm` is impossible regardless of which sorting algorithm you choose.
2. We can use `external merge sort ` .

### HDD

Random seek

```
10 ms
```

Sequential read

```
200 MB/s
```

Huge difference.

---

### SSD

Random access is much better than HDD,

but sequential throughput is still significantly higher.

---

So the answer should be:

> External Merge Sort performs large sequential reads and writes, minimizing expensive random disk seeks and maximizing disk throughput.



4. Each chunk can be sorted **independently** on separate CPU cores. The final merge can also be parallelized using multi-way merge techniques.