import unittest
from find_ball import Solution

class SolutionTest(unittest.TestCase):
    def test_findBall(self):
        solution = Solution()
        self.assertEqual([1,-1,-1,-1,-1], solution.findBall(grid = [[1,1,1,-1,-1],[1,1,1,-1,-1],[-1,-1,-1,1,1],[1,1,1,1,-1],[-1,-1,-1,-1,-1]]))
        self.assertEqual([-1], solution.findBall(grid = [[-1]]))
        self.assertEqual([0,1,2,3,4,-1], solution.findBall(grid = [[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1],[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1]]))

if __name__ == '__main__':
    unittest.main()