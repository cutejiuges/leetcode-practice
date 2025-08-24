from unittest import TestCase
from .find_diagonal_order import Solution


class TestSolution(TestCase):
    def test_find_diagonal_order(self):
        solution = Solution()
        self.assertEqual([1, 2, 4, 7, 5, 3, 6, 8, 9], solution.findDiagonalOrder(mat=[[1, 2, 3], [4, 5, 6], [7, 8, 9]]))
        self.assertEqual([1, 2, 3, 4], solution.findDiagonalOrder(mat=[[1, 2], [3, 4]]))
