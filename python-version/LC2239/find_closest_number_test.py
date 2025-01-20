import unittest
from find_closest_number import Solution

class SolutionTest(unittest.TestCase):
    def test_findClosestNumber(self):
        solution = Solution()
        self.assertEqual(1, solution.findClosestNumber(nums = [-4,-2,1,4,8]))
        self.assertEqual(1, solution.findClosestNumber(nums = [2,-1,1]))

if __name__ == '__main__':
    unittest.main()