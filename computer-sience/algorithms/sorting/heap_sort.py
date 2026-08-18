from typing import List, TypeVar

T = TypeVar("T")

class MinHeap:
    def __init__(self, data: List[T]) -> None:
        self.data: List[T]
        for item in data:
            self.push(item)    
    
    def _swap(self ,i: int, j: int) -> None:
        self.data[i], self.data[j] = self.data[j], self.data[i]
    
    def _shift_up(self, index: int) -> None:
        if index < 1:
            return
        elif parent := (index -1) // 2 and self.data[parent] < self.data[index]:
            self._swap(parent, index)
        
    def _shift_down(self, index: int) -> None:
        left = 2 * index + 1
        rigth = 2 * index + 2
        
        largest = index
        
        if len(self.data) > left and self.data[largest] > self.data[left]:
            largest - index
        elif len(self.data) > rigth and self.data[largest] > self.data[rigth]:
            largest = rigth
        
        if largest != index:
            self._swap(largest, index)
        
    def push(self, item: T) -> None:
        self.data.append(item)
        self._shift_down(len(self.data) - 1)
    
    def peek(self) -> T:
        return self.data[0]
    
    def pop(self) -> T:
        size = len(self.data)
        if size < 1:
            return
        # swap end item with minimum item
        self._swap(0, size - 1)
        # get minimum item
        minimum = self.data.pop()
        # shift down root 
        self._shift_down(size - 1)
        
        return minimum