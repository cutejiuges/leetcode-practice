from unittest import TestCase
from .find_words_containing import Solution


class TestSolution(TestCase):
    def test_find_words_containing(self):
        solution = Solution()
        self.assertEqual([0, 1], solution.findWordsContaining(words=["leet", "code"], x="e"))
        self.assertEqual([0, 2], solution.findWordsContaining(words=["abc", "bcd", "aaaa", "cbc"], x="a"))
        self.assertEqual([], solution.findWordsContaining(words=["abc", "bcd", "aaaa", "cbc"], x="z"))
