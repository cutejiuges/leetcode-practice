from unittest import TestCase
from .num_equiv_domino_pairs import Solution


class TestSolution(TestCase):
    def test_num_equiv_domino_pairs(self):
        solution = Solution()
        self.assertEqual(1, solution.numEquivDominoPairs(dominoes=[[1, 2], [2, 1], [3, 4], [5, 6]]))
        self.assertEqual(3, solution.numEquivDominoPairs(dominoes=[[1, 2], [1, 2], [1, 1], [1, 2], [2, 2]]))
