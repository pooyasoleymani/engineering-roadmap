---
Date: 2026-08-25
tags:
  - Cpp
---
---
At the highest level:

```
C++ source code
      |
      | compiler
      v
 object files
      |
      | linker
      v
 executable
      |
      v
    CPU
```

C++ is primarily a **compiled, statically typed language**.



```cpp
int main()
{
	return 0;
}
```

`main()` is the entry point of a C++ program. Execution begins there.


---
## Functions: the fundamental unit of behavior

Function represents an meaningful operation.

```cpp
double square(double x)
{
    return x * x;
}
```


---
## Modern C++ example

Let's combine today's concepts:

```cpp
#include <iostream>

double square(double value)
{
    return value * value;
}

void print_square(double value)
{
    std::cout << "The square of "
              << value
              << " is "
              << square(value)
              << '\n';
}

int main()
{
    constexpr double value{5.0};

    print_square(value);
}
```

Output:

```
The square of 5 is 25
```

Notice the choices:

- `std::cout` instead of `using namespace std`
- meaningful function names
- `double` explicitly where it communicates the domain
- `{}` initialization
- `constexpr` for a compile-time constant
- `'\n'` rather than `std::endl`

We will study why each of these choices matters.


---
