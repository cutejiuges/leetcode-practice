import unittest
from count_balls import Solution

class SolutionTest(unittest.TestCase):
    def test_counrBalls(self):
        solution = Solution()
        self.assertEqual(2, solution.countBalls(lowLimit = 1, highLimit = 10))
        self.assertEqual(2, solution.countBalls(lowLimit = 5, highLimit = 15))
        self.assertEqual(2, solution.countBalls(lowLimit = 19, highLimit = 28))

if __name__ == '__main__':
    unittest.main()