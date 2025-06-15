from unittest import TestCase
from .max_diff import Solution


class TestSolution(TestCase):
    def test_max_diff(self):
        solution = Solution()
        self.assertEqual(888, solution.maxDiff(555))
        self.assertEqual(8, solution.maxDiff(9))
        self.assertEqual(820000, solution.maxDiff(123456))
        self.assertEqual(80000, solution.maxDiff(10000))
        self.assertEqual(8700, solution.maxDiff(9288))
