# Interview Notes

## Q

Why is binary search fundamentally about eliminating impossible regions rather than "finding" an element?


## A 

Because binary search about maintaining invariant .


---

## Q 

What is a loop invariant?

## A

Processor that every time halve  invariant 

---

## Q 

Why does `low = mid` sometimes lead to an infinite loop?


## A

Because this is wrong invariant condition 


---

## Q

 Why do databases rely on `lower_bound` more often than ordinary binary search?


## A

it allows index structures like [[B-tree]] to instantly jump to the start of a range query and perform fast sequential scans

---

## Q

What is a monotonic predicate?


## A

it is technical method for answer to question with binary search and Once it becomes true, it stays true.

```text 
false false false true true true 
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

