import unittest
from max_distance import Solution

class SolutionTest(unittest.TestCase):
    def test_findSpecialInteger(self):
        solution = Solution()
        self.assertEqual(4, solution.maxDistance(arrays = [[1,2,3],[4,5],[1,2,3]]))
        self.assertEqual(0, solution.maxDistance(arrays = [[1],[1]]))

if __name__ == '__main__':
    unittest.main()