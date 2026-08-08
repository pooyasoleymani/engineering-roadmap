


## Q 

- Why partitioning the core operation of Quick Sort ? 


## A

Partitioning rearranges the array around a pivot so that every element on one side satisfies the ordering condition relative to the pivot, and the pivot reaches its final position. The recursive calls then solve the two remaining subproblems.



---

## Q 

- Compare Lomuto and Hoare partition schemes. What are the trade-offs?


## A


|                      | Hoare                          | Lomuto                         |
| -------------------- | ------------------------------ | ------------------------------ |
| Swaps                | Usually fewer                  | Usually more                   |
| Memory writes        | Usually fewer                  | Usually more                   |
| Implementation       | More subtle                    | Simpler                        |
| Pivot placement      | Not necessarily final position | Pivot ends in final position   |
| Typical teaching use | More advanced                  | Common textbook implementation |

---

## Q 

- Why does choosing the last element as pivot degrade performance on sorted input?


## A

If data is sorted choose last element as pivot

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

## Q

- Why does randomized pivot selection reduce the chance of worst-case behavior?


## A

It makes consistently bad partitions unlikely.

For example, a good partition might look like:

```
        pivot
       /     \
     45%     55%
```

while a bad partition looks like:

```
        pivot
       /
     99%
```

Randomization makes the sequence of bad partitions statistically unlikely.


---

## Q

- Why is Quick Sort generally faster than Merge Sort on modern CPUs?


## A

Sort often performs well because partitioning is mostly sequential and in-place, giving good cache locality and low auxiliary-memory traffic. Merge Sort requires additional memory and substantial copying during merging.


---

## Q 

- Why is Quick Sort not stable?


## A

Partitioning can move equal elements relative to each other.


---

## Q 

- What problem does Introsort solve?

## A

Introsort starts with Quick Sort:

```
Quick Sort
     ↓
monitor recursion depth
     ↓
too deep?
     ↓
Heap Sort
```

Heap Sort guarantees:

```
O(n log n)
```

So Introsort combines:

```
Quick Sort's average performance
+
Heap Sort's worst-case guarantee
```

---

## Q

- What improvements does PDQSort add over classic Quick Sort?


## A

- Better pivot selection
- Handles duplicates efficiently
- Detects nearly sorted input 
- Reduces branch mispredictions
- Avoids common Quick Sort worst cases



---

## Senior Engineer Challenge

You're building a log-processing service that sorts **100 million log entries** in memory.

Characteristics:

- 95% of the data is already sorted.
- Many entries have identical timestamps.
- Low memory overhead is required.
- Throughput is more important than latency.



## A

1. No I would prefer PDQSort because the workload is highly structured: 95% of the data is already sorted, and there are many duplicate keys. PDQSort is designed to perform well on such patterns while maintaining low memory overhead.
2. A naïve two-way Quick Sort can repeatedly produce terrible partitions.
For example:

```
5 | 5 5 5 5 5 5 5
```

then:

```
5 | 5 5 5 5 5
```

etc.

That can approach:

```
O(n²)
```

Good Quick Sort implementations therefore handle equal keys specially, often using **three-way partitioning** or related techniques.

or related techniques.
1. PDQSort :
	1. Handles duplicates efficiently
	2. Better pivot selection
	3. avoid Quick sort worst case
	4. Detects nearly sorted input 
	5. Reduces branch mispredictions
2. Stability is only useful if the original relative order of equal timestamps has semantic meaning. The presence of duplicates alone does not require a stable sort.
3. if dataset grows beyond RAM we most use external sorting strategy like merge sort.