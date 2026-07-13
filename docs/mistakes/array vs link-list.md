
Search in link-list is O(n)
Search in array is O(n)

## Trade-off

| Operation      | Array          | Linked List         |
| -------------- | -------------- | ------------------- |
| Index          | O(1)           | O(n)                |
| Search         | O(n)           | O(n)                |
| Insert End     | O(1) amortized | O(1)                |
| Insert Middle  | O(n)           | O(n) (after search) |
| Delete         | O(n)           | O(n) (after search) |
| Cache Friendly | Excellent      | Poor                |
| Memory Usage   | Low            | High                |
| Random Access  | Excellent      | Impossible          |

A senior engineer doesn't choose based only on asymptotic complexity. They also consider:

- Cache locality
- Memory overhead
- Allocation frequency
- Access patterns
- Expected workload