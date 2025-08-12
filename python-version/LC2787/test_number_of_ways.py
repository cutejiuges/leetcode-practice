from unittest import TestCase
from .number_of_ways import Solution


class TestSolution(TestCase):
    def test_number_of_ways(self):
        solution = Solution()
        self.assertEqual(1, solution.numberOfWays(10, 2))
        self.assertEqual(2, solution.numberOfWays(4, 1))
