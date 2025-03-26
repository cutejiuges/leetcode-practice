import unittest
from .minimum_sum import Solution


class TestSolution(unittest.TestCase):
    def test_minimum_sum(self):
        solution = Solution()
        self.assertEqual(18, solution.minimumSum(n=5, k=4))
        self.assertEqual(3, solution.minimumSum(n=2, k=6))


if __name__ == '__main__':
    unittest.main()
