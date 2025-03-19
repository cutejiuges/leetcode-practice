import unittest
from .find_matrix import Solution


class TestSolution(unittest.TestCase):
    def test_find_matrix(self):
        solution = Solution()
        self.assertEqual([[1, 3, 4, 2], [1, 3], [1]], solution.findMatrix(nums=[1, 3, 4, 1, 2, 3, 1]))
        self.assertEqual([[4, 3, 2, 1]], solution.findMatrix(nums=[1, 2, 3, 4]))


if __name__ == '__main__':
    unittest.main()
