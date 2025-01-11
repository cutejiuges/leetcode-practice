import unittest

class Solution:
    def generateKey(self, num1: int, num2: int, num3: int) -> int:
        pow, ans = 1, 0
        while num1 and num2 and num3:
            ans += min(num1 % 10, num2 % 10, num3 % 10) * pow
            pow *= 10
            num1, num2, num3 = num1 // 10, num2 // 10, num3 // 10
        return ans
    

class SolutionTest(unittest.TestCase):
    def test_generateKey(self):
        solution = Solution()
        self.assertEqual(solution.generateKey(1, 10, 1000), 0)
        self.assertEqual(solution.generateKey(987, 879, 798), 777)
        self.assertEqual(solution.generateKey(1, 2, 3), 1)

if __name__ == '__main__':
    unittest.main()