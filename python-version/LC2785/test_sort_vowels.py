from unittest import TestCase
from .sort_vowels import Solution


class TestSolution(TestCase):
    def test_sort_vowels(self):
        solution = Solution()
        self.assertEqual("lEOtcede", solution.sortVowels(s="lEetcOde"))
        self.assertEqual("lYmpH", solution.sortVowels(s="lYmpH"))
