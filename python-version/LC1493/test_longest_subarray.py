from unittest import TestCase
from .longest_subarray import Solution


class TestSolution(TestCase):
    def test_longest_subarray(self):
        solution = Solution()
        self.assertEqual(3, solution.longestSubarray(nums=[1, 1, 0, 1]))
        self.assertEqual(5, solution.longestSubarray(nums=[0, 1, 1, 1, 0, 1, 1, 0, 1]))
        self.assertEqual(2, solution.longestSubarray(nums=[1, 1, 1]))
