# Memory Layout Interview Notes

## Q

Why are arrays usually faster than linked lists?

## A

Arrays use contiguous memory, improving cache locality and enabling hardware prefetching.

Linked lists require pointer chasing, increasing cache misses.

---

## Q

What is a cache line?

## A

The smallest block transferred between RAM and CPU cache.

Usually 64 bytes.

---

## Q

What is spatial locality?

## A

Accessing nearby memory locations sequentially.

---

## Q

What is temporal locality?

## A

Recently accessed data is likely to be accessed again.

---

## Q

What is pointer chasing?

## A

Following pointers across scattered memory locations, often causing cache misses.

---

## Q

Why is heap allocation slower than stack allocation?

## A

Heap allocation requires dynamic memory management and often garbage collection.

Stack allocation is just adjusting the stack pointer.