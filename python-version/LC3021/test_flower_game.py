from unittest import TestCase
from .flower_game import Solution


class TestSolution(TestCase):
    def test_flower_game(self):
        solution = Solution()
        self.assertEqual(3, solution.flowerGame(3, 2))
        self.assertEqual(0, solution.flowerGame(1, 1))
