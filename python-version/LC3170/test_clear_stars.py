from unittest import TestCase
from .clear_stars import Solution


class TestSolution(TestCase):
    def test_clear_stars(self):
        solution = Solution()
        self.assertEqual("aab", solution.clearStars("aaba*"))
        self.assertEqual("abc", solution.clearStars("abc"))
        self.assertEqual("", solution.clearStars("d*o*"))
