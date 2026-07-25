#pragma once

#include <iostream>

namespace myArray {
    template<typename T>
    class Array {
    public:
        explicit Array(int size);
        Array(const Array& arr) = default;
        Array& operator=(const Array& arr) = default;
        Array& operator=(Array&& arr) = default;
        Array(Array&& arr) = default;
        ~Array() {
            delete[] elements_;
            std::cout << "array deallocated.\n";
        }

        const T& operator[](const int& index) const {
            if (index > size_) {
                throw std::runtime_error{"index out of range"};
            }
            return elements_[index];
        }

        T& operator[](const int& index) {
            if (index > size_) {
                throw std::runtime_error{"index out of range"};
            }
            return elements_[index];
        
            
        }
        const int& size() const noexcept {
            return size_;
        } 
        void push_back(T&& item) noexcept
        {
            if (size_ == capacity_) {
                auto newArray = new(elements_) T[capacity_ * 2];
                for (auto i=0;i<size_;++i) {
                    newArray[i] = elements_[i];
                }
                elements_ = newArray;
                newArray = nullptr;
            } 
            elements_[size_] = std::forward<T>(item);
            ++size_;
        }

    private:
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



}