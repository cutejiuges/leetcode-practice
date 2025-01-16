import unittest

from minimum_subarray_length import Solution

class TestSolution(unittest.TestCase):
    def test_inimumSubarrayLength(self):
        solution = Solution()
        self.assertEqual(1, solution.minimumSubarrayLength(nums = [1,2,3], k = 2))
        self.assertEqual(3, solution.minimumSubarrayLength(nums = [2,1,8], k = 10))
        self.assertEqual(1, solution.minimumSubarrayLength(nums = [1,2], k = 0))

if __name__ == '__main__':
    unittest.main()