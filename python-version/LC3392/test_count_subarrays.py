from unittest import TestCase
from .count_subarrays import Solution


class TestSolution(TestCase):
    def test_count_subarrays(self):
        solution = Solution()
        self.assertEqual(1, solution.countSubarrays(nums=[1, 2, 1, 4, 1]))
        self.assertEqual(0, solution.countSubarrays(nums=[1, 1, 1]))
