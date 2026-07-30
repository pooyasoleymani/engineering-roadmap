#include "binary_search.h"
#include <catch2/catch_test_macros.hpp>

TEST_CASE("binary_search", "[lower_bound]") {
    std::vector<int> arr {1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10}; 
    REQUIRE(Search::lower_bound(6, arr) == 5);
}

