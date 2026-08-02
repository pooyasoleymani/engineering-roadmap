#include <catch2/catch_test_macros.hpp>
#include <vector>
#include <algorithm>
#include "sorting.hpp"


TEST_CASE("Selection sort", "[sort-methods]") {
    std::vector<int> arr {1,23,56,45,23,5,6,7,2,22,33,33,55};
    auto expected = arr;
    std::sort(expected.begin(), expected.end());
    SelectionSort<int>(arr);
    REQUIRE(arr == expected);
};

TEST_CASE("bubble sort", "[sort-methods]") {
    std::vector<int> arr {1,23,56,45,23,5,6,7,2,22,33,33,55};
    auto expected = arr;
    std::sort(expected.begin(), expected.end());
    BubbleSort<int>(arr);
    REQUIRE(arr == expected);
}