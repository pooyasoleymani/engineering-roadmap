
## Q

 What is a stable sort?

## A

Preserves the relative order of elements with equal keys

---
## Q 

 Why do databases often require stable sorting?

## A

because database `GROUP BY` depend on stability

---
## Q

 What is an in-place algorithm?

## A

algorithms that don't use additional memory 

 ---
## Q
 
 Why can't comparison sorting beat **O(n log n)**?

## A

Binary decisions create a decision tree, and tree minimum height is

```
log₂(n!)
```

Using Stirling's approximation:

```
log₂(n!)

≈

n log₂ n
```

---
## Q 
 
 Why can Counting Sort achieve **O(n)**?

## A

Because it doesn't compare elements.

Instead it uses knowledge about the data.

---
## Q

 What makes an algorithm adaptive?

## A

Adaptive algorithms exploit existing order.

---
## Q

 Why do modern sorting algorithms care about CPU caches?

## A

Because complexity is not all thing and algorithms run on hardware 

---
## Q

 Why doesn't every language simply use Merge Sort?


## A

i think its about memory allocation  