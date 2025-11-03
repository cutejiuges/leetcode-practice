from unittest import TestCase
from .min_cost import Solution


class TestSolution(TestCase):
    def test_min_cost(self):
        solution = Solution()
        self.assertEqual(0, solution.minCost(colors="abc", needed_time=[1, 2, 3]))
        self.assertEqual(2, solution.minCost(colors="aabaa", needed_time=[1, 2, 3, 4, 1]))
        self.assertEqual(3, solution.minCost(colors="abaac", needed_time=[1, 2, 3, 4, 5]))
