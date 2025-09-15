from unittest import TestCase
from .can_be_typed_words import Solution


class TestSolution(TestCase):
    def test_can_be_typed_words(self):
        solution = Solution()
        self.assertEqual(1, solution.canBeTypedWords(text="hello world", broken_letters="ad"))
        self.assertEqual(1, solution.canBeTypedWords(text="leet code", broken_letters="lt"))
        self.assertEqual(0, solution.canBeTypedWords(text="leet code", broken_letters="e"))
