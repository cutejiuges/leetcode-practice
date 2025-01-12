import unittest

class Solution:
    def largestCombination(self, candidates: list[int]) -> int:
        # 求出数组的最大值
        mx = max(candidates)
        # 计算最大值的二进制位数
        maxLen = mx.bit_length()

        # 枚举二进制位，计算满足条件的数字个数
        ans = 0
        for i in range(maxLen):
            curCnt = 0
            for num in candidates:
                curCnt += num >> i & 1
            ans = max(ans, curCnt)
        return ans
    
class SolutionTest(unittest.TestCase):
    def TestLargestCombination(self):
        solution = Solution()
        self.assertEqual(4, solution.largestCombination([16,17,71,62,12,24,14]))
        self.assertEqual(2, solution.largestCombination([8, 8]))


if __name__ == '__main__':
    unittest.main()