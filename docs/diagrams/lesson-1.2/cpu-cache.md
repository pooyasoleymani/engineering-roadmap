```mermaid
graph LR

RAM --> CacheLine
CacheLine --> L1
L1 --> CPU

CacheLine["64 Byte Cache Line"]
```