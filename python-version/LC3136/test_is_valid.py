from unittest import TestCase
from .is_valid import Solution


class TestSolution(TestCase):
    def test_is_valid(self):
        solution = Solution()
        self.assertTrue(solution.isValid(word="234Adas"))
        self.assertFalse(solution.isValid(word="b3"))
        self.assertFalse(solution.isValid(word="a3$e"))
