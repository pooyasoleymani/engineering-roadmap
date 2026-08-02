## Definition

PDQSort (Pattern-Defeating `QuickSort`) is a fast, hybrid sorting algorithm created by Orson Peters that improves on `introsort` by adapting to existing patterns in data to achieve O(n) best-case and (O(n log n)) worst-case performance. It is unstable, operates in-place, and is used as the default unstable sort in languages like Rust (`slice::sort_unstable`) and Go (since version 1.19). 
Key Characteristics

- **Time Complexity**: O(n) best-case for structured inputs, \(O(n \log n)\) average and worst-case.
- **Memory**: In-place sorting that does not allocate extra memory.
- **Stability**: Unstable sorting algorithm (may reorder equal elements).
- **Hybrid Approach**: Combines `quicksort`, insertion sort, and `heapsort`. 

## Why It Is Fast

- **Pattern Detection**: It detects pre-sortedness, reverse-sortedness, and high counts of equal elements, switching strategies to defeat worst-case quadratic triggers.

- **Branchless Optimizations**: Specialized branchless partitioning paths for primitive or arithmetic types reduce CPU branch mispredictions.


## Advantages

- **O(n) Best Case**: Runs in linear time for fully sorted, reverse-sorted, or identical-element arrays.
- **No Worst Case**: Switches to **Heapsort** if it detects bad partitions, preventing the O(n²) trap of traditional Quicksort.
- **In-Place**: Operates with O(log n) auxiliary space, making it highly memory efficient.
- **Branchless Partitioning**: Boosts speed on modern CPUs by eliminating branch mispredictions for primitive data types.
- **Cache-Friendly**: Leverages the high cache locality inherent to Quicksort implementations.

## Disadvantages

- **Unstable**: Does not preserve the original relative order of duplicate elements.
- **Pivot Overhead**: Consumes extra CPU cycles choosing pivots (median-of-three or median-of-nine) on highly chaotic data.
- **Complex to Implement**: Combines three algorithms with strict fallback thresholds, making it harder to write and debug than standard Quicksort.
- **Type-Dependent Gain**: The branchless optimization benefits mostly primitive types (numbers); complex objects see smaller gains.