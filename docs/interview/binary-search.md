# Interview Notes

## Q

Why is binary search fundamentally about eliminating impossible regions rather than "finding" an element?


## A 

Binary search works by repeatedly eliminating half of the search space that cannot contain the answer. The loop invariant guarantees that if the answer exists, it always remains inside the current search interval. The invariant is the correctness mechanism; eliminating impossible regions is the strategy.

---

## Q 

What is a loop invariant?

## A

A condition that is true before the loop starts, remains true after every iteration, and is still true when the loop terminates.

---

## Q 

Why does `low = mid` sometimes lead to an infinite loop?


## A

The real reason is **lack of progress**


---

## Q

 Why do databases rely on `lower_bound` more often than ordinary binary search?


## A

it allows index structures like [[B-tree]] to instantly jump to the start of a range query and perform fast sequential scans

---

## Q

What is a monotonic predicate?


## A

A monotonic predicate is a function whose truth value changes only once.

Example:

```
False
False
False
True
True
True
```


---

## Q

Give three real-world optimization problems where binary search can be applied to the answer space.


## A


- cache size 
- buffer size 
- timeout
- maximum packet size
- number of servers 
- compensation level

If answer space is monotonic we can apply with binary search.

