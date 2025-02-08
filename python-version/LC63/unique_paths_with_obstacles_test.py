import unittest
from unique_paths_with_obstacles import Solution

class SolutionTest(unittest.TestCase):
    def test_uniquePathsWithObstacles(self):
        solution = Solution()
        self.assertEqual(2, solution.uniquePathsWithObstacles(obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]))
        self.assertEqual(1, solution.uniquePathsWithObstacles(obstacleGrid = [[0,1],[0,0]]))

if __name__ == '__main__':
    unittest.main()