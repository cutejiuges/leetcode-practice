from unittest import TestCase
from .difference_of_sums import Solution


class TestSolution(TestCase):
    def test_difference_of_sums(self):
        solution = Solution()
        self.assertEqual(19, solution.differenceOfSums(10, 3))
        self.assertEqual(15, solution.differenceOfSums(5, 6))
        self.assertEqual(-15, solution.differenceOfSums(5, 1))
