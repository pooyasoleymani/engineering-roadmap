from typing import MutableSequence

def _partition(data: MutableSequence, low: int, high: int) -> int:
    pivot = data[high]
    
    i = low - 1 # index
    
    for j in range(low, high):
        if data[j] < pivot:
            i += 1
            _swap(data, i, j)
    
    _swap(data, i + 1, high)
    # return posiotion
    return i + 1

def quick_sort(data: MutableSequence, low: int, high: int) -> None:
    if low < high:
        pos = _partition(data, low, high)
        quick_sort(data, low, pos -1)
        quick_sort(data, pos + 1, high)

def _swap(data: MutableSequence, i: int, j: int) -> None:
    data[i], data[j] = data[j], data[i]