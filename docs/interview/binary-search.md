# Interview Notes

## Q

Why isn't Big O enough?

## A

Because Big O ignores

- cache locality
- constant factors
- branch prediction
- compiler optimizations
- memory allocation
- IO latency

---

## Q

Why are arrays faster than linked lists?

## A

Arrays are contiguous.

CPU cache loads nearby elements together.

Linked lists require pointer chasing.

---

## Q

What is amortized analysis?

## A

Average cost over many operations.

Example

Dynamic array append.