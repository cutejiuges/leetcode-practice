from unittest import TestCase
from .count_squares import Solution


class TestSolution(TestCase):
    def test_count_squares(self):
        solution = Solution()
        self.assertEqual(15, solution.countSquares(matrix=[
            [0, 1, 1, 1],
            [1, 1, 1, 1],
            [0, 1, 1, 1]
        ]))
        self.assertEqual(7, solution.countSquares(matrix=[
            [1, 0, 1],
            [1, 1, 0],
            [1, 1, 0]
        ]))
