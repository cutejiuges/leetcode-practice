from unittest import TestCase
from .num_submat import Solution


class TestSolution(TestCase):
    def test_num_submat(self):
        solution = Solution()
        self.assertEqual(13, solution.numSubmat(mat=[[1, 0, 1], [1, 1, 0], [1, 1, 0]]))
        self.assertEqual(24, solution.numSubmat(mat=[[0, 1, 1, 0], [0, 1, 1, 1], [1, 1, 1, 0]]))
