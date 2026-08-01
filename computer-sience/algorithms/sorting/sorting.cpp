#include <vector>


template <typename T>
void swap(std::size_t i, std::size_t j) 
{
    auto temp{};
    
}


template <typename T>
void SelectionSort(std::vector<T> arr) 
{
    auto n = arr.size();
    for (auto i=0;i<n;i++) 
    {
        MinIdx = i;
        for (auto j=0;j<n-1;j++) 
        {
            if arr[MinIdx] > arr[j] 
            {
                MinIdx = j;
            }
        }
        swap(i, MinIdx);
    }
}