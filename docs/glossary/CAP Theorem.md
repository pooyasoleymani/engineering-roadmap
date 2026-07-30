## Definition

states that ==a distributed data store can simultaneously provide at most **two out of three** guarantees: **Consistency**, **Availability**, and **Partition Tolerance**==.

Because physical networks will inevitably experience delays or drops, you must always choose **Partition Tolerance**. Therefore, the practical trade-off is always between **Consistency** and **Availability**.


### The 3 Core Guarantees

- **Consistency (C):** Every read receives the most recent write or an error.
- **Availability (A):** Every non-failing node returns a non-error response, without guaranteeing it contains the most recent write.
- **Partition Tolerance (P):** The system continues to operate despite an arbitrary number of messages being dropped or delayed by the network between nodes.