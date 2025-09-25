from unittest import TestCase
from .minimum_total import Solution


class TestSolution(TestCase):
    def test_minimum_total(self):
        solution = Solution()
        self.assertEqual(11, solution.minimumTotal(triangle=[[2], [3, 4], [6, 5, 7], [4, 1, 8, 3]]))
        self.assertEqual(-10, solution.minimumTotal(triangle=[[-10]]))
