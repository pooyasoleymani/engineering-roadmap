
## Q

 What is a stable sort?

## A

Preserves the relative order of elements with equal keys

---
## Q 

 Why do databases often require stable sorting?

## A

Stable sorting is important because multiple sorting passes can be composed. For example, sort by Rating, then stably sort by Price. Equal-price products preserve their previous rating order. SQL operations such as `ORDER BY` on multiple columns often rely on this behavior, either directly or conceptually.

---
## Q

 What is an in-place algorithm?

## A

- Algorithms that don't use additional memory.
- Uses only **O(1) auxiliary memory** (ignoring recursion stack unless specified).

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

Sequential memory access improves cache locality, enables hardware prefetching, and reduces cache misses. Therefore, two algorithms with the same asymptotic complexity can have very different real-world performance.

---
## Q

 Why doesn't every language simply use Merge Sort?


## A


Additional reasons include:

- Extra memory (O(n))
- Allocation cost
- Copy overhead
- Cache behavior
- Some datasets are nearly sorted (Timsort excels)
- Some workloads favor in-place algorithms (Introsort, PDQSort)



---

## Senior Engineer Challenge

Imagine you're building an e-commerce platform.

Each product has:

```
type Product struct {
    ID       int
    Price    int
    Rating   float64
    Name     string
}
```

The product list is already sorted by **Rating**.

Now a new requirement arrives:

> Display products sorted by **Price**, but if two products have the same price, preserve their original rating order.

### Questions

1. Would you choose a **stable** or **unstable** sorting algorithm?
2. Why?
3. If your language's default sort is unstable, how could you still satisfy the requirement?
4. Would you sort twice or design a custom comparator?

This is a realistic problem you'll encounter in backend services, recommendation systems, and search ranking pipelines.

## A

1. We choose *Stable* sorting algorithm
2. because if sorting algorithm is *unstable* can't guarantee preserver original rating order.

Suppose we have:

```
Rating order

A
B
C
D
```

Prices:

```
A 100
B 200
C 100
D 300
```

Required result:

```
A 100
C 100
B 200
D 300
```

Notice

A stays before C.

---

If your language has **stable sort**:

One comparator is enough:

```
price < price
```

because stability preserves rating order.

---

If the sort is **unstable**:

A comparator

```
price < price
```

is **NOT enough**.

The sorting algorithm is allowed to swap equal-price items.

So

```
C
A
```

would also be valid.

---

How do we fix it?

### Option 1

Comparator

```
if(price != other.price)
    return price < other.price;

return rating > other.rating;
```

Now the comparator defines a total ordering.

No stability required.

---

### Option 2

Attach original position.

```
Price

↓

Original Index
```

Comparator

```
Price

↓

Original Position
```

Many production systems do exactly this.