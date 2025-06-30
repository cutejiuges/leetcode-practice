from unittest import TestCase
from .find_LHS import Solution


class TestSolution(TestCase):
    def test_find_lhs(self):
        solution = Solution()
        self.assertEqual(5, solution.findLHS(nums=[1, 3, 2, 2, 5, 2, 3, 7]))
        self.assertEqual(2, solution.findLHS(nums=[1, 2, 3, 4]))
        self.assertEqual(0, solution.findLHS(nums=[1, 1, 1, 1]))
