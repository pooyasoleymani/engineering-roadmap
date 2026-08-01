from typing import List, Optional, Sequence, Any

def bubble_sort(data: Optional[Sequence[Any]]) -> None:
    length = len(data)
    for i in range(1, length):
        for j in range(0, length-2):
            if data[j] > data[j+1]:
                data[j], data[j+1] = data[j+1], data[j]
                


if __name__ == "__main__":
    data = [1,23,56,45,23,5,6,7,2,22,33,33,55]
    bubble_sort(data=data)
    print(data)