#include "sorting.h"
#include <vector>


template <typename T>
void swap(std::size_t i, std::size_t j, std::vector<T>& arr) 
{
    std::vector<T> temp;
    temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}


template <typename T>
void SelectionSort(std::vector<T> arr) 
{
    auto n = arr.size();
    for (auto i=0;i<n;i++) 
    {
        auto MinIdx = i;
        for (auto j=0;j<n-1;j++) 
        {
            if (arr[MinIdx] > arr[j]) 
            {
                MinIdx = j;
            }
        }
        swap(i, MinIdx, arr);
    }
}


template <typename T>
void BubbleSort(std::vector<T> arr)
{
    auto n = arr.size();
    for (auto i=0;i<n-1;i++) 
    {
        for (auto j=i;j<n-1;j++) 
        {
            if (arr[j] > arr[j+1] )
            {
                swap(i, j, arr);
            }
        }
        
    }
}