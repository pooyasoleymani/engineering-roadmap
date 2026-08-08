


## Q 

- Why partitioning the core operation of Quick Sort ? 


## A

Because Without partitioning Quick Sort just recursion



---

## Q 

- Compare Lomuto and Hoare partition schemes. What are the trade-offs?


## A


|        | Hoare        | Lomuto             |
| ------ | ------------ | ------------------ |
| swaps  | If necessary | Very swaps         |
| memory | If necessary | Very memory writes |


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

If we have sorted elements randomize pivot reduce chance of select last element as pivot


---

## Q

- Why is Quick Sort generally faster than Merge Sort on modern CPUs?


## A

Quick Sort use contiguous memory then:

1. Better Cache locality
2. Better hardware prefetching


---

## Q 

- Why is Quick Sort not stable?


## A

Quick sort swap in every 