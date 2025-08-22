from unittest import TestCase
from .minimum_area import Solution


class TestSolution(TestCase):
    def test_minimum_area(self):
        solution = Solution()
        self.assertEqual(6, solution.minimumArea(grid=[[0, 1, 0], [1, 0, 1]]))
        self.assertEqual(1, solution.minimumArea(grid=[[0, 0], [1, 0]]))
