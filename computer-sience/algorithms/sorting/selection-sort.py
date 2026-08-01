from typing import Sequence, Optional, Any

def selection_sort(data: Optional[Sequence[Any]]) -> None:
    for i in range(1, len(data)):
        min_index = i
        for j in range(0, len(data)-1):
            if data[min_index] > data[j]:
                min_index = j
        data[j], data[i] = data[i], data[j]

if __name__ == "__main__":
    data = [1,23,56,45,23,5,6,7,2,22,33,33,55]
    selection_sort(data=data)
    print(data)