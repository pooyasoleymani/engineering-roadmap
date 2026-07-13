# Array Note

## Q

 Why is indexing O(1) ?

## A

Because of Address calculation CPU can calculate address of every index.



## Q

Why are arrays contiguous?

## A

Array 


## Q

Why are Python lists not linked lists?

## A

Because it have capacity 


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
i don't know


## Q

Why do most standard libraries implement dynamic arrays?

## A
because array close to hardware and its cache friendly, and use hardware prefetching 