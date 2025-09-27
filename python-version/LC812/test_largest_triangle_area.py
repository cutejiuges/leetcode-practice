from unittest import TestCase
from .largest_triangle_area import Solution


class TestSolution(TestCase):
    def test_largest_triangle_area(self):
        solution = Solution()
        self.assertTrue(
            abs(solution.largestTriangleArea(points=[[0, 0], [0, 1], [1, 0], [0, 2], [2, 0]]) - 2.00000) < 1e-5)
        self.assertTrue(abs(solution.largestTriangleArea(points=[[1, 0], [0, 0], [0, 1]]) - 0.50000) < 1e-5)
