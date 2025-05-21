from unittest import TestCase
from .min_zero_array import Solution


class TestSolution(TestCase):
    def test_min_zero_array(self):
        solution = Solution()
        self.assertEqual(2, solution.minZeroArray(nums=[2, 0, 2], queries=[[0, 2, 1], [0, 2, 1], [1, 1, 3]]))
        self.assertEqual(-1, solution.minZeroArray(nums=[4, 3, 2, 1], queries=[[1, 3, 2], [0, 2, 1]]))
        self.assertEqual(4, solution.minZeroArray(nums=[7, 6, 8], queries=[[0, 0, 2], [0, 1, 5], [2, 2, 5], [0, 2, 4]]))
