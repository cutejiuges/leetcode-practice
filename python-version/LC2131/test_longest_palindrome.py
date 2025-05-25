from unittest import TestCase
from .longest_palindrome import Solution


class TestSolution(TestCase):
    def test_longest_palindrome(self):
        solution = Solution()
        self.assertEqual(6, solution.longestPalindrome(words=["lc", "cl", "gg"]))
        self.assertEqual(8, solution.longestPalindrome(words=["ab", "ty", "yt", "lc", "cl", "ab"]))
        self.assertEqual(2, solution.longestPalindrome(words=["cc", "ll", "xx"]))
