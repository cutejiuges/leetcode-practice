from unittest import TestCase
from .count_subarrays import Solution


class TestSolution(TestCase):
    def test_count_subarrays(self):
        solution = Solution()
        self.assertEqual(6, solution.countSubarrays(nums=[2, 1, 4, 3, 5], k=10))
        self.assertEqual(5, solution.countSubarrays(nums=[1, 1, 1], k=5))
