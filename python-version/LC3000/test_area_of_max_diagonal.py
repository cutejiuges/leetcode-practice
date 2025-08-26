from unittest import TestCase
from .area_of_max_diagonal import Solution


class TestSolution(TestCase):
    def test_area_of_max_diagonal(self):
        solution = Solution()
        self.assertEqual(48, solution.areaOfMaxDiagonal(dimensions=[[9, 3], [8, 6]]))
        self.assertEqual(12, solution.areaOfMaxDiagonal(dimensions=[[3, 4], [4, 3]]))
