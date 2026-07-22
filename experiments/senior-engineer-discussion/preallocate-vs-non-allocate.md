---
Related: "[[dynamic-arrays-amortized-growth]]"
---
---
## Senior Engineer Discussion

Here's something to think about.

Suppose you're building a telemetry service that receives **100 million events per day**.

Two engineers propose:

**Engineer A**

```
events := []Event{}
```

**Engineer B**

```
events := make([]Event, 0, 5_000_000)
```

Both implementations are correct.

### Your challenge

Which implementation would you approve in a production code review, and what questions would you ask before making that decision?

Notice the subtlety: don't immediately pick Engineer B. A senior engineer first asks about the workload, memory budget, request lifecycle, and allocation patterns. Learning to ask those questions is as important as knowing the implementation details.


---

## Answer

Because a senior engineer doesn't answer before understanding the workload.

### 1.

Is

```
5 million
```

a guaranteed number?

or just an estimate?

---

### 2.

Will this slice live for

```
10 ms
```

or

```
24 hours
```

?

---

### 3.

Is memory plentiful?

Or are we inside a container with

```
512 MB RAM
```

?

---

### 4.

Does each request allocate

```
5 million events
```

?

or

Does the whole application process

```
5 million
```

per day?

Those are completely different.

---

### 5.

How big is one Event?

Suppose

```
type Event struct {
    Payload [4096]byte
}
```

Now

```
5,000,000 × 4 KB
```

≈ **20 GB**.

Preallocating would likely be disastrous.

If `Event` is tiny, the trade-off changes.

---

### 6.

How often is this code executed?

Once?

Every request?

Every second?

---

This is why engineering is about asking the right questions before optimizing.



---

Reference: [[lesson-1.3-arrays]]