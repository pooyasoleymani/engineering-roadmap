from unittest import TestCase
import unittest
from sorting import *
from merge_sort import mergeSort
from quick_sort import quick_sort, Lomuto

class TestSorting(TestCase):
    def setUp(self):
        self.data = [1,23,56,45,23,5,6,7,2,22,33,33,55]
        self.expected = [1, 2, 5, 6, 7, 22, 23, 23, 33, 33, 45, 55, 56]
    
    def test_bubble_sort(self) -> None:
        bubble_sort(self.data)
        self.assertEqual(self.data, self.expected)
        
        
    def test_selection_sort(self) -> None:
            selection_sort(self.data)
            self.assertEqual(self.data, self.expected)
            
    def test_irsertion_sort(self) -> None:
                insertion_sort(self.data)
                self.assertEqual(self.data, self.expected)
            

    def test_merge_sort(self) -> None:
                    mergeSort(self.data, 0, len(self.data) -1)
                    self.assertEqual(self.data, self.expected)
                    
                    
    def test_quick_sort(self) -> None:
        quick_sort(self.data, Lomuto)
        self.assertEqual(self.data, self.expected)
    
if __name__ == "__main__":
    unittest.main()