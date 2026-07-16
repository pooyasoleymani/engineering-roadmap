# Array Note

## Q

 Why is indexing O(1) ?

## A

Because of Address calculation CPU can calculate address of every index.



## Q

Why are arrays contiguous?

## A

- constant-time indexing
- CPU prefetching
- cache locality
- minimal metadata
- simple address calculation


## Q

Why are Python lists not linked lists?

## A

Because it have capacity and use index to access the elements without iteration


## Q

Why is insertion in the middle O(n)?

## A
Array for insert in middle need to free space for this reason shift all middle elements and after that put element on free space


## Q

Why does the CPU prefer arrays?

## A

Because CPU can predict address of every index of array and can use cache line , less memory overhead, cache miss , no need pointer chasing, and cache locality


## Q

When would you deliberately choose a linked list?

## A

- implementing intrusive lists inside kernels
- LRU cache implementations
- constant-time insertion/removal when the node is already known
- memory allocators
- some lock-free algorithms

Even the Linux kernel uses linked lists selectively because it controls memory layout carefully.


## Q

Why do most standard libraries implement dynamic arrays?

## A

because array close to hardware and its cache friendly, and use hardware prefetching 


---
## Mentor Challenge

Why doesn't Go expose the capacity of an array, but it does expose the capacity of a slice?

## A

Arrays have a compile-time fixed size, so "capacity" is identical to "length" and never changes. Slices are dynamic views over arrays, so Go exposes capacity because it determines whether append can reuse the existing backing array or must allocate a new one.