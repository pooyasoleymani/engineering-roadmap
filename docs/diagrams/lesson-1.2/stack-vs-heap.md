```mermaid
flowchart LR

Function --> Stack
Function --> Heap

Stack -->|"Automatic"| Return
Heap -->|"Garbage Collector"| Memory
```