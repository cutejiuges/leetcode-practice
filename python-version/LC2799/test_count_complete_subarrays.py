from unittest import TestCase
from .count_complete_subarrays import Solution


class TestSolution(TestCase):
    def test_count_complete_subarrays(self):
        solution = Solution()
        self.assertEqual(4, solution.countCompleteSubarrays(nums=[1, 3, 1, 2, 2]))
        self.assertEqual(10, solution.countCompleteSubarrays(nums=[5, 5, 5, 5]))
