from unittest import TestCase
from .has_increasing_subarrays import Solution


class TestSolution(TestCase):
    def test_has_increasing_subarrays(self):
        solution = Solution()
        self.assertTrue(solution.hasIncreasingSubarrays(nums=[2, 5, 7, 8, 9, 2, 3, 4, 3, 1], k=3))
        self.assertFalse(solution.hasIncreasingSubarrays(nums=[1, 2, 3, 4, 4, 4, 4, 5, 6, 7], k=5))
