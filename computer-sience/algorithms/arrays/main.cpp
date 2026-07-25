#include <iostream>
#include "array.hpp"


int main() 
{
    myArray::Array<int> a{2};
    a[0] = 10;
    a.push_back(2);
    std::cout << "size of array: " << a.size() << '\n';
    std::cout << a[0] << '\n'; 
    std::cout << a[1] << '\n'; 
    std::cout << a[2] << '\n'; 
    std::cout << a[3] << '\n'; 
    std::cout << a[4] << '\n';
}