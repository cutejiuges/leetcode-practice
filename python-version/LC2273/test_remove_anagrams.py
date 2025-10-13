from unittest import TestCase
from .remove_anagrams import Solution


class TestSolution(TestCase):
    def test_remove_anagrams(self):
        solution = Solution()
        self.assertEqual(["abba", "cd"], solution.removeAnagrams(words=["abba", "baba", "bbaa", "cd", "cd"]))
        self.assertEqual(["a", "b", "c", "d", "e"], solution.removeAnagrams(words=["a", "b", "c", "d", "e"]))
