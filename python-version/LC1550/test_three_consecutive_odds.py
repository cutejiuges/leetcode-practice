from unittest import TestCase
from .three_consecutive_odds import Solution


class TestSolution(TestCase):
    def test_three_consecutive_odds(self):
        solution = Solution()
        self.assertFalse(solution.threeConsecutiveOdds(arr=[2, 6, 4, 1]))
        self.assertTrue(solution.threeConsecutiveOdds(arr=[1, 2, 34, 3, 4, 5, 7, 23, 12]))
