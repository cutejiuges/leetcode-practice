from unittest import TestCase
from .possible_string_count import Solution


class TestSolution(TestCase):
    def test_possible_string_count(self):
        solution = Solution()
        self.assertEqual(5, solution.possibleStringCount(word="abbcccc"))
        self.assertEqual(1, solution.possibleStringCount(word="abcd"))
        self.assertEqual(4, solution.possibleStringCount(word="aaaa"))
