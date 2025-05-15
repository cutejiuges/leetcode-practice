from unittest import TestCase
from .get_longest_subsequence import Solution


class TestSolution(TestCase):
    def test_get_longest_subsequence(self):
        solution = Solution()
        self.assertEqual(["e", "b"], solution.getLongestSubsequence(words=["e", "a", "b"], groups=[0, 0, 1]))
        self.assertEqual(["a", "b", "c"],
                         solution.getLongestSubsequence(words=["a", "b", "c", "d"], groups=[1, 0, 1, 1]))
