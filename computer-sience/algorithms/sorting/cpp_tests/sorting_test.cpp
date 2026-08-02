#include <catch2/catch_test_macros.hpp>
#include <vector>
#include "sorting.hpp"


TEST_CASE("Selection sort", "[sort-methods]") {
    std::vector<int> arr {1,23,56,45,23,5,6,7,2,22,33,33,55};
    std::vector<int> sorted{1, 2, 5, 6, 7, 22, 23, 23, 33, 33, 45, 55, 56};

    SelectionSort<int>(arr);
    REQUIRE(arr == sorted);
};

TEST_CASE("bubble sort", "[sort-methods]") {
    std::vector<int> arr {1,23,56,45,23,5,6,7,2,22,33,33,55};
    std::vector<int> sorted{1, 2, 5, 6, 7, 22, 23, 23, 33, 33, 45, 55, 56};

    BubbleSort<int>(arr);
    REQUIRE(arr == sorted);
}