import unittest

class Solution:
    def waysToSplitArray(self, nums: list[int]) -> int:
        n = len(nums)
        left, right = 0, sum(nums)
        ans = 0
        for i in range(n-1):
            left += nums[i]
            right -= nums[i]
            if left >= right:
                ans += 1
        return ans
    
class SolutionTest(unittest.TestCase):
    def TestWaysToSplitArray(self):
        solution = Solution()
        self.assertEqual(2, solution.waysToSplitArray([10,4,-8,7]))
        self.assertEqual(2, solution.waysToSplitArray([2,3,1,0]))

if __name__ == '__main__':
    unittest.main()