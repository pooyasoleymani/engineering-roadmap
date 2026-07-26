# Ring Buffer

---
# Definition

A ring buffer (also called a circular array or circular queue) is a standard, fixed-size contiguous array that mathematically treats its memory as if the ends are tied together in a circle. It manages data using two indices (or pointers):

1. **Tail (Write):** Tracks where the next item should be inserted.    
2. **Head (Read):** Tracks where the next item should be removed.


When either pointer reaches the very end of the array, it doesn't stop or crash. Instead, it uses modulo arithmetic to immediately wrap back around to index `0`.

# Problem

It solves the Queue (First-In, First-Out) bottleneck. As we discussed, if you build a queue out of a standard dynamic array, removing the front item forces you to shift every other element to the left—a massive `O(n)` penalty.

If you build a queue out of a linked list, you get `O(1)` insertions and deletions, but you ruin your CPU cache performance (spatial locality) and waste memory on pointers. The ring buffer solves both problems simultaneously: it gives you `O(1)` queue operations while keeping all your data in a cache-friendly, tightly packed contiguous block.

# Use cases

- **Media Streaming:** Audio and video players use ring buffers to **preload** a few seconds of media. As you watch a frame, the read pointer moves forward, and the network downloader writes new frames into the empty space behind it.    
- **Producer-Consumer Multithreading:** Safely passing data between a thread that generates work (like a network listener) and a thread that processes it, without constantly locking and reallocating memory.
- **Hardware I/O:** Reading keystrokes from a keyboard or data packets from a network card. The hardware writes to the tail, and the OS reads from the head.
- **Logging Systems:** Maintaining a "rolling log" of the last 1,000 events. When it fills up, it just starts overwriting the oldest entries automatically.

# Advantages

- **O(1) Enqueue and Dequeue:** Adding and removing items is instantaneous. No elements ever have to be shifted.
- **Maximum Cache Efficiency:** Because it is just a standard array under the hood, the CPU can preload the sequence into `L1/L2` cache, making reads and writes incredibly fast.
- **Zero Allocation Overhead:** Once the buffer is created at startup, you never have to ask the OS for memory again. The pointers just run laps around the same block of RAM.

# Disadvantages

- **Fixed Capacity:** A true ring buffer cannot grow. If the Tail pointer catches up and laps the Head pointer, the buffer is full. You either have to drop the new data, freeze the writer until space opens up, or overwrite the oldest unread data (which is fine for a rolling log, but disastrous for a bank transaction queue).
- **Complex Resizing:** If you _do_ decide to make a dynamic ring buffer that resizes when full, copying the data to a new array is tricky. You can't just copy it linearly; you have to "unroll" the wrapped data so the Head starts at index `0` in the new array.
- **Slightly Complex Logic:** Calculating the indices requires division/modulo operations, and distinguishing between a completely full buffer and a completely empty buffer (since Head and Tail point to the same spot in both cases) requires an extra counter variable or a sacrificed array slot.

# Examples

- **C++:** `std::deque` is often implemented under the hood as a series of chunked circular arrays.
- **POSIX/Linux:** Network sockets heavily rely on circular buffers for incoming and outgoing packet queues.
- **Game Engines:** Storing fixed-size queues of recent input events or particle system data.