# Definition

A dynamic array is a resizable, contiguous data structure. Under the hood, it's just a standard fixed-size array, but it is wrapped in logic that tracks two things:

1. **Size:** The number of elements actually currently stored.
2. **Capacity:** The maximum number of elements the current underlying block of memory can hold.

When you try to add an element and the `size` equals the `capacity`, the dynamic array automatically allocates a larger block of memory behind the scenes.

# Problem

Standard arrays are rigid. To create one, you must tell the operating system exactly how much memory you need upfront (e.g., `int myArray[10]`). But in the real world, you rarely know exactly how much data you are going to process. If you guess too low, your program crashes or you can't load all the data. If you guess too high, you waste valuable RAM. Dynamic arrays solve this "fixed-size" constraint without losing the benefits of contiguous memory.

# Use cases

- **Default Collections:** Whenever you need a list of items and don't know the final count. This is why they are the default "go-to" collection in almost every modern language.
- **Buffers:** Reading data from a file or a network stream where the payload size is variable.
- **Stack Implementations:** A dynamic array is the perfect underlying structure for a LIFO (Last-In, First-Out) stack since adding/removing from the end is extremely fast.

# Advantages

- **Automatic Resizing:** You don't have to manually manage memory allocation, pointer tracking, or array copying.
- **O(1) Random Access:** Just like a static array, fetching the 1,000th element is instant because the memory is still perfectly contiguous.
- **Cache-Friendly:** Because it uses contiguous memory, iterating through a dynamic array is blazing fast for the CPU cache.

# Disadvantages

- **The Reallocation Penalty:** When the array hits its capacity limit, it has to find a brand new, larger block of contiguous memory, copy _every single existing element_ over to the new block, and delete the old one. This makes that specific insertion an $O(n)$ operation, causing a sudden (though brief) latency spike.
- **Memory Waste:** To avoid reallocating too often, dynamic arrays usually grow by a specific factor (often doubling in capacity, e.g., from 16 to 32 to 64). If you only end up needing 33 elements, you are locking up space for 64.
- **Slow Inserts/Deletes at the Front:** If you want to insert an item at index `0`, you have to shift every single other element in the array one spot to the right to make room.

# Examples

- **C++:** `std::vector` (usually doubles its capacity when full).
- **Java:** `ArrayList` (usually grows by 50% when full).
- **Python:** `list` (has a more complex growth pattern, but behaves as a dynamic array).
- **JavaScript:** The standard `Array` (engine implementations vary, but conceptually acts dynamically).