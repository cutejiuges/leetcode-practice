import unittest
from max_coins import Solution

class SolutionTest(unittest.TestCase):
    def test_maxCoins(self):
        solution = Solution()
        self.assertEqual(9, solution.maxCoins(piles = [2,4,1,2,7,8]))
        self.assertEqual(4, solution.maxCoins(piles = [2,4,5]))
        self.assertEqual(18, solution.maxCoins(piles = [9,8,7,6,5,1,2,3,4]))

if __name__ == '__main__':
    unittest.main()