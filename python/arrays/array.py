from __future__ import annotations
from typing import List, Sequence, Optional, Any, TypeVar


T = TypeVar("T")

class DynamicArray(Sequence[T]):
    def __init__(self, items: Optional[Sequence[T]]=None , *args) -> None:
        self.items: Sequence[T] = []
        if (items is not None) and len(self.items) < 8:
            self.capacity = 8
        for i,arg in enumerate(args):
            self.append(arg)
        self.size = len(self.items)

    def insert(self, index: int, item: T) -> None:
        self._shift_right(index)
        self.items[index] = item
    
    
    
    def append(self, item: T):
        if self.size == self.capacity:
            self.capacity *=2
        self.items[self.size] = item

    def _shift_right(self, index: int) -> None:
        if self.size == self.capacity:
            self.capacity *= 2
        i = self.size
        while i > index:
            self.items[i+1] = self.items[i]
            --i

    def __contains__(self, value):
            return value in self.items

    def __repr__(self):
        return f"{self.__class__.__name__}({self.items})"

if __name__ == "__main__":
    arr = DynamicArray(1, 2, 3, 4)
    print(repr(arr))