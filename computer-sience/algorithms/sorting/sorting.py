from typing import Sequence, Optional, Any

def bubble_sort(data: Optional[Sequence[Any]]) -> None:
    length = len(data)
    for i in range(1, length):
        for j in range(0, length-1):
            if data[j] > data[j+1]:
                data[j], data[j+1] = data[j+1], data[j]
                


def selection_sort(data: Optional[Sequence[Any]]) -> None:
    n = len(data)
    for i in range(n-1):
        min_index = i
        for j in range(i+1, n):
            if data[min_index] > data[j]:
                min_index = j
        data[min_index], data[i] = data[i], data[min_index]

def insertion_sort(data: Optional[Sequence[Any]]) -> None:
    n = len(data)
    for i in range(n-1):
        p = i
        while p > 0 and data[p+1] < data[p]:
            data[p], data[p+1] = data[p+1], data[p]
            p -=1
            

if __name__ == "__main__":
    data = [1,23,56,45,23,5,6,7,2,22,33,33,55]
    insertion_sort(data=data)
    print(data)