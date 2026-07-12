```mermaid
flowchart TD

CPU --> Registers
Registers --> L1
L1 --> L2
L2 --> L3
L3 --> RAM
RAM --> SSD
SSD --> HDD
```