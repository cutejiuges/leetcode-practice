from unittest import TestCase
from .match_players_and_traniners import Solution


class TestSolution(TestCase):
    def test_match_players_and_trainers(self):
        solution = Solution()
        self.assertEqual(2, solution.matchPlayersAndTrainers(players=[4, 7, 9], trainers=[8, 2, 5, 8]))
        self.assertEqual(1, solution.matchPlayersAndTrainers(players=[1, 1, 1], trainers=[10]))
