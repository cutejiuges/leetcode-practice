from unittest import TestCase
from .maximum_difference import Solution


class TestSolution(TestCase):
    def test_maximum_difference(self):
        solution = Solution()
        self.assertEqual(4, solution.maximumDifference([7, 1, 5, 4]))
        self.assertEqual(-1, solution.maximumDifference([9, 4, 3, 2]))
        self.assertEqual(9, solution.maximumDifference([1, 5, 2, 10]))
