
## Q

- Why isn't `append()` always O(1)?

## A

Because when reallocation occurs elements in old array copy into the new allocation array .

complexity:

```text
O(n)
``` 

---
## Q

- Why is the average append cost still O(1)?


## A

Most of operation in `append()` have O(1) and when reallocation occurs need to copy all elements this operation have O(n) in average O(1)  



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


---
## Q

- What is memory fragmentation?


## A 
memory fragmentation occurs when memory free between allocated memory and this memory can't use in contiguous memory.


---
## Q

- Why should you preallocate when possible?


## A

because preallocate memory no need to reallocation and effect on performance 

---
## Q

- When is doubling the best strategy?


## A

doubling in small array is best strategy and Fast growth

---
## Q

- When is doubling a poor strategy?


## A

In big size array  and need to lower  memory usage