from typing import List, Any

def merge(data: List[Any], left: int, mid: int, rigth: int) -> None:
    nl = mid - left + 1
    nr = rigth - mid
    
    # create temp arrays
    L: List[Any] = [0] * nl
    R: List[Any] = [0] * nr
    
    # copy data to temp arrays
    for i in range(nl):
        L[i] = data[left+i]
    
    for j in range(nr):
        R[j] = data[mid+j+1]
    
    i, j, k, = 0, 0, left
    
    # merge the temp arrays back
    while i< nl and j < nr:
        if L[i] <= R[j]:
            data[k] = L[i]
        else:
            