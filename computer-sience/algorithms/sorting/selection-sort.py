from typing import Sequence, Optional, Any

def selection_sort(data: Optional[Sequence[Any]]) -> None:
    n = len(data)
    for i in range(n-1):
        min_index = i
        for j in range(i+1, n):
            if data[min_index] > data[j]:
                min_index = j
        data[min_index], data[i] = data[i], data[min_index]

if __name__ == "__main__":
    data = [1,23,56,45,23,5,6,7,2,22,33,33,55]
    selection_sort(data=data)
    print(data)