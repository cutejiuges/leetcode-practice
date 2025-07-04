from unittest import TestCase
from .kth_character import Solution


class TestSolution(TestCase):
    def test_kth_character(self):
        solution = Solution()
        self.assertEqual('a', solution.kthCharacter(5, [0, 0, 0]))
        self.assertEqual('b', solution.kthCharacter(10, [0, 1, 0, 1]))
