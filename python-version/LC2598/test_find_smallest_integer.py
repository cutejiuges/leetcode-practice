from unittest import TestCase
from .find_smallest_integer import Solution


class TestSolution(TestCase):
    def test_find_smallest_integer(self):
        solution = Solution()
        self.assertEqual(4, solution.findSmallestInteger(nums=[1, -10, 7, 13, 6, 8], value=5))
        self.assertEqual(2, solution.findSmallestInteger(nums=[1, -10, 7, 13, 6, 8], value=7))
