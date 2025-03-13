from unittest import TestCase
import unittest
from .count_of_substrings import Solution


class TestSolution(TestCase):
    def test_count_of_substrings(self):
        solution = Solution()
        self.assertEqual(0, solution.countOfSubstrings(word="aeioqq", k=1))
        self.assertEqual(1, solution.countOfSubstrings(word="aeiou", k=0))
        self.assertEqual(3, solution.countOfSubstrings(word="ieaouqqieaouqq", k=1))


if __name__ == '__main__':
    unittest.main()
