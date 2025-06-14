from unittest import TestCase
from .min_max_difference import Solution


class TestSolution(TestCase):
    def test_min_max_difference(self):
        solution = Solution()
        self.assertEqual(99009, solution.minMaxDifference(num=11891))
        self.assertEqual(99, solution.minMaxDifference(num=90))
