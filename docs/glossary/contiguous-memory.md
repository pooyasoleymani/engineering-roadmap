

# Definition

Contiguous memory allocation is a memory management technique where a single, unbroken, sequential block of memory addresses is reserved for a data structure or process. If you ask the operating system for 100 bytes, it finds 100 bytes directly next to each other in RAM—no gaps, no jumping around.

# Problem

It primarily solves the overhead and inefficiency of scattered data. When data is fragmented across your RAM (like in a linked list), the CPU has to work much harder to chase pointers and fetch the next piece of information. By placing data side-by-side, contiguous memory solves the problem of **spatial locality**—when the CPU fetches one piece of data, it naturally grabs the adjacent data into its high-speed` L1/L2 `cache, drastically reducing slow trips to main memory.

# Use cases

- **Arrays and Vectors:** The backbone of most standard library collections (like `std::vector` in C++ or standard arrays in Java/Python).
- **High-Performance Computing & Gaming:** Where every microsecond counts. Cache misses from scattered memory can tank your frame rate or simulation speed.
- **Hardware Interfacing:** Graphics cards, network interfaces, and audio processors often require contiguous chunks to read data directly via Direct Memory Access (DMA) without CPU intervention.
- **Memory Arenas:** Custom allocators in complex systems often grab one massive contiguous block from the OS upfront and manage it internally to speed up runtime allocations.


# Advantages

- **O(1) Random Access:** Because the memory is unbroken, you can instantly calculate the exact address of any element using simple math: `Address = Start_Address + (Index * Element_Size)`.
- **Cache Friendliness:** Modern CPUs are heavily optimized for sequential reads. Contiguous memory results in significantly fewer cache misses.
- **Minimal Overhead:** There is no need to store extra pointers or metadata alongside your data to find the next element.

# Disadvantages

- **External Fragmentation:** Over time, as blocks of varying sizes are allocated and freed, your memory gets "Swiss-cheesed." You might have plenty of total free RAM, but if it's chopped up into small scattered pieces, a request for a large contiguous block will fail.
- **Expensive Resizing:** If your data grows beyond its allocated block and the adjacent memory is already taken, the system has to find a brand new, larger block elsewhere and copy _everything_ over (an $O(n)$ operation).
- **Pre-allocation Waste:** To avoid that expensive resizing, we often allocate more contiguous space than we currently need (capacity vs. size), which locks up memory that other programs can't use.

# Examples

- **C/C++ Arrays:** `int numbers[5] = {1, 2, 3, 4, 5};` guarantees those five integers sit right next to each other.
- **Strings:** Under the hood in most languages, the characters of a string are stored sequentially in memory.
- **Bitmap Images:** A raw image file is typically represented in memory as a massive, single-dimensional contiguous array of pixel color values.

---

## Related Topics

- Spatial Locality [[spatial-locality]]
- Temporal Locality [[temporal-locality]]
- Array [[array]]
- Cache miss [[cache-miss]]
- Cache line [[cache-line]]
- Pointer chasing [[docs/glossary/pointer-chasing|pointer-chasing]]