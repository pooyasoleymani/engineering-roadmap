
# Definition

A linked list is a linear data structure where elements are not stored in contiguous memory. Instead, the data is scattered wherever there is free space in RAM. Each element (called a **node**) contains two parts: the actual data, and a pointer (a memory address) linking to the next node in the sequence.

# Problem

It completely solves the expensive `O(n)` shifting problem of arrays. If you want to insert a new item at the very front of a dynamic array, you have to move every single existing item one spot to the right to make room.

In a linked list, inserting at the front is an `O(1)` operation. You just ask the OS for a tiny block of memory for your new node, put your data in it, and set its pointer to look at the old "head" of the list. No data gets shifted, and it takes the exact same amount of time whether the list has 10 items or 10 million.

# Use cases

- **Queues and Deques:** Scenarios where you are constantly adding and removing items from the ends of a collection.
- **Operating System Schedulers:** Managing lists of running processes where items are constantly cycling in and out.
- **Undo Functionality:** In applications like text editors, where each action is stored as a node pointing to the previous state.

# Advantages

- **O(1) Insertions/Deletions:** You can add or remove a node instantly at the front (and at the back, if you track the "tail" pointer). You can also insert instantly in the middle _if_ you already have a pointer to that specific spot.
- **No Reallocation Penalty:** The list grows node-by-node. It never hits a "capacity" limit where it has to freeze the program and copy massive blocks of memory to a new location.
- **No Wasted Capacity:** It uses exactly the amount of memory needed for the items currently in the list (plus the pointer overhead).

# Disadvantages

- **Terrible Cache Performance:** Because nodes are scattered randomly across your RAM, you lose all the benefits of spatial locality. When the CPU fetches a node, the next node isn't in the high-speed cache. This makes iterating through a linked list significantly slower in the real world than iterating through an array, even though both are theoretically `O(n)`.
- **O(n) Random Access:** You cannot instantly calculate the memory address of the 1,000th element. To find it, you have to start at the head node and follow 999 pointers, one by one.
- **Memory Overhead:** Every single piece of data requires extra memory to store the pointer. If you are storing a list of small types like booleans or small integers, the 64-bit memory address pointing to the next node might take up more RAM than the data itself!

# Examples

- **Singly Linked List:** Nodes only have one pointer looking forward to the next item.
- **Doubly Linked List:** Nodes have two pointers—one looking forward, and one looking backward. This takes up even more memory but allows you to traverse the list in both directions (like C++ `std::list`).