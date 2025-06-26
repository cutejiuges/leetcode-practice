from unittest import TestCase
from .longest_subsequence import Solution


class TestSolution(TestCase):
    def test_longest_subsequence(self):
        solution = Solution()
        self.assertEqual(5, solution.longestSubsequence(s="1001010", k=5))
        self.assertEqual(6, solution.longestSubsequence(s="00101001", k=1))
