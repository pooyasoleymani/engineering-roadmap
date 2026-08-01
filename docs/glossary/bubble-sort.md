## Definition

Bubble sort is ==a simple sorting algorithm that repeatedly compares adjacent elements and swaps them if they are in the wrong order==.

## How It Works

- **Compare pairs:** Look at two neighboring items in a list.
- **Swap if needed:** If the first item is bigger than the second item, switch their places.
- **Move down the list:** Go step-by-step to the end of the list so the largest value "bubbles" up to the last spot.
- **Repeat:** Do it again for the remaining items until no swaps are needed during a full pass.


## Performance and Features

- **Time complexity:** O(n²) for average and worst cases, making it slow for large lists.

- **Space complexity:** O(1) because it sorts items right in place without extra memory.

- **Stable:** Equal values keep their original order.


## Core Trade-Off
- **Simplicity vs. Efficiency**: You trade **speed** for **simplicity**. It requires minimal code and zero extra memory, but it becomes incredibly slow as your dataset grows.

## Advantages

- **Easy to understand**: The logic is highly intuitive for beginners learning algorithms.
- **Easy to implement**: It requires only a few lines of code and nested loops.
- **In-place sorting**: It operates directly on the original list, requiring an auxiliary space complexity of **O(1)**.
- **Stable sort**: It preserves the relative order of duplicate elements.
- **Fast for nearly-sorted data**: With a flag modifier, it can achieve a best-case time complexity of **O(n)** if the list is already sorted.

## Disadvantages

- **Poor time efficiency**: Its average and worst-case time complexity is **O(n²)**, making it highly inefficient for large datasets.
- **Excessive swaps**: It moves elements one step at a time, resulting in a massive number of write operations compared to algorithms like Quick Sort or Merge Sort.
- **Not scalable**: It does not scale well in real-world production applications.

