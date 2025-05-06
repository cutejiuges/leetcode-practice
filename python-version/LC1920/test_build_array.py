from unittest import TestCase
from .build_array import Solution


class TestSolution(TestCase):
    def test_build_array(self):
        solution = Solution()
        self.assertEqual([0, 1, 2, 4, 5, 3], solution.buildArray(nums=[0, 2, 1, 5, 3, 4]))
        self.assertEqual([4, 5, 0, 1, 2, 3], solution.buildArray(nums=[5, 0, 1, 2, 3, 4]))
