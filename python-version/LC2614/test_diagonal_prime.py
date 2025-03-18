import unittest
from .diagonal_prime import Solution


class TestSolution(unittest.TestCase):
    def test_diagonal_prime(self):
        solution = Solution()
        self.assertEqual(11, solution.diagonalPrime(nums=[[1, 2, 3], [5, 6, 7], [9, 10, 11]]))
        self.assertEqual(17, solution.diagonalPrime(nums=[[1, 2, 3], [5, 17, 7], [9, 11, 10]]))


if __name__ == '__main__':
    unittest.main()
