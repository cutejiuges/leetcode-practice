from unittest import TestCase
from .min_subarray import Solution


class TestSolution(TestCase):
    def test_min_subarray(self):
        solution = Solution()
        self.assertEqual(1, solution.minSubarray(nums=[3, 1, 4, 2], p=6))
        self.assertEqual(2, solution.minSubarray(nums=[6, 3, 5, 2], p=9))
        self.assertEqual(0, solution.minSubarray(nums=[1, 2, 3], p=3))
        self.assertEqual(-1, solution.minSubarray(nums=[1, 2, 3], p=7))
        self.assertEqual(0, solution.minSubarray(nums=[1000000000, 1000000000, 1000000000], p=3))
