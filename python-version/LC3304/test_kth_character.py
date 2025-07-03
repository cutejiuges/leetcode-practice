from unittest import TestCase
from .kth_character import Solution


class TestSolution(TestCase):
    def test_kth_character(self):
        solution = Solution()
        self.assertEqual('b', solution.kthCharacter(5))
        self.assertEqual('c', solution.kthCharacter(10))
