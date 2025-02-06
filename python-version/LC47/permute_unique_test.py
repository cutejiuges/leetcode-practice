import unittest
from permute_unique import Solution

class SolutionTest(unittest.TestCase):
    def test_permuteUnique(self):
        solution = Solution()
        self.assertEqual([[1,1,2],[1,2,1],[2,1,1]], solution.permuteUnique(nums = [1,1,2]))
        self.assertEqual([[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]], solution.permuteUnique(nums = [1,2,3]))

if __name__ == '__main__':
    unittest.main()