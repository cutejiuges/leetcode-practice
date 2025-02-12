import unittest
from minimum_size import Solution

class SolutionTest(unittest.TestCase):
    def test_minimumSize(self):
        solution = Solution()
        self.assertEqual(3, solution.minimumSize(nums = [9], maxOperations = 2))
        self.assertEqual(2, solution.minimumSize(nums = [2,4,8,2], maxOperations = 4))
        self.assertEqual(7, solution.minimumSize(nums = [7,17], maxOperations = 2))

if __name__ == '__main__':
    unittest.main()