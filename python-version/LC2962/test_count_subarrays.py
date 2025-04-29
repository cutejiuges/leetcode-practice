from unittest import TestCase
from .count_subarrays import Solution


class TestSolution(TestCase):
    def test_count_subarrays(self):
        solution = Solution()
        self.assertEqual(6, solution.countSubarrays(nums=[1, 3, 2, 3, 3], k=2))
        self.assertEqual(0, solution.countSubarrays(nums=[1, 4, 2, 1], k=3))
