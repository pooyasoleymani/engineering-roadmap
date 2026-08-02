---
Date: 2026-08-02
tags:
  - software_engineering
Related: "[[lesson-1.6.1-sorting]]"
---
---
## `std::sort`
- defined in header `<algorithm>`
- Sorts the elements in the range `[first, last)` in non-descending order. The order of equal elements is not guaranteed to be preserved.

```cpp
// constexpr C++20
template< class RandomIt >
void sort( RandomIt first, RandomIt last );

//constexpr  C++20
template< class RandomIt, class Compare >
void sort( RandomIt first, RandomIt last, Compare comp );

// C++17
template< class ExecutionPolicy, class RandomIt >
void sort( ExecutionPolicy&& policy,
           RandomIt first, RandomIt last );
           
// C++17
template< class ExecutionPolicy, class RandomIt, class Compare >
void sort( ExecutionPolicy&& policy,
           RandomIt first, RandomIt last, Compare comp );
```


### Parameters

|             |     |                                                                                                                                                                                                                                                                       |
| ----------- | --- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| first, last | -   | the pair of iterators defining the [range](https://en.cppreference.com/cpp/iterator#Ranges "cpp/iterator") of elements to sort                                                                                                                                        |
| policy      | -   | the [execution policy](https://en.cppreference.com/cpp/algorithm/execution_policy_tag_t "cpp/algorithm/execution policy tag t") to use                                                                                                                                |
| comp        | -   | comparison function object (i.e. an object that satisfies the requirements of [Compare](https://en.cppreference.com/cpp/named_req/Compare "cpp/named req/Compare")) which returns ​`true` if the first argument is _less_ than (i.e. is ordered _before_) the second. |

### Complexity

```
O(N.logN)
```


### Example

```cpp
#include <algorithm>
#include <array>
#include <functional>
#include <iostream>
#include <string_view>

int main()
{
    std::array<int, 10> s{5, 7, 4, 2, 8, 6, 1, 9, 0, 3};
    
    auto print = [&s](std::string_view const rem)
    {
        for (auto a : s)
            std::cout << a << ' ';
        std::cout << ": " << rem << '\n';
    };
    
    std::sort(s.begin(), s.end());
    print("sorted with the default operator<");
    
    std::sort(s.begin(), s.end(), std::greater<int>());
    print("sorted with the standard library compare function object");
    
    struct
    {
        bool operator()(int a, int b) const { return a < b; }
    }
    customLess;
    
    std::sort(s.begin(), s.end(), customLess);
    print("sorted with a custom function object");
    
    std::sort(s.begin(), s.end(), [](int a, int b)
                                  {
                                      return a > b;
                                  });
    print("sorted with a lambda expression");
```