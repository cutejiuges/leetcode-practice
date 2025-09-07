from unittest import TestCase
from .sum_zero import Solution


class TestSolution(TestCase):
    def test_sum_zero(self):
        solution = Solution()
        self.assertEqual(0, sum(solution.sumZero(3)))
        self.assertEqual(0, sum(solution.sumZero(1)))
        self.assertEqual(0, sum(solution.sumZero(5)))
        self.assertEqual(0, sum(solution.sumZero(6)))
