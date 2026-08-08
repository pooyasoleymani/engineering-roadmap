from typing import Iterable, Any, Sequence

def _partition(data: Iterable, low: int, high: int) -> int:
    pivot = data[high]
    
    i = low - 1 # index
    
    for j in range(low, high):
        if data[j] < pivot:
            i += 1
            _swap(data, i, j)
    
    _swap(data, i + 1, high)
    return 

def quick_sort(data: Iterable) -> None:
    pass

def _swap(data: Sequence[Any], i: int, j: int) -> None:
    data[i], data[j] = data[j], data[i]