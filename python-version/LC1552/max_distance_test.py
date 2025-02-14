import unittest
from max_distance import Solution

class SolutionTest(unittest.TestCase):
    def test_maxDistance(self):
        solution = Solution()
        self.assertEqual(3, solution.maxDistance(position = [1,2,3,4,7], m = 3))
        self.assertEqual(999999999, solution.maxDistance(position = [5,4,3,2,1,1000000000], m = 2))

if __name__ == '__main__':
    unittest.main()