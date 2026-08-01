## Definition
Selection sort is a simple, comparison-based sorting algorithm that splits a list into a sorted part and an unsorted part, repeatedly finding the smallest element to grow the sorted section.


## How It Works
**Find minimum:** Scan the unsorted part of the list to find the smallest value.
**Swap to front:** Swap that smallest value with the first element of the unsorted section.
**Shift boundary:** Move the boundary between the sorted and unsorted sections one position to the
right.
**Repeat:** Continue this process until the entire list is sorted.



## Performance and Features

- **Time complexity:** **O(n²)** in all cases (best, average, worst) because it always scans the entire unsorted list.
- **Space complexity:** **O(1)** because it sorts items in place without extra memory.
- **Unstable:** It can change the relative order of equal values during its swap step. 

## Core Trade-Off

- **Fewer Swaps vs. Fixed Time**: Selection sort minimizes write operations by making a maximum of **O(n)** swaps, but it sacrifices flexibility because it takes just as long to process a perfectly sorted list as it does a completely reversed one.
## Advantages

- **Minimal memory**: It operates entirely in place, using **O(1)** auxiliary space.
- **Low write operations**: It performs a maximum of `n-1` swaps, making it useful when memory write cycles are costly (like with Flash memory).
- **Simple logic**: It is straightforward to implement and debug. 
## Disadvantages

- **Inflexible performance**: Its best, worst, and average-case time complexities are all **O(n²)**.
- **No early exit**: It cannot detect if a list is already sorted, so it always runs the maximum number of comparisons.
- **Unstable sorting**: The swapping mechanic can easily shift identical items out of their original order.