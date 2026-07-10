from typing import List, Sequence, Optional, Any


class DynamicArray[T]:
    def __init__(items: Optional[Sequence[T]]=None) -> None:
        self.items = items or []
        self.capacity = 8
        self.size = len(self.item)
    
    def append(item: T):
        if self.size == self.capacity:
            self.capacity *=2
        self.items.append(item)