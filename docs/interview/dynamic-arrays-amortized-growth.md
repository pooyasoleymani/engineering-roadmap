
## Q

- Why isn't `append()` always O(1)?

## A

1. Allocate a new array.
2. Copy all existing elements.
3. Append the new element.
4. The old array becomes eligible for GC (Go) or deallocation (C++).


---
## Q

- Why is the average append cost still O(1)?


## A

Although some appends cost O(n) because of resizing, the total cost of many appends grows linearly with the number of inserted elements. Therefore, the amortized cost per append is O(1).



---
## Q

- Why doesn't Go always double slice capacity?


## A

Go have runtime grows strategy for small array `x2` and for big array `~x1.25` for reduce memory waste.


---
## Q

- Why does CPython use a smaller growth factor?


## A

Python don't use doubling for lower memory fragmentation. (Go and C++ have  shrink to fit)
CPython also tries to keep the allocator happy.
Growing by 12.5% instead of 100% makes it easier to reuse memory blocks.

---
## Q

- What is memory fragmentation?


## A 

There is enough total free memory, but not enough contiguous memory.


---
## Q

- Why should you preallocate when possible?


## A

because preallocate memory no need to reallocation and effect on performance .
- allocations
- copies
- GC pressure
- cache disruption

---
## Q

- When is doubling the best strategy?


## A

doubling in small array is best strategy and Fast growth.
- size unknown
- rapid growth
- throughput more important than memory

---
## Q

- When is doubling a poor strategy?


## A

In big size array  and need to lower  memory usage.
- very large slices
- memory constrained systems
- embedded systems



---
