#pragma once

#include <iostream>

namespace myArray {
    template<typename T>
    class Array {
    public:
        explicit Array(int size);
        Array(const Array& arr);
        Array& operator==(const Array& arr);
        Array& operator==(Array&& arr);
        Array(Array&& arr);

    private:
        ~Array() {
            delete[] elements_;
        }
        
        
        int size_;
        int capacity_;
        T *elements_; 
    };

    template<typename T>
    Array<T>::Array(int size):
    size_{size}, capacity_{8},elements_{new T{}}
    {
        std::cout << "new array allocated\n";
    };

    template <typename T>
    Array<T>::Array(const Array &arr)
    {

    }

}