#pragma once
#include <vector>


namespace Search {

    template<typename T>
    constexpr int lower_bound(int target, std::vector<T> arr) {
         int res{arr.size()};
         int low{0};
         int high{res -1};

        while low <= high {
            int mid {low + (high - low)/2 };

            if (target < arr[mid]) {
                res = mid;
                high = mid -1;
            }
            else {
                low = mid + 1;
            }
        }
        return res;
    }
}