from typing import List, Sequence, Optional


class DynamicArray:
    def __init__(items: Optional[Sequence]=None) -> None:
        self.items = items or []
        self.capacity = None
        self.size = len