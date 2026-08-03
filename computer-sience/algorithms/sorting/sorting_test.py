from unittest import TestCase
import unittest
from sorting import *
from merge_sort import mergeSort

class TestSorting(TestCase):
    def setUp(self):
        self.data = [1,23,56,45,23,5,6,7,2,22,33,33,55]
    
    def test_bubble_sort(self) -> None:
        soreted_data = [1, 2, 5, 6, 7, 22, 23, 23, 33, 33, 45, 55, 56]
        bubble_sort(self.data)
        self.assertEqual(self.data, soreted_data)
        
        
    def test_selection_sort(self) -> None:
            soreted_data = [1, 2, 5, 6, 7, 22, 23, 23, 33, 33, 45, 55, 56]
            selection_sort(self.data)
            self.assertEqual(self.data, soreted_data)
            
    def test_irsertion_sort(self) -> None:
                soreted_data = [1, 2, 5, 6, 7, 22, 23, 23, 33, 33, 45, 55, 56]
                insertion_sort(self.data)
                self.assertEqual(self.data, soreted_data)
            

    def test_merge_sort(self) -> None:
                    soreted_data = [1, 2, 5, 6, 7, 22, 23, 23, 33, 33, 45, 55, 56]
                    mergeSort(self.data, 0, len(self.data) -1)
                    self.assertEqual(self.data, soreted_data)
if __name__ == "__main__":
    unittest.main()