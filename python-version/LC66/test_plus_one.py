from unittest import TestCase
from .plus_one import Solution


class TestSolution(TestCase):
    def test_plus_one(self):
        solution = Solution()
        self.assertEqual([1, 2, 4], solution.plusOne(digits=[1, 2, 3]))
        self.assertEqual([4, 3, 2, 2], solution.plusOne(digits=[4, 3, 2, 1]))
        self.assertEqual([1, 0], solution.plusOne(digits=[9]))
