from unittest import TestCase
from .is_zero_array import Solution


class TestSolution(TestCase):
    def test_is_zero_array(self):
        solution = Solution()
        self.assertTrue(solution.isZeroArray(nums=[1, 0, 1], queries=[[0, 2]]))
        self.assertFalse(solution.isZeroArray(nums=[4, 3, 2, 1], queries=[[1, 3], [0, 2]]))
