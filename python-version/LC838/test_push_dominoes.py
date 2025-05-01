from unittest import TestCase
from .push_dominoes import Solution


class TestSolution(TestCase):
    def test_push_dominoes(self):
        solution = Solution()
        self.assertEqual("RR.L", solution.pushDominoes(dominoes="RR.L"))
        self.assertEqual("LL.RR.LLRRLL..", solution.pushDominoes(dominoes=".L.R...LR..L.."))
