from typing import List, TypeVar

T = TypeVar("T")

class MinHeap:
    def __init__(self, data: List[T]) -> None:
        self.data: List[T] = data or []
    
    def _swap(self ,i: int, j: int) -> None:
        self.data[i], self.data[j] = self.data[j], self.data[i]
    
    def _shift_up(self, index: int) -> None:
        parent = (index -1) // 2
        if self.data[parent] < self.data[index]:
            self._swap(parent, index)
            
    def _shift_down(self, index: int) -> None:
        left = 2 * index + 1
        rigth = 2 * index + 2
        
        if self.data[index] > self.data[left]:
            self._swap(index, left)
            return
        elif self.data[index] > self.data[rigth]:
            self._swap(index, rigth)
            return
        
    def push(self, item: T) -> None:
        pass