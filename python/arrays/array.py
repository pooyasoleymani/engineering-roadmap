from typing import List, Sequence, Optional, Any


class DynamicArray[T](Sequence[T]):
    def __init__(self, items: Optional[Sequence[T]]=None , *args) -> None:
        self.items: Sequence[T] = []
        if (items is not None) and len(self.items) < 8:
            self.capacity = 8
        for arg in args:
            self.items.append(arg)

        self.size = len(self.items)
    
    def append(self, item: T):
        if self.size == self.capacity:
            self.capacity *=2
        self.items.append(item)

    def __repr__(self):
        return f"{self.__class__.__name__}({self.items})"

if __name__ == "__main__":
    arr = DynamicArray(1, 2, 3, 4)
    print(repr(arr))