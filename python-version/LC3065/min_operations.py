import unittest

class Solution:
    def minOperations(self, nums: list[int], k: int) -> int:
        return sum(x < k for x in nums)
    

class SolutionTest(unittest.TestCase):
    def TestMinOperations(self):
        solution = Solution()
        self.assertEqual(3, solution.minOperations([2,11,10,1,3], 10))
        self.assertEqual(0, solution.minOperations([1,1,2,4,9], 1))
        self.assertEqual(4, solution.minOperations([1,1,2,4,9], 9))

if __name__ == '__main__':
    unittest.main()