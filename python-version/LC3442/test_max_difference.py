from unittest import TestCase
from .max_difference import Solution


class TestSolution(TestCase):
    def test_max_difference(self):
        solution = Solution()
        self.assertEqual(3, solution.maxDifference(s="aaaaabbc"))
        self.assertEqual(1, solution.maxDifference(s="abcabcab"))
