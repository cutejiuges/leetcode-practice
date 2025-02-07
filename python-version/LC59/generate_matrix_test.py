import unittest
from generate_matrix import Solution

class SolutionTest(unittest.TestCase):
    def test_generateMatrix(self):
        solution = Solution()
        self.assertEqual([[1,2,3],[8,9,4],[7,6,5]], solution.generateMatrix(n = 3))
        self.assertEqual([[1]], solution.generateMatrix(n = 1))

if __name__ == '__main__':
    unittest.main()