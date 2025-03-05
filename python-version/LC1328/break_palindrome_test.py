import unittest
from break_palindrome import Solution

class SolutionTest(unittest.TestCase):
    def test_breakPalindrome(self):
        solution = Solution()
        self.assertEqual("aaccba", solution.breakPalindrome(palindrome = "abccba"))
        self.assertEqual("", solution.breakPalindrome("a"))

if __name__ == '__main__':
    unittest.main()