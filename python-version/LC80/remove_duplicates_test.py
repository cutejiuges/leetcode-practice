import unittest
from remove_duplicates import Solution

class SolutionTest(unittest.TestCase):
    def test_removeDuplicates(self):
        solution = Solution()
        self.assertEqual(5, solution.removeDuplicates(nums = [1,1,1,2,2,3]))
        self.assertEqual(7, solution.removeDuplicates(nums = [0,0,1,1,1,1,2,3,3]))

if __name__ == '__main__':
    unittest.main()