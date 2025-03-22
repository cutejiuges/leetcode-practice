import unittest
from .row_and_maximum_ones import Solution


class TestSolution(unittest.TestCase):
    def test_row_and_maximum_ones(self):
        solution = Solution()
        self.assertEqual([0, 1], solution.rowAndMaximumOnes(mat=[[0, 1], [1, 0]]))
        self.assertEqual([1, 2], solution.rowAndMaximumOnes(mat=[[0, 0, 0], [0, 1, 1]]))
        self.assertEqual([1, 2], solution.rowAndMaximumOnes(mat=[[0, 0], [1, 1], [0, 0]]))


if __name__ == '__main__':
    unittest.main()
